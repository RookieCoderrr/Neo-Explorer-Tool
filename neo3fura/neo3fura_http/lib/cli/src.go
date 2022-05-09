package cli

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/joeqian10/neo3-gogogo/rpc"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	log2 "neo3fura_http/lib/log"
	"neo3fura_http/lib/type/h160"
	"neo3fura_http/var/stderr"
)

// T ...
type T struct {
	Redis     *redis.Client
	Db_online string
	C_online  *mongo.Client
	C_local   *mongo.Client
	Ctx       context.Context
	RpcCli    *rpc.RpcClient
	RpcPorts  []string
	NeoFs string
}

type Insert struct {
	Hash   h160.T
	Id 				int32
	UpdateCounter 	int32
}

func (me *T) QueryOne(args struct {
	Collection string
	Index      string
	Sort       bson.M
	Filter     bson.M
	Query      []string
}, ret *json.RawMessage) (map[string]interface{}, error) {
	var kvs string
	kvs = kvs + args.Collection
	kvs = kvs + args.Index
	for k, v := range args.Sort {
		kvs = kvs + k + fmt.Sprintf("%v", v)
	}
	for k, v := range args.Filter {
		kvs = kvs + k + fmt.Sprintf("%v", v)
	}
	for _, v := range args.Query {
		kvs = kvs + v
	}
	h := sha1.New()
	h.Write([]byte(kvs))
	hash := hex.EncodeToString(h.Sum(nil))
	val, err := me.Redis.Get(me.Ctx, hash).Result()
	// if sort != nil, it may have several results, we have to pick the sorted one
	if err != nil || len(args.Sort) != 0 || args.Index == "GetCandidateByAddress" || args.Index == "GetAssetInfoByContractHash" || args.Index == "GetVerifiedContractByContractHash" || args.Index == "GetVotesByCandidateAddress" {
		var result map[string]interface{}
		convert := make(map[string]interface{})
		collection := me.C_online.Database(me.Db_online).Collection(args.Collection)
		opts := options.FindOne().SetSort(args.Sort)
		err = collection.FindOne(me.Ctx, args.Filter, opts).Decode(&result)
		if err == mongo.ErrNoDocuments {
			return nil, stderr.ErrNotFound
		} else if err != nil {
			return nil, stderr.ErrFind1
		}
		if len(args.Query) == 0 {
			convert = result
		} else {
			for _, v := range args.Query {
				convert[v] = result[v]
			}
		}
		r, err := json.Marshal(convert)
		if err != nil {
			return nil, stderr.ErrFind2
		}
		err = me.Redis.Set(me.Ctx, hash, hex.EncodeToString(r), 0).Err()
		if err != nil {
			return nil, stderr.ErrFind3
		}
		*ret = json.RawMessage(r)
		return convert, nil
	} else {
		r, err := hex.DecodeString(val)
		if err != nil {
			return nil, stderr.ErrFind4
		}

		*ret = json.RawMessage(r)
		convert := make(map[string]interface{})
		err = json.Unmarshal(r, &convert)
		if convert["_id"] != nil {
			convert["_id"], err = primitive.ObjectIDFromHex(convert["_id"].(string))
		}
		if err != nil {
			return nil, stderr.ErrFind5
		}
		return convert, nil
	}
	return nil, nil
}

func (me *T) QueryAll(args struct {
	Collection string
	Index      string
	Sort       bson.M
	Filter     bson.M
	Query      []string
	Limit      int64
	Skip       int64
}, ret *json.RawMessage) ([]map[string]interface{}, int64, error) {

	//if args.Limit == 0 {
	//	args.Limit = limit2.DefaultLimit
	//} else if args.Limit > limit2.MaxLimit {
	//	args.Limit = limit2.MaxLimit
	//}

	var results []map[string]interface{}
	convert := make([]map[string]interface{}, 0)
	collection := me.C_online.Database(me.Db_online).Collection(args.Collection)
	op := options.Find()
	op.SetSort(args.Sort)
	op.SetLimit(args.Limit)
	op.SetSkip(args.Skip)
	co := options.CountOptions{}
	count, err := collection.CountDocuments(me.Ctx, args.Filter, &co)
	if err != nil {
		return nil, 0, stderr.ErrFind6
	}
	cursor, err := collection.Find(me.Ctx, args.Filter, op)
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			log2.Fatalf("Closing cursor error %v", err)
		}
	}(cursor, me.Ctx)
	if err == mongo.ErrNoDocuments {
		return nil, 0, stderr.ErrNotFound
	}
	if err != nil {
		return nil, 0, stderr.ErrFind7
	}
	if err = cursor.All(me.Ctx, &results); err != nil {
		return nil, 0, stderr.ErrFind8
	}
	for _, item := range results {
		if len(args.Query) == 0 {
			convert = append(convert, item)
		} else {
			temp := make(map[string]interface{})
			for _, v := range args.Query {
				temp[v] = item[v]
			}
			convert = append(convert, temp)
		}
	}
	r, err := json.Marshal(convert)
	if err != nil {
		return nil, 0, stderr.ErrFind9
	}
	*ret = json.RawMessage(r)
	return convert, count, nil
}

func (me *T) SaveJob(args struct {
	Collection string
	Data       bson.M
}) (bool, error) {
	collection := me.C_local.Database("job").Collection(args.Collection)
	_, err := collection.InsertOne(me.Ctx, args.Data)
	if err != nil {
		return false, stderr.ErrInsert
	}
	return true, nil
}

func (me *T) QueryOneJob(args struct {
	Collection string
	Filter     bson.M
}) (map[string]interface{}, error) {
	collection := me.C_local.Database("job").Collection(args.Collection)
	var result map[string]interface{}
	err := collection.FindOne(me.Ctx, args.Filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (me *T) QueryLastJob(args struct {
	Collection string
}) (map[string]interface{}, error) {
	collection := me.C_local.Database("job").Collection(args.Collection)
	var result map[string]interface{}
	opts := options.FindOne().SetSort(bson.M{"_id": -1})
	err := collection.FindOne(me.Ctx, bson.M{}, opts).Decode(&result)
	if err != nil {
		return nil, stderr.ErrFind10
	}
	return result, nil
}

func (me *T) QueryLastJobs(args struct {
	Collection string
	Index      string
	Sort       bson.M
	Filter     bson.M
	Query      []string
	Limit      int64
	Skip       int64
}) ([]map[string]interface{}, error) {
	collection := me.C_local.Database("job").Collection(args.Collection)
	var results []map[string]interface{}
	//
	op := options.Find()
	op.SetSort(args.Sort)
	op.SetLimit(args.Limit)
	op.SetSkip(args.Skip)
	cursor, err := collection.Find(me.Ctx, args.Filter, op)
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			log2.Fatalf("Closing cursor error %v", err)
		}
	}(cursor, me.Ctx)
	if err == mongo.ErrNoDocuments {
		return nil, stderr.ErrNotFound
	}
	if err != nil {
		return nil, stderr.ErrFind11
	}
	if err = cursor.All(me.Ctx, &results); err != nil {
		return nil, stderr.ErrFind12
	}
	return results, nil
}

func (me *T) QueryAggregate(args struct {
	Collection string
	Index      string
	Sort       bson.M
	Filter     bson.M
	Pipeline   []bson.M
	Query      []string
}, ret *json.RawMessage) ([]map[string]interface{}, error) {

	//for _, v := range args.Pipeline {
	//	limit := v["$limit"]
	//	if limit != nil {
	//		if limit.(int64) == 0 {
	//			v["$limit"] = limit2.DefaultLimit
	//		}
	//		if limit.(int64) > limit2.MaxLimit {
	//			v["$limit"] = limit2.MaxLimit
	//		}
	//	}
	//}

	var results []map[string]interface{}
	convert := make([]map[string]interface{}, 0)
	collection := me.C_online.Database(me.Db_online).Collection(args.Collection)
	op := options.AggregateOptions{}

	cursor, err := collection.Aggregate(me.Ctx, args.Pipeline, &op)

	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			log2.Fatalf("Closing cursor error %v", err)
		}
	}(cursor, me.Ctx)
	if err == mongo.ErrNoDocuments {
		return nil, stderr.ErrNotFound
	}
	if err != nil {
		return nil, stderr.ErrFind13
	}
	if err = cursor.All(me.Ctx, &results); err != nil {
		return nil, stderr.ErrFind14
	}

	for _, item := range results {
		if len(args.Query) == 0 {
			convert = append(convert, item)
		} else {
			temp := make(map[string]interface{})
			for _, v := range args.Query {
				temp[v] = item[v]
			}
			convert = append(convert, temp)
		}
	}

	r, err := json.Marshal(convert)
	if err != nil {
		return nil, stderr.ErrFind15
	}
	*ret = json.RawMessage(r)
	return convert, nil
}

func (me *T) QueryDocument(args struct {
	Collection string
	Index      string
	Sort       bson.M
	Filter     bson.M
}, ret *json.RawMessage) (map[string]interface{}, error) {
	co := options.CountOptions{}
	collection := me.C_online.Database(me.Db_online).Collection(args.Collection)
	count, err := collection.CountDocuments(me.Ctx, args.Filter, &co)
	if err == mongo.ErrNoDocuments {
		return nil, stderr.ErrNotFound
	}
	convert := make(map[string]interface{})
	convert["total counts"] = count
	r, err := json.Marshal(convert)
	if err != nil {
		return nil, stderr.ErrFind16
	}
	*ret = json.RawMessage(r)
	return convert, nil
}


// 去重查询统计
func (me *T) GetDistinctCount(args struct {
	Collection string
	Index      string
	Sort       bson.M
	Filter     bson.M
	Pipeline   []bson.M
	Query      []string
}, ret *json.RawMessage) (map[string]interface{}, error) {
	var results []map[string]interface{}
	convert := make(map[string]interface{})
	collection := me.C_online.Database(me.Db_online).Collection(args.Collection)
	op := options.AggregateOptions{}
	pipeline := bson.M{
		"$group": bson.M{"_id": "$hash"},
	}
	args.Pipeline = append(args.Pipeline, pipeline)
	args.Pipeline = append(args.Pipeline, bson.M{"$count": "count"})
	cursor, err := collection.Aggregate(me.Ctx, args.Pipeline, &op)

	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			log2.Fatalf("Closing cursor error %v", err)
		}
	}(cursor, me.Ctx)
	if err == mongo.ErrNoDocuments {
		return nil, stderr.ErrNotFound
	}
	if err != nil {
		return nil, stderr.ErrFind17
	}

	if err = cursor.All(me.Ctx, &results); err != nil {
		return nil, stderr.ErrFind18
	}

	convert["total"] = results[0]["count"]

	r, err := json.Marshal(convert)
	if err != nil {
		return nil, stderr.ErrFind19
	}
	*ret = json.RawMessage(r)

	return convert, nil

}


func (me *T) InsertDocument(args struct {
	Collection string
	Index      string
	Insert     *Insert
}, ret *json.RawMessage) (map[string]interface{}, error) {
	collection := me.C_online.Database(me.Db_online).Collection(args.Collection)
	_,err := collection.InsertOne(me.Ctx,&args.Insert)
	if err != nil {
		return nil, stderr.ErrInsertDocument
	}
	result := make(map[string]interface{})
	result["msg"] = "Insert document done!"
	r, err := json.Marshal(result)
	if err != nil {
		return nil, stderr.ErrInsertDocument
	}
	*ret = json.RawMessage(r)
	return result, nil
}

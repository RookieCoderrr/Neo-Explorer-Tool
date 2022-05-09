package api

import (
	"encoding/json"
	"neo3fura_http/lib/type/uintval"
	"neo3fura_http/var/stderr"

	"go.mongodb.org/mongo-driver/bson"
)

func (me *T) GetBlockByBlockHeight(args struct {
	BlockHeight uintval.T
	Filter      map[string]interface{}
	Raw         *map[string]interface{}
}, ret *json.RawMessage) error {
	if args.BlockHeight.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	r1, err := me.Client.QueryOne(struct {
		Collection string
		Index      string
		Sort       bson.M
		Filter     bson.M
		Query      []string
	}{
		Collection: "Block",
		Index:      "GetBlockByBlockHeight",
		Sort:       bson.M{},
		Filter:     bson.M{"index": args.BlockHeight},
		Query:      []string{},
	}, ret)
	if err != nil {
		return err
	}
	if args.Raw != nil {
		*args.Raw = r1
	}
	r1, err = me.Filter(r1, args.Filter)
	if err != nil {
		return err
	}
	r, err := json.Marshal(r1)
	if err != nil {
		return err
	}
	*ret = json.RawMessage(r)
	return nil
}

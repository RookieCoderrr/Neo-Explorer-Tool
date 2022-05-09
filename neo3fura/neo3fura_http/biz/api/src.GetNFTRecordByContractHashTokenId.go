package api

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"neo3fura_http/lib/type/h160"
	"neo3fura_http/lib/type/strval"
	"neo3fura_http/var/stderr"
	"strconv"
)

func (me *T) GetNFTRecordByContractHashTokenId(args struct {
	ContractHash h160.T
	MarketHash   h160.T
	TokenId      strval.T
	Filter       map[string]interface{}
}, ret *json.RawMessage) error {

	if args.ContractHash.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if len(args.TokenId) <= 0 {
		return stderr.ErrInvalidArgs
	}
	f := bson.M{"eventname": "Claim", "asset": args.ContractHash.Val(), "tokenid": args.TokenId.Val()}
	if len(args.MarketHash) > 0 {
		f["market"] = args.MarketHash.Val()
	}
	result := make([]map[string]interface{}, 0)

	r1, _, err := me.Client.QueryAll(struct {
		Collection string
		Index      string
		Sort       bson.M
		Filter     bson.M
		Query      []string
		Limit      int64
		Skip       int64
	}{
		Collection: "MarketNotification",
		Index:      "GetNFTRecordByContractHashTokenId",
		Sort:       bson.M{"timestamp": -1},
		Filter:     f,
		Query:      []string{},
	}, ret)
	if err != nil {
		return err
	}

	// 获取NFT的Transfer
	var raw2 []map[string]interface{}
	err3 := me.GetNep11TransferByContractHashTokenId(struct {
		ContractHash h160.T
		Limit        int64
		Skip         int64
		TokenId      strval.T
		Filter       map[string]interface{}
		Raw          *[]map[string]interface{}
	}{ContractHash: args.ContractHash, TokenId: args.TokenId, Raw: &raw2}, ret)
	if err3 != nil {
		return err3
	}

	for _, item := range raw2 {
		rr := make(map[string]interface{})
		rr["asset"] = item["contract"]
		rr["tokenid"] = item["tokenId"]
		rr["from"] = item["from"]
		rr["to"] = item["to"]
		rr["auctionAsset"] = "" //普通账户之间转账  无价格
		rr["auctionAmount"] = ""
		rr["timestamp"] = item["timestamp"]

		//筛选出从市场交易的nft 会有交易价格
		for _, i := range r1 {
			if item["txid"] == i["txid"] { //为了获取nft的交易价格
				extendData := i["extendData"].(string)
				var dat map[string]interface{}
				if err2 := json.Unmarshal([]byte(extendData), &dat); err2 == nil {
					bidAmount, err1 := strconv.ParseInt(dat["bidAmount"].(string), 10, 64)
					if err1 != nil {
						return err1
					}
					auctionAsset := dat["auctionAsset"]
					rr["auctionAsset"] = auctionAsset
					rr["auctionAmount"] = bidAmount

				} else {
					return err2
				}
			}
		}

		asset := item["contract"].(string)
		tokenid := item["tokenId"].(string)

		var raw3 map[string]interface{}
		err2 := getNFTProperties(strval.T(tokenid), h160.T(asset), me, ret, args.Filter, &raw3)
		if err2 != nil {
			rr["image"] = ""
			rr["name"] = ""
			rr["number"] = int64(-1)
			rr["properties"] = ""

		}

		rr["image"] = raw3["image"]
		rr["name"] = raw3["name"]
		rr["number"] = raw3["number"]
		rr["properties"] = raw3["properties"]

		result = append(result, rr)

	}

	num, err := strconv.ParseInt(strconv.Itoa(len(result)), 10, 64)
	if err != nil {
		return err
	}
	r2, err := me.FilterArrayAndAppendCount(result, num, args.Filter)
	if err != nil {
		return err
	}
	r, err := json.Marshal(r2)
	if err != nil {
		return err
	}

	*ret = json.RawMessage(r)
	return nil
}

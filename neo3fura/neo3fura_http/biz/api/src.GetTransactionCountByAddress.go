package api

import (
	"encoding/json"
	"neo3fura_http/lib/type/h160"
	"neo3fura_http/var/stderr"

	"go.mongodb.org/mongo-driver/bson"
)

func (me *T) GetTransactionCountByAddress(args struct {
	Address h160.T
	Filter  map[string]interface{}
}, ret *json.RawMessage) error {
	if args.Address.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	r1, err := me.Client.QueryDocument(struct {
		Collection string
		Index      string
		Sort       bson.M
		Filter     bson.M
	}{
		Collection: "Transaction",
		Index:      "GetTransactionCountByAddress",
		Sort:       bson.M{},
		Filter:     bson.M{"sender": args.Address.TransferAddress()},
	}, ret)
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

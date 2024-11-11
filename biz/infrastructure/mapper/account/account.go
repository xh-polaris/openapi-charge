package account

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Account struct {
	ID         primitive.ObjectID `bson:"_id" json:"id"`
	TxId       string             `bson:"tx_id" json:"txId"`
	Increment  int64              `bson:"increment" json:"increment"`
	MarginId   string             `bson:"margin_id" json:"marginId"`
	CreateTime time.Time          `bson:"create_time" json:"createTime"`
}

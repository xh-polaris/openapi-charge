package log

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Log struct {
	ID              primitive.ObjectID `bson:"_id" json:"id"`
	FullInterfaceId string             `bson:"full_interface_id" json:"fullInterfaceId"`
	UserId          string             `bson:"user_id" json:"userId"`
	KeyId           string             `bson:"key_id" json:"keyId"`
	Status          int64              `bson:"status" json:"status"`
	Info            string             `bson:"info" json:"info"`
	Count           int64              `bson:"count" json:"count"`
	Value           int64              `bson:"value" json:"value"`
	Timestamp       int64              `bson:"timestamp" json:"timestamp"`
	CreateTime      int64              `bson:"create_time" json:"createTime"`
}

package full

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Interface struct {
	ID              primitive.ObjectID `bson:"_id" json:"id"`
	BaseInterfaceId string             `bson:"base_interface_id" json:"baseInterfaceId"`
	UserId          string             `bson:"user_id" json:"userId"`
	ChargeType      int64              `bson:"charge_type" json:"chargeType"`
	Price           int64              `bson:"price" json:"price"`
	Margin          int64              `bson:"margin" json:"margin"`
	Status          int64              `bson:"status" json:"status"`
	CreateTime      int64              `bson:"create_time" json:"createTime"`
	UpdateTime      int64              `bson:"update_time" json:"updateTime"`
	DeleteTime      int64              `bson:"delete_time" json:"deleteTime"`
}

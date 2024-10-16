package full

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Interface struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	BaseInterfaceId string             `bson:"base_interface_id" json:"baseInterfaceId"`
	UserId          string             `bson:"user_id" json:"userId"`
	ChargeType      int64              `bson:"charge_type" json:"chargeType"`
	Price           int64              `bson:"price" json:"price"`
	Margin          int64              `bson:"margin" json:"margin"`
	Status          int64              `bson:"status" json:"status"`
	CreateTime      time.Time          `bson:"create_time" json:"createTime"`
	UpdateTime      time.Time          `bson:"update_time" json:"updateTime"`
	DeleteTime      time.Time          `bson:"delete_time,omitempty" json:"deleteTime,omitempty"`
}

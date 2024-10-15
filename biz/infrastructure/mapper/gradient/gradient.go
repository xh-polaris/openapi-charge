package gradient

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Discount struct {
	Num  int64 `bson:"num" json:"num"`
	Rate int64 `bson:"rate" json:"rate"`
	Low  int64 `bson:"low" json:"low"`
}

type Gradient struct {
	ID              primitive.ObjectID `bson:"_id" json:"id"`
	BaseInterfaceId string             `bson:"base_interface_id" json:"baseInterfaceId"`
	Discounts       []Discount         `bson:"discounts" json:"discounts"`
	Status          int64              `bson:"status" json:"status"`
	CreateTime      int64              `bson:"create_time" json:"createTime"`
	UpdateTime      int64              `bson:"update_time" json:"updateTime"`
	DeleteTime      int64              `bson:"delete_time" json:"deleteTime"`
}

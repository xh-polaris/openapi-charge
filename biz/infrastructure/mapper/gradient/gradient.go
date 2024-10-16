package gradient

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Discount struct {
	Num  int64 `bson:"num" json:"num"`
	Rate int64 `bson:"rate" json:"rate"`
	Low  int64 `bson:"low" json:"low"`
}

type Gradient struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	BaseInterfaceId string             `bson:"base_interface_id" json:"baseInterfaceId"`
	Discounts       []Discount         `bson:"discounts" json:"discounts"`
	Status          int64              `bson:"status" json:"status"`
	CreateTime      time.Time          `bson:"create_time" json:"createTime"`
	UpdateTime      time.Time          `bson:"update_time" json:"updateTime"`
	DeleteTime      time.Time          `bson:"delete_time,omitempty" json:"deleteTime,omitempty"`
}

package margin

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Margin struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FullInterfaceId string             `bson:"full_interface_id" json:"fullInterfaceId"`
	UserId          string             `bson:"user_id" json:"userId"`
	Margin          int64              `bson:"margin" json:"margin"`
	Status          int64              `bson:"status" json:"status"`
	CreateTime      time.Time          `bson:"create_time,omitempty" json:"createTime,omitempty"`
	UpdateTime      time.Time          `bson:"update_time,omitempty" json:"updateTime,omitempty"`
	DeleteTime      time.Time          `bson:"delete_time,omitempty" json:"deleteTime,omitempty"`
}

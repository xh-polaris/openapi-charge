package base

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Interface struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name       string             `bson:"name" json:"name"`
	Host       string             `bson:"host" json:"host"`
	Path       string             `bson:"path" json:"path"`
	Method     int64              `bson:"method" json:"method"`
	PassWay    int64              `bson:"passWay" json:"passWay"`
	Params     map[string]string  `bson:"params" json:"params"`
	Content    string             `bson:"content" json:"content"`
	Status     int64              `bson:"status" json:"status"`
	CreateTime time.Time          `bson:"create_time" json:"createTime"`
	UpdateTime time.Time          `bson:"update_time" json:"updateTime"`
	DeleteTime time.Time          `bson:"delete_time,omitempty" json:"deleteTime,omitempty"`
}

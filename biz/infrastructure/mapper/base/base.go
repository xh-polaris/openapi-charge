package base

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Interface struct {
	ID         primitive.ObjectID `bson:"_id" json:"id"`
	Name       string             `bson:"name" json:"name"`
	Host       string             `bson:"host" json:"host"`
	Path       string             `bson:"path" json:"path"`
	Method     int64              `bson:"method" json:"method"`
	PassWay    int64              `bson:"passWay" json:"passWay"`
	Params     map[string]string  `bson:"params" json:"params"`
	Content    string             `bson:"content" json:"content"`
	Status     int64              `bson:"status" json:"status"`
	CreateTime int64              `bson:"create_time" json:"createTime"`
	UpdateTime int64              `bson:"update_time" json:"updateTime"`
	DeleteTime int64              `bson:"delete_time" json:"deleteTime"`
}

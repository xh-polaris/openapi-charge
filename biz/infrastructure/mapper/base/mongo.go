package base

import (
	"context"
	"github.com/xh-polaris/openapi-charge/biz/infrastructure/config"
)

const (
	prefixKeyCacheKey = "cache:base"
	CollectionName    = "base"
)

type IMongoMapper interface {
	Insert(ctx context.Context, i *Interface) error
	Update(ctx context.Context, i *Interface) error
	FindAndCount(ctx context.Context, p *baisc.PaginationOptions) ([]*Interface, int64, error)
	Delete(ctx context.Context, id string) error
}

type MongoMapper struct {
	conn *monc.Model
}

func NewMongoMapper(config *config.Config) *MongoMapper {
	conn := monc.MustNewModel(config.Mongo.URL, config.Mongo.DB, CollectionName, config.Cache)
	return &MongoMapper{conn: conn}
}

func (m *MongoMapper) Insert(ctx context.Context, i *Interface) error {

}

func (m *MongoMapper) Update(ctx context.Context, i *Interface) error {

}

func (m *MongoMapper) Delete(ctx context.Context, id string) error {

}

func (m *MongoMapper) FindAndCount(ctx context.Context, p *basic.PainationOptions) ([]*Interface, int64, error) {

}

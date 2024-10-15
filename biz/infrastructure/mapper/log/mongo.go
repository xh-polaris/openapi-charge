package log

import (
	"context"
	"github.com/xh-polaris/openapi-charge/biz/infrastructure/config"
)

const (
	prefixKeyCacheKey = "cache:log"
	CollectionName    = "log"
)

type IMongoMapper interface {
	Insert(ctx context.Context, l *Log) error
	FindAndCount(ctx context.Context, p *baisc.PaginationOptions) ([]*Log, int64, error)
}

type MongoMapper struct {
	conn *monc.Model
}

func NewMongoMapper(config *config.Config) *MongoMapper {
	conn := monc.MustNewModel(config.Mongo.URL, config.Mongo.DB, CollectionName, config.Cache)
	return &MongoMapper{conn: conn}
}

func (m *MongoMapper) Insert(ctx context.Context, l *Log) error {
	//TODO implement me
	panic("implement me")
}

func (m *MongoMapper) FindAndCount(ctx context.Context, p *baisc.PaginationOptions) ([]*Log, int64, error) {
	//TODO implement me
	panic("implement me")
}

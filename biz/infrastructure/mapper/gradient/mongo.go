package gradient

import (
	"context"
	"github.com/xh-polaris/openapi-charge/biz/infrastructure/config"
)

const (
	prefixKeyCacheKey = "cache:gradient"
	CollectionName    = "gradient"
)

type IMongoMapper interface {
	Insert(ctx context.Context, g *Gradient) error
	Update(ctx context.Context, g *Gradient) error
	FindOne(ctx context.Context, basicInterfaceId string) ([]*Gradient, error)
	Delete(ctx context.Context, id string) error
}

type MongoMapper struct {
	conn *monc.Model
}

func NewMongoMapper(config *config.Config) *MongoMapper {
	conn := monc.MustNewModel(config.Mongo.URL, config.Mongo.DB, CollectionName, config.Cache)
	return &MongoMapper{conn: conn}
}

func (m *MongoMapper) Insert(ctx context.Context, g *Gradient) error {
	//TODO implement me
	panic("implement me")
}

func (m *MongoMapper) Update(ctx context.Context, g *Gradient) error {
	//TODO implement me
	panic("implement me")
}

func (m *MongoMapper) FindOne(ctx context.Context, basicInterfaceId string) ([]*Gradient, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MongoMapper) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

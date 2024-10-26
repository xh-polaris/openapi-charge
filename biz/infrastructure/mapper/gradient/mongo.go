package gradient

import (
	"context"
	"errors"
	"github.com/xh-polaris/openapi-charge/biz/infrastructure/config"
	"github.com/xh-polaris/openapi-charge/biz/infrastructure/consts"
	"github.com/zeromicro/go-zero/core/stores/monc"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	prefixKeyCacheKey = "cache:gradient"
	CollectionName    = "gradient"
)

type IMongoMapper interface {
	Insert(ctx context.Context, g *Gradient) error
	Update(ctx context.Context, g *Gradient) error
	FindOne(ctx context.Context, baseInterfaceId string) (*Gradient, error)
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
	if g.ID.IsZero() {
		g.ID = primitive.NewObjectID()
		g.CreateTime = time.Now()
		g.UpdateTime = g.CreateTime
	}
	_, err := m.conn.InsertOneNoCache(ctx, g)
	return err
}

func (m *MongoMapper) Update(ctx context.Context, g *Gradient) error {
	g.UpdateTime = time.Now()
	_, err := m.conn.UpdateByIDNoCache(ctx, g.ID, bson.M{consts.Set: g})
	return err
}

func (m *MongoMapper) FindOne(ctx context.Context, baseInterfaceId string) (*Gradient, error) {
	var g Gradient
	err := m.conn.FindOneNoCache(ctx, &g,
		bson.M{
			consts.FullInterfaceId: baseInterfaceId,
			consts.Status:          bson.M{consts.NotEqual: consts.DeleteStatus},
		})
	switch {
	case err == nil:
		return &g, nil
	case errors.Is(err, monc.ErrNotFound):
		return nil, consts.ErrNotFound
	default:
		return nil, err
	}
}

func (m *MongoMapper) Delete(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return consts.ErrInValidId
	}
	var g Gradient
	err = m.conn.FindOneNoCache(ctx, &g, bson.M{consts.ID: oid})
	if err != nil {
		return consts.ErrNotFound
	}
	now := time.Now()
	g.DeleteTime = now
	g.UpdateTime = now
	g.Status = consts.DeleteStatus
	_, err = m.conn.UpdateByIDNoCache(ctx, oid, bson.M{consts.Set: g})
	return err
}

package margin

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
	prefixKeyCacheKey = "cache:margin"
	CollectionName    = "margin"
)

type IMongoMapper interface {
	Insert(ctx context.Context, m *Margin) (string, error)
	FindOneByBaseInfIdAndUserId(ctx context.Context, fullInfId string, userId string) (*Margin, error)
	Update(ctx context.Context, m *Margin) error
	FindOne(ctx context.Context, id string) (*Margin, error)
	Delete(ctx context.Context, id string) error
	UpdateMargin(ctx context.Context, id string, increment int64) error
}

type MongoMapper struct {
	conn *monc.Model
}

func NewMongoMapper(config *config.Config) *MongoMapper {
	conn := monc.MustNewModel(config.Mongo.URL, config.Mongo.DB, CollectionName, config.Cache)
	return &MongoMapper{conn: conn}
}

func (mon *MongoMapper) Insert(ctx context.Context, m *Margin) (string, error) {
	if m.ID.IsZero() {
		m.ID = primitive.NewObjectID()
		m.CreateTime = time.Now()
		m.UpdateTime = m.CreateTime
	}
	key := prefixKeyCacheKey + m.ID.Hex()
	_, err := mon.conn.InsertOne(ctx, key, m)
	return m.ID.Hex(), err
}

func (mon *MongoMapper) Update(ctx context.Context, m *Margin) error {
	m.UpdateTime = time.Now()
	key := prefixKeyCacheKey + m.ID.Hex()
	_, err := mon.conn.UpdateByID(ctx, key, m.ID, bson.M{consts.Set: m})
	return err
}

func (m *MongoMapper) UpdateMargin(ctx context.Context, id string, increment int64) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return consts.ErrInValidId
	}
	key := prefixKeyCacheKey + id
	_, err = m.conn.UpdateByID(ctx, key, oid, bson.M{"$inc": bson.M{"margin": increment}})
	return err
}

func (mon *MongoMapper) FindOne(ctx context.Context, id string) (*Margin, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, consts.ErrInValidId
	}
	var m Margin
	key := prefixKeyCacheKey + id
	err = mon.conn.FindOne(ctx, key, &m, bson.M{
		consts.ID:     oid,
		consts.Status: bson.M{consts.NotEqual: consts.DeleteStatus},
	})
	switch {
	case err == nil:
		return &m, nil
	case errors.Is(err, monc.ErrNotFound):
		return nil, consts.ErrNotFound
	default:
		return nil, err
	}
}

func (mon *MongoMapper) FindOneByBaseInfIdAndUserId(ctx context.Context, fullInfId string, userId string) (*Margin, error) {
	var m Margin
	key := prefixKeyCacheKey + fullInfId + userId
	err := mon.conn.FindOne(ctx, key, &m, bson.M{
		consts.FullInterfaceId: fullInfId,
		consts.UserID:          userId,
		consts.Status:          bson.M{consts.NotEqual: consts.DeleteStatus},
	})
	switch {
	case err == nil:
		return &m, nil
	case errors.Is(err, monc.ErrNotFound):
		return nil, consts.ErrNotFound
	default:
		return nil, err
	}
}

func (mon *MongoMapper) Delete(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return consts.ErrInValidId
	}
	var m Margin
	key := prefixKeyCacheKey + id
	err = mon.conn.FindOne(ctx, key, &m, bson.M{consts.ID: oid})
	if err != nil {
		return consts.ErrNotFound
	}
	now := time.Now()
	m.DeleteTime = now
	m.UpdateTime = now
	m.Status = consts.DeleteStatus
	_, err = mon.conn.UpdateByID(ctx, key, oid, bson.M{consts.Set: m})
	return err
}

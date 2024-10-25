package full

import (
	"context"
	"errors"
	"github.com/xh-polaris/openapi-charge/biz/infrastructure/config"
	"github.com/xh-polaris/openapi-charge/biz/infrastructure/consts"
	util "github.com/xh-polaris/openapi-charge/biz/infrastructure/util/page"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/basic"
	"github.com/zeromicro/go-zero/core/stores/monc"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const (
	prefixKeyCacheKey = "cache:full"
	CollectionName    = "full"
)

type IMongoMapper interface {
	Insert(ctx context.Context, i *Interface) error
	Update(ctx context.Context, i *Interface) error
	UpdateMargin(ctx context.Context, id string, increment int64) error
	FindAndCountByUserId(ctx context.Context, userId string, p *basic.PaginationOptions) ([]*Interface, int64, error)
	FindOne(ctx context.Context, id string) (*Interface, error)
	FindOneByBaseInfIdAndUserId(ctx context.Context, baseInfId string, userId string) (*Interface, error)
	Delete(ctx context.Context, id string) error
}

type MongoMapper struct {
	conn *monc.Model
}

func NewMongoMapper(config *config.Config) *MongoMapper {
	conn := monc.MustNewModel(config.Mongo.URL, config.Mongo.DB, CollectionName, config.Cache)
	return &MongoMapper{conn: conn}
}

func (m *MongoMapper) Insert(ctx context.Context, i *Interface) (string, error) {
	if i.ID.IsZero() {
		i.ID = primitive.NewObjectID()
		i.CreateTime = time.Now()
		i.UpdateTime = i.CreateTime
	}
	key := prefixKeyCacheKey + i.ID.Hex()
	_, err := m.conn.InsertOne(ctx, key, i)
	return i.ID.Hex(), err
}

func (m *MongoMapper) Update(ctx context.Context, i *Interface) error {
	i.UpdateTime = time.Now()
	key := prefixKeyCacheKey + i.ID.Hex()
	_, err := m.conn.UpdateByID(ctx, key, i.ID, bson.M{consts.Set: i})
	return err
}

func (m *MongoMapper) UpdateMargin(ctx context.Context, id string, increment int64) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return consts.ErrInValidId
	}
	key := prefixKeyCacheKey + id
	_, err = m.conn.UpdateByID(ctx, key, oid, bson.M{"$inc": bson.M{"increment": increment}})
	return err
}

func (m *MongoMapper) FindAndCountByUserId(ctx context.Context, userId string, p *basic.PaginationOptions) ([]*Interface, int64, error) {
	skip, limit := util.ParsePageOpt(p)
	infs := make([]*Interface, 0)
	err := m.conn.Find(ctx, &infs,
		bson.M{
			consts.UserID: userId,
			consts.Status: bson.M{consts.NotEqual: consts.DeleteStatus}},
		&options.FindOptions{
			Skip:  &skip,
			Limit: &limit,
			Sort:  bson.M{consts.CreateTime: -1},
		})
	if err != nil {
		return nil, 0, consts.ErrNotFound
	}
	total, err := m.conn.CountDocuments(ctx, bson.M{
		consts.UserID: userId,
		consts.Status: bson.M{consts.NotEqual: consts.DeleteStatus},
	})
	if err != nil {
		return nil, 0, consts.ErrNotFound
	}
	return infs, total, nil
}

func (m *MongoMapper) Delete(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return consts.ErrInValidId
	}
	var i Interface
	key := prefixKeyCacheKey + id
	err = m.conn.FindOne(ctx, key, &i, bson.M{consts.ID: oid})
	if err != nil {
		return consts.ErrNotFound
	}
	now := time.Now()
	i.DeleteTime = now
	i.UpdateTime = now
	i.Status = consts.DeleteStatus
	_, err = m.conn.UpdateByID(ctx, key, oid, bson.M{consts.Set: i})
	return err
}

func (m *MongoMapper) FindOne(ctx context.Context, id string) (*Interface, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, consts.ErrInValidId
	}
	var inf Interface
	key := prefixKeyCacheKey + id
	err = m.conn.FindOne(ctx, key, &inf, bson.M{
		consts.ID:     oid,
		consts.Status: bson.M{consts.NotEqual: consts.DeleteStatus},
	})
	switch {
	case err == nil:
		return &inf, nil
	case errors.Is(err, monc.ErrNotFound):
		return nil, consts.ErrNotFound
	default:
		return nil, err
	}
}

func (m *MongoMapper) FindOneByBaseInfIdAndUserId(ctx context.Context, baseInfId string, userId string) (*Interface, error) {
	var inf Interface
	key := prefixKeyCacheKey + baseInfId + userId
	err := m.conn.FindOne(ctx, key, &inf, bson.M{
		consts.BasicInterfaceId: baseInfId,
		consts.UserID:           userId,
		consts.Status:           consts.EffectStatus,
	})
	switch {
	case err == nil:
		return &inf, nil
	case errors.Is(err, monc.ErrNotFound):
		return nil, consts.ErrNotFound
	default:
		return nil, err
	}
}

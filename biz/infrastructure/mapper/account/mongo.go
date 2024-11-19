package account

import (
	"context"
	"github.com/xh-polaris/openapi-charge/biz/infrastructure/config"
	"github.com/xh-polaris/openapi-charge/biz/infrastructure/consts"
	"github.com/zeromicro/go-zero/core/stores/monc"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	prefixKeyCacheKey = "cache:account"
	CollectionName    = "account"
)

type IMongoMapper interface {
	Insert(ctx context.Context, a *Account) error
	FindOneByTxId(ctx context.Context, txId string) (*Account, error)
	//FindAndCountByUserId(ctx context.Context, userId string, p *basic.PaginationOptions) ([]*Account, int64, error)
}

type MongoMapper struct {
	conn *monc.Model
}

func NewMongoMapper(config *config.Config) *MongoMapper {
	conn := monc.MustNewModel(config.Mongo.URL, config.Mongo.DB, CollectionName, config.Cache)
	return &MongoMapper{conn: conn}
}

func (m *MongoMapper) Insert(ctx context.Context, a *Account) error {
	if a.ID.IsZero() {
		a.ID = primitive.NewObjectID()
		a.CreateTime = time.Now()
	}
	_, err := m.conn.InsertOneNoCache(ctx, a)
	return err
}

func (m *MongoMapper) FindOneByTxId(ctx context.Context, txId string) (*Account, error) {
	var a Account
	err := m.conn.FindOneNoCache(ctx, &a,
		bson.M{
			consts.TxId: txId,
		})
	if err != nil {
		return nil, err
	}
	return &a, nil
}

package log

import (
	"context"
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
	prefixKeyCacheKey = "cache:log"
	CollectionName    = "log"
)

type IMongoMapper interface {
	Insert(ctx context.Context, l *Log) error
	FindAndCountByInfId(ctx context.Context, infId string, p *basic.PaginationOptions) ([]*Log, int64, error)
	FindAndCountByUserId(ctx context.Context, userId string, p *basic.PaginationOptions) ([]*Log, int64, error)
}

type MongoMapper struct {
	conn *monc.Model
}

func NewMongoMapper(config *config.Config) *MongoMapper {
	conn := monc.MustNewModel(config.Mongo.URL, config.Mongo.DB, CollectionName, config.Cache)
	return &MongoMapper{conn: conn}
}

func (m *MongoMapper) Insert(ctx context.Context, l *Log) error {
	if l.ID.IsZero() {
		l.ID = primitive.NewObjectID()
		l.CreateTime = time.Now()
	}
	key := prefixKeyCacheKey + l.ID.Hex()
	_, err := m.conn.InsertOne(ctx, key, l)
	return err
}

func (m *MongoMapper) FindAndCountByInfId(ctx context.Context, infId string, p *basic.PaginationOptions) ([]*Log, int64, error) {
	skip, limit := util.ParsePageOpt(p)
	logs := make([]*Log, 0, limit)
	err := m.conn.Find(ctx, &logs,
		bson.M{
			consts.FullInterfaceId: infId,
			consts.Status:          bson.M{consts.NotEqual: consts.DeleteStatus},
		}, &options.FindOptions{
			Skip:  &skip,
			Limit: &limit,
			Sort:  bson.M{consts.CreateTime: -1},
		})
	if err != nil {
		return nil, 0, err
	}

	total, err := m.conn.CountDocuments(ctx, bson.M{
		consts.FullInterfaceId: infId,
		consts.Status:          bson.M{consts.NotEqual: consts.DeleteStatus},
	})
	if err != nil {
		return nil, 0, err
	}
	return logs, total, nil
}

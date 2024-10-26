package base

import (
	"context"
	"errors"
	"fmt"
	"github.com/xh-polaris/openapi-charge/biz/infrastructure/config"
	"github.com/xh-polaris/openapi-charge/biz/infrastructure/consts"
	util "github.com/xh-polaris/openapi-charge/biz/infrastructure/util/page"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/basic"
	"github.com/zeromicro/go-zero/core/stores/monc"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/url"
	"time"
)

const (
	prefixKeyCacheKey = "cache:base"
	CollectionName    = "base"
)

type IMongoMapper interface {
	Insert(ctx context.Context, i *Interface) error
	Update(ctx context.Context, i *Interface) error
	FindAndCount(ctx context.Context, p *basic.PaginationOptions) ([]*Interface, int64, error)
	FindOne(ctx context.Context, id string) (i *Interface, err error)
	FindOneByURLAndMethod(ctx context.Context, rawURL string, method string) (i *Interface, err error)
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
	if i.ID.IsZero() {
		i.ID = primitive.NewObjectID()
		i.CreateTime = time.Now()
		i.UpdateTime = i.CreateTime
	}
	_, err := m.conn.InsertOneNoCache(ctx, i)
	return err
}

func (m *MongoMapper) Update(ctx context.Context, i *Interface) error {
	i.UpdateTime = time.Now()
	_, err := m.conn.UpdateByIDNoCache(ctx, i.ID, bson.M{"$set": i})
	return err
}

func (m *MongoMapper) Delete(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return consts.ErrInValidId
	}
	var i Interface
	err = m.conn.FindOneNoCache(ctx, &i, bson.M{consts.ID: oid})

	if err != nil {
		return consts.ErrNotFound
	}
	now := time.Now()
	i.DeleteTime = now
	i.UpdateTime = now
	i.Status = consts.DeleteStatus
	_, err = m.conn.UpdateByIDNoCache(ctx, oid, bson.M{"$set": i})
	return err
}

func (m *MongoMapper) FindAndCount(ctx context.Context, p *basic.PaginationOptions) ([]*Interface, int64, error) {
	skip, limit := util.ParsePageOpt(p)
	infs := make([]*Interface, 0)
	err := m.conn.Find(ctx, &infs,
		bson.M{
			consts.Status: bson.M{"$ne": consts.DeleteStatus}},
		&options.FindOptions{
			Skip:  &skip,
			Limit: &limit,
			Sort:  bson.M{consts.CreateTime: -1},
		})
	if err != nil {
		return nil, 0, err
	}

	total, err := m.conn.CountDocuments(ctx, bson.M{
		consts.Status: bson.M{"$ne": consts.DeleteStatus},
	})
	if err != nil {
		return nil, 0, err
	}
	return infs, total, nil
}

func (m *MongoMapper) FindOne(ctx context.Context, id string) (i *Interface, err error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, consts.ErrInValidId
	}
	var inf Interface
	err = m.conn.FindOneNoCache(ctx, &inf, bson.M{
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

func (m *MongoMapper) FindOneByURLAndMethod(ctx context.Context, rawURL string, method string) (i *Interface, err error) {
	host, path := parseURL(rawURL)

	var inf Interface
	err = m.conn.FindOneNoCache(ctx, &inf, bson.M{
		consts.Host:   host,
		consts.Path:   path,
		consts.Method: methodToEn(method),
		consts.Status: consts.EffectStatus,
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

func parseURL(rawURL string) (string, string) {
	// 解析 URL
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		fmt.Println("解析 URL 时发生错误:", err)
		return "", ""
	}

	// 获取 host 和 path
	host := parsedURL.Host
	path := parsedURL.Path

	return host, path
}

func methodToEn(method string) int64 {
	switch method {
	case "GET":
		return 0
	case "POST":
		return 1
	case "PUT":
		return 2
	default:
		return 0
	}
}

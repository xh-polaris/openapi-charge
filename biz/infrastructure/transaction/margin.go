package transaction

import (
	"context"
	"github.com/xh-polaris/openapi-charge/biz/infrastructure/config"
	"github.com/xh-polaris/openapi-charge/biz/infrastructure/consts"
	"github.com/xh-polaris/openapi-charge/biz/infrastructure/mapper/margin"
	"github.com/zeromicro/go-zero/core/stores/monc"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

const (
	MarginCollectionName = "margin"
)

type IMarginTransaction interface {
	UpdateMargin(ctx context.Context, id string, increment int64) error
}

type MarginTransaction struct {
	conn *monc.Model
}

func NewMarginTransaction(config *config.Config) *MarginTransaction {
	conn := monc.MustNewModel(config.Mongo.URL, config.Mongo.DB, MarginCollectionName, config.Cache)
	return &MarginTransaction{
		conn: conn,
	}
}

func (m *MarginTransaction) UpdateMargin(ctx context.Context, id string, increment int64) error {
	s, err := m.conn.StartSession()
	if err != nil {
		return err
	}
	defer s.EndSession(ctx)

	_, err = s.WithTransaction(ctx, func(sessionContext mongo.SessionContext) (interface{}, error) {
		// 查找余额
		oid, err2 := primitive.ObjectIDFromHex(id)
		if err2 != nil {
			return nil, consts.ErrInValidId
		}
		var aMargin margin.Margin
		err3 := m.conn.FindOneNoCache(ctx, &aMargin, bson.M{
			consts.ID:     oid,
			consts.Status: bson.M{consts.NotEqual: consts.DeleteStatus},
		})
		if err3 != nil {
			return nil, consts.ErrNotFound
		}

		// 判断是否足够
		if (increment > 0) || (increment+aMargin.Margin > 0) {
			// 余量足够
			_, err4 := m.conn.UpdateByIDNoCache(ctx, aMargin.ID, bson.M{
				"$inc": bson.M{
					"margin": increment,
				},
				"$set": bson.M{
					"update_time": time.Now(),
				},
			})
			if err4 != nil {
				return nil, consts.ErrUpdate
			}

			// TODO 新增流水
		}
		// 余量不足
		return aMargin, consts.ErrInsufficientMargin
	})
	return err
}

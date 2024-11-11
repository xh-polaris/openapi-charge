package service

import (
	"context"
	"github.com/google/wire"
	"github.com/xh-polaris/openapi-charge/biz/infrastructure/mapper/margin"
	"github.com/xh-polaris/openapi-charge/biz/infrastructure/transaction"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/openapi/charge"
	"strconv"
	"time"
)

type IMarginService interface {
	CreateMargin(ctx context.Context, req *charge.CreateMarginReq) (*charge.CreateMarginResp, error)
	UpdateMargin(ctx context.Context, req *charge.UpdateMarginReq) (*charge.UpdateMarginResp, error)
	GetMargin(ctx context.Context, req *charge.GetMarginReq) (*charge.GetMarginResp, error)
	DeleteMargin(ctx context.Context, req *charge.DeleteMarginReq) (*charge.DeleteMarginResp, error)
}

type MarginService struct {
	MarginMongoMapper *margin.MongoMapper
	MarginTransaction *transaction.MarginTransaction
}

var MarginServiceSet = wire.NewSet(
	wire.Struct(new(MarginService), "*"),
	wire.Bind(new(IMarginService), new(*MarginService)),
)

func (s *MarginService) CreateMargin(ctx context.Context, req *charge.CreateMarginReq) (*charge.CreateMarginResp, error) {
	oldMar, err := s.MarginMongoMapper.FindOneByBaseInfIdAndUserId(ctx, req.FullInterfaceId, req.UserId)
	if err == nil && oldMar != nil {
		return &charge.CreateMarginResp{
			Done:     false,
			Msg:      "已创建过接口余量",
			MarginId: oldMar.ID.Hex(),
		}, nil
	}

	now := time.Now()
	m := &margin.Margin{
		FullInterfaceId: req.FullInterfaceId,
		UserId:          req.UserId,
		Status:          0,
		CreateTime:      now,
		UpdateTime:      now,
	}
	marginId, err := s.MarginMongoMapper.Insert(ctx, m)
	if err != nil {
		return &charge.CreateMarginResp{
			Done:     false,
			Msg:      "创建接口余量失败",
			MarginId: marginId,
		}, err
	}
	return &charge.CreateMarginResp{
		Done:     true,
		Msg:      "创建接口余量成功",
		MarginId: marginId,
	}, nil

}
func (s *MarginService) GetMargin(ctx context.Context, req *charge.GetMarginReq) (*charge.GetMarginResp, error) {
	oldMar, err := s.MarginMongoMapper.FindOneByBaseInfIdAndUserId(ctx, req.FullInterfaceId, req.UserId)
	if err == nil && oldMar != nil {
		return &charge.GetMarginResp{
			Margin: &charge.Margin{
				Id:              oldMar.ID.Hex(),
				UserId:          oldMar.UserId,
				FullInterfaceId: oldMar.FullInterfaceId,
				Margin:          oldMar.Margin,
				CreateTime:      oldMar.CreateTime.Unix(),
				UpdateTime:      oldMar.UpdateTime.Unix(),
			},
		}, nil
	}
	return nil, err

}
func (s *MarginService) DeleteMargin(ctx context.Context, req *charge.DeleteMarginReq) (*charge.DeleteMarginResp, error) {
	err := s.MarginMongoMapper.Delete(ctx, req.Id)
	if err != nil {
		return &charge.DeleteMarginResp{
			Done: false,
			Msg:  "删除失败",
		}, err
	}
	return &charge.DeleteMarginResp{
		Done: true,
		Msg:  "删除成功",
	}, nil
}

func (s *MarginService) UpdateMargin(ctx context.Context, req *charge.UpdateMarginReq) (*charge.UpdateMarginResp, error) {
	var txId string
	if req.TxId != nil {
		txId = *req.TxId
	}

	err := s.MarginTransaction.UpdateMargin(ctx, req.Id, req.Increment, txId)
	if err != nil {
		return &charge.UpdateMarginResp{
			Done: false,
			Msg:  "更新余额失败",
		}, err
	}
	return &charge.UpdateMarginResp{
		Done: true,
		Msg:  "余额变化:" + formatIncrement(req.Increment),
	}, nil
}

func formatIncrement(increment int64) string {
	if increment >= 0 {
		return "+" + strconv.FormatInt(increment, 10)
	}
	return strconv.FormatInt(increment, 10)
}

package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/wire"
	"github.com/xh-polaris/openapi-charge/biz/infrastructure/consts"
	"github.com/xh-polaris/openapi-charge/biz/infrastructure/mapper/full"
	"github.com/xh-polaris/openapi-charge/biz/infrastructure/mapper/gradient"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/openapi/charge"
	"time"
)

type IGradientService interface {
	CreateGradient(ctx context.Context, req *charge.CreateGradientReq) (*charge.CreateGradientResp, error)
	UpdateGradient(ctx context.Context, req *charge.UpdateGradientReq) (*charge.UpdateGradientResp, error)
	GetGradient(ctx context.Context, req *charge.GetGradientReq) (*charge.GetGradientResp, error)
	GetAmount(ctx context.Context, req *charge.GetAmountReq) (*charge.GetAmountResp, error)
}

type GradientService struct {
	GradientMongoMapper      *gradient.MongoMapper
	FullInterfaceMongoMapper *full.MongoMapper
}

var GradientServiceSet = wire.NewSet(
	wire.Struct(new(GradientService), "*"),
	wire.Bind(new(IGradientService), new(*GradientService)),
)

func (s *GradientService) CreateGradient(ctx context.Context, req *charge.CreateGradientReq) (*charge.CreateGradientResp, error) {
	now := time.Now()
	g := &gradient.Gradient{
		BaseInterfaceId: req.BaseInterfaceId,
		Discounts:       ParseCDiscountsToDiscounts(req.Discounts),
		Status:          0,
		CreateTime:      now,
		UpdateTime:      now,
	}
	if s.GradientMongoMapper == nil {
		fmt.Printf("is NIl")
	}
	err := s.GradientMongoMapper.Insert(ctx, g)
	if err != nil {
		return &charge.CreateGradientResp{
			Done: false,
			Msg:  "创建梯度折扣失败",
		}, err
	}
	return &charge.CreateGradientResp{
		Done: true,
		Msg:  "创建梯度折扣成功",
	}, nil
}

func (s *GradientService) UpdateGradient(ctx context.Context, req *charge.UpdateGradientReq) (*charge.UpdateGradientResp, error) {
	g, err := s.GradientMongoMapper.FindOne(ctx, req.Id)
	if err != nil {
		return &charge.UpdateGradientResp{
			Done: false,
			Msg:  "梯度折扣不存在或已删除",
		}, err
	}
	g.Discounts = ParseCDiscountsToDiscounts(req.Discounts)
	g.Status = int64(req.Status)
	err = s.GradientMongoMapper.Update(ctx, g)
	if err != nil {
		return &charge.UpdateGradientResp{
			Done: false,
			Msg:  "更新梯度折扣失败",
		}, err
	}
	return &charge.UpdateGradientResp{
		Done: true,
		Msg:  "更新梯度折扣成功",
	}, nil
}

func (s *GradientService) GetGradient(ctx context.Context, req *charge.GetGradientReq) (*charge.GetGradientResp, error) {
	data, err := s.GradientMongoMapper.FindOneByBaseInfId(ctx, req.BaseInterfaceId)
	if err != nil {
		return nil, err
	}

	var g charge.Gradient
	g.Id = data.ID.Hex()
	g.Status = charge.InterfaceStatus(data.Status)
	g.CreateTime = data.CreateTime.Unix()
	g.UpdateTime = data.UpdateTime.Unix()
	g.BaseInterfaceId = data.BaseInterfaceId
	g.Discounts = ParseDiscountsToCDiscounts(data.Discounts)

	return &charge.GetGradientResp{
		Gradient: &g,
	}, nil
}

func (s *GradientService) GetAmount(ctx context.Context, req *charge.GetAmountReq) (*charge.GetAmountResp, error) {
	fullInfId := req.FullInfId
	increment := req.Increment

	fullInf, err := s.FullInterfaceMongoMapper.FindOne(ctx, fullInfId)
	if err != nil || fullInf == nil {
		return nil, err
	}

	aGradient, err := s.GradientMongoMapper.FindOneByBaseInfId(ctx, fullInf.BaseInterfaceId)
	isDiscount := true // 默认采用折扣
	// 未找到则不折扣
	if errors.Is(err, consts.ErrNotFound) {
		isDiscount = false
	} else if err != nil {
		return nil, err
	}

	var rate int64 = 100
	var originAmount = increment * fullInf.Price
	var amount = originAmount

	// 判断折扣是否正常状态
	if isDiscount && aGradient.Status == 0 {
		for _, discount := range aGradient.Discounts {
			if increment > discount.Low {
				rate = discount.Rate
			}
		}
		amount = originAmount * rate / 100
	}
	return &charge.GetAmountResp{
		Rate:         rate,
		OriginAmount: originAmount,
		Amount:       amount,
	}, nil
}

func ParseCDiscountsToDiscounts(cDiscounts []*charge.Discount) []gradient.Discount {
	discounts := make([]gradient.Discount, 0)
	for _, cDiscount := range cDiscounts {
		var dis = gradient.Discount{
			Num:  cDiscount.Num,
			Rate: cDiscount.Rate,
			Low:  cDiscount.Low,
		}
		discounts = append(discounts, dis)
	}
	return discounts
}

func ParseDiscountsToCDiscounts(Discounts []gradient.Discount) []*charge.Discount {
	cDiscounts := make([]*charge.Discount, 0)
	for _, discount := range Discounts {
		var cDis = &charge.Discount{
			Num:  discount.Num,
			Rate: discount.Rate,
			Low:  discount.Low,
		}
		cDiscounts = append(cDiscounts, cDis)
	}
	return cDiscounts
}

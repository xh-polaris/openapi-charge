package service

import (
	"context"
	"github.com/google/wire"
	"github.com/jinzhu/copier"
	"github.com/xh-polaris/openapi-charge/biz/infrastructure/mapper/full"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/openapi/charge"
	"strconv"
	"time"
)

type IFullInterfaceService interface {
	CreateFullInterface(ctx context.Context, req *charge.CreateFullInterfaceReq) (*charge.CreateFullInterfaceResp, error)
	UpdateFullInterface(ctx context.Context, req *charge.UpdateFullInterfaceReq) (*charge.UpdateFullInterfaceResp, error)
	UpdateMargin(ctx context.Context, req *charge.UpdateMarginReq) (*charge.UpdateMarginResp, error)
	DeleteFullInterface(ctx context.Context, req *charge.DeleteFullInterfaceReq) (*charge.DeleteFullInterfaceResp, error)
	GetFullInterface(ctx context.Context, req *charge.GetFullInterfaceReq) (*charge.GetFullInterfaceResp, error)
}

type FullInterfaceService struct {
	FullInterfaceMongoMapper *full.MongoMapper
}

var FullInterfaceServiceSet = wire.NewSet(
	wire.Struct(new(FullInterfaceService), "*"),
	wire.Bind(new(IFullInterfaceService), new(*FullInterfaceService)),
)

func (s *FullInterfaceService) CreateFullInterface(ctx context.Context, req *charge.CreateFullInterfaceReq) (*charge.CreateFullInterfaceResp, error) {
	now := time.Now()
	inf := &full.Interface{
		BaseInterfaceId: req.BaseInterfaceId,
		UserId:          req.UserId,
		ChargeType:      int64(req.ChargeType),
		Price:           req.Price,
		Margin:          req.Margin,
		Status:          0,
		CreateTime:      now,
		UpdateTime:      now,
	}
	err := s.FullInterfaceMongoMapper.Insert(ctx, inf)
	if err != nil {
		return &charge.CreateFullInterfaceResp{
			Done: false,
			Msg:  "创建完整接口失败",
		}, err
	}
	return &charge.CreateFullInterfaceResp{
		Done: true,
		Msg:  "创建完整接口成功",
	}, nil

}

func (s *FullInterfaceService) UpdateFullInterface(ctx context.Context, req *charge.UpdateFullInterfaceReq) (*charge.UpdateFullInterfaceResp, error) {
	inf, err := s.FullInterfaceMongoMapper.FindOne(ctx, req.Id)
	if err != nil {
		return &charge.UpdateFullInterfaceResp{
			Done: false,
			Msg:  "完整接口不存在或已删除",
		}, err
	}
	inf.ChargeType = int64(req.ChargeType)
	inf.Price = req.Price
	inf.Status = int64(req.Status)

	err = s.FullInterfaceMongoMapper.Update(ctx, inf)
	if err != nil {
		return &charge.UpdateFullInterfaceResp{
			Done: false,
			Msg:  "更新完整接口失败",
		}, err
	}
	return &charge.UpdateFullInterfaceResp{
		Done: true,
		Msg:  "更新完整接口成功",
	}, nil
}

func (s *FullInterfaceService) UpdateMargin(ctx context.Context, req *charge.UpdateMarginReq) (*charge.UpdateMarginResp, error) {
	inf, err := s.FullInterfaceMongoMapper.FindOne(ctx, req.Id)
	if err != nil {
		return &charge.UpdateMarginResp{
			Done: false,
			Msg:  "完整接口不存在或已删除",
		}, err
	}
	if req.Increment < 0 || inf.Margin+req.Increment < 0 {
		return &charge.UpdateMarginResp{
			Done: false,
			Msg:  "接口余量不足",
		}, err
	}
	err = s.FullInterfaceMongoMapper.UpdateMargin(ctx, req.Id, req.Increment)
	if err != nil {
		return &charge.UpdateMarginResp{
			Done: false,
			Msg:  "接口余量更新失败",
		}, err
	}
	return &charge.UpdateMarginResp{
		Done: true,
		Msg:  "接口余量" + strconv.FormatInt(req.Increment, 10),
	}, err
}

func (s *FullInterfaceService) DeleteFullInterface(ctx context.Context, req *charge.DeleteFullInterfaceReq) (*charge.DeleteFullInterfaceResp, error) {
	err := s.FullInterfaceMongoMapper.Delete(ctx, req.Id)
	if err != nil {
		return &charge.DeleteFullInterfaceResp{
			Done: false,
			Msg:  "删除完整接口失败",
		}, err
	}
	return &charge.DeleteFullInterfaceResp{
		Done: true,
		Msg:  "删除完整接口成功",
	}, err

}

func (s *FullInterfaceService) GetFullInterface(ctx context.Context, req *charge.GetFullInterfaceReq) (*charge.GetFullInterfaceResp, error) {
	userId := req.User.UserId
	data, total, err := s.FullInterfaceMongoMapper.FindAndCountByUserId(ctx, userId, req.PaginationOptions)
	if err != nil {
		return nil, err
	}
	infs := make([]*charge.FullInterface, 0)
	for _, val := range data {
		inf := &charge.FullInterface{}
		err := copier.Copy(inf, val)
		if err != nil {
			return nil, err
		}
		inf.Id = val.ID.Hex()
		inf.CreateTime = val.CreateTime.Unix()
		inf.UpdateTime = val.UpdateTime.Unix()
		inf.Status = charge.InterfaceStatus(val.Status)
		infs = append(infs, inf)
	}
	return &charge.GetFullInterfaceResp{
		FullInterfaces: infs,
		Total:          total,
	}, nil
}

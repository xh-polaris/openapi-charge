package service

import (
	"context"
	"errors"
	"github.com/google/wire"
	"github.com/jinzhu/copier"
	"github.com/xh-polaris/openapi-charge/biz/infrastructure/consts"
	"github.com/xh-polaris/openapi-charge/biz/infrastructure/mapper/base"
	"github.com/xh-polaris/openapi-charge/biz/infrastructure/mapper/full"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/openapi/charge"
	"time"
)

type IFullInterfaceService interface {
	CreateFullInterface(ctx context.Context, req *charge.CreateFullInterfaceReq) (*charge.CreateFullInterfaceResp, error)
	UpdateFullInterface(ctx context.Context, req *charge.UpdateFullInterfaceReq) (*charge.UpdateFullInterfaceResp, error)
	DeleteFullInterface(ctx context.Context, req *charge.DeleteFullInterfaceReq) (*charge.DeleteFullInterfaceResp, error)
	GetFullInterface(ctx context.Context, req *charge.GetFullInterfaceReq) (*charge.GetFullInterfaceResp, error)
	GetOneFullInterface(ctx context.Context, req *charge.GetOneFullInterfaceReq) (res *charge.GetOneFullInterfaceResp, err error)
	GetFullAndBaseInterfaceForCheck(ctx context.Context, req *charge.GetFullAndBaseInterfaceForCheckReq) (*charge.GetFullAndBaseInterfaceForCheckResp, error)
}

type FullInterfaceService struct {
	FullInterfaceMongoMapper *full.MongoMapper
	BaseInterfaceMongoMapper *base.MongoMapper
}

var FullInterfaceServiceSet = wire.NewSet(
	wire.Struct(new(FullInterfaceService), "*"),
	wire.Bind(new(IFullInterfaceService), new(*FullInterfaceService)),
)

func (s *FullInterfaceService) CreateFullInterface(ctx context.Context, req *charge.CreateFullInterfaceReq) (*charge.CreateFullInterfaceResp, error) {
	oldInf, err := s.FullInterfaceMongoMapper.FindOneByBaseInfIdAndUserId(ctx, req.BaseInterfaceId, req.UserId)
	if err == nil && oldInf != nil {
		return &charge.CreateFullInterfaceResp{
			Done:            false,
			Msg:             "已创建过完整接口",
			FullInterfaceId: oldInf.ID.Hex(),
		}, nil
	}

	now := time.Now()
	inf := &full.Interface{
		BaseInterfaceId: req.BaseInterfaceId,
		UserId:          req.UserId,
		ChargeType:      int64(req.ChargeType),
		Price:           req.Price,
		Status:          0,
		CreateTime:      now,
		UpdateTime:      now,
	}
	fullInterfaceId, err := s.FullInterfaceMongoMapper.Insert(ctx, inf)
	if err != nil {
		return &charge.CreateFullInterfaceResp{
			Done:            false,
			Msg:             "创建完整接口失败",
			FullInterfaceId: fullInterfaceId,
		}, err
	}
	return &charge.CreateFullInterfaceResp{
		Done:            true,
		Msg:             "创建完整接口成功",
		FullInterfaceId: fullInterfaceId,
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
	userId := req.UserId
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

func (s *FullInterfaceService) GetOneFullInterface(ctx context.Context, req *charge.GetOneFullInterfaceReq) (res *charge.GetOneFullInterfaceResp, err error) {
	id := req.Id
	inf, err := s.FullInterfaceMongoMapper.FindOne(ctx, id)
	if err != nil {
		return nil, err
	}
	if inf == nil {
		return nil, consts.ErrNotFound
	}
	res = &charge.GetOneFullInterfaceResp{
		Inf: &charge.FullInterface{
			Id:              inf.ID.Hex(),
			BaseInterfaceId: inf.BaseInterfaceId,
			UserId:          inf.UserId,
			ChargeType:      charge.ChargeType(inf.ChargeType),
			Price:           inf.Price,
			Status:          charge.InterfaceStatus(inf.Status),
			CreateTime:      inf.CreateTime.Unix(),
			UpdateTime:      inf.UpdateTime.Unix(),
		},
	}
	return res, nil
}

func (s *FullInterfaceService) GetFullAndBaseInterfaceForCheck(ctx context.Context, req *charge.GetFullAndBaseInterfaceForCheckReq) (*charge.GetFullAndBaseInterfaceForCheckResp, error) {
	url := req.Url
	userId := req.UserId
	method := req.Method

	// 获取基本接口
	baseInf, err := s.BaseInterfaceMongoMapper.FindOneByURLAndMethod(ctx, url, method)
	if err != nil {
		return &charge.GetFullAndBaseInterfaceForCheckResp{
			Id:                  "",
			BaseInterfaceId:     "",
			BaseInterfaceStatus: 0,
			UserId:              "",
			ChargeType:          0,
			Price:               0,
			Status:              0,
		}, consts.ErrNoBaseInf
	}

	// 获取完整接口
	fullInf, err := s.FullInterfaceMongoMapper.FindOneByBaseInfIdAndUserId(ctx, baseInf.ID.Hex(), userId)
	// 若没有找到则是用的模板
	if err != nil {
		if errors.Is(err, consts.ErrNotFound) {
			var templateUserId string
			if req.Role == 1 {
				templateUserId = consts.DeveloperRole
			} else {
				templateUserId = consts.EnterpriseRole
			}

			// 查找模板
			fullInf, err = s.FullInterfaceMongoMapper.FindOneByBaseInfIdAndUserId(ctx, baseInf.ID.Hex(), templateUserId)
		}
		if err != nil {
			return &charge.GetFullAndBaseInterfaceForCheckResp{
				Id:                  "",
				BaseInterfaceId:     "",
				BaseInterfaceStatus: 0,
				UserId:              "",
				ChargeType:          0,
				Price:               0,
				Status:              0,
			}, consts.ErrNoBaseInf
		}
	}

	return &charge.GetFullAndBaseInterfaceForCheckResp{
		Id:                  fullInf.ID.Hex(),
		BaseInterfaceId:     baseInf.ID.Hex(),
		BaseInterfaceStatus: baseInf.Status,
		UserId:              fullInf.UserId,
		ChargeType:          fullInf.ChargeType,
		Price:               fullInf.Price,
		Status:              fullInf.Status,
	}, nil

}

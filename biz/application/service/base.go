package service

import (
	"context"
	"github.com/google/wire"
	"github.com/jinzhu/copier"
	"github.com/xh-polaris/openapi-charge/biz/infrastructure/mapper/base"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/openapi/charge"
	"time"
)

type IBaseInterfaceService interface {
	CreateBaseInterface(ctx context.Context, req *charge.CreateBaseInterfaceReq) (*charge.CreateBaseInterfaceResp, error)
	UpdateBaseInterface(ctx context.Context, req *charge.UpdateBaseInterfaceReq) (*charge.UpdateBaseInterfaceResp, error)
	DeleteBaseInterface(ctx context.Context, req *charge.DeleteBaseInterfaceReq) (*charge.DeleteBaseInterfaceResp, error)
	GetBaseInterface(ctx context.Context, req *charge.GetBaseInterfaceReq) (*charge.GetBaseInterfaceResp, error)
}

type BaseInterfaceService struct {
	BaseInterfaceMongoMapper *base.MongoMapper
}

var BaseInterfaceServiceSet = wire.NewSet(
	wire.Struct(new(BaseInterfaceService), "*"),
	wire.Bind(new(IBaseInterfaceService), new(*BaseInterfaceService)),
)

func (s *BaseInterfaceService) CreateBaseInterface(ctx context.Context, req *charge.CreateBaseInterfaceReq) (*charge.CreateBaseInterfaceResp, error) {
	params := ParseParams(req.Params)
	now := time.Now()
	inf := &base.Interface{
		Name:       req.Name,
		Host:       req.Host,
		Path:       req.Path,
		Method:     int64(req.Method),
		PassWay:    int64(req.PassWay),
		Params:     params,
		Content:    req.Content,
		Status:     0,
		CreateTime: now,
		UpdateTime: now,
	}
	err := s.BaseInterfaceMongoMapper.Insert(ctx, inf)
	if err != nil {
		return &charge.CreateBaseInterfaceResp{
			Done: false,
			Msg:  "创建基础接口失败",
		}, err
	}
	return &charge.CreateBaseInterfaceResp{
		Done: true,
		Msg:  "创建基础接口成功",
	}, nil
}

func (s *BaseInterfaceService) UpdateBaseInterface(ctx context.Context, req *charge.UpdateBaseInterfaceReq) (*charge.UpdateBaseInterfaceResp, error) {
	inf, err := s.BaseInterfaceMongoMapper.FindOne(ctx, req.Id)
	if err != nil {
		return &charge.UpdateBaseInterfaceResp{
			Done: false,
			Msg:  "基础接口不存在或已删除",
		}, err
	}
	inf.Name = req.Name
	inf.Host = req.Host
	inf.Path = req.Path
	inf.Method = int64(req.Method)
	inf.PassWay = int64(req.PassWay)
	inf.Params = ParseParams(req.Params)
	inf.Content = req.Content
	inf.Status = int64(req.Status)

	err = s.BaseInterfaceMongoMapper.Update(ctx, inf)
	if err != nil {
		return &charge.UpdateBaseInterfaceResp{
			Done: false,
			Msg:  "更新基础接口失败",
		}, err
	}
	return &charge.UpdateBaseInterfaceResp{
		Done: true,
		Msg:  "更新基础接口成功",
	}, nil
}

func (s *BaseInterfaceService) DeleteBaseInterface(ctx context.Context, req *charge.DeleteBaseInterfaceReq) (*charge.DeleteBaseInterfaceResp, error) {
	err := s.BaseInterfaceMongoMapper.Delete(ctx, req.Id)
	if err != nil {
		return &charge.DeleteBaseInterfaceResp{
			Done: false,
			Msg:  "删除失败",
		}, err
	}
	return &charge.DeleteBaseInterfaceResp{
		Done: true,
		Msg:  "删除成功",
	}, nil
}

func (s *BaseInterfaceService) GetBaseInterface(ctx context.Context, req *charge.GetBaseInterfaceReq) (*charge.GetBaseInterfaceResp, error) {
	data, total, err := s.BaseInterfaceMongoMapper.FindAndCount(ctx, req.PaginationOptions)
	if err != nil {
		return nil, err
	}
	infs := make([]*charge.BaseInterface, 0)
	for _, val := range data {
		inf := &charge.BaseInterface{}
		err = copier.Copy(inf, val)
		if err != nil {
			return nil, err
		}
		inf.Id = val.ID.Hex()
		inf.CreateTime = val.CreateTime.Unix()
		inf.UpdateTime = val.UpdateTime.Unix()
		inf.Status = charge.InterfaceStatus(val.Status)
		infs = append(infs, inf)
	}
	return &charge.GetBaseInterfaceResp{
		BaseInterfaces: infs,
		Total:          total,
	}, nil
}

func ParseParams(parameters []*charge.Parameter) map[string]string {
	params := make(map[string]string)
	for _, param := range parameters {
		params[param.Name] = param.Type
	}
	return params
}

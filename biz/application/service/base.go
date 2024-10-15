package service

import (
	"context"
	"github.com/google/wire"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/openapi/charge"
)

type IBaseInterfaceService interface {
	CreateBaseInterface(ctx context.Context, req *charge.CreateBaseInterfaceReq) (*charge.CreateBaseInterfaceResp, error)
	UpdateBaseInterface(ctx context.Context, req *charge.UpdateBaseInterfaceReq) (*charge.UpdateBaseInterfaceResp, error)
	DeleteBaseInterface(ctx context.Context, req *charge.DeleteBaseInterfaceReq) (*charge.DeleteBaseInterfaceResp, error)
	GetBaseInterface(ctx context.Context, req *charge.GetBaseInterfaceReq) (*charge.GetBaseInterfaceResp, error)
}

type BaseInterfaceService struct {
}

var BaseInterfaceServiceSet = wire.NewSet(
	wire.Struct(new(BaseInterfaceService), "*"),
	wire.Bind(new(IBaseInterfaceService), new(*BaseInterfaceService)),
)

func (s *BaseInterfaceService) CreateBaseInterface(ctx context.Context, req *charge.CreateBaseInterfaceReq) (*charge.CreateBaseInterfaceResp, error) {
	//TODO implement me
	panic("implement me")
}

func (s *BaseInterfaceService) UpdateBaseInterface(ctx context.Context, req *charge.UpdateBaseInterfaceReq) (*charge.UpdateBaseInterfaceResp, error) {
	//TODO implement me
	panic("implement me")
}

func (s *BaseInterfaceService) DeleteBaseInterface(ctx context.Context, req *charge.DeleteBaseInterfaceReq) (*charge.DeleteBaseInterfaceResp, error) {
	//TODO implement me
	panic("implement me")
}

func (s *BaseInterfaceService) GetBaseInterface(ctx context.Context, req *charge.GetBaseInterfaceReq) (*charge.GetBaseInterfaceResp, error) {
	//TODO implement me
	panic("implement me")
}

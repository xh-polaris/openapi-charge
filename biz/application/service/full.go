package service

import (
	"context"
	"github.com/google/wire"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/openapi/charge"
)

type IFullInterfaceService interface {
	CreateFullInterface(ctx context.Context, req *charge.CreateFullInterfaceReq) (*charge.CreateFullInterfaceResp, error)
	UpdateFullInterface(ctx context.Context, req *charge.UpdateFullInterfaceReq) (*charge.UpdateFullInterfaceResp, error)
	UpdateMargin(ctx context.Context, req *charge.UpdateMarginReq) (*charge.UpdateMarginResp, error)
	DeleteFullInterface(ctx context.Context, req *charge.DeleteFullInterfaceReq) (*charge.DeleteFullInterfaceResp, error)
	GetFullInterface(ctx context.Context, req *charge.GetFullInterfaceReq) (*charge.GetFullInterfaceResp, error)
}

type FullInterfaceService struct {
}

var FullInterfaceServiceSet = wire.NewSet(
	wire.Struct(new(FullInterfaceService), "*"),
	wire.Bind(new(IFullInterfaceService), new(*FullInterfaceService)),
)

func (s *FullInterfaceService) CreateFullInterface(ctx context.Context, req *charge.CreateFullInterfaceReq) (*charge.CreateFullInterfaceResp, error) {
	//TODO implement me
	panic("implement me")
}

func (s *FullInterfaceService) UpdateFullInterface(ctx context.Context, req *charge.UpdateFullInterfaceReq) (*charge.UpdateFullInterfaceResp, error) {
	//TODO implement me
	panic("implement me")
}

func (s *FullInterfaceService) UpdateMargin(ctx context.Context, req *charge.UpdateMarginReq) (*charge.UpdateMarginResp, error) {
	//TODO implement me
	panic("implement me")
}

func (s *FullInterfaceService) DeleteFullInterface(ctx context.Context, req *charge.DeleteFullInterfaceReq) (*charge.DeleteFullInterfaceResp, error) {
	//TODO implement me
	panic("implement me")
}

func (s *FullInterfaceService) GetFullInterface(ctx context.Context, req *charge.GetFullInterfaceReq) (*charge.GetFullInterfaceResp, error) {
	//TODO implement me
	panic("implement me")
}

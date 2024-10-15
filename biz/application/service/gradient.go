package service

import (
	"context"
	"github.com/google/wire"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/openapi/charge"
)

type IGradientService interface {
	CreateGradient(ctx context.Context, req *charge.CreateGradientReq) (*charge.CreateGradientResp, error)
	UpdateGradient(ctx context.Context, req *charge.UpdateGradientReq) (*charge.UpdateGradientResp, error)
	GetGradient(ctx context.Context, req *charge.GetGradientReq) (*charge.GetGradientResp, error)
}

type GradientService struct {
}

var GradientServiceSet = wire.NewSet(
	wire.Struct(new(GradientService)),
	wire.Bind(new(IGradientService), new(*GradientService)),
)

func (s *GradientService) CreateGradient(ctx context.Context, req *charge.CreateGradientReq) (*charge.CreateGradientResp, error) {
	//TODO implement me
	panic("implement me")
}

func (s *GradientService) UpdateGradient(ctx context.Context, req *charge.UpdateGradientReq) (*charge.UpdateGradientResp, error) {
	//TODO implement me
	panic("implement me")
}

func (s *GradientService) GetGradient(ctx context.Context, req *charge.GetGradientReq) (*charge.GetGradientResp, error) {
	//TODO implement me
	panic("implement me")
}

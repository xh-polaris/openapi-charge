package service

import (
	"context"
	"github.com/google/wire"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/openapi/charge"
)

type ILogService interface {
	CreateLog(ctx context.Context, req *charge.CreateLogReq) (*charge.CreateLogResp, error)
	GetLog(ctx context.Context, req *charge.GetLogReq) (*charge.GetLogResp, error)
}

type LogService struct {
}

var LogServiceSet = wire.NewSet(
	wire.Struct(new(LogService), "*"),
	wire.Bind(new(ILogService), new(*LogService)),
)

func (s *LogService) CreateLog(ctx context.Context, req *charge.CreateLogReq) (res *charge.CreateLogResp, err error) {
	//TODO implement me
	panic("implement me")
}

func (s *LogService) GetLog(ctx context.Context, req *charge.GetLogReq) (res *charge.GetLogResp, err error) {
	//TODO implement me
	panic("implement me")
}

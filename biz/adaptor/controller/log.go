package controller

import (
	"context"
	"github.com/google/wire"
	"github.com/xh-polaris/openapi-charge/biz/application/service"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/openapi/charge"
)

type ILogController interface {
	CreateLog(ctx context.Context, req *charge.CreateLogReq) (res *charge.CreateLogResp, err error)
	GetLog(ctx context.Context, req *charge.GetLogReq) (res *charge.GetLogResp, err error)
	GetAccountByTxId(ctx context.Context, req *charge.GetAccountByTxIdReq) (res *charge.GetAccountByTxIdResp, err error)
}

type LogController struct {
	LogService     service.ILogService
	AccountService service.IAccountService
}

var LogControllerSet = wire.NewSet(
	wire.Struct(new(LogController), "*"),
	wire.Bind(new(ILogController), new(*LogController)),
)

func (c *LogController) CreateLog(ctx context.Context, req *charge.CreateLogReq) (res *charge.CreateLogResp, err error) {
	return c.LogService.CreateLog(ctx, req)
}
func (c *LogController) GetLog(ctx context.Context, req *charge.GetLogReq) (res *charge.GetLogResp, err error) {
	return c.LogService.GetLog(ctx, req)
}

func (c *LogController) GetAccountByTxId(ctx context.Context, req *charge.GetAccountByTxIdReq) (res *charge.GetAccountByTxIdResp, err error) {
	return c.AccountService.GetAccountByTxId(ctx, req)
}

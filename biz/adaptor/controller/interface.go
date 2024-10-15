package controller

import (
	"context"
	"github.com/google/wire"
	"github.com/xh-polaris/openapi-charge/biz/application/service"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/openapi/charge"
)

type IInterfaceController interface {
	CreateBaseInterface(ctx context.Context, req *charge.CreateBaseInterfaceReq) (res *charge.CreateBaseInterfaceResp, err error)
	UpdateBaseInterface(ctx context.Context, req *charge.UpdateBaseInterfaceReq) (res *charge.UpdateBaseInterfaceResp, err error)
	DeleteBaseInterface(ctx context.Context, req *charge.DeleteBaseInterfaceReq) (res *charge.DeleteBaseInterfaceResp, err error)
	GetBaseInterface(ctx context.Context, req *charge.GetBaseInterfaceReq) (res *charge.GetBaseInterfaceResp, err error)
	CreateFullInterface(ctx context.Context, req *charge.CreateFullInterfaceReq) (res *charge.CreateFullInterfaceResp, err error)
	UpdateFullInterface(ctx context.Context, req *charge.UpdateFullInterfaceReq) (res *charge.UpdateFullInterfaceResp, err error)
	UpdateMargin(ctx context.Context, req *charge.UpdateMarginReq) (res *charge.UpdateMarginResp, err error)
	DeleteFullInterface(ctx context.Context, req *charge.DeleteFullInterfaceReq) (res *charge.DeleteFullInterfaceResp, err error)
	GetFullInterface(ctx context.Context, req *charge.GetFullInterfaceReq) (res *charge.GetFullInterfaceResp, err error)
	CreateGradient(ctx context.Context, req *charge.CreateGradientReq) (res *charge.CreateGradientResp, err error)
	UpdateGradient(ctx context.Context, req *charge.UpdateGradientReq) (res *charge.UpdateGradientResp, err error)
	GetGradient(ctx context.Context, req *charge.GetGradientReq) (res *charge.GetGradientResp, err error)
}

type InterfaceController struct {
	BaseInterfaceService service.IBaseInterfaceService
	FullInterfaceService service.IFullInterfaceService
	GradientService      service.IGradientService
}

var InterfaceControllerSet = wire.NewSet(
	wire.Struct(new(InterfaceController), "*"),
	wire.Bind(new(IInterfaceController), new(*InterfaceController)),
)

func (c *InterfaceController) CreateBaseInterface(ctx context.Context, req *charge.CreateBaseInterfaceReq) (res *charge.CreateBaseInterfaceResp, err error) {
	return c.BaseInterfaceService.CreateBaseInterface(ctx, req)
}
func (c *InterfaceController) UpdateBaseInterface(ctx context.Context, req *charge.UpdateBaseInterfaceReq) (res *charge.UpdateBaseInterfaceResp, err error) {
	return c.BaseInterfaceService.UpdateBaseInterface(ctx, req)
}
func (c *InterfaceController) DeleteBaseInterface(ctx context.Context, req *charge.DeleteBaseInterfaceReq) (res *charge.DeleteBaseInterfaceResp, err error) {
	return c.BaseInterfaceService.DeleteBaseInterface(ctx, req)
}
func (c *InterfaceController) GetBaseInterface(ctx context.Context, req *charge.GetBaseInterfaceReq) (res *charge.GetBaseInterfaceResp, err error) {
	return c.BaseInterfaceService.GetBaseInterface(ctx, req)
}
func (c *InterfaceController) CreateFullInterface(ctx context.Context, req *charge.CreateFullInterfaceReq) (res *charge.CreateFullInterfaceResp, err error) {
	return c.FullInterfaceService.CreateFullInterface(ctx, req)
}
func (c *InterfaceController) UpdateFullInterface(ctx context.Context, req *charge.UpdateFullInterfaceReq) (res *charge.UpdateFullInterfaceResp, err error) {
	return c.FullInterfaceService.UpdateFullInterface(ctx, req)
}
func (c *InterfaceController) UpdateMargin(ctx context.Context, req *charge.UpdateMarginReq) (res *charge.UpdateMarginResp, err error) {
	return c.FullInterfaceService.UpdateMargin(ctx, req)
}
func (c *InterfaceController) DeleteFullInterface(ctx context.Context, req *charge.DeleteFullInterfaceReq) (res *charge.DeleteFullInterfaceResp, err error) {
	return c.FullInterfaceService.DeleteFullInterface(ctx, req)
}
func (c *InterfaceController) GetFullInterface(ctx context.Context, req *charge.GetFullInterfaceReq) (res *charge.GetFullInterfaceResp, err error) {
	return c.FullInterfaceService.GetFullInterface(ctx, req)
}
func (c *InterfaceController) CreateGradient(ctx context.Context, req *charge.CreateGradientReq) (res *charge.CreateGradientResp, err error) {
	return c.GradientService.CreateGradient(ctx, req)
}
func (c *InterfaceController) UpdateGradient(ctx context.Context, req *charge.UpdateGradientReq) (res *charge.UpdateGradientResp, err error) {
	return c.GradientService.UpdateGradient(ctx, req)
}
func (c *InterfaceController) GetGradient(ctx context.Context, req *charge.GetGradientReq) (res *charge.GetGradientResp, err error) {
	return c.GradientService.GetGradient(ctx, req)
}

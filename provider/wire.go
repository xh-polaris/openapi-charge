//go:build wireinject
// +build wireinject

package provider

import (
	"github.com/google/wire"
	"github.com/xh-polaris/openapi-charge/biz/adaptor"
	"github.com/xh-polaris/openapi-charge/provider"
)

func NewProvider() (*adaptor.ChargeServer, error) {
	wire.Build(
		wire.Struct(new(adaptor.ChargeServer), "*"),
		provider.ChargeServerProvider,
	)
	return nil, nil
}

//go:build wireinject
// +build wireinject

package provider

import (
	"github.com/google/wire"
	"github.com/xh-polaris/openapi-charge/biz/adaptor"
)

func NewProvider() (*adaptor.ChargeServer, error) {
	wire.Build(
		wire.Struct(new(adaptor.ChargeServer), "*"),
		AllProvider,
	)
	return nil, nil
}

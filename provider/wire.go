//go:build wireinject
// +build wireinject

package provider

import (
	"github.com/google/wire"
	"github.com/xhpolaris/opeanapi-user/biz/adaptor"
)

func NewProvider() (*adaptor.UserServer, error) {
	wire.Build(
		wire.Struct(new(adaptor.UserServer), "*"),
		AllProvider,
	)
	return nil, nil
}

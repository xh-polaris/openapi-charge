package provider

import (
	"github.com/google/wire"
	"github.com/xh-polaris/openapi-charge/biz/adaptor/controller"
	"github.com/xh-polaris/openapi-charge/biz/application/service"
	"github.com/xh-polaris/openapi-charge/biz/infrastructure/config"
	"github.com/xh-polaris/openapi-charge/biz/infrastructure/mapper/base"
	"github.com/xh-polaris/openapi-charge/biz/infrastructure/mapper/full"
	"github.com/xh-polaris/openapi-charge/biz/infrastructure/mapper/gradient"
	"github.com/xh-polaris/openapi-charge/biz/infrastructure/mapper/log"
)

var ChargeServerProvider = wire.NewSet(
	ControllerSet,
	ApplicationSet,
	InfrastructureSet,
)

var ControllerSet = wire.NewSet(
	controller.InterfaceControllerSet,
	controller.LogControllerSet,
)

var ApplicationSet = wire.NewSet(
	service.BaseInterfaceServiceSet,
	service.FullInterfaceServiceSet,
	service.GradientServiceSet,
	service.LogServiceSet,
	service.MarginServiceSet,
)

var InfrastructureSet = wire.NewSet(
	config.NewConfig,
	MapperSet,
)

var MapperSet = wire.NewSet(
	base.NewMongoMapper,
	full.NewMongoMapper,
	gradient.NewMongoMapper,
	log.NewMongoMapper,
)

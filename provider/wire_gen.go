// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package provider

import (
	"github.com/xh-polaris/openapi-charge/biz/adaptor"
	"github.com/xh-polaris/openapi-charge/biz/adaptor/controller"
	"github.com/xh-polaris/openapi-charge/biz/application/service"
	"github.com/xh-polaris/openapi-charge/biz/infrastructure/config"
	"github.com/xh-polaris/openapi-charge/biz/infrastructure/mapper/account"
	"github.com/xh-polaris/openapi-charge/biz/infrastructure/mapper/base"
	"github.com/xh-polaris/openapi-charge/biz/infrastructure/mapper/full"
	"github.com/xh-polaris/openapi-charge/biz/infrastructure/mapper/gradient"
	"github.com/xh-polaris/openapi-charge/biz/infrastructure/mapper/log"
	"github.com/xh-polaris/openapi-charge/biz/infrastructure/mapper/margin"
	"github.com/xh-polaris/openapi-charge/biz/infrastructure/transaction"
)

// Injectors from wire.go:

func NewProvider() (*adaptor.ChargeServer, error) {
	configConfig, err := config.NewConfig()
	if err != nil {
		return nil, err
	}
	mongoMapper := base.NewMongoMapper(configConfig)
	baseInterfaceService := &service.BaseInterfaceService{
		BaseInterfaceMongoMapper: mongoMapper,
	}
	fullMongoMapper := full.NewMongoMapper(configConfig)
	fullInterfaceService := &service.FullInterfaceService{
		FullInterfaceMongoMapper: fullMongoMapper,
		BaseInterfaceMongoMapper: mongoMapper,
	}
	gradientMongoMapper := gradient.NewMongoMapper(configConfig)
	gradientService := &service.GradientService{
		GradientMongoMapper:      gradientMongoMapper,
		FullInterfaceMongoMapper: fullMongoMapper,
	}
	marginMongoMapper := margin.NewMongoMapper(configConfig)
	marginTransaction := transaction.NewMarginTransaction(configConfig)
	marginService := &service.MarginService{
		MarginMongoMapper: marginMongoMapper,
		MarginTransaction: marginTransaction,
	}
	interfaceController := &controller.InterfaceController{
		BaseInterfaceService: baseInterfaceService,
		FullInterfaceService: fullInterfaceService,
		GradientService:      gradientService,
		MarginService:        marginService,
	}
	logMongoMapper := log.NewMongoMapper(configConfig)
	logService := &service.LogService{
		MarginService:   marginService,
		LogMongoMapper:  logMongoMapper,
		FullMongoMapper: fullMongoMapper,
	}
	accountMongoMapper := account.NewMongoMapper(configConfig)
	accountService := &service.AccountService{
		AccountMongoMapper: accountMongoMapper,
	}
	logController := &controller.LogController{
		LogService:     logService,
		AccountService: accountService,
	}
	chargeServer := &adaptor.ChargeServer{
		IInterfaceController: interfaceController,
		ILogController:       logController,
	}
	return chargeServer, nil
}

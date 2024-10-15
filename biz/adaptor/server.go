package adaptor

import "github.com/xh-polaris/openapi-charge/biz/adaptor/controller"

type ChargeServer struct {
	controller.IInterfaceController
	controller.ILogController
}

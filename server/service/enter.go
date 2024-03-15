package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/hrms"
	"github.com/flipped-aurora/gin-vue-admin/server/service/htxm"
	"github.com/flipped-aurora/gin-vue-admin/server/service/moneyms"
	"github.com/flipped-aurora/gin-vue-admin/server/service/msysmoney"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup    system.ServiceGroup
	ExampleServiceGroup   example.ServiceGroup
	HrmsServiceGroup      hrms.ServiceGroup
	MsysmoneyServiceGroup msysmoney.ServiceGroup
	MoneymsServiceGroup   moneyms.ServiceGroup
	HtxmServiceGroup      htxm.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)

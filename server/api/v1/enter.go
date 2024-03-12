package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/hrms"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/moneyms"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/msysmoney"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup    system.ApiGroup
	ExampleApiGroup   example.ApiGroup
	HrmsApiGroup      hrms.ApiGroup
	MsysmoneyApiGroup msysmoney.ApiGroup
	MoneymsApiGroup   moneyms.ApiGroup
}

var ApiGroupApp = new(ApiGroup)

package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/hrms"
	"github.com/flipped-aurora/gin-vue-admin/server/router/moneyms"
	"github.com/flipped-aurora/gin-vue-admin/server/router/msysmoney"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

type RouterGroup struct {
	System    system.RouterGroup
	Example   example.RouterGroup
	Hrms      hrms.RouterGroup
	Msysmoney msysmoney.RouterGroup
	Moneyms   moneyms.RouterGroup
}

var RouterGroupApp = new(RouterGroup)

package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type MsXmListSearch struct {

	//StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	//EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`

	Khmc      string     `json:"khmc" form:"khmc" `
	Xmmc      string     `json:"xmmc" form:"xmmc" `
	StartJssj *time.Time `json:"startJssj" form:"startJssj"`
	EndJssj   *time.Time `json:"endJssj" form:"endJssj"`
	Xmjl      string     `json:"xmjl" form:"xmjl" `
	Zxbm      string     `json:"zxbm" form:"zxbm" `
	Xmcs      string     `json:"xmcs" form:"xmcs" `
	Gdmll     *float64   `json:"gdmll" form:"gdmll" `
	Ztbj      string     `json:"ztbj" form:"ztbj" `
	request.PageInfo
	Sort  string `json:"sort" form:"sort"`
	Order string `json:"order" form:"order"`
}

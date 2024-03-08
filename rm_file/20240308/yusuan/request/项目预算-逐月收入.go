package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
	
)

type 项目预算-逐月收入Search struct{
    
        StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
        EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    
                      交付部  string `json:"交付部" form:"交付部" `
                      类型  string `json:"类型" form:"类型" `
                      项目名称  string `json:"项目名称" form:"项目名称" `
                      客户系  string `json:"客户系" form:"客户系" `
                      客户名称  string `json:"客户名称" form:"客户名称" `
    request.PageInfo
    Sort  string `json:"sort" form:"sort"`
    Order string `json:"order" form:"order"`
}

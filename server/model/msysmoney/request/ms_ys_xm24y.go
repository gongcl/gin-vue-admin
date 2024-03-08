package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
	
)

type MsYsXm24ySearch struct{
    
        StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
        EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    
                      Department  string `json:"Department" form:"Department" `
                      YusuanType  string `json:"YusuanType" form:"YusuanType" `
                      ProjectName  string `json:"ProjectName" form:"ProjectName" `
                      CustomGroup  string `json:"CustomGroup" form:"CustomGroup" `
                      CustomerName  string `json:"CustomerName" form:"CustomerName" `
                      TotalNR  *float64 `json:"TotalNR" form:"TotalNR" `
    request.PageInfo
    Sort  string `json:"sort" form:"sort"`
    Order string `json:"order" form:"order"`
}

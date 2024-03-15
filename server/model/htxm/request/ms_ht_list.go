package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
	
)

type MsHtListSearch struct{
    
                      Dept  string `json:"dept" form:"dept" `
                      Htmc  string `json:"htmc" form:"htmc" `
                StartJsrq  *time.Time  `json:"startJsrq" form:"startJsrq"`
                EndJsrq  *time.Time  `json:"endJsrq" form:"endJsrq"`
                      Sfbl  string `json:"sfbl" form:"sfbl" `
                      Sfel  string `json:"sfel" form:"sfel" `
                      Xmje  *float64 `json:"xmje" form:"xmje" `
    request.PageInfo
    Sort  string `json:"sort" form:"sort"`
    Order string `json:"order" form:"order"`
}

package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	
	
)

type MsYsnrXmSearch struct{
    
                      Department  string `json:"department" form:"department" `
                      ProjectName  string `json:"projectName" form:"projectName" `
                      CustomerName  string `json:"customerName" form:"customerName" `
                      NrTotal  *float64 `json:"nrTotal" form:"nrTotal" `
    request.PageInfo
}

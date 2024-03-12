package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	
	
)

type MsYsnrXmSearch struct{
    
                      CustomerName  string `json:"customerName" form:"customerName" `
                      Department  string `json:"department" form:"department" `
                      NrTotal  *float64 `json:"nrTotal" form:"nrTotal" `
                      NrType  string `json:"nrType" form:"nrType" `
                      ProjectName  string `json:"projectName" form:"projectName" `
    request.PageInfo
}

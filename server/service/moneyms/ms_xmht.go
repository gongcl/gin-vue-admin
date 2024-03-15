package moneyms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/moneyms"
	moneymsReq "github.com/flipped-aurora/gin-vue-admin/server/model/moneyms/request"
)

type MsXmHtService struct {
}

func (MsXmHtService *MsXmHtService) GetXmHtList(info moneymsReq.MsYsnrXmSearch) (list []moneyms.MsYsnrXm, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.MustGetGlobalDBByDBName("ms_db").Model(&moneyms.MsYsnrXm{})
	var msYsnrXms []moneyms.MsYsnrXm
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.CustomerName != "" {
		db = db.Where("customer_name LIKE ?", "%"+info.CustomerName+"%")
	}
	if info.Department != "" {
		db = db.Where("department LIKE ?", "%"+info.Department+"%")
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&msYsnrXms).Error
	return msYsnrXms, total, err
}

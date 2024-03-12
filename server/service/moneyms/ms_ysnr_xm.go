package moneyms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/moneyms"
    moneymsReq "github.com/flipped-aurora/gin-vue-admin/server/model/moneyms/request"
)

type MsYsnrXmService struct {
}

// CreateMsYsnrXm 创建项目预算表记录
// Author [piexlmax](https://github.com/piexlmax)
func (msYsnrXmService *MsYsnrXmService) CreateMsYsnrXm(msYsnrXm *moneyms.MsYsnrXm) (err error) {
	err = global.MustGetGlobalDBByDBName("ms_db").Create(msYsnrXm).Error
	return err
}

// DeleteMsYsnrXm 删除项目预算表记录
// Author [piexlmax](https://github.com/piexlmax)
func (msYsnrXmService *MsYsnrXmService)DeleteMsYsnrXm(projectName string) (err error) {
	err = global.MustGetGlobalDBByDBName("ms_db").Delete(&moneyms.MsYsnrXm{},"project_name = ?",projectName).Error
	return err
}

// DeleteMsYsnrXmByIds 批量删除项目预算表记录
// Author [piexlmax](https://github.com/piexlmax)
func (msYsnrXmService *MsYsnrXmService)DeleteMsYsnrXmByIds(projectNames []string) (err error) {
	err = global.MustGetGlobalDBByDBName("ms_db").Delete(&[]moneyms.MsYsnrXm{},"project_name in ?",projectNames).Error
	return err
}

// UpdateMsYsnrXm 更新项目预算表记录
// Author [piexlmax](https://github.com/piexlmax)
func (msYsnrXmService *MsYsnrXmService)UpdateMsYsnrXm(msYsnrXm moneyms.MsYsnrXm) (err error) {
	err = global.MustGetGlobalDBByDBName("ms_db").Save(&msYsnrXm).Error
	return err
}

// GetMsYsnrXm 根据projectName获取项目预算表记录
// Author [piexlmax](https://github.com/piexlmax)
func (msYsnrXmService *MsYsnrXmService)GetMsYsnrXm(projectName string) (msYsnrXm moneyms.MsYsnrXm, err error) {
	err = global.MustGetGlobalDBByDBName("ms_db").Where("project_name = ?", projectName).First(&msYsnrXm).Error
	return
}

// GetMsYsnrXmInfoList 分页获取项目预算表记录
// Author [piexlmax](https://github.com/piexlmax)
func (msYsnrXmService *MsYsnrXmService)GetMsYsnrXmInfoList(info moneymsReq.MsYsnrXmSearch) (list []moneyms.MsYsnrXm, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.MustGetGlobalDBByDBName("ms_db").Model(&moneyms.MsYsnrXm{})
    var msYsnrXms []moneyms.MsYsnrXm
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.CustomerName != "" {
        db = db.Where("customer_name LIKE ?","%"+ info.CustomerName+"%")
    }
    if info.Department != "" {
        db = db.Where("department LIKE ?","%"+ info.Department+"%")
    }
    if info.NrTotal != nil {
        db = db.Where("nr_total > ?",info.NrTotal)
    }
    if info.NrType != "" {
        db = db.Where("nr_type LIKE ?","%"+ info.NrType+"%")
    }
    if info.ProjectName != "" {
        db = db.Where("project_name LIKE ?","%"+ info.ProjectName+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&msYsnrXms).Error
	return  msYsnrXms, total, err
}

package ms_money

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/ms_money"
    ms_moneyReq "github.com/flipped-aurora/gin-vue-admin/server/model/ms_money/request"
)

type MsYsnrXmService struct {
}

// CreateMsYsnrXm 创建msYsnrXm表记录
// Author [piexlmax](https://github.com/piexlmax)
func (msYsnrXmService *MsYsnrXmService) CreateMsYsnrXm(msYsnrXm *ms_money.MsYsnrXm) (err error) {
	err = global.MustGetGlobalDBByDBName("ms_db").Create(msYsnrXm).Error
	return err
}

// DeleteMsYsnrXm 删除msYsnrXm表记录
// Author [piexlmax](https://github.com/piexlmax)
func (msYsnrXmService *MsYsnrXmService)DeleteMsYsnrXm(projectName string) (err error) {
	err = global.MustGetGlobalDBByDBName("ms_db").Delete(&ms_money.MsYsnrXm{},"project_name = ?",projectName).Error
	return err
}

// DeleteMsYsnrXmByIds 批量删除msYsnrXm表记录
// Author [piexlmax](https://github.com/piexlmax)
func (msYsnrXmService *MsYsnrXmService)DeleteMsYsnrXmByIds(projectNames []string) (err error) {
	err = global.MustGetGlobalDBByDBName("ms_db").Delete(&[]ms_money.MsYsnrXm{},"project_name in ?",projectNames).Error
	return err
}

// UpdateMsYsnrXm 更新msYsnrXm表记录
// Author [piexlmax](https://github.com/piexlmax)
func (msYsnrXmService *MsYsnrXmService)UpdateMsYsnrXm(msYsnrXm ms_money.MsYsnrXm) (err error) {
	err = global.MustGetGlobalDBByDBName("ms_db").Save(&msYsnrXm).Error
	return err
}

// GetMsYsnrXm 根据projectName获取msYsnrXm表记录
// Author [piexlmax](https://github.com/piexlmax)
func (msYsnrXmService *MsYsnrXmService)GetMsYsnrXm(projectName string) (msYsnrXm ms_money.MsYsnrXm, err error) {
	err = global.MustGetGlobalDBByDBName("ms_db").Where("project_name = ?", projectName).First(&msYsnrXm).Error
	return
}

// GetMsYsnrXmInfoList 分页获取msYsnrXm表记录
// Author [piexlmax](https://github.com/piexlmax)
func (msYsnrXmService *MsYsnrXmService)GetMsYsnrXmInfoList(info ms_moneyReq.MsYsnrXmSearch) (list []ms_money.MsYsnrXm, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.MustGetGlobalDBByDBName("ms_db").Model(&ms_money.MsYsnrXm{})
    var msYsnrXms []ms_money.MsYsnrXm
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Department != "" {
        db = db.Where("department LIKE ?","%"+ info.Department+"%")
    }
    if info.ProjectName != "" {
        db = db.Where("project_name LIKE ?","%"+ info.ProjectName+"%")
    }
    if info.CustomerName != "" {
        db = db.Where("customer_name LIKE ?","%"+ info.CustomerName+"%")
    }
    if info.NrTotal != nil {
        db = db.Where("nr_total > ?",info.NrTotal)
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

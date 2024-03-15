package htxm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/htxm"
	htxmReq "github.com/flipped-aurora/gin-vue-admin/server/model/htxm/request"
	"gorm.io/gorm"
)

type MsXmListService struct {
}

// CreateMsXmList 创建msXmList表记录
// Author [piexlmax](https://github.com/piexlmax)
func (msXmListService *MsXmListService) CreateMsXmList(msXmList *htxm.MsXmList) (err error) {
	err = global.MustGetGlobalDBByDBName("ms_db").Create(msXmList).Error
	return err
}

// DeleteMsXmList 删除msXmList表记录
// Author [piexlmax](https://github.com/piexlmax)
func (msXmListService *MsXmListService) DeleteMsXmList(ID string, userID uint) (err error) {
	err = global.MustGetGlobalDBByDBName("ms_db").Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&htxm.MsXmList{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&htxm.MsXmList{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteMsXmListByIds 批量删除msXmList表记录
// Author [piexlmax](https://github.com/piexlmax)
func (msXmListService *MsXmListService) DeleteMsXmListByIds(IDs []string, deleted_by uint) (err error) {
	err = global.MustGetGlobalDBByDBName("ms_db").Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&htxm.MsXmList{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&htxm.MsXmList{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateMsXmList 更新msXmList表记录
// Author [piexlmax](https://github.com/piexlmax)
func (msXmListService *MsXmListService) UpdateMsXmList(msXmList htxm.MsXmList) (err error) {
	err = global.MustGetGlobalDBByDBName("ms_db").Save(&msXmList).Error
	return err
}

// GetMsXmList 根据ID获取msXmList表记录
// Author [piexlmax](https://github.com/piexlmax)
func (msXmListService *MsXmListService) GetMsXmList(ID string) (msXmList htxm.MsXmList, err error) {
	err = global.MustGetGlobalDBByDBName("ms_db").Where("id = ?", ID).First(&msXmList).Error
	return
}

// GetMsXmListInfoList 分页获取msXmList表记录
// Author [piexlmax](https://github.com/piexlmax)
func (msXmListService *MsXmListService) GetMsXmListInfoList(info htxmReq.MsXmListSearch) (list []htxm.MsXmList, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.MustGetGlobalDBByDBName("ms_db").Model(&htxm.MsXmList{})
	var msXmLists []htxm.MsXmList
	// 如果有条件搜索 下方会自动创建搜索语句
	//if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
	// db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	//}
	if info.Khmc != "" {
		db = db.Where("khmc LIKE ?", "%"+info.Khmc+"%")
	}
	if info.Xmmc != "" {
		db = db.Where("xmmc LIKE ?", "%"+info.Xmmc+"%")
	}
	if info.StartJssj != nil && info.EndJssj != nil {
		db = db.Where("jssj BETWEEN ? AND ? ", info.StartJssj, info.EndJssj)
	}
	if info.Xmjl != "" {
		db = db.Where("xmjl LIKE ?", "%"+info.Xmjl+"%")
	}
	if info.Zxbm != "" {
		db = db.Where("zxbm LIKE ?", "%"+info.Zxbm+"%")
	}
	if info.Xmcs != "" {
		db = db.Where("xmcs LIKE ?", "%"+info.Xmcs+"%")
	}
	if info.Gdmll != nil {
		db = db.Where("gdmll < ?", info.Gdmll)
	}
	if info.Ztbj != "" {
		db = db.Where("ztbj LIKE ?", "%"+info.Ztbj+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["khmc"] = true
	if orderMap[info.Sort] {
		OrderStr = info.Sort
		if info.Order == "descending" {
			OrderStr = OrderStr + " desc"
		}
		db = db.Order(OrderStr)
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&msXmLists).Error
	return msXmLists, total, err
}

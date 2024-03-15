package htxm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/htxm"
    htxmReq "github.com/flipped-aurora/gin-vue-admin/server/model/htxm/request"
    "gorm.io/gorm"
)

type MsHtListService struct {
}

// CreateMsHtList 创建msHtList表记录
// Author [piexlmax](https://github.com/piexlmax)
func (msHtListService *MsHtListService) CreateMsHtList(msHtList *htxm.MsHtList) (err error) {
	err = global.MustGetGlobalDBByDBName("ms_db").Create(msHtList).Error
	return err
}

// DeleteMsHtList 删除msHtList表记录
// Author [piexlmax](https://github.com/piexlmax)
func (msHtListService *MsHtListService)DeleteMsHtList(htmc string,userID uint) (err error) {
	err = global.MustGetGlobalDBByDBName("ms_db").Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&htxm.MsHtList{}).Where("htmc = ?", htmc).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&htxm.MsHtList{},"htmc = ?",htmc).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteMsHtListByIds 批量删除msHtList表记录
// Author [piexlmax](https://github.com/piexlmax)
func (msHtListService *MsHtListService)DeleteMsHtListByIds(htmcs []string,deleted_by uint) (err error) {
	err = global.MustGetGlobalDBByDBName("ms_db").Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&htxm.MsHtList{}).Where("htmc in ?", htmcs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("htmc in ?", htmcs).Delete(&htxm.MsHtList{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateMsHtList 更新msHtList表记录
// Author [piexlmax](https://github.com/piexlmax)
func (msHtListService *MsHtListService)UpdateMsHtList(msHtList htxm.MsHtList) (err error) {
	err = global.MustGetGlobalDBByDBName("ms_db").Save(&msHtList).Error
	return err
}

// GetMsHtList 根据htmc获取msHtList表记录
// Author [piexlmax](https://github.com/piexlmax)
func (msHtListService *MsHtListService)GetMsHtList(htmc string) (msHtList htxm.MsHtList, err error) {
	err = global.MustGetGlobalDBByDBName("ms_db").Where("htmc = ?", htmc).First(&msHtList).Error
	return
}

// GetMsHtListInfoList 分页获取msHtList表记录
// Author [piexlmax](https://github.com/piexlmax)
func (msHtListService *MsHtListService)GetMsHtListInfoList(info htxmReq.MsHtListSearch) (list []htxm.MsHtList, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.MustGetGlobalDBByDBName("ms_db").Model(&htxm.MsHtList{})
    var msHtLists []htxm.MsHtList
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Dept != "" {
        db = db.Where("dept LIKE ?","%"+ info.Dept+"%")
    }
    if info.Htmc != "" {
        db = db.Where("htmc LIKE ?","%"+ info.Htmc+"%")
    }
        if info.StartJsrq != nil && info.EndJsrq != nil {
            db = db.Where("jsrq BETWEEN ? AND ? ",info.StartJsrq,info.EndJsrq)
        }
    if info.Sfbl != "" {
        db = db.Where("sfbl = ?",info.Sfbl)
    }
    if info.Sfel != "" {
        db = db.Where("sfel = ?",info.Sfel)
    }
    if info.Xmje != nil {
        db = db.Where("xmje > ?",info.Xmje)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
        var OrderStr string
        orderMap := make(map[string]bool)
         	orderMap["dept"] = true
         	orderMap["jsrq"] = true
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
	
	err = db.Find(&msHtLists).Error
	return  msHtLists, total, err
}

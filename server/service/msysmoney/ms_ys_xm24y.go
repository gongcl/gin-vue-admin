package msysmoney

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/msysmoney"
    msysmoneyReq "github.com/flipped-aurora/gin-vue-admin/server/model/msysmoney/request"
    "gorm.io/gorm"
)

type MsYsXm24yService struct {
}

// CreateMsYsXm24y 创建msYsXm24y表记录
// Author [piexlmax](https://github.com/piexlmax)
func (msYsXm24yService *MsYsXm24yService) CreateMsYsXm24y(msYsXm24y *msysmoney.MsYsXm24y) (err error) {
	err = global.GVA_DB.Create(msYsXm24y).Error
	return err
}

// DeleteMsYsXm24y 删除msYsXm24y表记录
// Author [piexlmax](https://github.com/piexlmax)
func (msYsXm24yService *MsYsXm24yService)DeleteMsYsXm24y(ID string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&msysmoney.MsYsXm24y{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&msysmoney.MsYsXm24y{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteMsYsXm24yByIds 批量删除msYsXm24y表记录
// Author [piexlmax](https://github.com/piexlmax)
func (msYsXm24yService *MsYsXm24yService)DeleteMsYsXm24yByIds(IDs []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&msysmoney.MsYsXm24y{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&msysmoney.MsYsXm24y{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateMsYsXm24y 更新msYsXm24y表记录
// Author [piexlmax](https://github.com/piexlmax)
func (msYsXm24yService *MsYsXm24yService)UpdateMsYsXm24y(msYsXm24y msysmoney.MsYsXm24y) (err error) {
	err = global.GVA_DB.Save(&msYsXm24y).Error
	return err
}

// GetMsYsXm24y 根据ID获取msYsXm24y表记录
// Author [piexlmax](https://github.com/piexlmax)
func (msYsXm24yService *MsYsXm24yService)GetMsYsXm24y(ID string) (msYsXm24y msysmoney.MsYsXm24y, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&msYsXm24y).Error
	return
}

// GetMsYsXm24yInfoList 分页获取msYsXm24y表记录
// Author [piexlmax](https://github.com/piexlmax)
func (msYsXm24yService *MsYsXm24yService)GetMsYsXm24yInfoList(info msysmoneyReq.MsYsXm24ySearch) (list []msysmoney.MsYsXm24y, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&msysmoney.MsYsXm24y{})
    var msYsXm24ys []msysmoney.MsYsXm24y
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Department != "" {
        db = db.Where("Department LIKE ?","%"+ info.Department+"%")
    }
    if info.YusuanType != "" {
        db = db.Where("Yusuan_type LIKE ?","%"+ info.YusuanType+"%")
    }
    if info.ProjectName != "" {
        db = db.Where("Project_name LIKE ?","%"+ info.ProjectName+"%")
    }
    if info.CustomGroup != "" {
        db = db.Where("Custom_group LIKE ?","%"+ info.CustomGroup+"%")
    }
    if info.CustomerName != "" {
        db = db.Where("Customer_name LIKE ?","%"+ info.CustomerName+"%")
    }
    if info.TotalNR != nil {
        db = db.Where("TotalNR > ?",info.TotalNR)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
        var OrderStr string
        orderMap := make(map[string]bool)
         	orderMap["Project_code"] = true
         	orderMap["Project_name"] = true
         	orderMap["Customer_name"] = true
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
	
	err = db.Find(&msYsXm24ys).Error
	return  msYsXm24ys, total, err
}

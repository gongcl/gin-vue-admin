package hrms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/hrms"
    hrmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/hrms/request"
    "gorm.io/gorm"
)

type HrEmployeeService struct {
}

// CreateHrEmployee 创建hrEmployee表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hrEmployeeService *HrEmployeeService) CreateHrEmployee(hrEmployee *hrms.HrEmployee) (err error) {
	err = global.GVA_DB.Create(hrEmployee).Error
	return err
}

// DeleteHrEmployee 删除hrEmployee表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hrEmployeeService *HrEmployeeService)DeleteHrEmployee(ID string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&hrms.HrEmployee{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&hrms.HrEmployee{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteHrEmployeeByIds 批量删除hrEmployee表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hrEmployeeService *HrEmployeeService)DeleteHrEmployeeByIds(IDs []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&hrms.HrEmployee{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&hrms.HrEmployee{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateHrEmployee 更新hrEmployee表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hrEmployeeService *HrEmployeeService)UpdateHrEmployee(hrEmployee hrms.HrEmployee) (err error) {
	err = global.GVA_DB.Save(&hrEmployee).Error
	return err
}

// GetHrEmployee 根据ID获取hrEmployee表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hrEmployeeService *HrEmployeeService)GetHrEmployee(ID string) (hrEmployee hrms.HrEmployee, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&hrEmployee).Error
	return
}

// GetHrEmployeeInfoList 分页获取hrEmployee表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hrEmployeeService *HrEmployeeService)GetHrEmployeeInfoList(info hrmsReq.HrEmployeeSearch) (list []hrms.HrEmployee, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&hrms.HrEmployee{})
    var hrEmployees []hrms.HrEmployee
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&hrEmployees).Error
	return  hrEmployees, total, err
}

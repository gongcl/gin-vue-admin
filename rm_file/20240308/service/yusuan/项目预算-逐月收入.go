package yusuan

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/yusuan"
    yusuanReq "github.com/flipped-aurora/gin-vue-admin/server/model/yusuan/request"
)

type 项目预算-逐月收入Service struct {
}

// Create项目预算-逐月收入 创建项目预算-逐月收入表记录
// Author [piexlmax](https://github.com/piexlmax)
func (项目预算Service *项目预算-逐月收入Service) Create项目预算-逐月收入(项目预算 *yusuan.项目预算-逐月收入) (err error) {
	err = global.GVA_DB.Create(项目预算).Error
	return err
}

// Delete项目预算-逐月收入 删除项目预算-逐月收入表记录
// Author [piexlmax](https://github.com/piexlmax)
func (项目预算Service *项目预算-逐月收入Service)Delete项目预算-逐月收入(ID string) (err error) {
	err = global.GVA_DB.Delete(&yusuan.项目预算-逐月收入{},"id = ?",ID).Error
	return err
}

// Delete项目预算-逐月收入ByIds 批量删除项目预算-逐月收入表记录
// Author [piexlmax](https://github.com/piexlmax)
func (项目预算Service *项目预算-逐月收入Service)Delete项目预算-逐月收入ByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]yusuan.项目预算-逐月收入{},"id in ?",IDs).Error
	return err
}

// Update项目预算-逐月收入 更新项目预算-逐月收入表记录
// Author [piexlmax](https://github.com/piexlmax)
func (项目预算Service *项目预算-逐月收入Service)Update项目预算-逐月收入(项目预算 yusuan.项目预算-逐月收入) (err error) {
	err = global.GVA_DB.Save(&项目预算).Error
	return err
}

// Get项目预算-逐月收入 根据ID获取项目预算-逐月收入表记录
// Author [piexlmax](https://github.com/piexlmax)
func (项目预算Service *项目预算-逐月收入Service)Get项目预算-逐月收入(ID string) (项目预算 yusuan.项目预算-逐月收入, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&项目预算).Error
	return
}

// Get项目预算-逐月收入InfoList 分页获取项目预算-逐月收入表记录
// Author [piexlmax](https://github.com/piexlmax)
func (项目预算Service *项目预算-逐月收入Service)Get项目预算-逐月收入InfoList(info yusuanReq.项目预算-逐月收入Search) (list []yusuan.项目预算-逐月收入, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&yusuan.项目预算-逐月收入{})
    var 项目预算s []yusuan.项目预算-逐月收入
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.交付部 != "" {
        db = db.Where("交付部 LIKE ?","%"+ info.交付部+"%")
    }
    if info.类型 != "" {
        db = db.Where("类型 LIKE ?","%"+ info.类型+"%")
    }
    if info.项目名称 != "" {
        db = db.Where("项目名称 LIKE ?","%"+ info.项目名称+"%")
    }
    if info.客户系 != "" {
        db = db.Where("客户系 LIKE ?","%"+ info.客户系+"%")
    }
    if info.客户名称 != "" {
        db = db.Where("客户名称 LIKE ?","%"+ info.客户名称+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
        var OrderStr string
        orderMap := make(map[string]bool)
         	orderMap["交付部"] = true
         	orderMap["类型"] = true
         	orderMap["项目名称"] = true
         	orderMap["客户名称"] = true
         	orderMap["预算NR"] = true
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
	
	err = db.Find(&项目预算s).Error
	return  项目预算s, total, err
}

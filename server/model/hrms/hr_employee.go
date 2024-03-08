// 自动生成模板HrEmployee
package hrms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
	
)

// hrEmployee表 结构体  HrEmployee
type HrEmployee struct {
 global.GVA_MODEL 
      Name  string `json:"name" form:"name" gorm:"column:name;comment:姓名;size:25;"`  //姓名 
      Status  *int `json:"status" form:"status" gorm:"column:status;comment:状态，0：在职，1：离职;size:10;"`  //状态，0：在职，1：离职 
      Type  string `json:"type" form:"type" gorm:"column:type;comment:员工类型;size:25;"`  //员工类型 
      Indate  *time.Time `json:"indate" form:"indate" gorm:"column:indate;comment:入职日期;"`  //入职日期 
      Outdate  *time.Time `json:"outdate" form:"outdate" gorm:"column:outdate;comment:离职日期;"`  //离职日期 
      Chargerate  *float64 `json:"chargerate" form:"chargerate" gorm:"column:chargerate;comment:单价;"`  //单价 
      Remark  string `json:"remark" form:"remark" gorm:"column:remark;comment:备注;size:255;"`  //备注 
      CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
      UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
      DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName hrEmployee表 HrEmployee自定义表名 hr_employee
func (HrEmployee) TableName() string {
  return "hr_employee"
}


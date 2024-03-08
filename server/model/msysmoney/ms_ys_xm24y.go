// 自动生成模板MsYsXm24y
package msysmoney

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// msYsXm24y表 结构体  MsYsXm24y
type MsYsXm24y struct {
 global.GVA_MODEL 
      Department  string `json:"Department" form:"Department" gorm:"column:Department;comment:;size:55;"`  //Department字段 
      YusuanType  string `json:"YusuanType" form:"YusuanType" gorm:"column:Yusuan_type;comment:;size:25;"`  //YusuanType字段 
      ProjectCode  string `json:"ProjectCode" form:"ProjectCode" gorm:"column:Project_code;comment:;size:25;"`  //ProjectCode字段 
      ProjectName  string `json:"ProjectName" form:"ProjectName" gorm:"column:Project_name;comment:;size:155;"`  //ProjectName字段 
      CustomGroup  string `json:"CustomGroup" form:"CustomGroup" gorm:"column:Custom_group;comment:;size:55;"`  //CustomGroup字段 
      CustomerName  string `json:"CustomerName" form:"CustomerName" gorm:"column:Customer_name;comment:;size:155;"`  //CustomerName字段 
      TotalNR  *float64 `json:"TotalNR" form:"TotalNR" gorm:"column:TotalNR;comment:;"`  //TotalNR字段 
      January  *float64 `json:"January" form:"January" gorm:"column:January;comment:;"`  //January字段 
      February  *float64 `json:"February" form:"February" gorm:"column:February;comment:;"`  //February字段 
      March  *float64 `json:"March" form:"March" gorm:"column:March;comment:;"`  //March字段 
      April  *float64 `json:"April" form:"April" gorm:"column:April;comment:;"`  //April字段 
      May  *float64 `json:"May" form:"May" gorm:"column:May;comment:;"`  //May字段 
      June  *float64 `json:"June" form:"June" gorm:"column:June;comment:;"`  //June字段 
      July  *float64 `json:"July" form:"July" gorm:"column:July;comment:;"`  //July字段 
      August  *float64 `json:"August" form:"August" gorm:"column:August;comment:;"`  //August字段 
      September  *float64 `json:"September" form:"September" gorm:"column:September;comment:;"`  //September字段 
      October  *float64 `json:"October" form:"October" gorm:"column:October;comment:;"`  //October字段 
      November  *float64 `json:"November" form:"November" gorm:"column:November;comment:;"`  //November字段 
      December  *float64 `json:"December" form:"December" gorm:"column:December;comment:;"`  //December字段 
      CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
      UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
      DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName msYsXm24y表 MsYsXm24y自定义表名 ms_ys_xm_24y
func (MsYsXm24y) TableName() string {
  return "ms_ys_xm_24y"
}


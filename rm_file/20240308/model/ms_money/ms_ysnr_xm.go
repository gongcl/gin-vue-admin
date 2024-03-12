// 自动生成模板MsYsnrXm
package ms_money

import (
	
	
	
)

// msYsnrXm表 结构体  MsYsnrXm
type MsYsnrXm struct {

      Department  string `json:"department" form:"department" gorm:"column:department;comment:;size:55;"`  //department字段 
      NrType  string `json:"nrType" form:"nrType" gorm:"column:nr_type;comment:;size:25;"`  //nrType字段 
      ProjectCode  string `json:"projectCode" form:"projectCode" gorm:"column:project_code;comment:;size:25;"`  //projectCode字段 
      ProjectName  string `json:"projectName" form:"projectName" gorm:"primarykey;column:project_name;comment:;size:55;"`  //projectName字段 
      CustomerGroup  string `json:"customerGroup" form:"customerGroup" gorm:"column:customer_group;comment:;size:25;"`  //customerGroup字段 
      CustomerName  string `json:"customerName" form:"customerName" gorm:"column:customer_name;comment:;size:55;"`  //customerName字段 
      NrTotal  *float64 `json:"nrTotal" form:"nrTotal" gorm:"column:nr_total;comment:;size:22;"`  //nrTotal字段 
      Month1  *float64 `json:"month1" form:"month1" gorm:"column:month_1;comment:;size:22;"`  //month1字段 
      Month2  *float64 `json:"month2" form:"month2" gorm:"column:month_2;comment:;size:22;"`  //month2字段 
      Month3  *float64 `json:"month3" form:"month3" gorm:"column:month_3;comment:;size:22;"`  //month3字段 
      Month4  *float64 `json:"month4" form:"month4" gorm:"column:month_4;comment:;size:22;"`  //month4字段 
      Month5  *float64 `json:"month5" form:"month5" gorm:"column:month_5;comment:;size:22;"`  //month5字段 
      Month6  *float64 `json:"month6" form:"month6" gorm:"column:month_6;comment:;size:22;"`  //month6字段 
      Month7  *float64 `json:"month7" form:"month7" gorm:"column:month_7;comment:;size:22;"`  //month7字段 
      Month8  *float64 `json:"month8" form:"month8" gorm:"column:month_8;comment:;size:22;"`  //month8字段 
      Month9  *float64 `json:"month9" form:"month9" gorm:"column:month_9;comment:;size:22;"`  //month9字段 
      Month10  *float64 `json:"month10" form:"month10" gorm:"column:month_10;comment:;size:22;"`  //month10字段 
      Month11  *float64 `json:"month11" form:"month11" gorm:"column:month_11;comment:;size:22;"`  //month11字段 
      Month12  *float64 `json:"month12" form:"month12" gorm:"column:month_12;comment:;size:22;"`  //month12字段 
}


// TableName msYsnrXm表 MsYsnrXm自定义表名 ms_ysnr_xm
func (MsYsnrXm) TableName() string {
  return "ms_ysnr_xm"
}


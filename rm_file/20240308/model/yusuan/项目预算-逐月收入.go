// 自动生成模板项目预算-逐月收入
package yusuan

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 项目预算-逐月收入表 结构体  项目预算-逐月收入
type 项目预算-逐月收入 struct {
 global.GVA_MODEL 
      交付部  string `json:"交付部" form:"交付部" gorm:"column:交付部;comment:;size:255;"`  //交付部字段 
      类型  string `json:"类型" form:"类型" gorm:"column:类型;comment:;size:255;"`  //类型字段 
      项目编码  string `json:"项目编码" form:"项目编码" gorm:"column:项目编码;comment:;size:255;"`  //项目编码字段 
      项目名称  string `json:"项目名称" form:"项目名称" gorm:"column:项目名称;comment:;size:255;"`  //项目名称字段 
      客户系  string `json:"客户系" form:"客户系" gorm:"column:客户系;comment:;size:255;"`  //客户系字段 
      客户名称  string `json:"客户名称" form:"客户名称" gorm:"column:客户名称;comment:;size:255;"`  //客户名称字段 
      预算NR  *float64 `json:"预算NR" form:"预算NR" gorm:"column:预算NR;comment:;size:22;"`  //预算NR字段 
      1月  *float64 `json:"1月" form:"1月" gorm:"column:1月;comment:;size:22;"`  //1月字段 
      2月  *float64 `json:"2月" form:"2月" gorm:"column:2月;comment:;size:22;"`  //2月字段 
      3月  *float64 `json:"3月" form:"3月" gorm:"column:3月;comment:;size:22;"`  //3月字段 
      4月  *float64 `json:"4月" form:"4月" gorm:"column:4月;comment:;size:22;"`  //4月字段 
      5月  *float64 `json:"5月" form:"5月" gorm:"column:5月;comment:;size:22;"`  //5月字段 
      6月  *float64 `json:"6月" form:"6月" gorm:"column:6月;comment:;size:22;"`  //6月字段 
      7月  *float64 `json:"7月" form:"7月" gorm:"column:7月;comment:;size:22;"`  //7月字段 
      8月  *float64 `json:"8月" form:"8月" gorm:"column:8月;comment:;size:22;"`  //8月字段 
      9月  *float64 `json:"9月" form:"9月" gorm:"column:9月;comment:;size:22;"`  //9月字段 
      10月  *float64 `json:"10月" form:"10月" gorm:"column:10月;comment:;size:22;"`  //10月字段 
      11月  *float64 `json:"11月" form:"11月" gorm:"column:11月;comment:;size:22;"`  //11月字段 
      12月  *float64 `json:"12月" form:"12月" gorm:"column:12月;comment:;size:22;"`  //12月字段 
}


// TableName 项目预算-逐月收入表 项目预算-逐月收入自定义表名 项目预算-逐月收入
func (项目预算-逐月收入) TableName() string {
  return "项目预算-逐月收入"
}


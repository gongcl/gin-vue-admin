// 自动生成模板MsHtList
package htxm

import (
	"time"
)

// msHtList表 结构体  MsHtList
type MsHtList struct {
	Dept      string     `json:"dept" form:"dept" gorm:"column:dept;comment:执行部门;size:55;"`                                //执行部门
	Kuxl      string     `json:"kuxl" form:"kuxl" gorm:"column:kuxl;comment:客户系;size:55;"`                                 //客户系
	Khmc      string     `json:"khmc" form:"khmc" gorm:"column:khmc;comment:客户名称;size:55;"`                                //客户名称
	Htbm      string     `json:"htbm" form:"htbm" gorm:"column:htbm;comment:合同编码;size:25;"`                                //合同编码
	Htmc      string     `json:"htmc" form:"htmc" gorm:"primarykey;column:htmc;comment:合同名称;size:255;" binding:"required"` //合同名称
	Htsx      string     `json:"htsx" form:"htsx" gorm:"column:htsx;comment:合同属性;size:25;"`                                //合同属性
	Jflx      string     `json:"jflx" form:"jflx" gorm:"column:jflx;comment:计费类型;size:25;"`                                //计费类型
	Jszq      string     `json:"jszq" form:"jszq" gorm:"column:jszq;comment:结算周期;size:25;"`                                //结算周期
	Jsbz      string     `json:"jsbz" form:"jsbz" gorm:"column:jsbz;comment:结算币种;size:25;"`                                //结算币种
	Xmje      *float64   `json:"xmje" form:"xmje" gorm:"column:xmje;comment:合同金额;size:22;"`                                //合同金额
	Sxrq      *time.Time `json:"sxrq" form:"sxrq" gorm:"column:sxrq;comment:生效日期;"`                                        //生效日期
	Jsrq      *time.Time `json:"jsrq" form:"jsrq" gorm:"column:jsrq;comment:结束日期;"`                                        //结束日期
	Hkzq      *int       `json:"hkzq" form:"hkzq" gorm:"column:hkzq;comment:回款周期;size:10;"`                                //回款周期
	Henh      *time.Time `json:"henh" form:"henh" gorm:"column:henh;comment:合同拿回;"`                                        //合同拿回
	Sfbl      string     `json:"sfbl" form:"sfbl" gorm:"column:sfbl;comment:是否B类;size:5;"`                                 //是否B类
	Sfel      string     `json:"sfel" form:"sfel" gorm:"column:sfel;comment:是否E类;size:5;"`                                 //是否E类
	Bzsm      string     `json:"bzsm" form:"bzsm" gorm:"column:bzsm;comment:备注说明;size:255;"`                               //是否E类
	CreatedBy uint       `gorm:"column:created_by;comment:创建者"`
	UpdatedBy uint       `gorm:"column:updated_by;comment:更新者"`
	DeletedBy uint       `gorm:"column:deleted_by;comment:删除者"`
}

// TableName msHtList表 MsHtList自定义表名 ms_ht_list
func (MsHtList) TableName() string {
	return "ms_ht_list"
}

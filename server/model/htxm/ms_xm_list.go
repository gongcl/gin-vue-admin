// 自动生成模板MsXmList
package htxm

import (
	"time"
)

// msXmList表 结构体  MsXmList
type MsXmList struct {
	Khmc      string     `json:"khmc" form:"khmc" gorm:"column:khmc;comment:客户名称;size:55;"`                               //客户名称
	Zcbm      string     `json:"zcbm" form:"zcbm" gorm:"primarykey;column:zcbm;comment:注册编码;size:25;" binding:"required"` //注册编码
	Htbm      string     `json:"htbm" form:"htbm" gorm:"column:htbm;comment:合同编码;size:25;"`                               //合同编码
	Htmc      string     `json:"htmc" form:"htmc" gorm:"column:htmc;comment:合同名称;size:255;"`                              //合同名称
	Htje      *float64   `json:"htje" form:"htje" gorm:"column:htje;comment:合同金额;size:22;"`                               //合同金额
	Xmbm      string     `json:"xmbm" form:"xmbm" gorm:"primarykey;column:xmbm;comment:项目编码;size:25;" binding:"required"` //项目编码
	Xmmc      string     `json:"xmmc" form:"xmmc" gorm:"column:xmmc;comment:项目名称;size:155;" binding:"required"`           //项目名称
	Xmje      *float64   `json:"xmje" form:"xmje" gorm:"column:xmje;comment:项目金额;size:22;"`                               //项目金额
	Kssj      *time.Time `json:"kssj" form:"kssj" gorm:"column:kssj;comment:项目开始日期;"`                                     //项目开始日期
	Jssj      *time.Time `json:"jssj" form:"jssj" gorm:"column:jssj;comment:项目结束日期;"`                                     //项目结束日期
	Xmjl      string     `json:"xmjl" form:"xmjl" gorm:"column:xmjl;comment:项目经理;size:25;"`                               //项目经理
	Zxbm      string     `json:"zxbm" form:"zxbm" gorm:"column:zxbm;comment:执行部门;size:155;"`                              //执行部门
	Xmzt      string     `json:"xmzt" form:"xmzt" gorm:"column:xmzt;comment:项目状态;size:25;"`                               //项目状态
	Xmcs      string     `json:"xmcs" form:"xmcs" gorm:"column:xmcs;comment:所在城市;size:25;"`                               //所在城市
	Site      string     `json:"site" form:"site" gorm:"column:site;comment:在岸离岸;size:25;"`                               //在岸离岸
	Jsms      string     `json:"jsms" form:"jsms" gorm:"column:jsms;comment:结算模式;size:25;"`                               //结算模式
	Srys      *float64   `json:"srys" form:"srys" gorm:"column:srys;comment:预算收入;size:22;"`                               //预算收入
	Mlys      *float64   `json:"mlys" form:"mlys" gorm:"column:mlys;comment:预算毛利;size:22;"`                               //预算毛利
	Mlyys     *float64   `json:"mlyys" form:"mlyys" gorm:"column:mlyys;comment:预算毛利率;size:22;"`                           //预算毛利率
	Gdsr      *float64   `json:"gdsr" form:"gdsr" gorm:"column:gdsr;comment:滚动收入;size:22;"`                               //滚动收入
	Gdml      *float64   `json:"gdml" form:"gdml" gorm:"column:gdml;comment:滚动毛利;size:22;"`                               //滚动毛利
	Gdmll     *float64   `json:"gdmll" form:"gdmll" gorm:"column:gdmll;comment:滚动毛利率;size:22;"`                           //滚动毛利率
	Gdrg      *float64   `json:"gdrg" form:"gdrg" gorm:"column:gdrg;comment:滚动人工;size:22;"`                               //滚动人工
	Gdqt      *float64   `json:"gdqt" form:"gdqt" gorm:"column:gdqt;comment:滚动其他成本;size:22;"`                             //滚动其他成本
	Ztbj      string     `json:"ztbj" form:"ztbj" gorm:"column:ztbj;comment:状态标记;size:255;"`                              //状态标记
	Bzsm      string     `json:"bzsm" form:"bzsm" gorm:"column:bzsm;comment:备注说明;size:255;"`                              //备注说明
	CreatedBy uint       `gorm:"column:created_by;comment:创建者"`
	UpdatedBy uint       `gorm:"column:updated_by;comment:更新者"`
	DeletedBy uint       `gorm:"column:deleted_by;comment:删除者"`
}

// TableName msXmList表 MsXmList自定义表名 ms_xm_list
func (MsXmList) TableName() string {
	return "ms_xm_list"
}

// 自动生成模板DCSSMTOutPut
package model

import (
	"gin-vue-admin/global"
      "time"
)

// 如果含有time.Time 请自行import time包
type DCSSMTOutPut struct {
      global.GVA_MODEL
      LineName  string `json:"LineName" form:"LineName" gorm:"column:LineName;comment:;type:varchar(50);size:50;"`
      MachineCode  int `json:"MachineCode" form:"MachineCode" gorm:"column:MachineCode;comment:;type:int;size:10;"`
      MachineName  string `json:"MachineName" form:"MachineName" gorm:"column:MachineName;comment:;type:varchar(50);size:50;"`
      Qty  int `json:"Qty" form:"Qty" gorm:"column:Qty;comment:;type:int;size:10;"`
      CreateTime  time.Time `json:"CreateTime" form:"CreateTime" gorm:"column:CreateTime;comment:;type:datetime;"`
      Remark  string `json:"Remark" form:"Remark" gorm:"column:Remark;comment:;type:varchar(50);size:50;"`
      Lane  int `json:"Lane" form:"Lane" gorm:"column:Lane;comment:;type:int;size:10;"`
      AddTime  time.Time `json:"AddTime" form:"AddTime" gorm:"column:AddTime;comment:;type:datetime;"`
      TableNo  string `json:"TableNo" form:"TableNo" gorm:"column:TableNo;comment:;type:varchar(50);size:50;"`
      BoardQty  int `json:"BoardQty" form:"BoardQty" gorm:"column:BoardQty;comment:;type:int;size:10;"`
}


func (DCSSMTOutPut) TableName() string {
  return "DCS_SMT_OutPut"
}


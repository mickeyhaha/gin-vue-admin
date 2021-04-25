// 自动生成模板DCSSMTConsumeAndReject
package model

import (
	"gin-vue-admin/global"
      "time"
)

// 如果含有time.Time 请自行import time包
type DCSSMTConsumeAndReject struct {
      global.GVA_MODEL
      LineName  string `json:"LineName" form:"LineName" gorm:"column:LineName;comment:;type:varchar(50);size:50;"`
      MachineCode  int `json:"MachineCode" form:"MachineCode" gorm:"column:MachineCode;comment:;type:int;size:10;"`
      MachineName  string `json:"MachineName" form:"MachineName" gorm:"column:MachineName;comment:;type:varchar(50);size:50;"`
      Lane  int `json:"Lane" form:"Lane" gorm:"column:Lane;comment:;type:int;size:10;"`
      TableNo  string `json:"TableNo" form:"TableNo" gorm:"column:TableNo;comment:;type:varchar(10);size:10;"`
      PickPos  int `json:"PickPos" form:"PickPos" gorm:"column:PickPos;comment:;type:int;size:10;"`
      SlotNo  string `json:"SlotNo" form:"SlotNo" gorm:"column:SlotNo;comment:;type:varchar(50);size:50;"`
      FeederNo  string `json:"FeederNo" form:"FeederNo" gorm:"column:FeederNo;comment:;type:varchar(50);size:50;"`
      MatrCode  string `json:"MatrCode" form:"MatrCode" gorm:"column:MatrCode;comment:;type:varchar(50);size:50;"`
      PlacedQty  int `json:"PlacedQty" form:"PlacedQty" gorm:"column:PlacedQty;comment:;type:int;size:10;"`
      PickError  int `json:"PickError" form:"PickError" gorm:"column:PickError;comment:;type:int;size:10;"`
      IdentError  int `json:"IdentError" form:"IdentError" gorm:"column:IdentError;comment:;type:int;size:10;"`
      OtherError  int `json:"OtherError" form:"OtherError" gorm:"column:OtherError;comment:;type:int;size:10;"`
      CreateTime  time.Time `json:"CreateTime" form:"CreateTime" gorm:"column:CreateTime;comment:;type:datetime;"`
      Remark  string `json:"Remark" form:"Remark" gorm:"column:Remark;comment:;type:varchar(50);size:50;"`
      AddTime  time.Time `json:"AddTime" form:"AddTime" gorm:"column:AddTime;comment:;type:datetime;"`

      // extra biz field
      ErrRate  float64 `json:"errRate" form:"errRate"`
      global.DeptFilter
}


func (DCSSMTConsumeAndReject) TableName() string {
  return "DCS_SMT_ConsumeAndReject"
}


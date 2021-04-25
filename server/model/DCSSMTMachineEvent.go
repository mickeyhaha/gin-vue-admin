// 自动生成模板DCSSMTMachineEvent
package model

import (
	"gin-vue-admin/global"
      "time"
)

// 如果含有time.Time 请自行import time包
type DCSSMTMachineEvent struct {
      global.GVA_MODEL
      LineName  string `json:"LineName" form:"LineName" gorm:"column:LineName;comment:;type:varchar(50);size:50;"`
      MachineCode  int `json:"MachineCode" form:"MachineCode" gorm:"column:MachineCode;comment:;type:int;size:10;"`
      MachineName  string `json:"MachineName" form:"MachineName" gorm:"column:MachineName;comment:;type:varchar(50);size:50;"`
      EventName  string `json:"EventName" form:"EventName" gorm:"column:EventName;comment:;type:varchar(5000);size:5000;"`
      EventRemark  string `json:"EventRemark" form:"EventRemark" gorm:"column:EventRemark;comment:;type:varchar(5000);size:5000;"`
      CreateTime  time.Time `json:"CreateTime" form:"CreateTime" gorm:"column:CreateTime;comment:;type:datetime;"`
      TableNo  string `json:"TableNo" form:"TableNo" gorm:"column:TableNo;comment:;type:varchar(50);size:50;"`
      Lane  int `json:"Lane" form:"Lane" gorm:"column:Lane;comment:;type:int;size:10;"`
      AddTime  time.Time `json:"AddTime" form:"AddTime" gorm:"column:AddTime;comment:;type:datetime;"`

      // extra biz field
      Count  int `json:"count" form:"count"`
      Total  float64 `json:"total" form:"total"`
      global.DeptFilter
}


func (DCSSMTMachineEvent) TableName() string {
  return "DCS_SMT_MachineEvent"
}


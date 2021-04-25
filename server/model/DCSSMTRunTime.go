// 自动生成模板DCSSMTRunTime
package model

import (
	"gin-vue-admin/global"
      "time"
)

// 如果含有time.Time 请自行import time包
type DCSSMTRunTime struct {
      global.GVA_MODEL
      LineName  string `json:"LineName" form:"LineName" gorm:"column:LineName;comment:;type:varchar(50);size:50;"`
      MachineCode  int `json:"MachineCode" form:"MachineCode" gorm:"column:MachineCode;comment:;type:int;size:10;"`
      MachineName  string `json:"MachineName" form:"MachineName" gorm:"column:MachineName;comment:;type:varchar(50);size:50;"`
      Lane  int `json:"Lane" form:"Lane" gorm:"column:Lane;comment:;type:int;size:10;"`
      SetupName  string `json:"SetupName" form:"SetupName" gorm:"column:SetupName;comment:;type:varchar(50);size:50;"`
      TimeCode  string `json:"TimeCode" form:"TimeCode" gorm:"column:TimeCode;comment:;type:varchar(50);size:50;"`
      TimeValue  float64 `json:"TimeValue" form:"TimeValue" gorm:"column:TimeValue;comment:;type:decimal;size:18,2;"`
      CreateTime  time.Time `json:"CreateTime" form:"CreateTime" gorm:"column:CreateTime;comment:;type:datetime;"`
      Remark  string `json:"Remark" form:"Remark" gorm:"column:Remark;comment:;type:varchar(50);size:50;"`
      AddTime  time.Time `json:"AddTime" form:"AddTime" gorm:"column:AddTime;comment:;type:datetime;"`

      // extra biz field
      Count  int `json:"count" form:"count"`
      Total  float64 `json:"total" form:"total"`
      global.DeptFilter
}


func (DCSSMTRunTime) TableName() string {
  return "DCS_SMT_RunTime"
}


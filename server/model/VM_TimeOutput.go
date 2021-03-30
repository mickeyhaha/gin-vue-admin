// 自动生成模板VM_TimeOutput
package model

import (
	"gin-vue-admin/global"
	"time"
)

// 如果含有time.Time 请自行import time包
type VM_TimeOutput struct {
      global.GVA_MODEL
      Dest  int `json:"dest" form:"dest" gorm:"column:Dest;comment:目标产量;type:int;"`
      Real  string `json:"real" form:"real" gorm:"column:Real;comment:实际产量逗号隔开;type:varchar(500);size:500;"`
      Hour  int `json:"hour" form:"hour" gorm:"column:Hour;comment:小时数;type:int;"`
      BeginTime  time.Time `json:"beginTime" form:"beginTime" gorm:"column:BeginTime;comment:开始时间;type:datetime;"`
      LineID  int `json:"lineID" form:"lineID" gorm:"column:LineID;comment:LineID;type:int;"`
      ManuDest  *bool `json:"manuDest" form:"manuDest" gorm:"column:ManuDest;comment:制令单目标数量;type:tinyint;"`
}


func (VM_TimeOutput) TableName() string {
  return "VM_TimeOutput"
}


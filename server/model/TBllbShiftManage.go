// 自动生成模板TBllbShiftManage
package model

import (
	"gin-vue-admin/global"
      "time"
)

// 如果含有time.Time 请自行import time包
type TBllbShiftManage struct {
      global.GVA_MODEL
      ShiftManageCode  string `json:"ShiftManageCode" form:"ShiftManageCode" gorm:"column:ShiftManageCode;comment:;type:varchar(50);size:50;"`
      ShiftManageName  string `json:"ShiftManageName" form:"ShiftManageName" gorm:"column:ShiftManageName;comment:;type:varchar(50);size:50;"`
      WorkShopCode  string `json:"WorkShopCode" form:"WorkShopCode" gorm:"column:WorkShopCode;comment:;type:varchar(50);size:50;"`
      StartTime  string `json:"StartTime" form:"StartTime" gorm:"column:StartTime;comment:;type:varchar(50);size:50;"`
      EndTime  string `json:"EndTime" form:"EndTime" gorm:"column:EndTime;comment:;type:varchar(50);size:50;"`
      Creator  string `json:"Creator" form:"Creator" gorm:"column:Creator;comment:;type:varchar(50);size:50;"`
      Createtime  time.Time `json:"Createtime" form:"Createtime" gorm:"column:Createtime;comment:;type:datetime;"`
      Updator  string `json:"Updator" form:"Updator" gorm:"column:Updator;comment:;type:varchar(50);size:50;"`
      Updatetime  time.Time `json:"Updatetime" form:"Updatetime" gorm:"column:Updatetime;comment:;type:datetime;"`
      TotalMente  float64 `json:"TotalMente" form:"TotalMente" gorm:"column:TotalMente;comment:;type:decimal;size:18,2;"`
}


func (TBllbShiftManage) TableName() string {
  return "T_Bllb_ShiftManage"
}


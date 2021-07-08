// 自动生成模板TRSBusiProduceRecords
package model

import (
	"gin-vue-admin/global"
      "time"
)

// 如果含有time.Time 请自行import time包
type TRSBusiProduceRecords struct {
      global.GVA_MODEL
      LineID  int `json:"LineID" form:"LineID" gorm:"column:LineID;comment:;type:int;size:10;"`
      WorkOrderNo  string `json:"WorkOrderNo" form:"WorkOrderNo" gorm:"column:WorkOrderNo;comment:;type:varchar(50);size:50;"`
      SetupName  string `json:"SetupName" form:"SetupName" gorm:"column:SetupName;comment:;type:varchar(100);size:100;"`
      Qty  int `json:"Qty" form:"Qty" gorm:"column:Qty;comment:;type:int;size:10;"`
      CreateTime  time.Time `json:"CreateTime" form:"CreateTime" gorm:"column:CreateTime;comment:;type:datetime;"`
      Remark  string `json:"Remark" form:"Remark" gorm:"column:Remark;comment:;type:varchar(100);size:100;"`
      MOrderNo  string `json:"MOrderNo" form:"MOrderNo" gorm:"column:MOrderNo;comment:;type:varchar(50);size:50;"`
}


func (TRSBusiProduceRecords) TableName() string {
  return "TRS_Busi_ProduceRecords"
}


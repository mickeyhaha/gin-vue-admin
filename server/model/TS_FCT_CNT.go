// 自动生成模板TS_FCT_CNT
package model

import (
	"gin-vue-admin/global"
      "time"
)

// 如果含有time.Time 请自行import time包
type TS_FCT_CNT struct {
      global.GVA_MODEL
      Count  int `json:"count" form:"count" gorm:"column:count;comment:Count;type:bigint;"`
      ErrCount  int `json:"errCount" form:"errCount" gorm:"column:ErrCount;comment:ErrCount;type:bigint;"`
      CreateTime  time.Time `json:"createTime" form:"createTime" gorm:"column:CreateTime;comment:CreateTime;type:datetime;"`
      LineID  int `json:"lineID" form:"lineID" gorm:"column:LineID;comment:LineID;type:int;"`
      OrderNo  string `json:"orderNo" form:"orderNo" gorm:"column:OrderNo;comment:OrderNo;type:varchar(50);size:50;"`
      IssueName  string `json:"issueName" form:"issueName" gorm:"column:IssueName;comment:IssueName;type:varchar(50);size:50;"`
      Result  int `json:"result" form:"result" gorm:"column:Result;comment:Result;type:int;"`
      FCTID  int `json:"fCTID" form:"fCTID" gorm:"column:FCTID;comment:FCTID;type:int;"`
}


func (TS_FCT_CNT) TableName() string {
  return "TS_FCT_CNT"
}


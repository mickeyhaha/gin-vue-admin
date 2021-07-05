// 自动生成模板PUBMOrderProduce2
package model

import (
	"gin-vue-admin/global"
	"time"
)

// 如果含有time.Time 请自行import time包
type PUBMOrderProduce2 struct {
      global.GVA_MODEL
      LineName  string `json:"LineName" form:"LineName" gorm:"column:LineName;comment:;type:varchar(50);size:50;"`
      MOrderNo  string `json:"MOrderNo" form:"MOrderNo" gorm:"column:MOrderNo;comment:;type:varchar(50);size:50;"`
      WorkOrderNo  string `json:"WorkOrderNo" form:"WorkOrderNo" gorm:"column:WorkOrderNo;comment:;type:varchar(50);size:50;"`
      Customer  string `json:"Customer" form:"Customer" gorm:"column:Customer;comment:;type:varchar(50);size:50;"`
      CustOrderNo  string `json:"CustOrderNo" form:"CustOrderNo" gorm:"column:CustOrderNo;comment:;type:varchar(50);size:50;"`
      MachineType  string `json:"MachineType" form:"MachineType" gorm:"column:MachineType;comment:;type:varchar(50);size:50;"`
      BOMNo  string `json:"BOMNo" form:"BOMNo" gorm:"column:BOMNo;comment:BOM编号;type:varchar(50);size:50;"`
      BOMVersion  string `json:"BOMVersion" form:"BOMVersion" gorm:"column:BOMVersion;comment:BOM版本;type:varchar(50);size:50;"`
      Product  string `json:"Product" form:"Product" gorm:"column:Product;comment:产品编号;type:varchar(50);size:50;"`
      CycleTime  int `json:"CycleTime" form:"CycleTime" gorm:"column:CycleTime;comment:周期;type:bigint;size:19;"`
      PasteSide  string `json:"PasteSide" form:"PasteSide" gorm:"column:PasteSide;comment:版面;type:varchar(50);size:50;"`
      PanelCount  int `json:"PanelCount" form:"PanelCount" gorm:"column:PanelCount;comment:拼版数;type:bigint;size:19;"`
      Qty  int `json:"Qty" form:"Qty" gorm:"column:Qty;comment:计划数量;type:bigint;size:19;"`
      QtyCompleted  int `json:"QtyCompleted" form:"QtyCompleted" gorm:"column:QtyCompleted;comment:;type:bigint;size:19;"`
      BeginTime  time.Time `json:"BeginTime" form:"BeginTime" gorm:"column:BeginTime;comment:开始时间;type:datetime;"`
      EndTime  time.Time `json:"EndTime" form:"EndTime" gorm:"column:EndTime;comment:结束时间;type:datetime;"`
      Status  int `json:"Status" form:"Status" gorm:"column:Status;comment:;type:bigint;size:19;"`
      Remark  string `json:"Remark" form:"Remark" gorm:"column:Remark;comment:备注;type:varchar(50);size:50;"`
      JITRunning  int `json:"JITRunning" form:"JITRunning" gorm:"column:JITRunning;comment:是否启动JIT拉动;type:bigint;size:19;"`
      CreateTime  time.Time `json:"CreateTime" form:"CreateTime" gorm:"column:CreateTime;comment:;type:datetime;"`
      LMTime  time.Time `json:"LMTime" form:"LMTime" gorm:"column:LMTime;comment:最后时间;type:datetime;"`
      PanelType  int `json:"PanelType" form:"PanelType" gorm:"column:PanelType;comment:PCB类型;type:bigint;size:19;"`

      // extra biz field
      LineID  int `json:"lineID" form:"lineID" gorm:"column:LineID;comment:;type:int;"`
      Count  int `json:"count" form:"count" gorm:"column:count;comment:Count;type:bigint;"`
      ErrCount  int `json:"errCount" form:"errCount" gorm:"column:ErrCount;comment:ErrCount;type:bigint;"`
      AoiErrRate  string `json:"aoiErrCount" form:"aoiErrCount"`
      TotalMente  float64 `json:"totalMente" form:"totalMente"`
      Balance  string `json:"balance" form:"balance"`
      global.DeptFilter
}


func (PUBMOrderProduce2) TableName() string {
  return "PUB_MOrderProduce"
}


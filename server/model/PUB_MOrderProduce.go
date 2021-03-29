// 自动生成模板PUB_MOrderProduce
package model

import (
	"gin-vue-admin/global"
	"time"
)

// 如果含有time.Time 请自行import time包
type PUB_MOrderProduce struct {
      global.GVA_MODEL
      LineName  string `json:"lineName" form:"lineName" gorm:"column:LineName;comment:;type:varchar(50);size:50;"`
      MOrderNo  string `json:"mOrderNo" form:"mOrderNo" gorm:"column:MOrderNo;comment:;type:varchar(50);size:50;"`
      WorkOrderNo  string `json:"workOrderNo" form:"workOrderNo" gorm:"column:WorkOrderNo;comment:;type:varchar(50);size:50;"`
      Customer  string `json:"customer" form:"customer" gorm:"column:Customer;comment:;type:varchar(50);size:50;"`
      CustOrderNo  string `json:"custOrderNo" form:"custOrderNo" gorm:"column:CustOrderNo;comment:;type:varchar(50);size:50;"`
      MachineType  string `json:"machineType" form:"machineType" gorm:"column:MachineType;comment:;type:varchar(50);size:50;"`
      BOMNo  string `json:"bOMNo" form:"bOMNo" gorm:"column:BOMNo;comment:BOM编号;type:varchar(50);size:50;"`
      BOMVersion  string `json:"bOMVersion" form:"bOMVersion" gorm:"column:BOMVersion;comment:BOM版本;type:varchar(50);size:50;"`
      Product  string `json:"product" form:"product" gorm:"column:Product;comment:产品编号;type:varchar(50);size:50;"`
      CycleTime  int `json:"cycleTime" form:"cycleTime" gorm:"column:CycleTime;comment:周期;type:int;"`
      PasteSide  string `json:"pasteSide" form:"pasteSide" gorm:"column:PasteSide;comment:版面;type:varchar(50);size:50;"`
      PanelCount  int `json:"panelCount" form:"panelCount" gorm:"column:PanelCount;comment:拼版数;type:int;"`
      Qty  int `json:"qty" form:"qty" gorm:"column:Qty;comment:计划数量;type:int;"`
      QtyCompleted  int `json:"qtyCompleted" form:"qtyCompleted" gorm:"column:QtyCompleted;comment:;type:int;"`
      BeginTime  time.Time `json:"beginTime" form:"beginTime" gorm:"column:BeginTime;comment:开始时间;type:datetime;"`
      EndTime  time.Time `json:"endTime" form:"endTime" gorm:"column:EndTime;comment:结束时间;type:datetime;"`
      Status  int `json:"status" form:"status" gorm:"column:Status;comment:;type:int;"`
      Remark  string `json:"remark" form:"remark" gorm:"column:Remark;comment:备注;type:varchar(50);size:50;"`
      JITRunning  int `json:"jITRunning" form:"jITRunning" gorm:"column:JITRunning;comment:是否启动JIT拉动;type:int;"`
      CreateTime  time.Time `json:"createTime" form:"createTime" gorm:"column:CreateTime;comment:;type:datetime;"`
      LMTime  time.Time `json:"lMTime" form:"lMTime" gorm:"column:LMTime;comment:最后时间;type:datetime;"`
      PanelType  int `json:"panelType" form:"panelType" gorm:"column:PanelType;comment:PCB类型;type:int;"`
      //SampleRate  string `json:"sampleRate" form:"sampleRate" gorm:"column:SampleRate;comment:样品单;type:varchar;"`
      //SideIndex  string `json:"sideIndex" form:"sideIndex" gorm:"column:SideIndex;comment:SideIndex;type:varchar;"`
      //Lane1  *bool `json:"lane1" form:"lane1" gorm:"column:Lane1;comment:;type:tinyint;"`
      //Lane2  *bool `json:"lane2" form:"lane2" gorm:"column:Lane2;comment:轨道二;type:tinyint;"`
}


func (PUB_MOrderProduce) TableName() string {
  return "PUB_MOrderProduce"
}


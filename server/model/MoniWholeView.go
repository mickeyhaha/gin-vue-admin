// 自动生成模板MoniWholeView
package model

import (
	"gin-vue-admin/global"
	"time"
)

// 如果含有time.Time 请自行import time包
type MoniWholeView struct {
      global.GVA_MODEL
      LineID  int `json:"lineID" form:"lineID" gorm:"column:LineID;comment:LineID;type:int;"`
      MachineCode  int `json:"machineCode" form:"machineCode" gorm:"column:MachineCode;comment:机器;type:int;"`
      TableNo  string `json:"tableNo" form:"tableNo" gorm:"column:TableNo;comment:台车;type:varchar(10);size:10;"`
      BCT  string `json:"bCT" form:"bCT" gorm:"column:BCT;comment:BCT;type:varchar(50);size:50;"`
      TrackNo  int `json:"trackNo" form:"trackNo" gorm:"column:TrackNo;comment:TrackNo;type:int;"`
      Division  int `json:"division" form:"division" gorm:"column:division;comment:Division;type:int;"`
      PickPos  int `json:"pickPos" form:"pickPos" gorm:"column:PickPos;comment:吸取位置;type:int;"`
      FeederType  string `json:"feederType" form:"feederType" gorm:"column:feeder_type;comment:FeederType;type:varchar(50);size:50;"`
      FeederNo  string `json:"feederNo" form:"feederNo" gorm:"column:FeederNo;comment:FeederNo;type:varchar(50);size:50;"`
      MatrCode  string `json:"matrCode" form:"matrCode" gorm:"column:MatrCode;comment:料号;type:varchar(50);size:50;"`
      PartStatus  int `json:"partStatus" form:"partStatus" gorm:"column:PartStatus;comment:PartStatus;type:int;"`
      CycleTime  float64 `json:"cycleTime" form:"cycleTime" gorm:"column:CycleTime;comment:CycleTime;type:float;"`
      PartPos  string `json:"partPos" form:"partPos" gorm:"column:PartPos;comment:PartPos;type:varchar(200);size:200;"`
      PlacedQty  int `json:"placedQty" form:"placedQty" gorm:"column:PlacedQty;comment:PlacedQty;type:int;"`
      PickError  int `json:"pickError" form:"pickError" gorm:"column:PickError;comment:吸取错误;type:int;"`
      IdentError  int `json:"identError" form:"identError" gorm:"column:IdentError;comment:IdentError;type:int;"`
      OtherError  int `json:"otherError" form:"otherError" gorm:"column:OtherError;comment:OtherError;type:int;"`
      TotalConsume  int `json:"totalConsume" form:"totalConsume" gorm:"column:TotalConsume;comment:TotalConsume;type:int;"`
      UniqueCode  string `json:"uniqueCode" form:"uniqueCode" gorm:"column:UniqueCode;comment:UniqueCode;type:varchar(100);size:100;"`
      LeftQty  int `json:"leftQty" form:"leftQty" gorm:"column:LeftQty;comment:LeftQty;type:int;"`
      StandbyCode  string `json:"standbyCode" form:"standbyCode" gorm:"column:StandbyCode;comment:;type:varchar(100);size:100;"`
      StandbyQty  int `json:"standbyQty" form:"standbyQty" gorm:"column:StandbyQty;comment:StandbyQty;type:int;"`
      TotalLeftQty  int `json:"totalLeftQty" form:"totalLeftQty" gorm:"column:TotalLeftQty;comment:TotalLeftQty;type:int;"`
      StopTrail  *bool `json:"stopTrail" form:"stopTrail" gorm:"column:StopTrail;comment:StopTrail;type:tinyint;"`
      TotalError  int `json:"totalError" form:"totalError" gorm:"column:TotalError;comment:TotalError;type:int;"`
      ChangeTime  time.Time `json:"changeTime" form:"changeTime" gorm:"column:ChangeTime;comment:ChangeTime;type:datetime;"`
      RejectRate  float64 `json:"rejectRate" form:"rejectRate" gorm:"column:RejectRate;comment:RejectRate;type:float;"`
      LeftTime  float64 `json:"leftTime" form:"leftTime" gorm:"column:LeftTime;comment:LeftTime;type:float;"`
      UseLevel  int `json:"useLevel" form:"useLevel" gorm:"column:UseLevel;comment:UseLevel;type:int;"`
      Scanner  string `json:"scanner" form:"scanner" gorm:"column:Scanner;comment:Scanner;type:varchar(50);size:50;"`
      ScanTime  time.Time `json:"scanTime" form:"scanTime" gorm:"column:ScanTime;comment:ScanTime;type:datetime;"`
      Remark  string `json:"remark" form:"remark" gorm:"column:Remark;comment:Remark;type:varchar(50);size:50;"`
      IsTray  *bool `json:"isTray" form:"isTray" gorm:"column:IsTray;comment:IsTray;type:tinyint;"`
}


func (MoniWholeView) TableName() string {
  return "MoniWholeView"
}


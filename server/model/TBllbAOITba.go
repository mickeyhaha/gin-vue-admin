// 自动生成模板TBllbAOITba
package model

import (
	"gin-vue-admin/global"
      "time"
)

// 如果含有time.Time 请自行import time包
type TBllbAOITba struct {
      global.GVA_MODEL
      LineID  int `json:"LineID" form:"LineID" gorm:"column:LineID;comment:;type:int;size:10;"`
      MachineType  string `json:"MachineType" form:"MachineType" gorm:"column:MachineType;comment:;type:varchar(50);size:50;"`
      OrderNo  string `json:"OrderNo" form:"OrderNo" gorm:"column:OrderNo;comment:;type:varchar(50);size:50;"`
      Program  string `json:"Program" form:"Program" gorm:"column:Program;comment:;type:varchar(50);size:50;"`
      PCBCode  string `json:"PCBCode" form:"PCBCode" gorm:"column:PCBCode;comment:;type:varchar(50);size:50;"`
      Result *bool  `json:"Result" form:"Result" gorm:"column:Result;comment:;type:bit;"`
      CreateTime  time.Time `json:"CreateTime" form:"CreateTime" gorm:"column:CreateTime;comment:;type:datetime;"`
      SysOwner  string `json:"SysOwner" form:"SysOwner" gorm:"column:SysOwner;comment:;type:varchar(50);size:50;"`
      Ext1  string `json:"Ext1" form:"Ext1" gorm:"column:Ext1;comment:;type:varchar(50);size:50;"`
      Ext2  string `json:"Ext2" form:"Ext2" gorm:"column:Ext2;comment:;type:varchar(50);size:50;"`
      Ext3  string `json:"Ext3" form:"Ext3" gorm:"column:Ext3;comment:;type:varchar(50);size:50;"`
      Ext4  string `json:"Ext4" form:"Ext4" gorm:"column:Ext4;comment:;type:varchar(50);size:50;"`
      Ext5  string `json:"Ext5" form:"Ext5" gorm:"column:Ext5;comment:;type:varchar(50);size:50;"`
      Ext6  string `json:"Ext6" form:"Ext6" gorm:"column:Ext6;comment:;type:varchar(50);size:50;"`
      Ext7  string `json:"Ext7" form:"Ext7" gorm:"column:Ext7;comment:;type:varchar(50);size:50;"`
      Ext8  string `json:"Ext8" form:"Ext8" gorm:"column:Ext8;comment:;type:varchar(50);size:50;"`
      Ext9  string `json:"Ext9" form:"Ext9" gorm:"column:Ext9;comment:;type:varchar(50);size:50;"`
      Ext10  string `json:"Ext10" form:"Ext10" gorm:"column:Ext10;comment:;type:varchar(50);size:50;"`
      InsType  int `json:"InsType" form:"InsType" gorm:"column:InsType;comment:;type:int;size:10;"`
}


func (TBllbAOITba) TableName() string {
  return "T_Bllb_AOI_tba"
}


// 自动生成模板TBllbAOIRepairTbar
package model

import (
	"gin-vue-admin/global"
      "time"
)

// 如果含有time.Time 请自行import time包
type TBllbAOIRepairTbar struct {
      global.GVA_MODEL
      AOIID  int `json:"AOIID" form:"AOIID" gorm:"column:AOIID;comment:;type:int;size:10;"`
      PCBCode  string `json:"PCBCode" form:"PCBCode" gorm:"column:PCBCode;comment:;type:varchar(50);size:50;"`
      IssueCode  string `json:"IssueCode" form:"IssueCode" gorm:"column:IssueCode;comment:;type:varchar(50);size:50;"`
      IssueName  string `json:"IssueName" form:"IssueName" gorm:"column:IssueName;comment:;type:varchar(50);size:50;"`
      Position  string `json:"Position" form:"Position" gorm:"column:Position;comment:;type:varchar(50);size:50;"`
      OtherInfo  string `json:"OtherInfo" form:"OtherInfo" gorm:"column:OtherInfo;comment:;type:varchar(50);size:50;"`
      CreateTime  time.Time `json:"CreateTime" form:"CreateTime" gorm:"column:CreateTime;comment:;type:datetime;"`
      Completed *bool  `json:"Completed" form:"Completed" gorm:"column:Completed;comment:;type:bit;"`
      RepairMan  string `json:"RepairMan" form:"RepairMan" gorm:"column:RepairMan;comment:;type:varchar(50);size:50;"`
      RepairTime  time.Time `json:"RepairTime" form:"RepairTime" gorm:"column:RepairTime;comment:;type:datetime;"`
      Remark  string `json:"Remark" form:"Remark" gorm:"column:Remark;comment:;type:varchar(50);size:50;"`
      IssueCase  string `json:"IssueCase" form:"IssueCase" gorm:"column:IssueCase;comment:;type:varchar(50);size:50;"`
      ResponsibileDept  string `json:"ResponsibileDept" form:"ResponsibileDept" gorm:"column:ResponsibileDept;comment:;type:varchar(50);size:50;"`
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
}


func (TBllbAOIRepairTbar) TableName() string {
  return "T_Bllb_AOIRepair_tbar"
}


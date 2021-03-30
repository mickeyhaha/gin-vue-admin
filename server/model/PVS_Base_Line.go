// 自动生成模板PVS_Base_Line
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type PVS_Base_Line struct {
      global.GVA_MODEL
      LineID  int `json:"lineID" form:"lineID" gorm:"column:LineID;comment:;type:int;"`
      LineName  string `json:"lineName" form:"lineName" gorm:"column:LineName;comment:LineName;type:varchar(50);size:50;"`
      LineType  string `json:"lineType" form:"lineType" gorm:"column:LineType;comment:LineType;type:varchar(50);size:50;"`
      ReadFilePath  string `json:"readFilePath" form:"readFilePath" gorm:"column:ReadFilePath;comment:程式读取路径;type:varchar(200);size:200;"`
      LineName2  string `json:"lineName2" form:"lineName2" gorm:"column:LineName2;comment:LineName2;type:varchar(100);size:100;"`
      MOrderNo  string `json:"mOrderNo" form:"mOrderNo" gorm:"column:MOrderNo;comment:制令单号;type:varchar(50);size:50;"`
      Remark  string `json:"remark" form:"remark" gorm:"column:Remark;comment:Remark;type:varchar(200);size:200;"`
      Warnning  *bool `json:"warnning" form:"warnning" gorm:"column:Warnning;comment:是否报警;type:tinyint;"`
      WarnningStr  string `json:"warnningStr" form:"warnningStr" gorm:"column:WarnningStr;comment:WarnningStr;type:varchar(50);size:50;"`
}


func (PVS_Base_Line) TableName() string {
  return "PVS_Base_Line"
}


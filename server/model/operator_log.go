// 自动生成模板OperatorLog
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type OperatorLog struct {
      global.GVA_MODEL
      Op_id  string `json:"op_id" form:"op_id" gorm:"column:op_id;comment:操作员ID;type:varchar(128);"`
      Op_name  string `json:"op_name" form:"op_name" gorm:"column:op_name;comment:操作员姓名;type:varchar(128);size:128;"`
      Action  string `json:"action" form:"action" gorm:"column:action;comment:操作;type:varchar(128);size:128;"`
      Status  int `json:"status" form:"status" gorm:"column:status;comment:status;type:int;"`
      Result  int `json:"result" form:"result" gorm:"column:result;comment:result;type:int;"`
}



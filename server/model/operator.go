// 自动生成模板Operator
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type Operator struct {
      global.GVA_MODEL
      Operator_id  string `json:"operator_id" form:"operator_id" gorm:"column:operator_id;comment:操作员id;type:varchar(128);size:128;"`
      Operator_name  string `json:"operator_name" form:"operator_name" gorm:"column:operator_name;comment:operator_name;type:varchar(128);size:128;"`
      Status  int `json:"status" form:"status" gorm:"column:status;comment:status;type:int;"`
}



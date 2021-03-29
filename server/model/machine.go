// 自动生成模板Machine
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type Machine struct {
      global.GVA_MODEL
      MachineId  string `json:"machineId" form:"machineId" gorm:"column:machine_id;comment:;type:varchar(255);size:255;"`
      MachineName  string `json:"machineName" form:"machineName" gorm:"column:machine_name;comment:;type:varchar(255);size:255;"`
}



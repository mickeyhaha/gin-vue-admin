// 自动生成模板DeviceLog
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type DeviceLog struct {
      global.GVA_MODEL
      DeviceId  string `json:"deviceId" form:"deviceId" gorm:"column:device_id;comment:deviceId;type:varchar(128);size:128;"`
      Temp  float64 `json:"temp" form:"temp" gorm:"column:temp;comment:温度;type:float;"`
      Data  string `json:"data" form:"data" gorm:"column:data;comment:上报数据;type:varchar(128);size:128;"`
}


func (DeviceLog) TableName() string {
  return "device_log"
}


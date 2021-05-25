// 自动生成模板Device
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type Device struct {
      global.GVA_MODEL
      DeviceId  string `json:"deviceId" form:"deviceId" gorm:"column:device_id;comment:;type:varchar(128);size:128;"`
      DeviceName  string `json:"deviceName" form:"deviceName" gorm:"column:device_name;comment:;type:varchar(128);size:128;"`
      DisplayId  string `json:"displayId" form:"displayId" gorm:"column:display_id;comment:;type:varchar(128);size:128;"`
      DisplayName  string `json:"displayName" form:"displayName" gorm:"column:display_name;comment:;type:varchar(128);size:128;"`
      BarCode  string `json:"barCode" form:"barCode" gorm:"column:bar_code;comment:;type:varchar(128);size:128;"`
      Longitude  float64 `json:"longitude" form:"longitude" gorm:"column:longitude;comment:;type:double;size:22;"`
      Latitude  float64 `json:"latitude" form:"latitude" gorm:"column:latitude;comment:;type:double;size:22;"`
      PlaceId  int `json:"placeId" form:"placeId" gorm:"column:placeId;comment:;type:bigint;size:19;"`
      PlaceName  int `json:"placeName" form:"placeName" gorm:"column:place_name;comment:;type:bigint;size:19;"`
      Status  int `json:"status" form:"status" gorm:"column:status;comment:;type:int;size:10;"`
      Temp  float64 `json:"temp" form:"temp" gorm:"column:temp;comment:;type:float;"`
      TempMax  float64 `json:"tempMax" form:"tempMax" gorm:"column:temp_max;comment:;type:float;"`
      TempMin  float64 `json:"tempMin" form:"tempMin" gorm:"column:temp_min;comment:;type:float;"`
      TempAvg  float64 `json:"tempAvg" form:"tempAvg" gorm:"column:temp_avg;comment:;type:float;"`
      Threshold  float64 `json:"threshold" form:"threshold" gorm:"column:threshold;comment:;type:float;"`
      OpId  string `json:"opId" form:"opId" gorm:"column:op_id;comment:;type:varchar(128);"`
      OpName  string `json:"opName" form:"opName" gorm:"column:op_name;comment:opName;type:varchar(128);size:128;"`
}


func (Device) TableName() string {
  return "device"
}


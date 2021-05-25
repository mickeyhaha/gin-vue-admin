// 自动生成模板DevicePlace
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type DevicePlace struct {
      global.GVA_MODEL
      PlaceId  string `json:"placeId" form:"placeId" gorm:"column:place_id;comment:设备ID;type:varchar(128);size:128;"`
      PlaceName  string `json:"placeName" form:"placeName" gorm:"column:place_name;comment:位置名称;type:varchar(128);size:128;"`
      Status  int `json:"status" form:"status" gorm:"column:status;comment:状态;type:int;"`
      Longitude  float64 `json:"longitude" form:"longitude" gorm:"column:longitude;comment:经度;type:float;"`
      Latitude  float64 `json:"latitude" form:"latitude" gorm:"column:latitude;comment:纬度;type:float;"`
}



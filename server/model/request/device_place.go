package request

import "gin-vue-admin/model"

type DevicePlaceSearch struct{
    model.DevicePlace
    PageInfo
}
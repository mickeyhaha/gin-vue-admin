package request

import "gin-vue-admin/model"

type DeviceSearch struct{
    model.Device
    PageInfo
}
package request

import "gin-vue-admin/model"

type DeviceLogSearch struct{
    model.DeviceLog
    PageInfo
}
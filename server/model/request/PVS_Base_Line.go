package request

import "gin-vue-admin/model"

type PVS_Base_LineSearch struct{
    model.PVS_Base_Line
    PageInfo
}
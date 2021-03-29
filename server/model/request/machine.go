package request

import "gin-vue-admin/model"

type MachineSearch struct{
    model.Machine
    PageInfo
}
package request

import "gin-vue-admin/model"

type VM_TimeOutputSearch struct{
    model.VM_TimeOutput
    PageInfo
}
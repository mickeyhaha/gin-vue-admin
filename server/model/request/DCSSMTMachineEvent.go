package request

import "gin-vue-admin/model"

type DCSSMTMachineEventSearch struct{
    model.DCSSMTMachineEvent
    PageInfo
}
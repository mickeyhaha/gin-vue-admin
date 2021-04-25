package request

import "gin-vue-admin/model"

type DCSSMTRunTimeSearch struct{
    model.DCSSMTRunTime
    PageInfo
}
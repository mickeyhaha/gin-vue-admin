package request

import "gin-vue-admin/model"

type DCSSMTOutPutSearch struct{
    model.DCSSMTOutPut
    PageInfo
}
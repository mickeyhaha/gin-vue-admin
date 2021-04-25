package request

import "gin-vue-admin/model"

type DCSSMTRejectSearch struct{
    model.DCSSMTReject
    PageInfo
}
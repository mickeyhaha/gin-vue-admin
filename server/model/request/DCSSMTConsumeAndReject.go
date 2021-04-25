package request

import "gin-vue-admin/model"

type DCSSMTConsumeAndRejectSearch struct{
    model.DCSSMTConsumeAndReject
    PageInfo
}
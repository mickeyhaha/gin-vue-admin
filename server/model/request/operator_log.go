package request

import "gin-vue-admin/model"

type OperatorLogSearch struct{
    model.OperatorLog
    PageInfo
}
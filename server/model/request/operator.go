package request

import "gin-vue-admin/model"

type OperatorSearch struct{
    model.Operator
    PageInfo
}
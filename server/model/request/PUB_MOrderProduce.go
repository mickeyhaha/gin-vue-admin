package request

import "gin-vue-admin/model"

type PUB_MOrderProduceSearch struct{
    model.PUB_MOrderProduce
    PageInfo
}
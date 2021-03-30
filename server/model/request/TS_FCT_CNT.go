package request

import "gin-vue-admin/model"

type TS_FCT_CNTSearch struct{
    model.TS_FCT_CNT
    PageInfo
}
package request

import "gin-vue-admin/model"

type TS_VI_CNTSearch struct{
    model.TS_VI_CNT
    PageInfo
}
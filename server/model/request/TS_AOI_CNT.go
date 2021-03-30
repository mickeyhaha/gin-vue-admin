package request

import "gin-vue-admin/model"

type TS_AOI_CNTSearch struct{
    model.TS_AOI_CNT
    PageInfo
}
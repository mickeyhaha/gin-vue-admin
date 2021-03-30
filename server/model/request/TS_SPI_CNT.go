package request

import "gin-vue-admin/model"

type TS_SPI_CNTSearch struct{
    model.TS_SPI_CNT
    PageInfo
}
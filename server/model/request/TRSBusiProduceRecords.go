package request

import "gin-vue-admin/model"

type TRSBusiProduceRecordsSearch struct{
    model.TRSBusiProduceRecords
    PageInfo
}
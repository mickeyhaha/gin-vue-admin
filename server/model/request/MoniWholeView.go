package request

import "gin-vue-admin/model"

type MoniWholeViewSearch struct{
    model.MoniWholeView
    PageInfo
}
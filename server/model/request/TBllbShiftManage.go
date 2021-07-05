package request

import "gin-vue-admin/model"

type TBllbShiftManageSearch struct{
    model.TBllbShiftManage
    PageInfo
}
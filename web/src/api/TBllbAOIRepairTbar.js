import service from '@/utils/request'

// @Tags TBllbAOIRepairTbar
// @Summary 创建TBllbAOIRepairTbar
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TBllbAOIRepairTbar true "创建TBllbAOIRepairTbar"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /TBART/createTBllbAOIRepairTbar [post]
export const createTBllbAOIRepairTbar = (data) => {
     return service({
         url: "/TBART/createTBllbAOIRepairTbar",
         method: 'post',
         data
     })
 }


// @Tags TBllbAOIRepairTbar
// @Summary 删除TBllbAOIRepairTbar
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TBllbAOIRepairTbar true "删除TBllbAOIRepairTbar"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /TBART/deleteTBllbAOIRepairTbar [delete]
 export const deleteTBllbAOIRepairTbar = (data) => {
     return service({
         url: "/TBART/deleteTBllbAOIRepairTbar",
         method: 'delete',
         data
     })
 }

// @Tags TBllbAOIRepairTbar
// @Summary 删除TBllbAOIRepairTbar
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TBllbAOIRepairTbar"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /TBART/deleteTBllbAOIRepairTbar [delete]
 export const deleteTBllbAOIRepairTbarByIds = (data) => {
     return service({
         url: "/TBART/deleteTBllbAOIRepairTbarByIds",
         method: 'delete',
         data
     })
 }

// @Tags TBllbAOIRepairTbar
// @Summary 更新TBllbAOIRepairTbar
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TBllbAOIRepairTbar true "更新TBllbAOIRepairTbar"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /TBART/updateTBllbAOIRepairTbar [put]
 export const updateTBllbAOIRepairTbar = (data) => {
     return service({
         url: "/TBART/updateTBllbAOIRepairTbar",
         method: 'put',
         data
     })
 }


// @Tags TBllbAOIRepairTbar
// @Summary 用id查询TBllbAOIRepairTbar
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TBllbAOIRepairTbar true "用id查询TBllbAOIRepairTbar"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /TBART/findTBllbAOIRepairTbar [get]
 export const findTBllbAOIRepairTbar = (params) => {
     return service({
         url: "/TBART/findTBllbAOIRepairTbar",
         method: 'get',
         params
     })
 }


// @Tags TBllbAOIRepairTbar
// @Summary 分页获取TBllbAOIRepairTbar列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取TBllbAOIRepairTbar列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /TBART/getTBllbAOIRepairTbarList [get]
 export const getTBllbAOIRepairTbarList = (params) => {
     return service({
         url: "/TBART/getTBllbAOIRepairTbarList",
         method: 'get',
         params
     })
 }
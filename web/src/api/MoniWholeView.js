import service from '@/utils/request'

// @Tags MoniWholeView
// @Summary 创建MoniWholeView
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MoniWholeView true "创建MoniWholeView"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /MWV/createMoniWholeView [post]
export const createMoniWholeView = (data) => {
     return service({
         url: "/MWV/createMoniWholeView",
         method: 'post',
         data
     })
 }


// @Tags MoniWholeView
// @Summary 删除MoniWholeView
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MoniWholeView true "删除MoniWholeView"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /MWV/deleteMoniWholeView [delete]
 export const deleteMoniWholeView = (data) => {
     return service({
         url: "/MWV/deleteMoniWholeView",
         method: 'delete',
         data
     })
 }

// @Tags MoniWholeView
// @Summary 删除MoniWholeView
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除MoniWholeView"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /MWV/deleteMoniWholeView [delete]
 export const deleteMoniWholeViewByIds = (data) => {
     return service({
         url: "/MWV/deleteMoniWholeViewByIds",
         method: 'delete',
         data
     })
 }

// @Tags MoniWholeView
// @Summary 更新MoniWholeView
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MoniWholeView true "更新MoniWholeView"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /MWV/updateMoniWholeView [put]
 export const updateMoniWholeView = (data) => {
     return service({
         url: "/MWV/updateMoniWholeView",
         method: 'put',
         data
     })
 }


// @Tags MoniWholeView
// @Summary 用id查询MoniWholeView
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MoniWholeView true "用id查询MoniWholeView"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /MWV/findMoniWholeView [get]
 export const findMoniWholeView = (params) => {
     return service({
         url: "/MWV/findMoniWholeView",
         method: 'get',
         params
     })
 }


// @Tags MoniWholeView
// @Summary 分页获取MoniWholeView列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取MoniWholeView列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /MWV/getMoniWholeViewList [get]
 export const getMoniWholeViewList = (params) => {
     return service({
         url: "/MWV/getMoniWholeViewList",
         method: 'get',
         params
     })
 }

export const getLackWarnings = (params) => {
    return service({
        url: "/MWV/getMoniWholeViewList",
        method: 'get',
        params
    })
}
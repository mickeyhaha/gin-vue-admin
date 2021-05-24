import service from '@/utils/request'

// @Tags OperatorLog
// @Summary 创建OperatorLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OperatorLog true "创建OperatorLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /opLog/createOperatorLog [post]
export const createOperatorLog = (data) => {
     return service({
         url: "/opLog/createOperatorLog",
         method: 'post',
         data
     })
 }


// @Tags OperatorLog
// @Summary 删除OperatorLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OperatorLog true "删除OperatorLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /opLog/deleteOperatorLog [delete]
 export const deleteOperatorLog = (data) => {
     return service({
         url: "/opLog/deleteOperatorLog",
         method: 'delete',
         data
     })
 }

// @Tags OperatorLog
// @Summary 删除OperatorLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除OperatorLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /opLog/deleteOperatorLog [delete]
 export const deleteOperatorLogByIds = (data) => {
     return service({
         url: "/opLog/deleteOperatorLogByIds",
         method: 'delete',
         data
     })
 }

// @Tags OperatorLog
// @Summary 更新OperatorLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OperatorLog true "更新OperatorLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /opLog/updateOperatorLog [put]
 export const updateOperatorLog = (data) => {
     return service({
         url: "/opLog/updateOperatorLog",
         method: 'put',
         data
     })
 }


// @Tags OperatorLog
// @Summary 用id查询OperatorLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OperatorLog true "用id查询OperatorLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /opLog/findOperatorLog [get]
 export const findOperatorLog = (params) => {
     return service({
         url: "/opLog/findOperatorLog",
         method: 'get',
         params
     })
 }


// @Tags OperatorLog
// @Summary 分页获取OperatorLog列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取OperatorLog列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /opLog/getOperatorLogList [get]
 export const getOperatorLogList = (params) => {
     return service({
         url: "/opLog/getOperatorLogList",
         method: 'get',
         params
     })
 }
import service from '@/utils/request'

// @Tags Operator
// @Summary 创建Operator
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Operator true "创建Operator"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /opt/createOperator [post]
export const createOperator = (data) => {
     return service({
         url: "/opt/createOperator",
         method: 'post',
         data
     })
 }


// @Tags Operator
// @Summary 删除Operator
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Operator true "删除Operator"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /opt/deleteOperator [delete]
 export const deleteOperator = (data) => {
     return service({
         url: "/opt/deleteOperator",
         method: 'delete',
         data
     })
 }

// @Tags Operator
// @Summary 删除Operator
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Operator"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /opt/deleteOperator [delete]
 export const deleteOperatorByIds = (data) => {
     return service({
         url: "/opt/deleteOperatorByIds",
         method: 'delete',
         data
     })
 }

// @Tags Operator
// @Summary 更新Operator
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Operator true "更新Operator"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /opt/updateOperator [put]
 export const updateOperator = (data) => {
     return service({
         url: "/opt/updateOperator",
         method: 'put',
         data
     })
 }


// @Tags Operator
// @Summary 用id查询Operator
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Operator true "用id查询Operator"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /opt/findOperator [get]
 export const findOperator = (params) => {
     return service({
         url: "/opt/findOperator",
         method: 'get',
         params
     })
 }


// @Tags Operator
// @Summary 分页获取Operator列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Operator列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /opt/getOperatorList [get]
 export const getOperatorList = (params) => {
     return service({
         url: "/opt/getOperatorList",
         method: 'get',
         params
     })
 }
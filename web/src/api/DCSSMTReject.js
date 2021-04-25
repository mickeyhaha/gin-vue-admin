import service from '@/utils/request'

// @Tags DCSSMTReject
// @Summary 创建DCSSMTReject
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DCSSMTReject true "创建DCSSMTReject"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /DSR/createDCSSMTReject [post]
export const createDCSSMTReject = (data) => {
     return service({
         url: "/DSR/createDCSSMTReject",
         method: 'post',
         data
     })
 }


// @Tags DCSSMTReject
// @Summary 删除DCSSMTReject
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DCSSMTReject true "删除DCSSMTReject"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /DSR/deleteDCSSMTReject [delete]
 export const deleteDCSSMTReject = (data) => {
     return service({
         url: "/DSR/deleteDCSSMTReject",
         method: 'delete',
         data
     })
 }

// @Tags DCSSMTReject
// @Summary 删除DCSSMTReject
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除DCSSMTReject"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /DSR/deleteDCSSMTReject [delete]
 export const deleteDCSSMTRejectByIds = (data) => {
     return service({
         url: "/DSR/deleteDCSSMTRejectByIds",
         method: 'delete',
         data
     })
 }

// @Tags DCSSMTReject
// @Summary 更新DCSSMTReject
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DCSSMTReject true "更新DCSSMTReject"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /DSR/updateDCSSMTReject [put]
 export const updateDCSSMTReject = (data) => {
     return service({
         url: "/DSR/updateDCSSMTReject",
         method: 'put',
         data
     })
 }


// @Tags DCSSMTReject
// @Summary 用id查询DCSSMTReject
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DCSSMTReject true "用id查询DCSSMTReject"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /DSR/findDCSSMTReject [get]
 export const findDCSSMTReject = (params) => {
     return service({
         url: "/DSR/findDCSSMTReject",
         method: 'get',
         params
     })
 }


// @Tags DCSSMTReject
// @Summary 分页获取DCSSMTReject列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取DCSSMTReject列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /DSR/getDCSSMTRejectList [get]
 export const getDCSSMTRejectList = (params) => {
     return service({
         url: "/DSR/getDCSSMTRejectList",
         method: 'get',
         params
     })
 }

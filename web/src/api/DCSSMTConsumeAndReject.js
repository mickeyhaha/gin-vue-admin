import service from '@/utils/request'

// @Tags DCSSMTConsumeAndReject
// @Summary 创建DCSSMTConsumeAndReject
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DCSSMTConsumeAndReject true "创建DCSSMTConsumeAndReject"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /DSCR/createDCSSMTConsumeAndReject [post]
export const createDCSSMTConsumeAndReject = (data) => {
     return service({
         url: "/DSCR/createDCSSMTConsumeAndReject",
         method: 'post',
         data
     })
 }


// @Tags DCSSMTConsumeAndReject
// @Summary 删除DCSSMTConsumeAndReject
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DCSSMTConsumeAndReject true "删除DCSSMTConsumeAndReject"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /DSCR/deleteDCSSMTConsumeAndReject [delete]
 export const deleteDCSSMTConsumeAndReject = (data) => {
     return service({
         url: "/DSCR/deleteDCSSMTConsumeAndReject",
         method: 'delete',
         data
     })
 }

// @Tags DCSSMTConsumeAndReject
// @Summary 删除DCSSMTConsumeAndReject
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除DCSSMTConsumeAndReject"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /DSCR/deleteDCSSMTConsumeAndReject [delete]
 export const deleteDCSSMTConsumeAndRejectByIds = (data) => {
     return service({
         url: "/DSCR/deleteDCSSMTConsumeAndRejectByIds",
         method: 'delete',
         data
     })
 }

// @Tags DCSSMTConsumeAndReject
// @Summary 更新DCSSMTConsumeAndReject
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DCSSMTConsumeAndReject true "更新DCSSMTConsumeAndReject"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /DSCR/updateDCSSMTConsumeAndReject [put]
 export const updateDCSSMTConsumeAndReject = (data) => {
     return service({
         url: "/DSCR/updateDCSSMTConsumeAndReject",
         method: 'put',
         data
     })
 }


// @Tags DCSSMTConsumeAndReject
// @Summary 用id查询DCSSMTConsumeAndReject
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DCSSMTConsumeAndReject true "用id查询DCSSMTConsumeAndReject"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /DSCR/findDCSSMTConsumeAndReject [get]
 export const findDCSSMTConsumeAndReject = (params) => {
     return service({
         url: "/DSCR/findDCSSMTConsumeAndReject",
         method: 'get',
         params
     })
 }


// @Tags DCSSMTConsumeAndReject
// @Summary 分页获取DCSSMTConsumeAndReject列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取DCSSMTConsumeAndReject列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /DSCR/getDCSSMTConsumeAndRejectList [get]
 export const getDCSSMTConsumeAndRejectList = (params) => {
     return service({
         url: "/DSCR/getDCSSMTConsumeAndRejectList",
         method: 'get',
         params
     })
 }
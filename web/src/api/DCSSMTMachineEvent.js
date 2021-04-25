import service from '@/utils/request'

// @Tags DCSSMTMachineEvent
// @Summary 创建DCSSMTMachineEvent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DCSSMTMachineEvent true "创建DCSSMTMachineEvent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /DSME/createDCSSMTMachineEvent [post]
export const createDCSSMTMachineEvent = (data) => {
     return service({
         url: "/DSME/createDCSSMTMachineEvent",
         method: 'post',
         data
     })
 }


// @Tags DCSSMTMachineEvent
// @Summary 删除DCSSMTMachineEvent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DCSSMTMachineEvent true "删除DCSSMTMachineEvent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /DSME/deleteDCSSMTMachineEvent [delete]
 export const deleteDCSSMTMachineEvent = (data) => {
     return service({
         url: "/DSME/deleteDCSSMTMachineEvent",
         method: 'delete',
         data
     })
 }

// @Tags DCSSMTMachineEvent
// @Summary 删除DCSSMTMachineEvent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除DCSSMTMachineEvent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /DSME/deleteDCSSMTMachineEvent [delete]
 export const deleteDCSSMTMachineEventByIds = (data) => {
     return service({
         url: "/DSME/deleteDCSSMTMachineEventByIds",
         method: 'delete',
         data
     })
 }

// @Tags DCSSMTMachineEvent
// @Summary 更新DCSSMTMachineEvent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DCSSMTMachineEvent true "更新DCSSMTMachineEvent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /DSME/updateDCSSMTMachineEvent [put]
 export const updateDCSSMTMachineEvent = (data) => {
     return service({
         url: "/DSME/updateDCSSMTMachineEvent",
         method: 'put',
         data
     })
 }


// @Tags DCSSMTMachineEvent
// @Summary 用id查询DCSSMTMachineEvent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DCSSMTMachineEvent true "用id查询DCSSMTMachineEvent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /DSME/findDCSSMTMachineEvent [get]
 export const findDCSSMTMachineEvent = (params) => {
     return service({
         url: "/DSME/findDCSSMTMachineEvent",
         method: 'get',
         params
     })
 }


// @Tags DCSSMTMachineEvent
// @Summary 分页获取DCSSMTMachineEvent列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取DCSSMTMachineEvent列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /DSME/getDCSSMTMachineEventList [get]
 export const getDCSSMTMachineEventList = (params) => {
     return service({
         url: "/DSME/getDCSSMTMachineEventList",
         method: 'get',
         params
     })
 }
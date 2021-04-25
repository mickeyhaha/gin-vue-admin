import service from '@/utils/request'

// @Tags DCSSMTRunTime
// @Summary 创建DCSSMTRunTime
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DCSSMTRunTime true "创建DCSSMTRunTime"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /DSRT/createDCSSMTRunTime [post]
export const createDCSSMTRunTime = (data) => {
     return service({
         url: "/DSRT/createDCSSMTRunTime",
         method: 'post',
         data
     })
 }


// @Tags DCSSMTRunTime
// @Summary 删除DCSSMTRunTime
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DCSSMTRunTime true "删除DCSSMTRunTime"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /DSRT/deleteDCSSMTRunTime [delete]
 export const deleteDCSSMTRunTime = (data) => {
     return service({
         url: "/DSRT/deleteDCSSMTRunTime",
         method: 'delete',
         data
     })
 }

// @Tags DCSSMTRunTime
// @Summary 删除DCSSMTRunTime
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除DCSSMTRunTime"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /DSRT/deleteDCSSMTRunTime [delete]
 export const deleteDCSSMTRunTimeByIds = (data) => {
     return service({
         url: "/DSRT/deleteDCSSMTRunTimeByIds",
         method: 'delete',
         data
     })
 }

// @Tags DCSSMTRunTime
// @Summary 更新DCSSMTRunTime
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DCSSMTRunTime true "更新DCSSMTRunTime"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /DSRT/updateDCSSMTRunTime [put]
 export const updateDCSSMTRunTime = (data) => {
     return service({
         url: "/DSRT/updateDCSSMTRunTime",
         method: 'put',
         data
     })
 }


// @Tags DCSSMTRunTime
// @Summary 用id查询DCSSMTRunTime
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DCSSMTRunTime true "用id查询DCSSMTRunTime"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /DSRT/findDCSSMTRunTime [get]
 export const findDCSSMTRunTime = (params) => {
     return service({
         url: "/DSRT/findDCSSMTRunTime",
         method: 'get',
         params
     })
 }


// @Tags DCSSMTRunTime
// @Summary 分页获取DCSSMTRunTime列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取DCSSMTRunTime列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /DSRT/getDCSSMTRunTimeList [get]
 export const getDCSSMTRunTimeList = (params) => {
     return service({
         url: "/DSRT/getDCSSMTRunTimeList",
         method: 'get',
         params
     })
 }
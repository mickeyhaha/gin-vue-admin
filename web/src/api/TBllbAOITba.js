import service from '@/utils/request'

// @Tags TBllbAOITba
// @Summary 创建TBllbAOITba
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TBllbAOITba true "创建TBllbAOITba"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /TBAT/createTBllbAOITba [post]
export const createTBllbAOITba = (data) => {
     return service({
         url: "/TBAT/createTBllbAOITba",
         method: 'post',
         data
     })
 }


// @Tags TBllbAOITba
// @Summary 删除TBllbAOITba
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TBllbAOITba true "删除TBllbAOITba"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /TBAT/deleteTBllbAOITba [delete]
 export const deleteTBllbAOITba = (data) => {
     return service({
         url: "/TBAT/deleteTBllbAOITba",
         method: 'delete',
         data
     })
 }

// @Tags TBllbAOITba
// @Summary 删除TBllbAOITba
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TBllbAOITba"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /TBAT/deleteTBllbAOITba [delete]
 export const deleteTBllbAOITbaByIds = (data) => {
     return service({
         url: "/TBAT/deleteTBllbAOITbaByIds",
         method: 'delete',
         data
     })
 }

// @Tags TBllbAOITba
// @Summary 更新TBllbAOITba
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TBllbAOITba true "更新TBllbAOITba"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /TBAT/updateTBllbAOITba [put]
 export const updateTBllbAOITba = (data) => {
     return service({
         url: "/TBAT/updateTBllbAOITba",
         method: 'put',
         data
     })
 }


// @Tags TBllbAOITba
// @Summary 用id查询TBllbAOITba
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TBllbAOITba true "用id查询TBllbAOITba"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /TBAT/findTBllbAOITba [get]
 export const findTBllbAOITba = (params) => {
     return service({
         url: "/TBAT/findTBllbAOITba",
         method: 'get',
         params
     })
 }


// @Tags TBllbAOITba
// @Summary 分页获取TBllbAOITba列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取TBllbAOITba列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /TBAT/getTBllbAOITbaList [get]
 export const getTBllbAOITbaList = (params) => {
     return service({
         url: "/TBAT/getTBllbAOITbaList",
         method: 'get',
         params
     })
 }
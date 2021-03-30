import service from '@/utils/request'

// @Tags TS_AOI_CNT
// @Summary 创建TS_AOI_CNT
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TS_AOI_CNT true "创建TS_AOI_CNT"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /TAC/createTS_AOI_CNT [post]
export const createTS_AOI_CNT = (data) => {
     return service({
         url: "/TAC/createTS_AOI_CNT",
         method: 'post',
         data
     })
 }


// @Tags TS_AOI_CNT
// @Summary 删除TS_AOI_CNT
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TS_AOI_CNT true "删除TS_AOI_CNT"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /TAC/deleteTS_AOI_CNT [delete]
 export const deleteTS_AOI_CNT = (data) => {
     return service({
         url: "/TAC/deleteTS_AOI_CNT",
         method: 'delete',
         data
     })
 }

// @Tags TS_AOI_CNT
// @Summary 删除TS_AOI_CNT
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TS_AOI_CNT"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /TAC/deleteTS_AOI_CNT [delete]
 export const deleteTS_AOI_CNTByIds = (data) => {
     return service({
         url: "/TAC/deleteTS_AOI_CNTByIds",
         method: 'delete',
         data
     })
 }

// @Tags TS_AOI_CNT
// @Summary 更新TS_AOI_CNT
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TS_AOI_CNT true "更新TS_AOI_CNT"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /TAC/updateTS_AOI_CNT [put]
 export const updateTS_AOI_CNT = (data) => {
     return service({
         url: "/TAC/updateTS_AOI_CNT",
         method: 'put',
         data
     })
 }


// @Tags TS_AOI_CNT
// @Summary 用id查询TS_AOI_CNT
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TS_AOI_CNT true "用id查询TS_AOI_CNT"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /TAC/findTS_AOI_CNT [get]
 export const findTS_AOI_CNT = (params) => {
     return service({
         url: "/TAC/findTS_AOI_CNT",
         method: 'get',
         params
     })
 }


// @Tags TS_AOI_CNT
// @Summary 分页获取TS_AOI_CNT列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取TS_AOI_CNT列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /TAC/getTS_AOI_CNTList [get]
 export const getTS_AOI_CNTList = (params) => {
     return service({
         url: "/TAC/getTS_AOI_CNTList",
         method: 'get',
         params
     })
 }
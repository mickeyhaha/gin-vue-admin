import service from '@/utils/request'

// @Tags TS_FCT_CNT
// @Summary 创建TS_FCT_CNT
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TS_FCT_CNT true "创建TS_FCT_CNT"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /TFC/createTS_FCT_CNT [post]
export const createTS_FCT_CNT = (data) => {
     return service({
         url: "/TFC/createTS_FCT_CNT",
         method: 'post',
         data
     })
 }


// @Tags TS_FCT_CNT
// @Summary 删除TS_FCT_CNT
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TS_FCT_CNT true "删除TS_FCT_CNT"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /TFC/deleteTS_FCT_CNT [delete]
 export const deleteTS_FCT_CNT = (data) => {
     return service({
         url: "/TFC/deleteTS_FCT_CNT",
         method: 'delete',
         data
     })
 }

// @Tags TS_FCT_CNT
// @Summary 删除TS_FCT_CNT
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TS_FCT_CNT"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /TFC/deleteTS_FCT_CNT [delete]
 export const deleteTS_FCT_CNTByIds = (data) => {
     return service({
         url: "/TFC/deleteTS_FCT_CNTByIds",
         method: 'delete',
         data
     })
 }

// @Tags TS_FCT_CNT
// @Summary 更新TS_FCT_CNT
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TS_FCT_CNT true "更新TS_FCT_CNT"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /TFC/updateTS_FCT_CNT [put]
 export const updateTS_FCT_CNT = (data) => {
     return service({
         url: "/TFC/updateTS_FCT_CNT",
         method: 'put',
         data
     })
 }


// @Tags TS_FCT_CNT
// @Summary 用id查询TS_FCT_CNT
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TS_FCT_CNT true "用id查询TS_FCT_CNT"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /TFC/findTS_FCT_CNT [get]
 export const findTS_FCT_CNT = (params) => {
     return service({
         url: "/TFC/findTS_FCT_CNT",
         method: 'get',
         params
     })
 }


// @Tags TS_FCT_CNT
// @Summary 分页获取TS_FCT_CNT列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取TS_FCT_CNT列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /TFC/getTS_FCT_CNTList [get]
 export const getTS_FCT_CNTList = (params) => {
     return service({
         url: "/TFC/getTS_FCT_CNTList",
         method: 'get',
         params
     })
 }
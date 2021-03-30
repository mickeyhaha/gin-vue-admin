import service from '@/utils/request'

// @Tags TS_VI_CNT
// @Summary 创建TS_VI_CNT
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TS_VI_CNT true "创建TS_VI_CNT"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /TVC/createTS_VI_CNT [post]
export const createTS_VI_CNT = (data) => {
     return service({
         url: "/TVC/createTS_VI_CNT",
         method: 'post',
         data
     })
 }


// @Tags TS_VI_CNT
// @Summary 删除TS_VI_CNT
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TS_VI_CNT true "删除TS_VI_CNT"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /TVC/deleteTS_VI_CNT [delete]
 export const deleteTS_VI_CNT = (data) => {
     return service({
         url: "/TVC/deleteTS_VI_CNT",
         method: 'delete',
         data
     })
 }

// @Tags TS_VI_CNT
// @Summary 删除TS_VI_CNT
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TS_VI_CNT"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /TVC/deleteTS_VI_CNT [delete]
 export const deleteTS_VI_CNTByIds = (data) => {
     return service({
         url: "/TVC/deleteTS_VI_CNTByIds",
         method: 'delete',
         data
     })
 }

// @Tags TS_VI_CNT
// @Summary 更新TS_VI_CNT
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TS_VI_CNT true "更新TS_VI_CNT"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /TVC/updateTS_VI_CNT [put]
 export const updateTS_VI_CNT = (data) => {
     return service({
         url: "/TVC/updateTS_VI_CNT",
         method: 'put',
         data
     })
 }


// @Tags TS_VI_CNT
// @Summary 用id查询TS_VI_CNT
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TS_VI_CNT true "用id查询TS_VI_CNT"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /TVC/findTS_VI_CNT [get]
 export const findTS_VI_CNT = (params) => {
     return service({
         url: "/TVC/findTS_VI_CNT",
         method: 'get',
         params
     })
 }


// @Tags TS_VI_CNT
// @Summary 分页获取TS_VI_CNT列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取TS_VI_CNT列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /TVC/getTS_VI_CNTList [get]
 export const getTS_VI_CNTList = (params) => {
     return service({
         url: "/TVC/getTS_VI_CNTList",
         method: 'get',
         params
     })
 }
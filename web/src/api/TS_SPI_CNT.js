import service from '@/utils/request'

// @Tags TS_SPI_CNT
// @Summary 创建TS_SPI_CNT
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TS_SPI_CNT true "创建TS_SPI_CNT"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /TSC/createTS_SPI_CNT [post]
export const createTS_SPI_CNT = (data) => {
     return service({
         url: "/TSC/createTS_SPI_CNT",
         method: 'post',
         data
     })
 }


// @Tags TS_SPI_CNT
// @Summary 删除TS_SPI_CNT
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TS_SPI_CNT true "删除TS_SPI_CNT"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /TSC/deleteTS_SPI_CNT [delete]
 export const deleteTS_SPI_CNT = (data) => {
     return service({
         url: "/TSC/deleteTS_SPI_CNT",
         method: 'delete',
         data
     })
 }

// @Tags TS_SPI_CNT
// @Summary 删除TS_SPI_CNT
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TS_SPI_CNT"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /TSC/deleteTS_SPI_CNT [delete]
 export const deleteTS_SPI_CNTByIds = (data) => {
     return service({
         url: "/TSC/deleteTS_SPI_CNTByIds",
         method: 'delete',
         data
     })
 }

// @Tags TS_SPI_CNT
// @Summary 更新TS_SPI_CNT
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TS_SPI_CNT true "更新TS_SPI_CNT"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /TSC/updateTS_SPI_CNT [put]
 export const updateTS_SPI_CNT = (data) => {
     return service({
         url: "/TSC/updateTS_SPI_CNT",
         method: 'put',
         data
     })
 }


// @Tags TS_SPI_CNT
// @Summary 用id查询TS_SPI_CNT
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TS_SPI_CNT true "用id查询TS_SPI_CNT"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /TSC/findTS_SPI_CNT [get]
 export const findTS_SPI_CNT = (params) => {
     return service({
         url: "/TSC/findTS_SPI_CNT",
         method: 'get',
         params
     })
 }


// @Tags TS_SPI_CNT
// @Summary 分页获取TS_SPI_CNT列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取TS_SPI_CNT列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /TSC/getTS_SPI_CNTList [get]
 export const getTS_SPI_CNTList = (params) => {
     return service({
         url: "/TSC/getTS_SPI_CNTList",
         method: 'get',
         params
     })
 }
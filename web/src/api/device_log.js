import service from '@/utils/request'

// @Tags DeviceLog
// @Summary 创建DeviceLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DeviceLog true "创建DeviceLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /deviceLog/createDeviceLog [post]
export const createDeviceLog = (data) => {
     return service({
         url: "/deviceLog/createDeviceLog",
         method: 'post',
         data
     })
 }


// @Tags DeviceLog
// @Summary 删除DeviceLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DeviceLog true "删除DeviceLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /deviceLog/deleteDeviceLog [delete]
 export const deleteDeviceLog = (data) => {
     return service({
         url: "/deviceLog/deleteDeviceLog",
         method: 'delete',
         data
     })
 }

// @Tags DeviceLog
// @Summary 删除DeviceLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除DeviceLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /deviceLog/deleteDeviceLog [delete]
 export const deleteDeviceLogByIds = (data) => {
     return service({
         url: "/deviceLog/deleteDeviceLogByIds",
         method: 'delete',
         data
     })
 }

// @Tags DeviceLog
// @Summary 更新DeviceLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DeviceLog true "更新DeviceLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /deviceLog/updateDeviceLog [put]
 export const updateDeviceLog = (data) => {
     return service({
         url: "/deviceLog/updateDeviceLog",
         method: 'put',
         data
     })
 }


// @Tags DeviceLog
// @Summary 用id查询DeviceLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DeviceLog true "用id查询DeviceLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /deviceLog/findDeviceLog [get]
 export const findDeviceLog = (params) => {
     return service({
         url: "/deviceLog/findDeviceLog",
         method: 'get',
         params
     })
 }


// @Tags DeviceLog
// @Summary 分页获取DeviceLog列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取DeviceLog列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /deviceLog/getDeviceLogList [get]
 export const getDeviceLogList = (params) => {
     return service({
         url: "/deviceLog/getDeviceLogList",
         method: 'get',
         params
     })
 }
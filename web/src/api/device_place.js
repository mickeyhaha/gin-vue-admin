import service from '@/utils/request'

// @Tags DevicePlace
// @Summary 创建DevicePlace
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DevicePlace true "创建DevicePlace"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /place/createDevicePlace [post]
export const createDevicePlace = (data) => {
     return service({
         url: "/place/createDevicePlace",
         method: 'post',
         data
     })
 }


// @Tags DevicePlace
// @Summary 删除DevicePlace
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DevicePlace true "删除DevicePlace"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /place/deleteDevicePlace [delete]
 export const deleteDevicePlace = (data) => {
     return service({
         url: "/place/deleteDevicePlace",
         method: 'delete',
         data
     })
 }

// @Tags DevicePlace
// @Summary 删除DevicePlace
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除DevicePlace"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /place/deleteDevicePlace [delete]
 export const deleteDevicePlaceByIds = (data) => {
     return service({
         url: "/place/deleteDevicePlaceByIds",
         method: 'delete',
         data
     })
 }

// @Tags DevicePlace
// @Summary 更新DevicePlace
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DevicePlace true "更新DevicePlace"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /place/updateDevicePlace [put]
 export const updateDevicePlace = (data) => {
     return service({
         url: "/place/updateDevicePlace",
         method: 'put',
         data
     })
 }


// @Tags DevicePlace
// @Summary 用id查询DevicePlace
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DevicePlace true "用id查询DevicePlace"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /place/findDevicePlace [get]
 export const findDevicePlace = (params) => {
     return service({
         url: "/place/findDevicePlace",
         method: 'get',
         params
     })
 }


// @Tags DevicePlace
// @Summary 分页获取DevicePlace列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取DevicePlace列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /place/getDevicePlaceList [get]
 export const getDevicePlaceList = (params) => {
     return service({
         url: "/place/getDevicePlaceList",
         method: 'get',
         params
     })
 }
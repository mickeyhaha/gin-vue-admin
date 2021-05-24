import service from '@/utils/request'

// @Tags Device
// @Summary 创建Device
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Device true "创建Device"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dev/createDevice [post]
export const createDevice = (data) => {
     return service({
         url: "/dev/createDevice",
         method: 'post',
         data
     })
 }


// @Tags Device
// @Summary 删除Device
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Device true "删除Device"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dev/deleteDevice [delete]
 export const deleteDevice = (data) => {
     return service({
         url: "/dev/deleteDevice",
         method: 'delete',
         data
     })
 }

// @Tags Device
// @Summary 删除Device
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Device"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dev/deleteDevice [delete]
 export const deleteDeviceByIds = (data) => {
     return service({
         url: "/dev/deleteDeviceByIds",
         method: 'delete',
         data
     })
 }

// @Tags Device
// @Summary 更新Device
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Device true "更新Device"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /dev/updateDevice [put]
 export const updateDevice = (data) => {
     return service({
         url: "/dev/updateDevice",
         method: 'put',
         data
     })
 }


// @Tags Device
// @Summary 用id查询Device
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Device true "用id查询Device"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /dev/findDevice [get]
 export const findDevice = (params) => {
     return service({
         url: "/dev/findDevice",
         method: 'get',
         params
     })
 }


// @Tags Device
// @Summary 分页获取Device列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Device列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dev/getDeviceList [get]
 export const getDeviceList = (params) => {
     return service({
         url: "/dev/getDeviceList",
         method: 'get',
         params
     })
 }
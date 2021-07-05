import service from '@/utils/request'

// @Tags TBllbShiftManage
// @Summary 创建TBllbShiftManage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TBllbShiftManage true "创建TBllbShiftManage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /TSM/createTBllbShiftManage [post]
export const createTBllbShiftManage = (data) => {
     return service({
         url: "/TSM/createTBllbShiftManage",
         method: 'post',
         data
     })
 }


// @Tags TBllbShiftManage
// @Summary 删除TBllbShiftManage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TBllbShiftManage true "删除TBllbShiftManage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /TSM/deleteTBllbShiftManage [delete]
 export const deleteTBllbShiftManage = (data) => {
     return service({
         url: "/TSM/deleteTBllbShiftManage",
         method: 'delete',
         data
     })
 }

// @Tags TBllbShiftManage
// @Summary 删除TBllbShiftManage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TBllbShiftManage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /TSM/deleteTBllbShiftManage [delete]
 export const deleteTBllbShiftManageByIds = (data) => {
     return service({
         url: "/TSM/deleteTBllbShiftManageByIds",
         method: 'delete',
         data
     })
 }

// @Tags TBllbShiftManage
// @Summary 更新TBllbShiftManage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TBllbShiftManage true "更新TBllbShiftManage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /TSM/updateTBllbShiftManage [put]
 export const updateTBllbShiftManage = (data) => {
     return service({
         url: "/TSM/updateTBllbShiftManage",
         method: 'put',
         data
     })
 }


// @Tags TBllbShiftManage
// @Summary 用id查询TBllbShiftManage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TBllbShiftManage true "用id查询TBllbShiftManage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /TSM/findTBllbShiftManage [get]
 export const findTBllbShiftManage = (params) => {
     return service({
         url: "/TSM/findTBllbShiftManage",
         method: 'get',
         params
     })
 }


// @Tags TBllbShiftManage
// @Summary 分页获取TBllbShiftManage列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取TBllbShiftManage列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /TSM/getTBllbShiftManageList [get]
 export const getTBllbShiftManageList = (params) => {
     return service({
         url: "/TSM/getTBllbShiftManageList",
         method: 'get',
         params
     })
 }
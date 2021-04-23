import service from '@/utils/request'

// @Tags DCSSMTOutPut
// @Summary 创建DCSSMTOutPut
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DCSSMTOutPut true "创建DCSSMTOutPut"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /DSO/createDCSSMTOutPut [post]
export const createDCSSMTOutPut = (data) => {
     return service({
         url: "/DSO/createDCSSMTOutPut",
         method: 'post',
         data
     })
 }


// @Tags DCSSMTOutPut
// @Summary 删除DCSSMTOutPut
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DCSSMTOutPut true "删除DCSSMTOutPut"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /DSO/deleteDCSSMTOutPut [delete]
 export const deleteDCSSMTOutPut = (data) => {
     return service({
         url: "/DSO/deleteDCSSMTOutPut",
         method: 'delete',
         data
     })
 }

// @Tags DCSSMTOutPut
// @Summary 删除DCSSMTOutPut
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除DCSSMTOutPut"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /DSO/deleteDCSSMTOutPut [delete]
 export const deleteDCSSMTOutPutByIds = (data) => {
     return service({
         url: "/DSO/deleteDCSSMTOutPutByIds",
         method: 'delete',
         data
     })
 }

// @Tags DCSSMTOutPut
// @Summary 更新DCSSMTOutPut
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DCSSMTOutPut true "更新DCSSMTOutPut"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /DSO/updateDCSSMTOutPut [put]
 export const updateDCSSMTOutPut = (data) => {
     return service({
         url: "/DSO/updateDCSSMTOutPut",
         method: 'put',
         data
     })
 }


// @Tags DCSSMTOutPut
// @Summary 用id查询DCSSMTOutPut
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DCSSMTOutPut true "用id查询DCSSMTOutPut"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /DSO/findDCSSMTOutPut [get]
 export const findDCSSMTOutPut = (params) => {
     return service({
         url: "/DSO/findDCSSMTOutPut",
         method: 'get',
         params
     })
 }


// @Tags DCSSMTOutPut
// @Summary 分页获取DCSSMTOutPut列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取DCSSMTOutPut列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /DSO/getDCSSMTOutPutList [get]
 export const getDCSSMTOutPutList = (params) => {
     return service({
         url: "/DSO/getDCSSMTOutPutList",
         method: 'get',
         params
     })
 }
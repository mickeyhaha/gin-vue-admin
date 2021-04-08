import service from '@/utils/request'

// @Tags PUBMOrderProduce2
// @Summary 创建PUBMOrderProduce2
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PUBMOrderProduce2 true "创建PUBMOrderProduce2"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /PUBMOrderProduce/createPUBMOrderProduce2 [post]
export const createPUBMOrderProduce2 = (data) => {
     return service({
         url: "/PUBMOrderProduce/createPUBMOrderProduce2",
         method: 'post',
         data
     })
 }


// @Tags PUBMOrderProduce2
// @Summary 删除PUBMOrderProduce2
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PUBMOrderProduce2 true "删除PUBMOrderProduce2"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /PUBMOrderProduce/deletePUBMOrderProduce2 [delete]
 export const deletePUBMOrderProduce2 = (data) => {
     return service({
         url: "/PUBMOrderProduce/deletePUBMOrderProduce2",
         method: 'delete',
         data
     })
 }

// @Tags PUBMOrderProduce2
// @Summary 删除PUBMOrderProduce2
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PUBMOrderProduce2"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /PUBMOrderProduce/deletePUBMOrderProduce2 [delete]
 export const deletePUBMOrderProduce2ByIds = (data) => {
     return service({
         url: "/PUBMOrderProduce/deletePUBMOrderProduce2ByIds",
         method: 'delete',
         data
     })
 }

// @Tags PUBMOrderProduce2
// @Summary 更新PUBMOrderProduce2
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PUBMOrderProduce2 true "更新PUBMOrderProduce2"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /PUBMOrderProduce/updatePUBMOrderProduce2 [put]
 export const updatePUBMOrderProduce2 = (data) => {
     return service({
         url: "/PUBMOrderProduce/updatePUBMOrderProduce2",
         method: 'put',
         data
     })
 }


// @Tags PUBMOrderProduce2
// @Summary 用id查询PUBMOrderProduce2
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PUBMOrderProduce2 true "用id查询PUBMOrderProduce2"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /PUBMOrderProduce/findPUBMOrderProduce2 [get]
 export const findPUBMOrderProduce2 = (params) => {
     return service({
         url: "/PUBMOrderProduce/findPUBMOrderProduce2",
         method: 'get',
         params
     })
 }


// @Tags PUBMOrderProduce2
// @Summary 分页获取PUBMOrderProduce2列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取PUBMOrderProduce2列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /PUBMOrderProduce/getPUBMOrderProduce2List [get]
 export const getPUBMOrderProduce2List = (params) => {
     return service({
         url: "/PUBMOrderProduce/getPUBMOrderProduce2List",
         method: 'get',
         params
     })
 }

export const getLineCurrOrderList = (params) => {
    return service({
        url: "/PUBMOrderProduce/getLineCurrOrderList",
        method: 'get',
        params
    })
}
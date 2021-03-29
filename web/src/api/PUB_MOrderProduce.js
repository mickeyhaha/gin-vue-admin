import service from '@/utils/request'

// @Tags PUB_MOrderProduce
// @Summary 创建PUB_MOrderProduce
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PUB_MOrderProduce true "创建PUB_MOrderProduce"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /PMOP/createPUB_MOrderProduce [post]
export const createPUB_MOrderProduce = (data) => {
     return service({
         url: "/PMOP/createPUB_MOrderProduce",
         method: 'post',
         data
     })
 }


// @Tags PUB_MOrderProduce
// @Summary 删除PUB_MOrderProduce
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PUB_MOrderProduce true "删除PUB_MOrderProduce"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /PMOP/deletePUB_MOrderProduce [delete]
 export const deletePUB_MOrderProduce = (data) => {
     return service({
         url: "/PMOP/deletePUB_MOrderProduce",
         method: 'delete',
         data
     })
 }

// @Tags PUB_MOrderProduce
// @Summary 删除PUB_MOrderProduce
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PUB_MOrderProduce"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /PMOP/deletePUB_MOrderProduce [delete]
 export const deletePUB_MOrderProduceByIds = (data) => {
     return service({
         url: "/PMOP/deletePUB_MOrderProduceByIds",
         method: 'delete',
         data
     })
 }

// @Tags PUB_MOrderProduce
// @Summary 更新PUB_MOrderProduce
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PUB_MOrderProduce true "更新PUB_MOrderProduce"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /PMOP/updatePUB_MOrderProduce [put]
 export const updatePUB_MOrderProduce = (data) => {
     return service({
         url: "/PMOP/updatePUB_MOrderProduce",
         method: 'put',
         data
     })
 }


// @Tags PUB_MOrderProduce
// @Summary 用id查询PUB_MOrderProduce
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PUB_MOrderProduce true "用id查询PUB_MOrderProduce"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /PMOP/findPUB_MOrderProduce [get]
 export const findPUB_MOrderProduce = (params) => {
     return service({
         url: "/PMOP/findPUB_MOrderProduce",
         method: 'get',
         params
     })
 }


// @Tags PUB_MOrderProduce
// @Summary 分页获取PUB_MOrderProduce列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取PUB_MOrderProduce列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /PMOP/getPUB_MOrderProduceList [get]
 export const getPUB_MOrderProduceList = (params) => {
     return service({
         url: "/PMOP/getPUB_MOrderProduceList",
         method: 'get',
         params
     })
 }
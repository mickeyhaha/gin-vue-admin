import service from '@/utils/request'

// @Tags TRSBusiProduceRecords
// @Summary 创建TRSBusiProduceRecords
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TRSBusiProduceRecords true "创建TRSBusiProduceRecords"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /TBPR/createTRSBusiProduceRecords [post]
export const createTRSBusiProduceRecords = (data) => {
     return service({
         url: "/TBPR/createTRSBusiProduceRecords",
         method: 'post',
         data
     })
 }


// @Tags TRSBusiProduceRecords
// @Summary 删除TRSBusiProduceRecords
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TRSBusiProduceRecords true "删除TRSBusiProduceRecords"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /TBPR/deleteTRSBusiProduceRecords [delete]
 export const deleteTRSBusiProduceRecords = (data) => {
     return service({
         url: "/TBPR/deleteTRSBusiProduceRecords",
         method: 'delete',
         data
     })
 }

// @Tags TRSBusiProduceRecords
// @Summary 删除TRSBusiProduceRecords
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TRSBusiProduceRecords"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /TBPR/deleteTRSBusiProduceRecords [delete]
 export const deleteTRSBusiProduceRecordsByIds = (data) => {
     return service({
         url: "/TBPR/deleteTRSBusiProduceRecordsByIds",
         method: 'delete',
         data
     })
 }

// @Tags TRSBusiProduceRecords
// @Summary 更新TRSBusiProduceRecords
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TRSBusiProduceRecords true "更新TRSBusiProduceRecords"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /TBPR/updateTRSBusiProduceRecords [put]
 export const updateTRSBusiProduceRecords = (data) => {
     return service({
         url: "/TBPR/updateTRSBusiProduceRecords",
         method: 'put',
         data
     })
 }


// @Tags TRSBusiProduceRecords
// @Summary 用id查询TRSBusiProduceRecords
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TRSBusiProduceRecords true "用id查询TRSBusiProduceRecords"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /TBPR/findTRSBusiProduceRecords [get]
 export const findTRSBusiProduceRecords = (params) => {
     return service({
         url: "/TBPR/findTRSBusiProduceRecords",
         method: 'get',
         params
     })
 }


// @Tags TRSBusiProduceRecords
// @Summary 分页获取TRSBusiProduceRecords列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取TRSBusiProduceRecords列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /TBPR/getTRSBusiProduceRecordsList [get]
 export const getTRSBusiProduceRecordsList = (params) => {
     return service({
         url: "/TBPR/getTRSBusiProduceRecordsList",
         method: 'get',
         params
     })
 }
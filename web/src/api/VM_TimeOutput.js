import service from '@/utils/request'

// @Tags VM_TimeOutput
// @Summary 创建VM_TimeOutput
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.VM_TimeOutput true "创建VM_TimeOutput"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /VTO/createVM_TimeOutput [post]
export const createVM_TimeOutput = (data) => {
     return service({
         url: "/VTO/createVM_TimeOutput",
         method: 'post',
         data
     })
 }


// @Tags VM_TimeOutput
// @Summary 删除VM_TimeOutput
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.VM_TimeOutput true "删除VM_TimeOutput"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /VTO/deleteVM_TimeOutput [delete]
 export const deleteVM_TimeOutput = (data) => {
     return service({
         url: "/VTO/deleteVM_TimeOutput",
         method: 'delete',
         data
     })
 }

// @Tags VM_TimeOutput
// @Summary 删除VM_TimeOutput
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除VM_TimeOutput"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /VTO/deleteVM_TimeOutput [delete]
 export const deleteVM_TimeOutputByIds = (data) => {
     return service({
         url: "/VTO/deleteVM_TimeOutputByIds",
         method: 'delete',
         data
     })
 }

// @Tags VM_TimeOutput
// @Summary 更新VM_TimeOutput
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.VM_TimeOutput true "更新VM_TimeOutput"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /VTO/updateVM_TimeOutput [put]
 export const updateVM_TimeOutput = (data) => {
     return service({
         url: "/VTO/updateVM_TimeOutput",
         method: 'put',
         data
     })
 }


// @Tags VM_TimeOutput
// @Summary 用id查询VM_TimeOutput
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.VM_TimeOutput true "用id查询VM_TimeOutput"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /VTO/findVM_TimeOutput [get]
 export const findVM_TimeOutput = (params) => {
     return service({
         url: "/VTO/findVM_TimeOutput",
         method: 'get',
         params
     })
 }


// @Tags VM_TimeOutput
// @Summary 分页获取VM_TimeOutput列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取VM_TimeOutput列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /VTO/getVM_TimeOutputList [get]
 export const getVM_TimeOutputList = (params) => {
     return service({
         url: "/VTO/getVM_TimeOutputList",
         method: 'get',
         params
     })
 }
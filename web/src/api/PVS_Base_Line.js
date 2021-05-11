import service from '@/utils/request'
import { deprecate } from 'util';

// @Tags PVS_Base_Line
// @Summary 创建PVS_Base_Line
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PVS_Base_Line true "创建PVS_Base_Line"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /PBL/createPVS_Base_Line [post]
export const createPVS_Base_Line = (data) => {
     return service({
         url: "/PBL/createPVS_Base_Line",
         method: 'post',
         data
     })
 }


// @Tags PVS_Base_Line
// @Summary 删除PVS_Base_Line
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PVS_Base_Line true "删除PVS_Base_Line"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /PBL/deletePVS_Base_Line [delete]
 export const deletePVS_Base_Line = (data) => {
     return service({
         url: "/PBL/deletePVS_Base_Line",
         method: 'delete',
         data
     })
 }

// @Tags PVS_Base_Line
// @Summary 删除PVS_Base_Line
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PVS_Base_Line"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /PBL/deletePVS_Base_Line [delete]
 export const deletePVS_Base_LineByIds = (data) => {
     return service({
         url: "/PBL/deletePVS_Base_LineByIds",
         method: 'delete',
         data
     })
 }

// @Tags PVS_Base_Line
// @Summary 更新PVS_Base_Line
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PVS_Base_Line true "更新PVS_Base_Line"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /PBL/updatePVS_Base_Line [put]
 export const updatePVS_Base_Line = (data) => {
     return service({
         url: "/PBL/updatePVS_Base_Line",
         method: 'put',
         data
     })
 }


// @Tags PVS_Base_Line
// @Summary 用id查询PVS_Base_Line
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PVS_Base_Line true "用id查询PVS_Base_Line"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /PBL/findPVS_Base_Line [get]
 export const findPVS_Base_Line = (params) => {
     return service({
         url: "/PBL/findPVS_Base_Line",
         method: 'get',
         params
     })
 }


// @Tags PVS_Base_Line
// @Summary 分页获取PVS_Base_Line列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取PVS_Base_Line列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /PBL/getPVS_Base_LineList [get]
 export const getPVS_Base_LineList = (params) => {
     return service({
         url: "/PBL/getPVS_Base_LineList",
         method: 'get',
         params
     })
 }

 // All smt service here
export const getDeptLineSummary = (params) => {
    return service({
        url: "/PBL/getDeptLineSummary",
        method: 'get',
        params
    })
}

export const getWorkOrderListByLine = (params) => {
    console.log(params)
    return service({
        url: "/PBL/getWorkOrderListByLine",
        method: 'get',
        params
    })
}

export const getDCSSMTOutPutList4Chart = (params) => {
    return service({
        url: "/PBL/getDCSSMTOutPutList4Chart",
        method: 'get',
        params
    })
}

export const getPUBMOrderProduce2InfoList4Chart = (params) => {
    return service({
        url: "/PBL/getPUBMOrderProduce2InfoList4Chart",
        method: 'get',
        params
    })
}

// @deprecated
export const getDCSSMTRejectRate4Chart = (params) => {
    return service({
        url: "/PBL/getDCSSMTRejectRate4Chart",
        method: 'get',
        params
    })
}

export const getDCSSMTConsumeAndRejectRate4Chart = (params) => {
    return service({
        url: "/PBL/getDCSSMTConsumeAndRejectRate4Chart",
        method: 'get',
        params
    })
}

export const getDCSSMTMachineEvent4Chart = (params) => {
    return service({
        url: "/PBL/getDCSSMTMachineEvent4Chart",
        method: 'get',
        params
    })
}

export const getDCSSMTRunTime4Chart = (params) => {
    return service({
        url: "/PBL/getDCSSMTRunTime4Chart",
        method: 'get',
        params
    })
}

export const getDCSSMTOutPutList4ChartDash = (params) => {
    return service({
        url: "/PBL/getDCSSMTOutPutList4ChartDash",
        method: 'get',
        params
    })
}

export const getPUBMOrderProduce2InfoList4ChartDash = (params) => {
    return service({
        url: "/PBL/getPUBMOrderProduce2InfoList4ChartDash",
        method: 'get',
        params
    })
}

export const getDCSSMTConsumeAndRejectRate4ChartDash = (params) => {
    return service({
        url: "/PBL/getDCSSMTConsumeAndRejectRate4ChartDash",
        method: 'get',
        params
    })
}

export const getDCSSMTMachineEvent4ChartDash = (params) => {
    return service({
        url: "/PBL/getDCSSMTMachineEvent4ChartDash",
        method: 'get',
        params
    })
}

export const getDCSSMTRunTime4ChartDash = (params) => {
    return service({
        url: "/PBL/getDCSSMTRunTime4ChartDash",
        method: 'get',
        params
    })
}
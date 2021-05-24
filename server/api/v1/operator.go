package v1

import (
	"gin-vue-admin/global"
    "gin-vue-admin/model"
    "gin-vue-admin/model/request"
    "gin-vue-admin/model/response"
    "gin-vue-admin/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

// @Tags Operator
// @Summary 创建Operator
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Operator true "创建Operator"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /opt/createOperator [post]
func CreateOperator(c *gin.Context) {
	var opt model.Operator
	_ = c.ShouldBindJSON(&opt)
	if err := service.CreateOperator(opt); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Operator
// @Summary 删除Operator
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Operator true "删除Operator"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /opt/deleteOperator [delete]
func DeleteOperator(c *gin.Context) {
	var opt model.Operator
	_ = c.ShouldBindJSON(&opt)
	if err := service.DeleteOperator(opt); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Operator
// @Summary 批量删除Operator
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Operator"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /opt/deleteOperatorByIds [delete]
func DeleteOperatorByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteOperatorByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags Operator
// @Summary 更新Operator
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Operator true "更新Operator"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /opt/updateOperator [put]
func UpdateOperator(c *gin.Context) {
	var opt model.Operator
	_ = c.ShouldBindJSON(&opt)
	if err := service.UpdateOperator(opt); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Operator
// @Summary 用id查询Operator
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Operator true "用id查询Operator"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /opt/findOperator [get]
func FindOperator(c *gin.Context) {
	var opt model.Operator
	_ = c.ShouldBindQuery(&opt)
	if err, reopt := service.GetOperator(opt.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reopt": reopt}, c)
	}
}

// @Tags Operator
// @Summary 分页获取Operator列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.OperatorSearch true "分页获取Operator列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /opt/getOperatorList [get]
func GetOperatorList(c *gin.Context) {
	var pageInfo request.OperatorSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetOperatorInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败", zap.Any("err", err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}

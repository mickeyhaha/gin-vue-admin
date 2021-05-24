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

// @Tags OperatorLog
// @Summary 创建OperatorLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OperatorLog true "创建OperatorLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /opLog/createOperatorLog [post]
func CreateOperatorLog(c *gin.Context) {
	var opLog model.OperatorLog
	_ = c.ShouldBindJSON(&opLog)
	if err := service.CreateOperatorLog(opLog); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags OperatorLog
// @Summary 删除OperatorLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OperatorLog true "删除OperatorLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /opLog/deleteOperatorLog [delete]
func DeleteOperatorLog(c *gin.Context) {
	var opLog model.OperatorLog
	_ = c.ShouldBindJSON(&opLog)
	if err := service.DeleteOperatorLog(opLog); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags OperatorLog
// @Summary 批量删除OperatorLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除OperatorLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /opLog/deleteOperatorLogByIds [delete]
func DeleteOperatorLogByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteOperatorLogByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags OperatorLog
// @Summary 更新OperatorLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OperatorLog true "更新OperatorLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /opLog/updateOperatorLog [put]
func UpdateOperatorLog(c *gin.Context) {
	var opLog model.OperatorLog
	_ = c.ShouldBindJSON(&opLog)
	if err := service.UpdateOperatorLog(opLog); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags OperatorLog
// @Summary 用id查询OperatorLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OperatorLog true "用id查询OperatorLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /opLog/findOperatorLog [get]
func FindOperatorLog(c *gin.Context) {
	var opLog model.OperatorLog
	_ = c.ShouldBindQuery(&opLog)
	if err, reopLog := service.GetOperatorLog(opLog.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reopLog": reopLog}, c)
	}
}

// @Tags OperatorLog
// @Summary 分页获取OperatorLog列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.OperatorLogSearch true "分页获取OperatorLog列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /opLog/getOperatorLogList [get]
func GetOperatorLogList(c *gin.Context) {
	var pageInfo request.OperatorLogSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetOperatorLogInfoList(pageInfo); err != nil {
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

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

// @Tags DCSSMTReject
// @Summary 创建DCSSMTReject
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DCSSMTReject true "创建DCSSMTReject"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /DSR/createDCSSMTReject [post]
func CreateDCSSMTReject(c *gin.Context) {
	var DSR model.DCSSMTReject
	_ = c.ShouldBindJSON(&DSR)
	if err := service.CreateDCSSMTReject(DSR); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags DCSSMTReject
// @Summary 删除DCSSMTReject
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DCSSMTReject true "删除DCSSMTReject"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /DSR/deleteDCSSMTReject [delete]
func DeleteDCSSMTReject(c *gin.Context) {
	var DSR model.DCSSMTReject
	_ = c.ShouldBindJSON(&DSR)
	if err := service.DeleteDCSSMTReject(DSR); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags DCSSMTReject
// @Summary 批量删除DCSSMTReject
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除DCSSMTReject"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /DSR/deleteDCSSMTRejectByIds [delete]
func DeleteDCSSMTRejectByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteDCSSMTRejectByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags DCSSMTReject
// @Summary 更新DCSSMTReject
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DCSSMTReject true "更新DCSSMTReject"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /DSR/updateDCSSMTReject [put]
func UpdateDCSSMTReject(c *gin.Context) {
	var DSR model.DCSSMTReject
	_ = c.ShouldBindJSON(&DSR)
	if err := service.UpdateDCSSMTReject(DSR); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags DCSSMTReject
// @Summary 用id查询DCSSMTReject
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DCSSMTReject true "用id查询DCSSMTReject"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /DSR/findDCSSMTReject [get]
func FindDCSSMTReject(c *gin.Context) {
	var DSR model.DCSSMTReject
	_ = c.ShouldBindQuery(&DSR)
	if err, reDSR := service.GetDCSSMTReject(DSR.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reDSR": reDSR}, c)
	}
}

// @Tags DCSSMTReject
// @Summary 分页获取DCSSMTReject列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.DCSSMTRejectSearch true "分页获取DCSSMTReject列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /DSR/getDCSSMTRejectList [get]
func GetDCSSMTRejectList(c *gin.Context) {
	var pageInfo request.DCSSMTRejectSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetDCSSMTRejectInfoList(pageInfo); err != nil {
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

func GetDCSSMTRejectRate4Chart(c *gin.Context) {
	var pageInfo request.DCSSMTRejectSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetDCSSMTRejectRate4Chart(pageInfo); err != nil {
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
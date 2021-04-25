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

// @Tags DCSSMTConsumeAndReject
// @Summary 创建DCSSMTConsumeAndReject
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DCSSMTConsumeAndReject true "创建DCSSMTConsumeAndReject"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /DSCR/createDCSSMTConsumeAndReject [post]
func CreateDCSSMTConsumeAndReject(c *gin.Context) {
	var DSCR model.DCSSMTConsumeAndReject
	_ = c.ShouldBindJSON(&DSCR)
	if err := service.CreateDCSSMTConsumeAndReject(DSCR); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags DCSSMTConsumeAndReject
// @Summary 删除DCSSMTConsumeAndReject
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DCSSMTConsumeAndReject true "删除DCSSMTConsumeAndReject"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /DSCR/deleteDCSSMTConsumeAndReject [delete]
func DeleteDCSSMTConsumeAndReject(c *gin.Context) {
	var DSCR model.DCSSMTConsumeAndReject
	_ = c.ShouldBindJSON(&DSCR)
	if err := service.DeleteDCSSMTConsumeAndReject(DSCR); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags DCSSMTConsumeAndReject
// @Summary 批量删除DCSSMTConsumeAndReject
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除DCSSMTConsumeAndReject"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /DSCR/deleteDCSSMTConsumeAndRejectByIds [delete]
func DeleteDCSSMTConsumeAndRejectByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteDCSSMTConsumeAndRejectByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags DCSSMTConsumeAndReject
// @Summary 更新DCSSMTConsumeAndReject
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DCSSMTConsumeAndReject true "更新DCSSMTConsumeAndReject"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /DSCR/updateDCSSMTConsumeAndReject [put]
func UpdateDCSSMTConsumeAndReject(c *gin.Context) {
	var DSCR model.DCSSMTConsumeAndReject
	_ = c.ShouldBindJSON(&DSCR)
	if err := service.UpdateDCSSMTConsumeAndReject(DSCR); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags DCSSMTConsumeAndReject
// @Summary 用id查询DCSSMTConsumeAndReject
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DCSSMTConsumeAndReject true "用id查询DCSSMTConsumeAndReject"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /DSCR/findDCSSMTConsumeAndReject [get]
func FindDCSSMTConsumeAndReject(c *gin.Context) {
	var DSCR model.DCSSMTConsumeAndReject
	_ = c.ShouldBindQuery(&DSCR)
	if err, reDSCR := service.GetDCSSMTConsumeAndReject(DSCR.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reDSCR": reDSCR}, c)
	}
}

// @Tags DCSSMTConsumeAndReject
// @Summary 分页获取DCSSMTConsumeAndReject列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.DCSSMTConsumeAndRejectSearch true "分页获取DCSSMTConsumeAndReject列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /DSCR/getDCSSMTConsumeAndRejectList [get]
func GetDCSSMTConsumeAndRejectList(c *gin.Context) {
	var pageInfo request.DCSSMTConsumeAndRejectSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetDCSSMTConsumeAndRejectInfoList(pageInfo); err != nil {
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

func GetDCSSMTConsumeAndRejectRate4Chart(c *gin.Context) {
	var pageInfo request.DCSSMTConsumeAndRejectSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetDCSSMTConsumeAndRejectRate4Chart(pageInfo); err != nil {
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

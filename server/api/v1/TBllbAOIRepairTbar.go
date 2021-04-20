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

// @Tags TBllbAOIRepairTbar
// @Summary 创建TBllbAOIRepairTbar
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TBllbAOIRepairTbar true "创建TBllbAOIRepairTbar"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /TBART/createTBllbAOIRepairTbar [post]
func CreateTBllbAOIRepairTbar(c *gin.Context) {
	var TBART model.TBllbAOIRepairTbar
	_ = c.ShouldBindJSON(&TBART)
	if err := service.CreateTBllbAOIRepairTbar(TBART); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags TBllbAOIRepairTbar
// @Summary 删除TBllbAOIRepairTbar
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TBllbAOIRepairTbar true "删除TBllbAOIRepairTbar"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /TBART/deleteTBllbAOIRepairTbar [delete]
func DeleteTBllbAOIRepairTbar(c *gin.Context) {
	var TBART model.TBllbAOIRepairTbar
	_ = c.ShouldBindJSON(&TBART)
	if err := service.DeleteTBllbAOIRepairTbar(TBART); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags TBllbAOIRepairTbar
// @Summary 批量删除TBllbAOIRepairTbar
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TBllbAOIRepairTbar"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /TBART/deleteTBllbAOIRepairTbarByIds [delete]
func DeleteTBllbAOIRepairTbarByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteTBllbAOIRepairTbarByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags TBllbAOIRepairTbar
// @Summary 更新TBllbAOIRepairTbar
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TBllbAOIRepairTbar true "更新TBllbAOIRepairTbar"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /TBART/updateTBllbAOIRepairTbar [put]
func UpdateTBllbAOIRepairTbar(c *gin.Context) {
	var TBART model.TBllbAOIRepairTbar
	_ = c.ShouldBindJSON(&TBART)
	if err := service.UpdateTBllbAOIRepairTbar(TBART); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags TBllbAOIRepairTbar
// @Summary 用id查询TBllbAOIRepairTbar
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TBllbAOIRepairTbar true "用id查询TBllbAOIRepairTbar"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /TBART/findTBllbAOIRepairTbar [get]
func FindTBllbAOIRepairTbar(c *gin.Context) {
	var TBART model.TBllbAOIRepairTbar
	_ = c.ShouldBindQuery(&TBART)
	if err, reTBART := service.GetTBllbAOIRepairTbar(TBART.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reTBART": reTBART}, c)
	}
}

// @Tags TBllbAOIRepairTbar
// @Summary 分页获取TBllbAOIRepairTbar列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.TBllbAOIRepairTbarSearch true "分页获取TBllbAOIRepairTbar列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /TBART/getTBllbAOIRepairTbarList [get]
func GetTBllbAOIRepairTbarList(c *gin.Context) {
	var pageInfo request.TBllbAOIRepairTbarSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetTBllbAOIRepairTbarInfoList(pageInfo); err != nil {
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

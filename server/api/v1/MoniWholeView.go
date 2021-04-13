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

// @Tags MoniWholeView
// @Summary 创建MoniWholeView
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MoniWholeView true "创建MoniWholeView"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /MWV/createMoniWholeView [post]
func CreateMoniWholeView(c *gin.Context) {
	var MWV model.MoniWholeView
	_ = c.ShouldBindJSON(&MWV)
	if err := service.CreateMoniWholeView(MWV); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags MoniWholeView
// @Summary 删除MoniWholeView
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MoniWholeView true "删除MoniWholeView"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /MWV/deleteMoniWholeView [delete]
func DeleteMoniWholeView(c *gin.Context) {
	var MWV model.MoniWholeView
	_ = c.ShouldBindJSON(&MWV)
	if err := service.DeleteMoniWholeView(MWV); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags MoniWholeView
// @Summary 批量删除MoniWholeView
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除MoniWholeView"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /MWV/deleteMoniWholeViewByIds [delete]
func DeleteMoniWholeViewByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteMoniWholeViewByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags MoniWholeView
// @Summary 更新MoniWholeView
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MoniWholeView true "更新MoniWholeView"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /MWV/updateMoniWholeView [put]
func UpdateMoniWholeView(c *gin.Context) {
	var MWV model.MoniWholeView
	_ = c.ShouldBindJSON(&MWV)
	if err := service.UpdateMoniWholeView(MWV); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags MoniWholeView
// @Summary 用id查询MoniWholeView
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MoniWholeView true "用id查询MoniWholeView"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /MWV/findMoniWholeView [get]
func FindMoniWholeView(c *gin.Context) {
	var MWV model.MoniWholeView
	_ = c.ShouldBindQuery(&MWV)
	if err, reMWV := service.GetMoniWholeView(MWV.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reMWV": reMWV}, c)
	}
}

// @Tags MoniWholeView
// @Summary 分页获取MoniWholeView列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.MoniWholeViewSearch true "分页获取MoniWholeView列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /MWV/getMoniWholeViewList [get]
func GetMoniWholeViewList(c *gin.Context) {
	var pageInfo request.MoniWholeViewSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetMoniWholeViewInfoList(pageInfo); err != nil {
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

func GetRejectRateList(c *gin.Context) {
	var pageInfo request.MoniWholeViewSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetRejectRateList(pageInfo); err != nil {
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
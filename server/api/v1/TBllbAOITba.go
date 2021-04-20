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

// @Tags TBllbAOITba
// @Summary 创建TBllbAOITba
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TBllbAOITba true "创建TBllbAOITba"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /TBAT/createTBllbAOITba [post]
func CreateTBllbAOITba(c *gin.Context) {
	var TBAT model.TBllbAOITba
	_ = c.ShouldBindJSON(&TBAT)
	if err := service.CreateTBllbAOITba(TBAT); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags TBllbAOITba
// @Summary 删除TBllbAOITba
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TBllbAOITba true "删除TBllbAOITba"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /TBAT/deleteTBllbAOITba [delete]
func DeleteTBllbAOITba(c *gin.Context) {
	var TBAT model.TBllbAOITba
	_ = c.ShouldBindJSON(&TBAT)
	if err := service.DeleteTBllbAOITba(TBAT); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags TBllbAOITba
// @Summary 批量删除TBllbAOITba
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TBllbAOITba"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /TBAT/deleteTBllbAOITbaByIds [delete]
func DeleteTBllbAOITbaByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteTBllbAOITbaByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags TBllbAOITba
// @Summary 更新TBllbAOITba
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TBllbAOITba true "更新TBllbAOITba"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /TBAT/updateTBllbAOITba [put]
func UpdateTBllbAOITba(c *gin.Context) {
	var TBAT model.TBllbAOITba
	_ = c.ShouldBindJSON(&TBAT)
	if err := service.UpdateTBllbAOITba(TBAT); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags TBllbAOITba
// @Summary 用id查询TBllbAOITba
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TBllbAOITba true "用id查询TBllbAOITba"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /TBAT/findTBllbAOITba [get]
func FindTBllbAOITba(c *gin.Context) {
	var TBAT model.TBllbAOITba
	_ = c.ShouldBindQuery(&TBAT)
	if err, reTBAT := service.GetTBllbAOITba(TBAT.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reTBAT": reTBAT}, c)
	}
}

// @Tags TBllbAOITba
// @Summary 分页获取TBllbAOITba列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.TBllbAOITbaSearch true "分页获取TBllbAOITba列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /TBAT/getTBllbAOITbaList [get]
func GetTBllbAOITbaList(c *gin.Context) {
	var pageInfo request.TBllbAOITbaSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetTBllbAOITbaInfoList(pageInfo); err != nil {
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

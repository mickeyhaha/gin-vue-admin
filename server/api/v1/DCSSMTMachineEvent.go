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

// @Tags DCSSMTMachineEvent
// @Summary 创建DCSSMTMachineEvent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DCSSMTMachineEvent true "创建DCSSMTMachineEvent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /DSME/createDCSSMTMachineEvent [post]
func CreateDCSSMTMachineEvent(c *gin.Context) {
	var DSME model.DCSSMTMachineEvent
	_ = c.ShouldBindJSON(&DSME)
	if err := service.CreateDCSSMTMachineEvent(DSME); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags DCSSMTMachineEvent
// @Summary 删除DCSSMTMachineEvent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DCSSMTMachineEvent true "删除DCSSMTMachineEvent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /DSME/deleteDCSSMTMachineEvent [delete]
func DeleteDCSSMTMachineEvent(c *gin.Context) {
	var DSME model.DCSSMTMachineEvent
	_ = c.ShouldBindJSON(&DSME)
	if err := service.DeleteDCSSMTMachineEvent(DSME); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags DCSSMTMachineEvent
// @Summary 批量删除DCSSMTMachineEvent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除DCSSMTMachineEvent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /DSME/deleteDCSSMTMachineEventByIds [delete]
func DeleteDCSSMTMachineEventByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteDCSSMTMachineEventByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags DCSSMTMachineEvent
// @Summary 更新DCSSMTMachineEvent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DCSSMTMachineEvent true "更新DCSSMTMachineEvent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /DSME/updateDCSSMTMachineEvent [put]
func UpdateDCSSMTMachineEvent(c *gin.Context) {
	var DSME model.DCSSMTMachineEvent
	_ = c.ShouldBindJSON(&DSME)
	if err := service.UpdateDCSSMTMachineEvent(DSME); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags DCSSMTMachineEvent
// @Summary 用id查询DCSSMTMachineEvent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DCSSMTMachineEvent true "用id查询DCSSMTMachineEvent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /DSME/findDCSSMTMachineEvent [get]
func FindDCSSMTMachineEvent(c *gin.Context) {
	var DSME model.DCSSMTMachineEvent
	_ = c.ShouldBindQuery(&DSME)
	if err, reDSME := service.GetDCSSMTMachineEvent(DSME.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reDSME": reDSME}, c)
	}
}

// @Tags DCSSMTMachineEvent
// @Summary 分页获取DCSSMTMachineEvent列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.DCSSMTMachineEventSearch true "分页获取DCSSMTMachineEvent列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /DSME/getDCSSMTMachineEventList [get]
func GetDCSSMTMachineEventList(c *gin.Context) {
	var pageInfo request.DCSSMTMachineEventSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetDCSSMTMachineEventInfoList(pageInfo); err != nil {
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

func GetDCSSMTMachineEvent4Chart(c *gin.Context) {
	var pageInfo request.DCSSMTMachineEventSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetDCSSMTMachineEvent4Chart(pageInfo); err != nil {
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

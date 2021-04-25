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

// @Tags DCSSMTRunTime
// @Summary 创建DCSSMTRunTime
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DCSSMTRunTime true "创建DCSSMTRunTime"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /DSRT/createDCSSMTRunTime [post]
func CreateDCSSMTRunTime(c *gin.Context) {
	var DSRT model.DCSSMTRunTime
	_ = c.ShouldBindJSON(&DSRT)
	if err := service.CreateDCSSMTRunTime(DSRT); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags DCSSMTRunTime
// @Summary 删除DCSSMTRunTime
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DCSSMTRunTime true "删除DCSSMTRunTime"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /DSRT/deleteDCSSMTRunTime [delete]
func DeleteDCSSMTRunTime(c *gin.Context) {
	var DSRT model.DCSSMTRunTime
	_ = c.ShouldBindJSON(&DSRT)
	if err := service.DeleteDCSSMTRunTime(DSRT); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags DCSSMTRunTime
// @Summary 批量删除DCSSMTRunTime
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除DCSSMTRunTime"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /DSRT/deleteDCSSMTRunTimeByIds [delete]
func DeleteDCSSMTRunTimeByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteDCSSMTRunTimeByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags DCSSMTRunTime
// @Summary 更新DCSSMTRunTime
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DCSSMTRunTime true "更新DCSSMTRunTime"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /DSRT/updateDCSSMTRunTime [put]
func UpdateDCSSMTRunTime(c *gin.Context) {
	var DSRT model.DCSSMTRunTime
	_ = c.ShouldBindJSON(&DSRT)
	if err := service.UpdateDCSSMTRunTime(DSRT); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags DCSSMTRunTime
// @Summary 用id查询DCSSMTRunTime
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DCSSMTRunTime true "用id查询DCSSMTRunTime"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /DSRT/findDCSSMTRunTime [get]
func FindDCSSMTRunTime(c *gin.Context) {
	var DSRT model.DCSSMTRunTime
	_ = c.ShouldBindQuery(&DSRT)
	if err, reDSRT := service.GetDCSSMTRunTime(DSRT.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reDSRT": reDSRT}, c)
	}
}

// @Tags DCSSMTRunTime
// @Summary 分页获取DCSSMTRunTime列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.DCSSMTRunTimeSearch true "分页获取DCSSMTRunTime列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /DSRT/getDCSSMTRunTimeList [get]
func GetDCSSMTRunTimeList(c *gin.Context) {
	var pageInfo request.DCSSMTRunTimeSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetDCSSMTRunTimeInfoList(pageInfo); err != nil {
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

func GetDCSSMTRunTime4Chart(c *gin.Context) {
	var pageInfo request.DCSSMTRunTimeSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetDCSSMTRunTime4Chart(pageInfo); err != nil {
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
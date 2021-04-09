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

// @Tags TS_AOI_CNT
// @Summary 创建TS_AOI_CNT
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TS_AOI_CNT true "创建TS_AOI_CNT"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /TAC/createTS_AOI_CNT [post]
func CreateTS_AOI_CNT(c *gin.Context) {
	var TAC model.TS_AOI_CNT
	_ = c.ShouldBindJSON(&TAC)
	if err := service.CreateTS_AOI_CNT(TAC); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags TS_AOI_CNT
// @Summary 删除TS_AOI_CNT
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TS_AOI_CNT true "删除TS_AOI_CNT"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /TAC/deleteTS_AOI_CNT [delete]
func DeleteTS_AOI_CNT(c *gin.Context) {
	var TAC model.TS_AOI_CNT
	_ = c.ShouldBindJSON(&TAC)
	if err := service.DeleteTS_AOI_CNT(TAC); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags TS_AOI_CNT
// @Summary 批量删除TS_AOI_CNT
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TS_AOI_CNT"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /TAC/deleteTS_AOI_CNTByIds [delete]
func DeleteTS_AOI_CNTByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteTS_AOI_CNTByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags TS_AOI_CNT
// @Summary 更新TS_AOI_CNT
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TS_AOI_CNT true "更新TS_AOI_CNT"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /TAC/updateTS_AOI_CNT [put]
func UpdateTS_AOI_CNT(c *gin.Context) {
	var TAC model.TS_AOI_CNT
	_ = c.ShouldBindJSON(&TAC)
	if err := service.UpdateTS_AOI_CNT(TAC); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags TS_AOI_CNT
// @Summary 用id查询TS_AOI_CNT
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TS_AOI_CNT true "用id查询TS_AOI_CNT"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /TAC/findTS_AOI_CNT [get]
func FindTS_AOI_CNT(c *gin.Context) {
	var TAC model.TS_AOI_CNT
	_ = c.ShouldBindQuery(&TAC)
	if err, reTAC := service.GetTS_AOI_CNT(TAC.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reTAC": reTAC}, c)
	}
}

// @Tags TS_AOI_CNT
// @Summary 分页获取TS_AOI_CNT列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.TS_AOI_CNTSearch true "分页获取TS_AOI_CNT列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /TAC/getTS_AOI_CNTList [get]
func GetTS_AOI_CNTList(c *gin.Context) {
	var pageInfo request.TS_AOI_CNTSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetTS_AOI_CNTInfoList(pageInfo); err != nil {
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

func GetTS_AOI_CNTList4Chart(c *gin.Context) {
	var pageInfo request.TS_AOI_CNTSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetTS_AOI_CNTInfoList4Chart(pageInfo); err != nil {
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

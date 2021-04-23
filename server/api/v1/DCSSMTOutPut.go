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

// @Tags DCSSMTOutPut
// @Summary 创建DCSSMTOutPut
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DCSSMTOutPut true "创建DCSSMTOutPut"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /DSO/createDCSSMTOutPut [post]
func CreateDCSSMTOutPut(c *gin.Context) {
	var DSO model.DCSSMTOutPut
	_ = c.ShouldBindJSON(&DSO)
	if err := service.CreateDCSSMTOutPut(DSO); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags DCSSMTOutPut
// @Summary 删除DCSSMTOutPut
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DCSSMTOutPut true "删除DCSSMTOutPut"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /DSO/deleteDCSSMTOutPut [delete]
func DeleteDCSSMTOutPut(c *gin.Context) {
	var DSO model.DCSSMTOutPut
	_ = c.ShouldBindJSON(&DSO)
	if err := service.DeleteDCSSMTOutPut(DSO); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags DCSSMTOutPut
// @Summary 批量删除DCSSMTOutPut
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除DCSSMTOutPut"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /DSO/deleteDCSSMTOutPutByIds [delete]
func DeleteDCSSMTOutPutByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteDCSSMTOutPutByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags DCSSMTOutPut
// @Summary 更新DCSSMTOutPut
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DCSSMTOutPut true "更新DCSSMTOutPut"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /DSO/updateDCSSMTOutPut [put]
func UpdateDCSSMTOutPut(c *gin.Context) {
	var DSO model.DCSSMTOutPut
	_ = c.ShouldBindJSON(&DSO)
	if err := service.UpdateDCSSMTOutPut(DSO); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags DCSSMTOutPut
// @Summary 用id查询DCSSMTOutPut
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DCSSMTOutPut true "用id查询DCSSMTOutPut"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /DSO/findDCSSMTOutPut [get]
func FindDCSSMTOutPut(c *gin.Context) {
	var DSO model.DCSSMTOutPut
	_ = c.ShouldBindQuery(&DSO)
	if err, reDSO := service.GetDCSSMTOutPut(DSO.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reDSO": reDSO}, c)
	}
}

// @Tags DCSSMTOutPut
// @Summary 分页获取DCSSMTOutPut列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.DCSSMTOutPutSearch true "分页获取DCSSMTOutPut列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /DSO/getDCSSMTOutPutList [get]
func GetDCSSMTOutPutList(c *gin.Context) {
	var pageInfo request.DCSSMTOutPutSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetDCSSMTOutPutInfoList(pageInfo); err != nil {
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

func GetDCSSMTOutPutList4Chart(c *gin.Context) {
	var pageInfo request.DCSSMTOutPutSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetDCSSMTOutPutInfoList4Chart(pageInfo); err != nil {
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

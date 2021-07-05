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

// @Tags TBllbShiftManage
// @Summary 创建TBllbShiftManage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TBllbShiftManage true "创建TBllbShiftManage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /TSM/createTBllbShiftManage [post]
func CreateTBllbShiftManage(c *gin.Context) {
	var TSM model.TBllbShiftManage
	_ = c.ShouldBindJSON(&TSM)
	if err := service.CreateTBllbShiftManage(TSM); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags TBllbShiftManage
// @Summary 删除TBllbShiftManage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TBllbShiftManage true "删除TBllbShiftManage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /TSM/deleteTBllbShiftManage [delete]
func DeleteTBllbShiftManage(c *gin.Context) {
	var TSM model.TBllbShiftManage
	_ = c.ShouldBindJSON(&TSM)
	if err := service.DeleteTBllbShiftManage(TSM); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags TBllbShiftManage
// @Summary 批量删除TBllbShiftManage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TBllbShiftManage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /TSM/deleteTBllbShiftManageByIds [delete]
func DeleteTBllbShiftManageByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteTBllbShiftManageByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags TBllbShiftManage
// @Summary 更新TBllbShiftManage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TBllbShiftManage true "更新TBllbShiftManage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /TSM/updateTBllbShiftManage [put]
func UpdateTBllbShiftManage(c *gin.Context) {
	var TSM model.TBllbShiftManage
	_ = c.ShouldBindJSON(&TSM)
	if err := service.UpdateTBllbShiftManage(TSM); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags TBllbShiftManage
// @Summary 用id查询TBllbShiftManage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TBllbShiftManage true "用id查询TBllbShiftManage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /TSM/findTBllbShiftManage [get]
func FindTBllbShiftManage(c *gin.Context) {
	var TSM model.TBllbShiftManage
	_ = c.ShouldBindQuery(&TSM)
	if err, reTSM := service.GetTBllbShiftManage(TSM.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reTSM": reTSM}, c)
	}
}

// @Tags TBllbShiftManage
// @Summary 分页获取TBllbShiftManage列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.TBllbShiftManageSearch true "分页获取TBllbShiftManage列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /TSM/getTBllbShiftManageList [get]
func GetTBllbShiftManageList(c *gin.Context) {
	var pageInfo request.TBllbShiftManageSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetTBllbShiftManageInfoList(pageInfo); err != nil {
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

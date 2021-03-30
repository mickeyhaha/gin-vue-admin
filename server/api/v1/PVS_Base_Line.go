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

// @Tags PVS_Base_Line
// @Summary 创建PVS_Base_Line
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PVS_Base_Line true "创建PVS_Base_Line"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /PBL/createPVS_Base_Line [post]
func CreatePVS_Base_Line(c *gin.Context) {
	var PBL model.PVS_Base_Line
	_ = c.ShouldBindJSON(&PBL)
	if err := service.CreatePVS_Base_Line(PBL); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags PVS_Base_Line
// @Summary 删除PVS_Base_Line
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PVS_Base_Line true "删除PVS_Base_Line"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /PBL/deletePVS_Base_Line [delete]
func DeletePVS_Base_Line(c *gin.Context) {
	var PBL model.PVS_Base_Line
	_ = c.ShouldBindJSON(&PBL)
	if err := service.DeletePVS_Base_Line(PBL); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags PVS_Base_Line
// @Summary 批量删除PVS_Base_Line
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PVS_Base_Line"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /PBL/deletePVS_Base_LineByIds [delete]
func DeletePVS_Base_LineByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeletePVS_Base_LineByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags PVS_Base_Line
// @Summary 更新PVS_Base_Line
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PVS_Base_Line true "更新PVS_Base_Line"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /PBL/updatePVS_Base_Line [put]
func UpdatePVS_Base_Line(c *gin.Context) {
	var PBL model.PVS_Base_Line
	_ = c.ShouldBindJSON(&PBL)
	if err := service.UpdatePVS_Base_Line(PBL); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags PVS_Base_Line
// @Summary 用id查询PVS_Base_Line
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PVS_Base_Line true "用id查询PVS_Base_Line"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /PBL/findPVS_Base_Line [get]
func FindPVS_Base_Line(c *gin.Context) {
	var PBL model.PVS_Base_Line
	_ = c.ShouldBindQuery(&PBL)
	if err, rePBL := service.GetPVS_Base_Line(PBL.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rePBL": rePBL}, c)
	}
}

// @Tags PVS_Base_Line
// @Summary 分页获取PVS_Base_Line列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PVS_Base_LineSearch true "分页获取PVS_Base_Line列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /PBL/getPVS_Base_LineList [get]
func GetPVS_Base_LineList(c *gin.Context) {
	var pageInfo request.PVS_Base_LineSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetPVS_Base_LineInfoList(pageInfo); err != nil {
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

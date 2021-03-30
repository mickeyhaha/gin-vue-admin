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

// @Tags VM_TimeOutput
// @Summary 创建VM_TimeOutput
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.VM_TimeOutput true "创建VM_TimeOutput"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /VTO/createVM_TimeOutput [post]
func CreateVM_TimeOutput(c *gin.Context) {
	var VTO model.VM_TimeOutput
	_ = c.ShouldBindJSON(&VTO)
	if err := service.CreateVM_TimeOutput(VTO); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags VM_TimeOutput
// @Summary 删除VM_TimeOutput
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.VM_TimeOutput true "删除VM_TimeOutput"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /VTO/deleteVM_TimeOutput [delete]
func DeleteVM_TimeOutput(c *gin.Context) {
	var VTO model.VM_TimeOutput
	_ = c.ShouldBindJSON(&VTO)
	if err := service.DeleteVM_TimeOutput(VTO); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags VM_TimeOutput
// @Summary 批量删除VM_TimeOutput
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除VM_TimeOutput"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /VTO/deleteVM_TimeOutputByIds [delete]
func DeleteVM_TimeOutputByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteVM_TimeOutputByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags VM_TimeOutput
// @Summary 更新VM_TimeOutput
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.VM_TimeOutput true "更新VM_TimeOutput"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /VTO/updateVM_TimeOutput [put]
func UpdateVM_TimeOutput(c *gin.Context) {
	var VTO model.VM_TimeOutput
	_ = c.ShouldBindJSON(&VTO)
	if err := service.UpdateVM_TimeOutput(VTO); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags VM_TimeOutput
// @Summary 用id查询VM_TimeOutput
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.VM_TimeOutput true "用id查询VM_TimeOutput"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /VTO/findVM_TimeOutput [get]
func FindVM_TimeOutput(c *gin.Context) {
	var VTO model.VM_TimeOutput
	_ = c.ShouldBindQuery(&VTO)
	if err, reVTO := service.GetVM_TimeOutput(VTO.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reVTO": reVTO}, c)
	}
}

// @Tags VM_TimeOutput
// @Summary 分页获取VM_TimeOutput列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.VM_TimeOutputSearch true "分页获取VM_TimeOutput列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /VTO/getVM_TimeOutputList [get]
func GetVM_TimeOutputList(c *gin.Context) {
	var pageInfo request.VM_TimeOutputSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetVM_TimeOutputInfoList(pageInfo); err != nil {
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

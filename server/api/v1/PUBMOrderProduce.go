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

// @Tags PUBMOrderProduce2
// @Summary 创建PUBMOrderProduce2
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PUBMOrderProduce2 true "创建PUBMOrderProduce2"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /PUBMOrderProduce/createPUBMOrderProduce2 [post]
func CreatePUBMOrderProduce2(c *gin.Context) {
	var PUBMOrderProduce model.PUBMOrderProduce2
	_ = c.ShouldBindJSON(&PUBMOrderProduce)
	if err := service.CreatePUBMOrderProduce2(PUBMOrderProduce); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags PUBMOrderProduce2
// @Summary 删除PUBMOrderProduce2
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PUBMOrderProduce2 true "删除PUBMOrderProduce2"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /PUBMOrderProduce/deletePUBMOrderProduce2 [delete]
func DeletePUBMOrderProduce2(c *gin.Context) {
	var PUBMOrderProduce model.PUBMOrderProduce2
	_ = c.ShouldBindJSON(&PUBMOrderProduce)
	if err := service.DeletePUBMOrderProduce2(PUBMOrderProduce); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags PUBMOrderProduce2
// @Summary 批量删除PUBMOrderProduce2
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PUBMOrderProduce2"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /PUBMOrderProduce/deletePUBMOrderProduce2ByIds [delete]
func DeletePUBMOrderProduce2ByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeletePUBMOrderProduce2ByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags PUBMOrderProduce2
// @Summary 更新PUBMOrderProduce2
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PUBMOrderProduce2 true "更新PUBMOrderProduce2"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /PUBMOrderProduce/updatePUBMOrderProduce2 [put]
func UpdatePUBMOrderProduce2(c *gin.Context) {
	var PUBMOrderProduce model.PUBMOrderProduce2
	_ = c.ShouldBindJSON(&PUBMOrderProduce)
	if err := service.UpdatePUBMOrderProduce2(PUBMOrderProduce); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags PUBMOrderProduce2
// @Summary 用id查询PUBMOrderProduce2
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PUBMOrderProduce2 true "用id查询PUBMOrderProduce2"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /PUBMOrderProduce/findPUBMOrderProduce2 [get]
func FindPUBMOrderProduce2(c *gin.Context) {
	var PUBMOrderProduce model.PUBMOrderProduce2
	_ = c.ShouldBindQuery(&PUBMOrderProduce)
	if err, rePUBMOrderProduce := service.GetPUBMOrderProduce2(PUBMOrderProduce.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rePUBMOrderProduce": rePUBMOrderProduce}, c)
	}
}

// @Tags PUBMOrderProduce2
// @Summary 分页获取PUBMOrderProduce2列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PUBMOrderProduce2Search true "分页获取PUBMOrderProduce2列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /PUBMOrderProduce/getPUBMOrderProduce2List [get]
func GetPUBMOrderProduce2List(c *gin.Context) {
	var pageInfo request.PUBMOrderProduce2Search
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetPUBMOrderProduce2InfoList(pageInfo); err != nil {
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

func GetLineCurrOrderList(c *gin.Context) {
	orderInfo := request.PUBMOrderProduce2Search{}
	orderInfo.Status = 2
	if err, orders, total := service.GetPUBMOrderProduce2InfoList(orderInfo); err == nil {
		response.OkWithDetailed(response.PageResult{
			List:     orders,
			Total:    total,
			Page:     orderInfo.Page,
			PageSize: orderInfo.PageSize,
		}, "获取成功", c)
	} else {
		global.GVA_LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	}
}

func GetPUBMOrderProduce2InfoList4Chart(c *gin.Context) {
	var pageInfo request.PUBMOrderProduce2Search
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetPUBMOrderProduce2InfoList4Chart(pageInfo); err != nil {
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


func GetPUBMOrderProduce2InfoList4ChartDash(c *gin.Context) {
	var pageInfo request.PUBMOrderProduce2Search
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetPUBMOrderProduce2InfoList4ChartDash(pageInfo); err != nil {
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
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

// @Tags PUB_MOrderProduce
// @Summary 创建PUB_MOrderProduce
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PUB_MOrderProduce true "创建PUB_MOrderProduce"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /PMOP/createPUB_MOrderProduce [post]
func CreatePUB_MOrderProduce(c *gin.Context) {
	var PMOP model.PUB_MOrderProduce
	_ = c.ShouldBindJSON(&PMOP)
	if err := service.CreatePUB_MOrderProduce(PMOP); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags PUB_MOrderProduce
// @Summary 删除PUB_MOrderProduce
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PUB_MOrderProduce true "删除PUB_MOrderProduce"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /PMOP/deletePUB_MOrderProduce [delete]
func DeletePUB_MOrderProduce(c *gin.Context) {
	var PMOP model.PUB_MOrderProduce
	_ = c.ShouldBindJSON(&PMOP)
	if err := service.DeletePUB_MOrderProduce(PMOP); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags PUB_MOrderProduce
// @Summary 批量删除PUB_MOrderProduce
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PUB_MOrderProduce"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /PMOP/deletePUB_MOrderProduceByIds [delete]
func DeletePUB_MOrderProduceByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeletePUB_MOrderProduceByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags PUB_MOrderProduce
// @Summary 更新PUB_MOrderProduce
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PUB_MOrderProduce true "更新PUB_MOrderProduce"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /PMOP/updatePUB_MOrderProduce [put]
func UpdatePUB_MOrderProduce(c *gin.Context) {
	var PMOP model.PUB_MOrderProduce
	_ = c.ShouldBindJSON(&PMOP)
	if err := service.UpdatePUB_MOrderProduce(PMOP); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags PUB_MOrderProduce
// @Summary 用id查询PUB_MOrderProduce
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PUB_MOrderProduce true "用id查询PUB_MOrderProduce"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /PMOP/findPUB_MOrderProduce [get]
func FindPUB_MOrderProduce(c *gin.Context) {
	var PMOP model.PUB_MOrderProduce
	_ = c.ShouldBindQuery(&PMOP)
	if err, rePMOP := service.GetPUB_MOrderProduce(PMOP.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rePMOP": rePMOP}, c)
	}
}

// @Tags PUB_MOrderProduce
// @Summary 分页获取PUB_MOrderProduce列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PUB_MOrderProduceSearch true "分页获取PUB_MOrderProduce列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /PMOP/getPUB_MOrderProduceList [get]
func GetPUB_MOrderProduceList(c *gin.Context) {
	var pageInfo request.PUB_MOrderProduceSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetPUB_MOrderProduceInfoList(pageInfo); err != nil {
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

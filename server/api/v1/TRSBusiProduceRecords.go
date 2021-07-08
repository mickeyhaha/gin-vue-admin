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

// @Tags TRSBusiProduceRecords
// @Summary 创建TRSBusiProduceRecords
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TRSBusiProduceRecords true "创建TRSBusiProduceRecords"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /TBPR/createTRSBusiProduceRecords [post]
func CreateTRSBusiProduceRecords(c *gin.Context) {
	var TBPR model.TRSBusiProduceRecords
	_ = c.ShouldBindJSON(&TBPR)
	if err := service.CreateTRSBusiProduceRecords(TBPR); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags TRSBusiProduceRecords
// @Summary 删除TRSBusiProduceRecords
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TRSBusiProduceRecords true "删除TRSBusiProduceRecords"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /TBPR/deleteTRSBusiProduceRecords [delete]
func DeleteTRSBusiProduceRecords(c *gin.Context) {
	var TBPR model.TRSBusiProduceRecords
	_ = c.ShouldBindJSON(&TBPR)
	if err := service.DeleteTRSBusiProduceRecords(TBPR); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags TRSBusiProduceRecords
// @Summary 批量删除TRSBusiProduceRecords
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TRSBusiProduceRecords"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /TBPR/deleteTRSBusiProduceRecordsByIds [delete]
func DeleteTRSBusiProduceRecordsByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteTRSBusiProduceRecordsByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags TRSBusiProduceRecords
// @Summary 更新TRSBusiProduceRecords
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TRSBusiProduceRecords true "更新TRSBusiProduceRecords"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /TBPR/updateTRSBusiProduceRecords [put]
func UpdateTRSBusiProduceRecords(c *gin.Context) {
	var TBPR model.TRSBusiProduceRecords
	_ = c.ShouldBindJSON(&TBPR)
	if err := service.UpdateTRSBusiProduceRecords(TBPR); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags TRSBusiProduceRecords
// @Summary 用id查询TRSBusiProduceRecords
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TRSBusiProduceRecords true "用id查询TRSBusiProduceRecords"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /TBPR/findTRSBusiProduceRecords [get]
func FindTRSBusiProduceRecords(c *gin.Context) {
	var TBPR model.TRSBusiProduceRecords
	_ = c.ShouldBindQuery(&TBPR)
	if err, reTBPR := service.GetTRSBusiProduceRecords(TBPR.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reTBPR": reTBPR}, c)
	}
}

// @Tags TRSBusiProduceRecords
// @Summary 分页获取TRSBusiProduceRecords列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.TRSBusiProduceRecordsSearch true "分页获取TRSBusiProduceRecords列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /TBPR/getTRSBusiProduceRecordsList [get]
func GetTRSBusiProduceRecordsList(c *gin.Context) {
	var pageInfo request.TRSBusiProduceRecordsSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetTRSBusiProduceRecordsInfoList(pageInfo); err != nil {
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

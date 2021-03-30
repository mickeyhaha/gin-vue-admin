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

// @Tags TS_SPI_CNT
// @Summary 创建TS_SPI_CNT
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TS_SPI_CNT true "创建TS_SPI_CNT"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /TSC/createTS_SPI_CNT [post]
func CreateTS_SPI_CNT(c *gin.Context) {
	var TSC model.TS_SPI_CNT
	_ = c.ShouldBindJSON(&TSC)
	if err := service.CreateTS_SPI_CNT(TSC); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags TS_SPI_CNT
// @Summary 删除TS_SPI_CNT
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TS_SPI_CNT true "删除TS_SPI_CNT"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /TSC/deleteTS_SPI_CNT [delete]
func DeleteTS_SPI_CNT(c *gin.Context) {
	var TSC model.TS_SPI_CNT
	_ = c.ShouldBindJSON(&TSC)
	if err := service.DeleteTS_SPI_CNT(TSC); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags TS_SPI_CNT
// @Summary 批量删除TS_SPI_CNT
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TS_SPI_CNT"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /TSC/deleteTS_SPI_CNTByIds [delete]
func DeleteTS_SPI_CNTByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteTS_SPI_CNTByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags TS_SPI_CNT
// @Summary 更新TS_SPI_CNT
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TS_SPI_CNT true "更新TS_SPI_CNT"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /TSC/updateTS_SPI_CNT [put]
func UpdateTS_SPI_CNT(c *gin.Context) {
	var TSC model.TS_SPI_CNT
	_ = c.ShouldBindJSON(&TSC)
	if err := service.UpdateTS_SPI_CNT(TSC); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags TS_SPI_CNT
// @Summary 用id查询TS_SPI_CNT
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TS_SPI_CNT true "用id查询TS_SPI_CNT"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /TSC/findTS_SPI_CNT [get]
func FindTS_SPI_CNT(c *gin.Context) {
	var TSC model.TS_SPI_CNT
	_ = c.ShouldBindQuery(&TSC)
	if err, reTSC := service.GetTS_SPI_CNT(TSC.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reTSC": reTSC}, c)
	}
}

// @Tags TS_SPI_CNT
// @Summary 分页获取TS_SPI_CNT列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.TS_SPI_CNTSearch true "分页获取TS_SPI_CNT列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /TSC/getTS_SPI_CNTList [get]
func GetTS_SPI_CNTList(c *gin.Context) {
	var pageInfo request.TS_SPI_CNTSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetTS_SPI_CNTInfoList(pageInfo); err != nil {
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

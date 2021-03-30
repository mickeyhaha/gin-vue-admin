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

// @Tags TS_FCT_CNT
// @Summary 创建TS_FCT_CNT
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TS_FCT_CNT true "创建TS_FCT_CNT"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /TFC/createTS_FCT_CNT [post]
func CreateTS_FCT_CNT(c *gin.Context) {
	var TFC model.TS_FCT_CNT
	_ = c.ShouldBindJSON(&TFC)
	if err := service.CreateTS_FCT_CNT(TFC); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags TS_FCT_CNT
// @Summary 删除TS_FCT_CNT
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TS_FCT_CNT true "删除TS_FCT_CNT"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /TFC/deleteTS_FCT_CNT [delete]
func DeleteTS_FCT_CNT(c *gin.Context) {
	var TFC model.TS_FCT_CNT
	_ = c.ShouldBindJSON(&TFC)
	if err := service.DeleteTS_FCT_CNT(TFC); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags TS_FCT_CNT
// @Summary 批量删除TS_FCT_CNT
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TS_FCT_CNT"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /TFC/deleteTS_FCT_CNTByIds [delete]
func DeleteTS_FCT_CNTByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteTS_FCT_CNTByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags TS_FCT_CNT
// @Summary 更新TS_FCT_CNT
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TS_FCT_CNT true "更新TS_FCT_CNT"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /TFC/updateTS_FCT_CNT [put]
func UpdateTS_FCT_CNT(c *gin.Context) {
	var TFC model.TS_FCT_CNT
	_ = c.ShouldBindJSON(&TFC)
	if err := service.UpdateTS_FCT_CNT(TFC); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags TS_FCT_CNT
// @Summary 用id查询TS_FCT_CNT
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TS_FCT_CNT true "用id查询TS_FCT_CNT"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /TFC/findTS_FCT_CNT [get]
func FindTS_FCT_CNT(c *gin.Context) {
	var TFC model.TS_FCT_CNT
	_ = c.ShouldBindQuery(&TFC)
	if err, reTFC := service.GetTS_FCT_CNT(TFC.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reTFC": reTFC}, c)
	}
}

// @Tags TS_FCT_CNT
// @Summary 分页获取TS_FCT_CNT列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.TS_FCT_CNTSearch true "分页获取TS_FCT_CNT列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /TFC/getTS_FCT_CNTList [get]
func GetTS_FCT_CNTList(c *gin.Context) {
	var pageInfo request.TS_FCT_CNTSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetTS_FCT_CNTInfoList(pageInfo); err != nil {
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

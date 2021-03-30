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

// @Tags TS_VI_CNT
// @Summary 创建TS_VI_CNT
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TS_VI_CNT true "创建TS_VI_CNT"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /TVC/createTS_VI_CNT [post]
func CreateTS_VI_CNT(c *gin.Context) {
	var TVC model.TS_VI_CNT
	_ = c.ShouldBindJSON(&TVC)
	if err := service.CreateTS_VI_CNT(TVC); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags TS_VI_CNT
// @Summary 删除TS_VI_CNT
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TS_VI_CNT true "删除TS_VI_CNT"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /TVC/deleteTS_VI_CNT [delete]
func DeleteTS_VI_CNT(c *gin.Context) {
	var TVC model.TS_VI_CNT
	_ = c.ShouldBindJSON(&TVC)
	if err := service.DeleteTS_VI_CNT(TVC); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags TS_VI_CNT
// @Summary 批量删除TS_VI_CNT
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TS_VI_CNT"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /TVC/deleteTS_VI_CNTByIds [delete]
func DeleteTS_VI_CNTByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteTS_VI_CNTByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags TS_VI_CNT
// @Summary 更新TS_VI_CNT
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TS_VI_CNT true "更新TS_VI_CNT"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /TVC/updateTS_VI_CNT [put]
func UpdateTS_VI_CNT(c *gin.Context) {
	var TVC model.TS_VI_CNT
	_ = c.ShouldBindJSON(&TVC)
	if err := service.UpdateTS_VI_CNT(TVC); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags TS_VI_CNT
// @Summary 用id查询TS_VI_CNT
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TS_VI_CNT true "用id查询TS_VI_CNT"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /TVC/findTS_VI_CNT [get]
func FindTS_VI_CNT(c *gin.Context) {
	var TVC model.TS_VI_CNT
	_ = c.ShouldBindQuery(&TVC)
	if err, reTVC := service.GetTS_VI_CNT(TVC.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reTVC": reTVC}, c)
	}
}

// @Tags TS_VI_CNT
// @Summary 分页获取TS_VI_CNT列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.TS_VI_CNTSearch true "分页获取TS_VI_CNT列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /TVC/getTS_VI_CNTList [get]
func GetTS_VI_CNTList(c *gin.Context) {
	var pageInfo request.TS_VI_CNTSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetTS_VI_CNTInfoList(pageInfo); err != nil {
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

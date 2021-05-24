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

// @Tags DeviceLog
// @Summary 创建DeviceLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DeviceLog true "创建DeviceLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /deviceLog/createDeviceLog [post]
func CreateDeviceLog(c *gin.Context) {
	var deviceLog model.DeviceLog
	_ = c.ShouldBindJSON(&deviceLog)
	if err := service.CreateDeviceLog(deviceLog); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags DeviceLog
// @Summary 删除DeviceLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DeviceLog true "删除DeviceLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /deviceLog/deleteDeviceLog [delete]
func DeleteDeviceLog(c *gin.Context) {
	var deviceLog model.DeviceLog
	_ = c.ShouldBindJSON(&deviceLog)
	if err := service.DeleteDeviceLog(deviceLog); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags DeviceLog
// @Summary 批量删除DeviceLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除DeviceLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /deviceLog/deleteDeviceLogByIds [delete]
func DeleteDeviceLogByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteDeviceLogByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags DeviceLog
// @Summary 更新DeviceLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DeviceLog true "更新DeviceLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /deviceLog/updateDeviceLog [put]
func UpdateDeviceLog(c *gin.Context) {
	var deviceLog model.DeviceLog
	_ = c.ShouldBindJSON(&deviceLog)
	if err := service.UpdateDeviceLog(deviceLog); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags DeviceLog
// @Summary 用id查询DeviceLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DeviceLog true "用id查询DeviceLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /deviceLog/findDeviceLog [get]
func FindDeviceLog(c *gin.Context) {
	var deviceLog model.DeviceLog
	_ = c.ShouldBindQuery(&deviceLog)
	if err, redeviceLog := service.GetDeviceLog(deviceLog.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"redeviceLog": redeviceLog}, c)
	}
}

// @Tags DeviceLog
// @Summary 分页获取DeviceLog列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.DeviceLogSearch true "分页获取DeviceLog列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /deviceLog/getDeviceLogList [get]
func GetDeviceLogList(c *gin.Context) {
	var pageInfo request.DeviceLogSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetDeviceLogInfoList(pageInfo); err != nil {
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

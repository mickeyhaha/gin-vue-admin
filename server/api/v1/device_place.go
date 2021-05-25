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

// @Tags DevicePlace
// @Summary 创建DevicePlace
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DevicePlace true "创建DevicePlace"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /place/createDevicePlace [post]
func CreateDevicePlace(c *gin.Context) {
	var place model.DevicePlace
	_ = c.ShouldBindJSON(&place)
	if err := service.CreateDevicePlace(place); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags DevicePlace
// @Summary 删除DevicePlace
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DevicePlace true "删除DevicePlace"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /place/deleteDevicePlace [delete]
func DeleteDevicePlace(c *gin.Context) {
	var place model.DevicePlace
	_ = c.ShouldBindJSON(&place)
	if err := service.DeleteDevicePlace(place); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags DevicePlace
// @Summary 批量删除DevicePlace
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除DevicePlace"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /place/deleteDevicePlaceByIds [delete]
func DeleteDevicePlaceByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteDevicePlaceByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags DevicePlace
// @Summary 更新DevicePlace
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DevicePlace true "更新DevicePlace"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /place/updateDevicePlace [put]
func UpdateDevicePlace(c *gin.Context) {
	var place model.DevicePlace
	_ = c.ShouldBindJSON(&place)
	if err := service.UpdateDevicePlace(place); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags DevicePlace
// @Summary 用id查询DevicePlace
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DevicePlace true "用id查询DevicePlace"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /place/findDevicePlace [get]
func FindDevicePlace(c *gin.Context) {
	var place model.DevicePlace
	_ = c.ShouldBindQuery(&place)
	if err, replace := service.GetDevicePlace(place.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"replace": replace}, c)
	}
}

// @Tags DevicePlace
// @Summary 分页获取DevicePlace列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.DevicePlaceSearch true "分页获取DevicePlace列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /place/getDevicePlaceList [get]
func GetDevicePlaceList(c *gin.Context) {
	var pageInfo request.DevicePlaceSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetDevicePlaceInfoList(pageInfo); err != nil {
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

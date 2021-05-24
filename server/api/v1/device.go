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

// @Tags Device
// @Summary 创建Device
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Device true "创建Device"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dev/createDevice [post]
func CreateDevice(c *gin.Context) {
	var dev model.Device
	_ = c.ShouldBindJSON(&dev)
	if err := service.CreateDevice(dev); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Device
// @Summary 删除Device
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Device true "删除Device"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dev/deleteDevice [delete]
func DeleteDevice(c *gin.Context) {
	var dev model.Device
	_ = c.ShouldBindJSON(&dev)
	if err := service.DeleteDevice(dev); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Device
// @Summary 批量删除Device
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Device"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /dev/deleteDeviceByIds [delete]
func DeleteDeviceByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteDeviceByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags Device
// @Summary 更新Device
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Device true "更新Device"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /dev/updateDevice [put]
func UpdateDevice(c *gin.Context) {
	var dev model.Device
	_ = c.ShouldBindJSON(&dev)
	if err := service.UpdateDevice(dev); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Device
// @Summary 用id查询Device
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Device true "用id查询Device"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /dev/findDevice [get]
func FindDevice(c *gin.Context) {
	var dev model.Device
	_ = c.ShouldBindQuery(&dev)
	if err, redev := service.GetDevice(dev.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"redev": redev}, c)
	}
}

// @Tags Device
// @Summary 分页获取Device列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.DeviceSearch true "分页获取Device列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dev/getDeviceList [get]
func GetDeviceList(c *gin.Context) {
	var pageInfo request.DeviceSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetDeviceInfoList(pageInfo); err != nil {
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

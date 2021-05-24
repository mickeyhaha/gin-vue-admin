package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitDeviceLogRouter(Router *gin.RouterGroup) {
	DeviceLogRouter := Router.Group("deviceLog").Use(middleware.OperationRecord())
	{
		DeviceLogRouter.POST("createDeviceLog", v1.CreateDeviceLog)   // 新建DeviceLog
		DeviceLogRouter.DELETE("deleteDeviceLog", v1.DeleteDeviceLog) // 删除DeviceLog
		DeviceLogRouter.DELETE("deleteDeviceLogByIds", v1.DeleteDeviceLogByIds) // 批量删除DeviceLog
		DeviceLogRouter.PUT("updateDeviceLog", v1.UpdateDeviceLog)    // 更新DeviceLog
		DeviceLogRouter.GET("findDeviceLog", v1.FindDeviceLog)        // 根据ID获取DeviceLog
		DeviceLogRouter.GET("getDeviceLogList", v1.GetDeviceLogList)  // 获取DeviceLog列表
	}
}

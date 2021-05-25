package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitDevicePlaceRouter(Router *gin.RouterGroup) {
	DevicePlaceRouter := Router.Group("place").Use(middleware.OperationRecord())
	{
		DevicePlaceRouter.POST("createDevicePlace", v1.CreateDevicePlace)   // 新建DevicePlace
		DevicePlaceRouter.DELETE("deleteDevicePlace", v1.DeleteDevicePlace) // 删除DevicePlace
		DevicePlaceRouter.DELETE("deleteDevicePlaceByIds", v1.DeleteDevicePlaceByIds) // 批量删除DevicePlace
		DevicePlaceRouter.PUT("updateDevicePlace", v1.UpdateDevicePlace)    // 更新DevicePlace
		DevicePlaceRouter.GET("findDevicePlace", v1.FindDevicePlace)        // 根据ID获取DevicePlace
		DevicePlaceRouter.GET("getDevicePlaceList", v1.GetDevicePlaceList)  // 获取DevicePlace列表
	}
}

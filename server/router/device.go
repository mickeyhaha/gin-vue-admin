package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitDeviceRouter(Router *gin.RouterGroup) {
	DeviceRouter := Router.Group("dev").Use(middleware.OperationRecord())
	{
		DeviceRouter.POST("createDevice", v1.CreateDevice)   // 新建Device
		DeviceRouter.DELETE("deleteDevice", v1.DeleteDevice) // 删除Device
		DeviceRouter.DELETE("deleteDeviceByIds", v1.DeleteDeviceByIds) // 批量删除Device
		DeviceRouter.PUT("updateDevice", v1.UpdateDevice)    // 更新Device
		DeviceRouter.GET("findDevice", v1.FindDevice)        // 根据ID获取Device
		DeviceRouter.GET("getDeviceList", v1.GetDeviceList)  // 获取Device列表
	}
}

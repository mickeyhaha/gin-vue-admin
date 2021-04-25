package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitDCSSMTMachineEventRouter(Router *gin.RouterGroup) {
	DCSSMTMachineEventRouter := Router.Group("DSME").Use(middleware.OperationRecord())
	{
		DCSSMTMachineEventRouter.POST("createDCSSMTMachineEvent", v1.CreateDCSSMTMachineEvent)   // 新建DCSSMTMachineEvent
		DCSSMTMachineEventRouter.DELETE("deleteDCSSMTMachineEvent", v1.DeleteDCSSMTMachineEvent) // 删除DCSSMTMachineEvent
		DCSSMTMachineEventRouter.DELETE("deleteDCSSMTMachineEventByIds", v1.DeleteDCSSMTMachineEventByIds) // 批量删除DCSSMTMachineEvent
		DCSSMTMachineEventRouter.PUT("updateDCSSMTMachineEvent", v1.UpdateDCSSMTMachineEvent)    // 更新DCSSMTMachineEvent
		DCSSMTMachineEventRouter.GET("findDCSSMTMachineEvent", v1.FindDCSSMTMachineEvent)        // 根据ID获取DCSSMTMachineEvent
		DCSSMTMachineEventRouter.GET("getDCSSMTMachineEventList", v1.GetDCSSMTMachineEventList)  // 获取DCSSMTMachineEvent列表
	}
}

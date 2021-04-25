package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitDCSSMTRunTimeRouter(Router *gin.RouterGroup) {
	DCSSMTRunTimeRouter := Router.Group("DSRT").Use(middleware.OperationRecord())
	{
		DCSSMTRunTimeRouter.POST("createDCSSMTRunTime", v1.CreateDCSSMTRunTime)   // 新建DCSSMTRunTime
		DCSSMTRunTimeRouter.DELETE("deleteDCSSMTRunTime", v1.DeleteDCSSMTRunTime) // 删除DCSSMTRunTime
		DCSSMTRunTimeRouter.DELETE("deleteDCSSMTRunTimeByIds", v1.DeleteDCSSMTRunTimeByIds) // 批量删除DCSSMTRunTime
		DCSSMTRunTimeRouter.PUT("updateDCSSMTRunTime", v1.UpdateDCSSMTRunTime)    // 更新DCSSMTRunTime
		DCSSMTRunTimeRouter.GET("findDCSSMTRunTime", v1.FindDCSSMTRunTime)        // 根据ID获取DCSSMTRunTime
		DCSSMTRunTimeRouter.GET("getDCSSMTRunTimeList", v1.GetDCSSMTRunTimeList)  // 获取DCSSMTRunTime列表
	}
}

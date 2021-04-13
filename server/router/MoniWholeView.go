package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitMoniWholeViewRouter(Router *gin.RouterGroup) {
	MoniWholeViewRouter := Router.Group("MWV").Use(middleware.OperationRecord())
	{
		MoniWholeViewRouter.POST("createMoniWholeView", v1.CreateMoniWholeView)   // 新建MoniWholeView
		MoniWholeViewRouter.DELETE("deleteMoniWholeView", v1.DeleteMoniWholeView) // 删除MoniWholeView
		MoniWholeViewRouter.DELETE("deleteMoniWholeViewByIds", v1.DeleteMoniWholeViewByIds) // 批量删除MoniWholeView
		MoniWholeViewRouter.PUT("updateMoniWholeView", v1.UpdateMoniWholeView)    // 更新MoniWholeView
		MoniWholeViewRouter.GET("findMoniWholeView", v1.FindMoniWholeView)        // 根据ID获取MoniWholeView
		MoniWholeViewRouter.GET("getMoniWholeViewList", v1.GetMoniWholeViewList)  // 获取MoniWholeView列表
		MoniWholeViewRouter.GET("getRejectRateList", v1.GetRejectRateList)  // 获取MoniWholeView列表
	}
}

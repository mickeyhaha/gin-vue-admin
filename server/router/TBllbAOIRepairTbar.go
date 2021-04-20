package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitTBllbAOIRepairTbarRouter(Router *gin.RouterGroup) {
	TBllbAOIRepairTbarRouter := Router.Group("TBART").Use(middleware.OperationRecord())
	{
		TBllbAOIRepairTbarRouter.POST("createTBllbAOIRepairTbar", v1.CreateTBllbAOIRepairTbar)   // 新建TBllbAOIRepairTbar
		TBllbAOIRepairTbarRouter.DELETE("deleteTBllbAOIRepairTbar", v1.DeleteTBllbAOIRepairTbar) // 删除TBllbAOIRepairTbar
		TBllbAOIRepairTbarRouter.DELETE("deleteTBllbAOIRepairTbarByIds", v1.DeleteTBllbAOIRepairTbarByIds) // 批量删除TBllbAOIRepairTbar
		TBllbAOIRepairTbarRouter.PUT("updateTBllbAOIRepairTbar", v1.UpdateTBllbAOIRepairTbar)    // 更新TBllbAOIRepairTbar
		TBllbAOIRepairTbarRouter.GET("findTBllbAOIRepairTbar", v1.FindTBllbAOIRepairTbar)        // 根据ID获取TBllbAOIRepairTbar
		TBllbAOIRepairTbarRouter.GET("getTBllbAOIRepairTbarList", v1.GetTBllbAOIRepairTbarList)  // 获取TBllbAOIRepairTbar列表
	}
}

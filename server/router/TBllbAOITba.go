package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitTBllbAOITbaRouter(Router *gin.RouterGroup) {
	TBllbAOITbaRouter := Router.Group("TBAT").Use(middleware.OperationRecord())
	{
		TBllbAOITbaRouter.POST("createTBllbAOITba", v1.CreateTBllbAOITba)   // 新建TBllbAOITba
		TBllbAOITbaRouter.DELETE("deleteTBllbAOITba", v1.DeleteTBllbAOITba) // 删除TBllbAOITba
		TBllbAOITbaRouter.DELETE("deleteTBllbAOITbaByIds", v1.DeleteTBllbAOITbaByIds) // 批量删除TBllbAOITba
		TBllbAOITbaRouter.PUT("updateTBllbAOITba", v1.UpdateTBllbAOITba)    // 更新TBllbAOITba
		TBllbAOITbaRouter.GET("findTBllbAOITba", v1.FindTBllbAOITba)        // 根据ID获取TBllbAOITba
		TBllbAOITbaRouter.GET("getTBllbAOITbaList", v1.GetTBllbAOITbaList)  // 获取TBllbAOITba列表
	}
}

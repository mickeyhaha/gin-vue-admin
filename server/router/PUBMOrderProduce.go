package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitPUBMOrderProduce2Router(Router *gin.RouterGroup) {
	PUBMOrderProduce2Router := Router.Group("PUBMOrderProduce").Use(middleware.OperationRecord())
	{
		PUBMOrderProduce2Router.POST("createPUBMOrderProduce2", v1.CreatePUBMOrderProduce2)   // 新建PUBMOrderProduce2
		PUBMOrderProduce2Router.DELETE("deletePUBMOrderProduce2", v1.DeletePUBMOrderProduce2) // 删除PUBMOrderProduce2
		PUBMOrderProduce2Router.DELETE("deletePUBMOrderProduce2ByIds", v1.DeletePUBMOrderProduce2ByIds) // 批量删除PUBMOrderProduce2
		PUBMOrderProduce2Router.PUT("updatePUBMOrderProduce2", v1.UpdatePUBMOrderProduce2)    // 更新PUBMOrderProduce2
		PUBMOrderProduce2Router.GET("findPUBMOrderProduce2", v1.FindPUBMOrderProduce2)        // 根据ID获取PUBMOrderProduce2
		PUBMOrderProduce2Router.GET("getPUBMOrderProduce2List", v1.GetPUBMOrderProduce2List)  // 获取PUBMOrderProduce2列表
	}
}

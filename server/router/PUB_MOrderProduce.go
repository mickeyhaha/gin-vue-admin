package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitPUB_MOrderProduceRouter(Router *gin.RouterGroup) {
	PUB_MOrderProduceRouter := Router.Group("PMOP").Use(middleware.OperationRecord())
	{
		PUB_MOrderProduceRouter.POST("createPUB_MOrderProduce", v1.CreatePUB_MOrderProduce)   // 新建PUB_MOrderProduce
		PUB_MOrderProduceRouter.DELETE("deletePUB_MOrderProduce", v1.DeletePUB_MOrderProduce) // 删除PUB_MOrderProduce
		PUB_MOrderProduceRouter.DELETE("deletePUB_MOrderProduceByIds", v1.DeletePUB_MOrderProduceByIds) // 批量删除PUB_MOrderProduce
		PUB_MOrderProduceRouter.PUT("updatePUB_MOrderProduce", v1.UpdatePUB_MOrderProduce)    // 更新PUB_MOrderProduce
		PUB_MOrderProduceRouter.GET("findPUB_MOrderProduce", v1.FindPUB_MOrderProduce)        // 根据ID获取PUB_MOrderProduce
		PUB_MOrderProduceRouter.GET("getPUB_MOrderProduceList", v1.GetPUB_MOrderProduceList)  // 获取PUB_MOrderProduce列表
	}
}

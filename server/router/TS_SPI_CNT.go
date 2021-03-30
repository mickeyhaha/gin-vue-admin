package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitTS_SPI_CNTRouter(Router *gin.RouterGroup) {
	TS_SPI_CNTRouter := Router.Group("TSC").Use(middleware.OperationRecord())
	{
		TS_SPI_CNTRouter.POST("createTS_SPI_CNT", v1.CreateTS_SPI_CNT)   // 新建TS_SPI_CNT
		TS_SPI_CNTRouter.DELETE("deleteTS_SPI_CNT", v1.DeleteTS_SPI_CNT) // 删除TS_SPI_CNT
		TS_SPI_CNTRouter.DELETE("deleteTS_SPI_CNTByIds", v1.DeleteTS_SPI_CNTByIds) // 批量删除TS_SPI_CNT
		TS_SPI_CNTRouter.PUT("updateTS_SPI_CNT", v1.UpdateTS_SPI_CNT)    // 更新TS_SPI_CNT
		TS_SPI_CNTRouter.GET("findTS_SPI_CNT", v1.FindTS_SPI_CNT)        // 根据ID获取TS_SPI_CNT
		TS_SPI_CNTRouter.GET("getTS_SPI_CNTList", v1.GetTS_SPI_CNTList)  // 获取TS_SPI_CNT列表
	}
}

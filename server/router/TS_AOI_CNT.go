package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitTS_AOI_CNTRouter(Router *gin.RouterGroup) {
	TS_AOI_CNTRouter := Router.Group("TAC").Use(middleware.OperationRecord())
	{
		TS_AOI_CNTRouter.POST("createTS_AOI_CNT", v1.CreateTS_AOI_CNT)   // 新建TS_AOI_CNT
		TS_AOI_CNTRouter.DELETE("deleteTS_AOI_CNT", v1.DeleteTS_AOI_CNT) // 删除TS_AOI_CNT
		TS_AOI_CNTRouter.DELETE("deleteTS_AOI_CNTByIds", v1.DeleteTS_AOI_CNTByIds) // 批量删除TS_AOI_CNT
		TS_AOI_CNTRouter.PUT("updateTS_AOI_CNT", v1.UpdateTS_AOI_CNT)    // 更新TS_AOI_CNT
		TS_AOI_CNTRouter.GET("findTS_AOI_CNT", v1.FindTS_AOI_CNT)        // 根据ID获取TS_AOI_CNT
		TS_AOI_CNTRouter.GET("getTS_AOI_CNTList", v1.GetTS_AOI_CNTList)  // 获取TS_AOI_CNT列表
		TS_AOI_CNTRouter.GET("getTS_AOI_CNTList4Chart", v1.GetTS_AOI_CNTList4Chart)  // 获取TS_AOI_CNT列表4Chart
	}
}

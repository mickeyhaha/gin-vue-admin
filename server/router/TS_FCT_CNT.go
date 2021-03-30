package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitTS_FCT_CNTRouter(Router *gin.RouterGroup) {
	TS_FCT_CNTRouter := Router.Group("TFC").Use(middleware.OperationRecord())
	{
		TS_FCT_CNTRouter.POST("createTS_FCT_CNT", v1.CreateTS_FCT_CNT)   // 新建TS_FCT_CNT
		TS_FCT_CNTRouter.DELETE("deleteTS_FCT_CNT", v1.DeleteTS_FCT_CNT) // 删除TS_FCT_CNT
		TS_FCT_CNTRouter.DELETE("deleteTS_FCT_CNTByIds", v1.DeleteTS_FCT_CNTByIds) // 批量删除TS_FCT_CNT
		TS_FCT_CNTRouter.PUT("updateTS_FCT_CNT", v1.UpdateTS_FCT_CNT)    // 更新TS_FCT_CNT
		TS_FCT_CNTRouter.GET("findTS_FCT_CNT", v1.FindTS_FCT_CNT)        // 根据ID获取TS_FCT_CNT
		TS_FCT_CNTRouter.GET("getTS_FCT_CNTList", v1.GetTS_FCT_CNTList)  // 获取TS_FCT_CNT列表
	}
}

package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitTS_VI_CNTRouter(Router *gin.RouterGroup) {
	TS_VI_CNTRouter := Router.Group("TVC").Use(middleware.OperationRecord())
	{
		TS_VI_CNTRouter.POST("createTS_VI_CNT", v1.CreateTS_VI_CNT)   // 新建TS_VI_CNT
		TS_VI_CNTRouter.DELETE("deleteTS_VI_CNT", v1.DeleteTS_VI_CNT) // 删除TS_VI_CNT
		TS_VI_CNTRouter.DELETE("deleteTS_VI_CNTByIds", v1.DeleteTS_VI_CNTByIds) // 批量删除TS_VI_CNT
		TS_VI_CNTRouter.PUT("updateTS_VI_CNT", v1.UpdateTS_VI_CNT)    // 更新TS_VI_CNT
		TS_VI_CNTRouter.GET("findTS_VI_CNT", v1.FindTS_VI_CNT)        // 根据ID获取TS_VI_CNT
		TS_VI_CNTRouter.GET("getTS_VI_CNTList", v1.GetTS_VI_CNTList)  // 获取TS_VI_CNT列表
	}
}

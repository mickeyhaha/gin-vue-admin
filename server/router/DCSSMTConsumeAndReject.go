package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitDCSSMTConsumeAndRejectRouter(Router *gin.RouterGroup) {
	DCSSMTConsumeAndRejectRouter := Router.Group("DSCR").Use(middleware.OperationRecord())
	{
		DCSSMTConsumeAndRejectRouter.POST("createDCSSMTConsumeAndReject", v1.CreateDCSSMTConsumeAndReject)   // 新建DCSSMTConsumeAndReject
		DCSSMTConsumeAndRejectRouter.DELETE("deleteDCSSMTConsumeAndReject", v1.DeleteDCSSMTConsumeAndReject) // 删除DCSSMTConsumeAndReject
		DCSSMTConsumeAndRejectRouter.DELETE("deleteDCSSMTConsumeAndRejectByIds", v1.DeleteDCSSMTConsumeAndRejectByIds) // 批量删除DCSSMTConsumeAndReject
		DCSSMTConsumeAndRejectRouter.PUT("updateDCSSMTConsumeAndReject", v1.UpdateDCSSMTConsumeAndReject)    // 更新DCSSMTConsumeAndReject
		DCSSMTConsumeAndRejectRouter.GET("findDCSSMTConsumeAndReject", v1.FindDCSSMTConsumeAndReject)        // 根据ID获取DCSSMTConsumeAndReject
		DCSSMTConsumeAndRejectRouter.GET("getDCSSMTConsumeAndRejectList", v1.GetDCSSMTConsumeAndRejectList)  // 获取DCSSMTConsumeAndReject列表
	}
}

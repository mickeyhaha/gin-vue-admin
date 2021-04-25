package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitDCSSMTRejectRouter(Router *gin.RouterGroup) {
	DCSSMTRejectRouter := Router.Group("DSR").Use(middleware.OperationRecord())
	{
		DCSSMTRejectRouter.POST("createDCSSMTReject", v1.CreateDCSSMTReject)   // 新建DCSSMTReject
		DCSSMTRejectRouter.DELETE("deleteDCSSMTReject", v1.DeleteDCSSMTReject) // 删除DCSSMTReject
		DCSSMTRejectRouter.DELETE("deleteDCSSMTRejectByIds", v1.DeleteDCSSMTRejectByIds) // 批量删除DCSSMTReject
		DCSSMTRejectRouter.PUT("updateDCSSMTReject", v1.UpdateDCSSMTReject)    // 更新DCSSMTReject
		DCSSMTRejectRouter.GET("findDCSSMTReject", v1.FindDCSSMTReject)        // 根据ID获取DCSSMTReject
		DCSSMTRejectRouter.GET("getDCSSMTRejectList", v1.GetDCSSMTRejectList)  // 获取DCSSMTReject列表
	}
}

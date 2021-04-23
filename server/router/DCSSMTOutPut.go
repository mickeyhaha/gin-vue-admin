package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitDCSSMTOutPutRouter(Router *gin.RouterGroup) {
	DCSSMTOutPutRouter := Router.Group("DSO").Use(middleware.OperationRecord())
	{
		DCSSMTOutPutRouter.POST("createDCSSMTOutPut", v1.CreateDCSSMTOutPut)   // 新建DCSSMTOutPut
		DCSSMTOutPutRouter.DELETE("deleteDCSSMTOutPut", v1.DeleteDCSSMTOutPut) // 删除DCSSMTOutPut
		DCSSMTOutPutRouter.DELETE("deleteDCSSMTOutPutByIds", v1.DeleteDCSSMTOutPutByIds) // 批量删除DCSSMTOutPut
		DCSSMTOutPutRouter.PUT("updateDCSSMTOutPut", v1.UpdateDCSSMTOutPut)    // 更新DCSSMTOutPut
		DCSSMTOutPutRouter.GET("findDCSSMTOutPut", v1.FindDCSSMTOutPut)        // 根据ID获取DCSSMTOutPut
		DCSSMTOutPutRouter.GET("getDCSSMTOutPutList", v1.GetDCSSMTOutPutList)  // 获取DCSSMTOutPut列表
	}
}

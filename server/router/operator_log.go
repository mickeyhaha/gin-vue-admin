package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitOperatorLogRouter(Router *gin.RouterGroup) {
	OperatorLogRouter := Router.Group("opLog").Use(middleware.OperationRecord())
	{
		OperatorLogRouter.POST("createOperatorLog", v1.CreateOperatorLog)   // 新建OperatorLog
		OperatorLogRouter.DELETE("deleteOperatorLog", v1.DeleteOperatorLog) // 删除OperatorLog
		OperatorLogRouter.DELETE("deleteOperatorLogByIds", v1.DeleteOperatorLogByIds) // 批量删除OperatorLog
		OperatorLogRouter.PUT("updateOperatorLog", v1.UpdateOperatorLog)    // 更新OperatorLog
		OperatorLogRouter.GET("findOperatorLog", v1.FindOperatorLog)        // 根据ID获取OperatorLog
		OperatorLogRouter.GET("getOperatorLogList", v1.GetOperatorLogList)  // 获取OperatorLog列表
	}
}

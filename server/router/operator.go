package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitOperatorRouter(Router *gin.RouterGroup) {
	OperatorRouter := Router.Group("opt").Use(middleware.OperationRecord())
	{
		OperatorRouter.POST("createOperator", v1.CreateOperator)   // 新建Operator
		OperatorRouter.DELETE("deleteOperator", v1.DeleteOperator) // 删除Operator
		OperatorRouter.DELETE("deleteOperatorByIds", v1.DeleteOperatorByIds) // 批量删除Operator
		OperatorRouter.PUT("updateOperator", v1.UpdateOperator)    // 更新Operator
		OperatorRouter.GET("findOperator", v1.FindOperator)        // 根据ID获取Operator
		OperatorRouter.GET("getOperatorList", v1.GetOperatorList)  // 获取Operator列表
	}
}

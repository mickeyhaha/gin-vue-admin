package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitTBllbShiftManageRouter(Router *gin.RouterGroup) {
	TBllbShiftManageRouter := Router.Group("TSM").Use(middleware.OperationRecord())
	{
		TBllbShiftManageRouter.POST("createTBllbShiftManage", v1.CreateTBllbShiftManage)   // 新建TBllbShiftManage
		TBllbShiftManageRouter.DELETE("deleteTBllbShiftManage", v1.DeleteTBllbShiftManage) // 删除TBllbShiftManage
		TBllbShiftManageRouter.DELETE("deleteTBllbShiftManageByIds", v1.DeleteTBllbShiftManageByIds) // 批量删除TBllbShiftManage
		TBllbShiftManageRouter.PUT("updateTBllbShiftManage", v1.UpdateTBllbShiftManage)    // 更新TBllbShiftManage
		TBllbShiftManageRouter.GET("findTBllbShiftManage", v1.FindTBllbShiftManage)        // 根据ID获取TBllbShiftManage
		TBllbShiftManageRouter.GET("getTBllbShiftManageList", v1.GetTBllbShiftManageList)  // 获取TBllbShiftManage列表
	}
}

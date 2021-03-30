package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitVM_TimeOutputRouter(Router *gin.RouterGroup) {
	VM_TimeOutputRouter := Router.Group("VTO").Use(middleware.OperationRecord())
	{
		VM_TimeOutputRouter.POST("createVM_TimeOutput", v1.CreateVM_TimeOutput)   // 新建VM_TimeOutput
		VM_TimeOutputRouter.DELETE("deleteVM_TimeOutput", v1.DeleteVM_TimeOutput) // 删除VM_TimeOutput
		VM_TimeOutputRouter.DELETE("deleteVM_TimeOutputByIds", v1.DeleteVM_TimeOutputByIds) // 批量删除VM_TimeOutput
		VM_TimeOutputRouter.PUT("updateVM_TimeOutput", v1.UpdateVM_TimeOutput)    // 更新VM_TimeOutput
		VM_TimeOutputRouter.GET("findVM_TimeOutput", v1.FindVM_TimeOutput)        // 根据ID获取VM_TimeOutput
		VM_TimeOutputRouter.GET("getVM_TimeOutputList", v1.GetVM_TimeOutputList)  // 获取VM_TimeOutput列表
	}
}

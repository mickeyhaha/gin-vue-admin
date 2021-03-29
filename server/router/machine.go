package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitMachineRouter(Router *gin.RouterGroup) {
	MachineRouter := Router.Group("machine").Use(middleware.OperationRecord())
	{
		MachineRouter.POST("createMachine", v1.CreateMachine)   // 新建Machine
		MachineRouter.DELETE("deleteMachine", v1.DeleteMachine) // 删除Machine
		MachineRouter.DELETE("deleteMachineByIds", v1.DeleteMachineByIds) // 批量删除Machine
		MachineRouter.PUT("updateMachine", v1.UpdateMachine)    // 更新Machine
		MachineRouter.GET("findMachine", v1.FindMachine)        // 根据ID获取Machine
		MachineRouter.GET("getMachineList", v1.GetMachineList)  // 获取Machine列表
		
		MachineRouter.GET("getFlopData", v1.GetFlopData)  // getFlopData
	}
}

package v1

import (
	"gin-vue-admin/global"
    "gin-vue-admin/model"
    "gin-vue-admin/model/request"
    "gin-vue-admin/model/response"
    "gin-vue-admin/service"
    "github.com/gin-gonic/gin"
	"go.uber.org/zap"
	//"math/rand"

	//"math/rand"
)

// @Tags Machine
// @Summary 创建Machine
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Machine true "创建Machine"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /machine/createMachine [post]
func CreateMachine(c *gin.Context) {
	var machine model.Machine
	_ = c.ShouldBindJSON(&machine)
	if err := service.CreateMachine(machine); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Machine
// @Summary 删除Machine
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Machine true "删除Machine"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /machine/deleteMachine [delete]
func DeleteMachine(c *gin.Context) {
	var machine model.Machine
	_ = c.ShouldBindJSON(&machine)
	if err := service.DeleteMachine(machine); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Machine
// @Summary 批量删除Machine
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Machine"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /machine/deleteMachineByIds [delete]
func DeleteMachineByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteMachineByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags Machine
// @Summary 更新Machine
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Machine true "更新Machine"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /machine/updateMachine [put]
func UpdateMachine(c *gin.Context) {
	var machine model.Machine
	_ = c.ShouldBindJSON(&machine)
	if err := service.UpdateMachine(machine); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Machine
// @Summary 用id查询Machine
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Machine true "用id查询Machine"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /machine/findMachine [get]
func FindMachine(c *gin.Context) {
	var machine model.Machine
	_ = c.ShouldBindQuery(&machine)
	if err, remachine := service.GetMachine(machine.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"remachine": remachine}, c)
	}
}

// @Tags Machine
// @Summary 分页获取Machine列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.MachineSearch true "分页获取Machine列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /machine/getMachineList [get]
func GetMachineList(c *gin.Context) {
	var pageInfo request.MachineSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetMachineInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败", zap.Any("err", err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}

func randomExtend(minNum, maxNum int64) int64 {
	return 10
}

func GetFlopData(c *gin.Context) {
	var pageInfo request.MachineSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetMachineInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败", zap.Any("err", err))
       response.FailWithMessage("获取失败", c)
   } else {
       response.OkWithDetailed(response.PageResult{
           List:     list,
           Total:    total,
           Page:     pageInfo.Page,
           PageSize: pageInfo.PageSize,
       }, "获取成功", c)
   }
}

package v1

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags PVS_Base_Line
// @Summary 创建PVS_Base_Line
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PVS_Base_Line true "创建PVS_Base_Line"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /PBL/createPVS_Base_Line [post]
func CreatePVS_Base_Line(c *gin.Context) {
	var PBL model.PVS_Base_Line
	_ = c.ShouldBindJSON(&PBL)
	if err := service.CreatePVS_Base_Line(PBL); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags PVS_Base_Line
// @Summary 删除PVS_Base_Line
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PVS_Base_Line true "删除PVS_Base_Line"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /PBL/deletePVS_Base_Line [delete]
func DeletePVS_Base_Line(c *gin.Context) {
	var PBL model.PVS_Base_Line
	_ = c.ShouldBindJSON(&PBL)
	if err := service.DeletePVS_Base_Line(PBL); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags PVS_Base_Line
// @Summary 批量删除PVS_Base_Line
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PVS_Base_Line"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /PBL/deletePVS_Base_LineByIds [delete]
func DeletePVS_Base_LineByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeletePVS_Base_LineByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags PVS_Base_Line
// @Summary 更新PVS_Base_Line
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PVS_Base_Line true "更新PVS_Base_Line"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /PBL/updatePVS_Base_Line [put]
func UpdatePVS_Base_Line(c *gin.Context) {
	var PBL model.PVS_Base_Line
	_ = c.ShouldBindJSON(&PBL)
	if err := service.UpdatePVS_Base_Line(PBL); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags PVS_Base_Line
// @Summary 用id查询PVS_Base_Line
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PVS_Base_Line true "用id查询PVS_Base_Line"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /PBL/findPVS_Base_Line [get]
func FindPVS_Base_Line(c *gin.Context) {
	var PBL model.PVS_Base_Line
	_ = c.ShouldBindQuery(&PBL)
	if err, rePBL := service.GetPVS_Base_Line(PBL.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rePBL": rePBL}, c)
	}
}

// @Tags PVS_Base_Line
// @Summary 分页获取PVS_Base_Line列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PVS_Base_LineSearch true "分页获取PVS_Base_Line列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /PBL/getPVS_Base_LineList [get]
func GetPVS_Base_LineList(c *gin.Context) {
	var pageInfo request.PVS_Base_LineSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetPVS_Base_LineInfoList(pageInfo); err != nil {
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

func GetDeptLineSummary(c *gin.Context) {
	var pageInfo request.PUBMOrderProduce2Search
	pageInfo.Status = 2
	if err, line, total := service.GetPUBMOrderProduce2InfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		// Fill aoi info
		// TODO 打开
		if err, aoiInfoList, _ := service.GetTS_AOI_CNTInfoListByLine(); err == nil {
			lineArr := line.([]model.PUBMOrderProduce2)
			for i := 0; i < len(lineArr); i++ {
				for j := 0; j < len(aoiInfoList); j++  {
					if lineArr[i].LineID == aoiInfoList[j].LineID {
						lineArr[i].ErrCount = aoiInfoList[j].ErrCount
						lineArr[i].Count = aoiInfoList[j].Count
						if lineArr[i].Count != 0 {
							lineArr[i].AoiErrRate = (float64(lineArr[i].ErrCount) / float64(lineArr[i].Count))
						}
						break
					}
				}
			}
		}
		response.OkWithDetailed(response.PageResult{
			List:     line,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

func GetWorkOrderListByLine(c *gin.Context) {
	var pageInfo request.PUBMOrderProduce2Search
	_ = c.ShouldBindQuery(&pageInfo)

	if err, lines, total := service.GetPUBMOrderProduce2InfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		orderArr := lines.([]model.PUBMOrderProduce2)
		lineOrderMap := make(map[string]model.PUBMOrderProduce2, 0)
		for i:=0; i< len(orderArr); i++ {
			if _, ok := lineOrderMap[orderArr[i].WorkOrderNo]; !ok {
				lineOrderMap[orderArr[i].WorkOrderNo] = orderArr[i]
			}
		}
		lineArr := make([]model.PUBMOrderProduce2, 0)
		for _, value := range lineOrderMap {
			lineArr = append(lineArr, value)
		}
		response.OkWithDetailed(response.PageResult{
			List:     lineArr,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}


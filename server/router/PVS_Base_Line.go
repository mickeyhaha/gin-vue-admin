package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitPVS_Base_LineRouter(Router *gin.RouterGroup) {
	PVS_Base_LineRouter := Router.Group("PBL").Use(middleware.OperationRecord())
	{
		PVS_Base_LineRouter.POST("createPVS_Base_Line", v1.CreatePVS_Base_Line)   // 新建PVS_Base_Line
		PVS_Base_LineRouter.DELETE("deletePVS_Base_Line", v1.DeletePVS_Base_Line) // 删除PVS_Base_Line
		PVS_Base_LineRouter.DELETE("deletePVS_Base_LineByIds", v1.DeletePVS_Base_LineByIds) // 批量删除PVS_Base_Line
		PVS_Base_LineRouter.PUT("updatePVS_Base_Line", v1.UpdatePVS_Base_Line)    // 更新PVS_Base_Line
		PVS_Base_LineRouter.GET("findPVS_Base_Line", v1.FindPVS_Base_Line)        // 根据ID获取PVS_Base_Line
		PVS_Base_LineRouter.GET("getPVS_Base_LineList", v1.GetPVS_Base_LineList)  // 获取PVS_Base_Line列表

		// All smt router here
		PVS_Base_LineRouter.GET("getDeptLineSummary", v1.GetDeptLineSummary)
		PVS_Base_LineRouter.GET("getWorkOrderListByLine", v1.GetWorkOrderListByLine)
		PVS_Base_LineRouter.GET("getDCSSMTOutPutList4Chart", v1.GetDCSSMTOutPutList4Chart)  // TODO 获取产量old
		PVS_Base_LineRouter.GET("getPUBMOrderProduce2InfoList4Chart", v1.GetPUBMOrderProduce2InfoList4Chart)  // 获取产量
		PVS_Base_LineRouter.GET("getDCSSMTConsumeAndRejectRate4Chart", v1.GetDCSSMTConsumeAndRejectRate4Chart)  // 获取抛料率
		//PVS_Base_LineRouter.GET("getDCSSMTConsumeAndRejectRate4Chart", v1.GetRejectRateList4Chart)  // 获取抛料率
		PVS_Base_LineRouter.GET("getDCSSMTMachineEvent4Chart", v1.GetDCSSMTMachineEvent4Chart)  // 获取事件异常
		PVS_Base_LineRouter.GET("getDCSSMTRunTime4Chart", v1.GetDCSSMTRunTime4Chart)  // 获取停机时长 & 停机分布
		PVS_Base_LineRouter.GET("getTS_AOI_CNTList4Chart", v1.GetTS_AOI_CNTList4Chart)  // 获取TS_AOI_CNT列表4Chart

		PVS_Base_LineRouter.GET("getTS_AOI_Spc4Chart", v1.GetTS_AOI_Spc4Chart)  // 获取TS_AOI_CNT列表SPC数据分布图
		//PVS_Base_LineRouter.GET("getTS_AOI_SpcAvg4Chart", v1.GetTS_AOI_SpcAvg4Chart)  // 获取TS_AOI_CNT列表SPC数据均值图
		//PVS_Base_LineRouter.GET("getTS_AOI_SpcRage4Chart", v1.GetTS_AOI_SpcRage4Chart)  // 获取TS_AOI_CNT列表SPC数据极差图
		//PVS_Base_LineRouter.GET("getTS_AOI_SpcVariance4Chart", v1.GetTS_AOI_SpcVariance4Chart)  // 获取TS_AOI_CNT列表SPC数据方差图

		// chart service for dept_dashboard
		//PVS_Base_LineRouter.GET("getPUBMOrderProduce2InfoList4ChartDash", v1.GetPUBMOrderProduce2InfoList4ChartDash)  // 获取产量
		//PVS_Base_LineRouter.GET("getDCSSMTConsumeAndRejectRate4ChartDash", v1.GetDCSSMTConsumeAndRejectRate4ChartDash)  // 获取抛料率
		//PVS_Base_LineRouter.GET("getDCSSMTMachineEvent4ChartDash", v1.GetDCSSMTMachineEvent4ChartDash)  // 获取事件异常
		//PVS_Base_LineRouter.GET("getDCSSMTRunTime4ChartDash", v1.GetDCSSMTRunTime4ChartDash)  // 获取停机时长 & 停机分布
		//PVS_Base_LineRouter.GET("getTS_AOI_CNTList4ChartDash", v1.GetTS_AOI_CNTList4ChartDash)  // 获取TS_AOI_CNT列表4Chart
	}
}

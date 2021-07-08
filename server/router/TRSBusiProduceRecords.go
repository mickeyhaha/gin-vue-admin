package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitTRSBusiProduceRecordsRouter(Router *gin.RouterGroup) {
	TRSBusiProduceRecordsRouter := Router.Group("TBPR").Use(middleware.OperationRecord())
	{
		TRSBusiProduceRecordsRouter.POST("createTRSBusiProduceRecords", v1.CreateTRSBusiProduceRecords)   // 新建TRSBusiProduceRecords
		TRSBusiProduceRecordsRouter.DELETE("deleteTRSBusiProduceRecords", v1.DeleteTRSBusiProduceRecords) // 删除TRSBusiProduceRecords
		TRSBusiProduceRecordsRouter.DELETE("deleteTRSBusiProduceRecordsByIds", v1.DeleteTRSBusiProduceRecordsByIds) // 批量删除TRSBusiProduceRecords
		TRSBusiProduceRecordsRouter.PUT("updateTRSBusiProduceRecords", v1.UpdateTRSBusiProduceRecords)    // 更新TRSBusiProduceRecords
		TRSBusiProduceRecordsRouter.GET("findTRSBusiProduceRecords", v1.FindTRSBusiProduceRecords)        // 根据ID获取TRSBusiProduceRecords
		TRSBusiProduceRecordsRouter.GET("getTRSBusiProduceRecordsList", v1.GetTRSBusiProduceRecordsList)  // 获取TRSBusiProduceRecords列表
	}
}

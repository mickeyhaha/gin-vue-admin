package main

import (
	"gin-vue-admin/core"
	"gin-vue-admin/global"
	"gin-vue-admin/initialize"
)

// @title Swagger Example API
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
// @BasePath /
func main() {
	global.GVA_VP = core.Viper()      // 初始化Viper
	global.GVA_LOG = core.Zap()       // 初始化zap日志库
	global.GVA_DB = initialize.Gorm("mysql") // gorm连接数据库
	global.GVA_DB_MSSQL = initialize.Gorm("mssql") // gorm连接数据库
	if global.GVA_DB != nil {
		initialize.MysqlTables(global.GVA_DB) // 初始化表
		initialize.MssqlTables(global.GVA_DB_MSSQL) // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.GVA_DB.DB()
		dbMssql, _ := global.GVA_DB_MSSQL.DB()
		defer db.Close()
		defer dbMssql.Close()
	}
	core.RunWindowsServer()
}

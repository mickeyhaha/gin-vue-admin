package global

import (
	"go.uber.org/zap"

	"gin-vue-admin/config"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	GVA_DB     *gorm.DB
	GVA_DB_MSSQL     *gorm.DB
	GVA_REDIS  *redis.Client
	GVA_CONFIG config.Server
	GVA_VP     *viper.Viper
	//GVA_LOG    *oplogging.Logger
	GVA_LOG    *zap.Logger
)

//const DateBaseFmt = "2006-01-02 15:04:05"
const DateBaseFmt = "01/02"
const DateTimeBaseFmt = "01/02 15:04"
const Shift_Day_Begin_Hour = 8
const Shift_Day_End_Hour = 18
const Shift_Night_Begin_Hour = 19
const Shift_Night_End_Hour = 7
package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateDeviceLog
//@description: 创建DeviceLog记录
//@param: deviceLog model.DeviceLog
//@return: err error

func CreateDeviceLog(deviceLog model.DeviceLog) (err error) {
	err = global.GVA_DB.Create(&deviceLog).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteDeviceLog
//@description: 删除DeviceLog记录
//@param: deviceLog model.DeviceLog
//@return: err error

func DeleteDeviceLog(deviceLog model.DeviceLog) (err error) {
	err = global.GVA_DB.Delete(&deviceLog).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteDeviceLogByIds
//@description: 批量删除DeviceLog记录
//@param: ids request.IdsReq
//@return: err error

func DeleteDeviceLogByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.DeviceLog{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateDeviceLog
//@description: 更新DeviceLog记录
//@param: deviceLog *model.DeviceLog
//@return: err error

func UpdateDeviceLog(deviceLog model.DeviceLog) (err error) {
	err = global.GVA_DB.Save(&deviceLog).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetDeviceLog
//@description: 根据id获取DeviceLog记录
//@param: id uint
//@return: err error, deviceLog model.DeviceLog

func GetDeviceLog(id uint) (err error, deviceLog model.DeviceLog) {
	err = global.GVA_DB.Where("id = ?", id).First(&deviceLog).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetDeviceLogInfoList
//@description: 分页获取DeviceLog记录
//@param: info request.DeviceLogSearch
//@return: err error, list interface{}, total int64

func GetDeviceLogInfoList(info request.DeviceLogSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.DeviceLog{})
    var deviceLogs []model.DeviceLog
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.DeviceId != "" {
        db = db.Where("`device_id` = ?",info.DeviceId)
    }
    if info.Temp != 0 {
        db = db.Where("`temp` > ?",info.Temp)
    }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&deviceLogs).Error
	return err, deviceLogs, total
}
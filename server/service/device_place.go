package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateDevicePlace
//@description: 创建DevicePlace记录
//@param: place model.DevicePlace
//@return: err error

func CreateDevicePlace(place model.DevicePlace) (err error) {
	err = global.GVA_DB.Create(&place).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteDevicePlace
//@description: 删除DevicePlace记录
//@param: place model.DevicePlace
//@return: err error

func DeleteDevicePlace(place model.DevicePlace) (err error) {
	err = global.GVA_DB.Delete(&place).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteDevicePlaceByIds
//@description: 批量删除DevicePlace记录
//@param: ids request.IdsReq
//@return: err error

func DeleteDevicePlaceByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.DevicePlace{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateDevicePlace
//@description: 更新DevicePlace记录
//@param: place *model.DevicePlace
//@return: err error

func UpdateDevicePlace(place model.DevicePlace) (err error) {
	err = global.GVA_DB.Save(&place).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetDevicePlace
//@description: 根据id获取DevicePlace记录
//@param: id uint
//@return: err error, place model.DevicePlace

func GetDevicePlace(id uint) (err error, place model.DevicePlace) {
	err = global.GVA_DB.Where("id = ?", id).First(&place).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetDevicePlaceInfoList
//@description: 分页获取DevicePlace记录
//@param: info request.DevicePlaceSearch
//@return: err error, list interface{}, total int64

func GetDevicePlaceInfoList(info request.DevicePlaceSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.DevicePlace{})
    var places []model.DevicePlace
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.PlaceId != "" {
        db = db.Where("`place_id` LIKE ?","%"+ info.PlaceId+"%")
    }
    if info.PlaceName != "" {
        db = db.Where("`place_name` LIKE ?","%"+ info.PlaceName+"%")
    }
    if info.Status != 0 {
        db = db.Where("`status` = ?",info.Status)
    }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&places).Error
	return err, places, total
}
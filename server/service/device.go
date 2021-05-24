package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateDevice
//@description: 创建Device记录
//@param: dev model.Device
//@return: err error

func CreateDevice(dev model.Device) (err error) {
	err = global.GVA_DB.Create(&dev).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteDevice
//@description: 删除Device记录
//@param: dev model.Device
//@return: err error

func DeleteDevice(dev model.Device) (err error) {
	err = global.GVA_DB.Delete(&dev).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteDeviceByIds
//@description: 批量删除Device记录
//@param: ids request.IdsReq
//@return: err error

func DeleteDeviceByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Device{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateDevice
//@description: 更新Device记录
//@param: dev *model.Device
//@return: err error

func UpdateDevice(dev model.Device) (err error) {
	err = global.GVA_DB.Save(&dev).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetDevice
//@description: 根据id获取Device记录
//@param: id uint
//@return: err error, dev model.Device

func GetDevice(id uint) (err error, dev model.Device) {
	err = global.GVA_DB.Where("id = ?", id).First(&dev).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetDeviceInfoList
//@description: 分页获取Device记录
//@param: info request.DeviceSearch
//@return: err error, list interface{}, total int64

func GetDeviceInfoList(info request.DeviceSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.Device{})
    var devs []model.Device
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.DeviceId != "" {
        db = db.Where("`device_id` LIKE ?","%"+ info.DeviceId+"%")
    }
    if info.DeviceName != "" {
        db = db.Where("`device_name` LIKE ?","%"+ info.DeviceName+"%")
    }
    if info.DisplayId != "" {
        db = db.Where("`display_id` LIKE ?","%"+ info.DisplayId+"%")
    }
    if info.DisplayName != "" {
        db = db.Where("`display_name` LIKE ?","%"+ info.DisplayName+"%")
    }
    if info.BarCode != "" {
        db = db.Where("`bar_code` LIKE ?","%"+ info.BarCode+"%")
    }
    if info.OpName != "" {
        db = db.Where("`op_name` LIKE ?","%"+ info.OpName+"%")
    }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&devs).Error
	return err, devs, total
}
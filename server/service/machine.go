package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateMachine
//@description: 创建Machine记录
//@param: machine model.Machine
//@return: err error

func CreateMachine(machine model.Machine) (err error) {
	err = global.GVA_DB.Create(&machine).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteMachine
//@description: 删除Machine记录
//@param: machine model.Machine
//@return: err error

func DeleteMachine(machine model.Machine) (err error) {
	err = global.GVA_DB.Delete(&machine).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteMachineByIds
//@description: 批量删除Machine记录
//@param: ids request.IdsReq
//@return: err error

func DeleteMachineByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Machine{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateMachine
//@description: 更新Machine记录
//@param: machine *model.Machine
//@return: err error

func UpdateMachine(machine model.Machine) (err error) {
	err = global.GVA_DB.Save(&machine).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetMachine
//@description: 根据id获取Machine记录
//@param: id uint
//@return: err error, machine model.Machine

func GetMachine(id uint) (err error, machine model.Machine) {
	err = global.GVA_DB.Where("id = ?", id).First(&machine).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetMachineInfoList
//@description: 分页获取Machine记录
//@param: info request.MachineSearch
//@return: err error, list interface{}, total int64

func GetMachineInfoList(info request.MachineSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.Machine{})
    var machines []model.Machine
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.MachineId != "" {
        db = db.Where("`machine_id` LIKE ?","%"+ info.MachineId+"%")
    }
    if info.MachineName != "" {
        db = db.Where("`machine_name` LIKE ?","%"+ info.MachineName+"%")
    }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&machines).Error
	return err, machines, total
}
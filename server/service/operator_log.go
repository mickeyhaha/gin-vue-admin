package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateOperatorLog
//@description: 创建OperatorLog记录
//@param: opLog model.OperatorLog
//@return: err error

func CreateOperatorLog(opLog model.OperatorLog) (err error) {
	err = global.GVA_DB.Create(&opLog).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteOperatorLog
//@description: 删除OperatorLog记录
//@param: opLog model.OperatorLog
//@return: err error

func DeleteOperatorLog(opLog model.OperatorLog) (err error) {
	err = global.GVA_DB.Delete(&opLog).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteOperatorLogByIds
//@description: 批量删除OperatorLog记录
//@param: ids request.IdsReq
//@return: err error

func DeleteOperatorLogByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.OperatorLog{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateOperatorLog
//@description: 更新OperatorLog记录
//@param: opLog *model.OperatorLog
//@return: err error

func UpdateOperatorLog(opLog model.OperatorLog) (err error) {
	err = global.GVA_DB.Save(&opLog).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetOperatorLog
//@description: 根据id获取OperatorLog记录
//@param: id uint
//@return: err error, opLog model.OperatorLog

func GetOperatorLog(id uint) (err error, opLog model.OperatorLog) {
	err = global.GVA_DB.Where("id = ?", id).First(&opLog).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetOperatorLogInfoList
//@description: 分页获取OperatorLog记录
//@param: info request.OperatorLogSearch
//@return: err error, list interface{}, total int64

func GetOperatorLogInfoList(info request.OperatorLogSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.OperatorLog{})
    var opLogs []model.OperatorLog
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Op_name != "" {
        db = db.Where("`op_name` LIKE ?","%"+ info.Op_name+"%")
    }
    if info.Action != "" {
        db = db.Where("`action` = ?",info.Action)
    }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&opLogs).Error
	return err, opLogs, total
}
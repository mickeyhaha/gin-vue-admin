package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateOperator
//@description: 创建Operator记录
//@param: opt model.Operator
//@return: err error

func CreateOperator(opt model.Operator) (err error) {
	err = global.GVA_DB.Create(&opt).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteOperator
//@description: 删除Operator记录
//@param: opt model.Operator
//@return: err error

func DeleteOperator(opt model.Operator) (err error) {
	err = global.GVA_DB.Delete(&opt).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteOperatorByIds
//@description: 批量删除Operator记录
//@param: ids request.IdsReq
//@return: err error

func DeleteOperatorByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Operator{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateOperator
//@description: 更新Operator记录
//@param: opt *model.Operator
//@return: err error

func UpdateOperator(opt model.Operator) (err error) {
	err = global.GVA_DB.Save(&opt).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetOperator
//@description: 根据id获取Operator记录
//@param: id uint
//@return: err error, opt model.Operator

func GetOperator(id uint) (err error, opt model.Operator) {
	err = global.GVA_DB.Where("id = ?", id).First(&opt).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetOperatorInfoList
//@description: 分页获取Operator记录
//@param: info request.OperatorSearch
//@return: err error, list interface{}, total int64

func GetOperatorInfoList(info request.OperatorSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.Operator{})
    var opts []model.Operator
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Operator_id != "" {
        db = db.Where("`operator_id` LIKE ?","%"+ info.Operator_id+"%")
    }
    if info.Operator_name != "" {
        db = db.Where("`operator_name` LIKE ?","%"+ info.Operator_name+"%")
    }
    if info.Status != 0 {
        db = db.Where("`status` = ?",info.Status)
    }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&opts).Error
	return err, opts, total
}
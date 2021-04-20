package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateTBllbAOITba
//@description: 创建TBllbAOITba记录
//@param: TBAT model.TBllbAOITba
//@return: err error

func CreateTBllbAOITba(TBAT model.TBllbAOITba) (err error) {
	err = global.GVA_DB_MSSQL.Create(&TBAT).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteTBllbAOITba
//@description: 删除TBllbAOITba记录
//@param: TBAT model.TBllbAOITba
//@return: err error

func DeleteTBllbAOITba(TBAT model.TBllbAOITba) (err error) {
	err = global.GVA_DB_MSSQL.Delete(&TBAT).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteTBllbAOITbaByIds
//@description: 批量删除TBllbAOITba记录
//@param: ids request.IdsReq
//@return: err error

func DeleteTBllbAOITbaByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB_MSSQL.Delete(&[]model.TBllbAOITba{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateTBllbAOITba
//@description: 更新TBllbAOITba记录
//@param: TBAT *model.TBllbAOITba
//@return: err error

func UpdateTBllbAOITba(TBAT model.TBllbAOITba) (err error) {
	err = global.GVA_DB_MSSQL.Save(&TBAT).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetTBllbAOITba
//@description: 根据id获取TBllbAOITba记录
//@param: id uint
//@return: err error, TBAT model.TBllbAOITba

func GetTBllbAOITba(id uint) (err error, TBAT model.TBllbAOITba) {
	err = global.GVA_DB_MSSQL.Where("id = ?", id).First(&TBAT).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetTBllbAOITbaInfoList
//@description: 分页获取TBllbAOITba记录
//@param: info request.TBllbAOITbaSearch
//@return: err error, list interface{}, total int64

func GetTBllbAOITbaInfoList(info request.TBllbAOITbaSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB_MSSQL.Model(&model.TBllbAOITba{})
    var TBATs []model.TBllbAOITba
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&TBATs).Error
	return err, TBATs, total
}
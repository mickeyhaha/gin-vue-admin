package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateTBllbAOIRepairTbar
//@description: 创建TBllbAOIRepairTbar记录
//@param: TBART model.TBllbAOIRepairTbar
//@return: err error

func CreateTBllbAOIRepairTbar(TBART model.TBllbAOIRepairTbar) (err error) {
	err = global.GVA_DB_MSSQL.Create(&TBART).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteTBllbAOIRepairTbar
//@description: 删除TBllbAOIRepairTbar记录
//@param: TBART model.TBllbAOIRepairTbar
//@return: err error

func DeleteTBllbAOIRepairTbar(TBART model.TBllbAOIRepairTbar) (err error) {
	err = global.GVA_DB_MSSQL.Delete(&TBART).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteTBllbAOIRepairTbarByIds
//@description: 批量删除TBllbAOIRepairTbar记录
//@param: ids request.IdsReq
//@return: err error

func DeleteTBllbAOIRepairTbarByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB_MSSQL.Delete(&[]model.TBllbAOIRepairTbar{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateTBllbAOIRepairTbar
//@description: 更新TBllbAOIRepairTbar记录
//@param: TBART *model.TBllbAOIRepairTbar
//@return: err error

func UpdateTBllbAOIRepairTbar(TBART model.TBllbAOIRepairTbar) (err error) {
	err = global.GVA_DB_MSSQL.Save(&TBART).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetTBllbAOIRepairTbar
//@description: 根据id获取TBllbAOIRepairTbar记录
//@param: id uint
//@return: err error, TBART model.TBllbAOIRepairTbar

func GetTBllbAOIRepairTbar(id uint) (err error, TBART model.TBllbAOIRepairTbar) {
	err = global.GVA_DB_MSSQL.Where("id = ?", id).First(&TBART).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetTBllbAOIRepairTbarInfoList
//@description: 分页获取TBllbAOIRepairTbar记录
//@param: info request.TBllbAOIRepairTbarSearch
//@return: err error, list interface{}, total int64

func GetTBllbAOIRepairTbarInfoList(info request.TBllbAOIRepairTbarSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB_MSSQL.Model(&model.TBllbAOIRepairTbar{})
    var TBARTs []model.TBllbAOIRepairTbar
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&TBARTs).Error
	return err, TBARTs, total
}
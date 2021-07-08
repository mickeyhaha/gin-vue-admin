package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateTRSBusiProduceRecords
//@description: 创建TRSBusiProduceRecords记录
//@param: TBPR model.TRSBusiProduceRecords
//@return: err error

func CreateTRSBusiProduceRecords(TBPR model.TRSBusiProduceRecords) (err error) {
	err = global.GVA_DB.Create(&TBPR).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteTRSBusiProduceRecords
//@description: 删除TRSBusiProduceRecords记录
//@param: TBPR model.TRSBusiProduceRecords
//@return: err error

func DeleteTRSBusiProduceRecords(TBPR model.TRSBusiProduceRecords) (err error) {
	err = global.GVA_DB.Delete(&TBPR).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteTRSBusiProduceRecordsByIds
//@description: 批量删除TRSBusiProduceRecords记录
//@param: ids request.IdsReq
//@return: err error

func DeleteTRSBusiProduceRecordsByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.TRSBusiProduceRecords{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateTRSBusiProduceRecords
//@description: 更新TRSBusiProduceRecords记录
//@param: TBPR *model.TRSBusiProduceRecords
//@return: err error

func UpdateTRSBusiProduceRecords(TBPR model.TRSBusiProduceRecords) (err error) {
	err = global.GVA_DB.Save(&TBPR).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetTRSBusiProduceRecords
//@description: 根据id获取TRSBusiProduceRecords记录
//@param: id uint
//@return: err error, TBPR model.TRSBusiProduceRecords

func GetTRSBusiProduceRecords(id uint) (err error, TBPR model.TRSBusiProduceRecords) {
	err = global.GVA_DB.Where("id = ?", id).First(&TBPR).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetTRSBusiProduceRecordsInfoList
//@description: 分页获取TRSBusiProduceRecords记录
//@param: info request.TRSBusiProduceRecordsSearch
//@return: err error, list interface{}, total int64

func GetTRSBusiProduceRecordsInfoList(info request.TRSBusiProduceRecordsSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.TRSBusiProduceRecords{})
    var TBPRs []model.TRSBusiProduceRecords
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&TBPRs).Error
	return err, TBPRs, total
}
package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateTS_FCT_CNT
//@description: 创建TS_FCT_CNT记录
//@param: TFC model.TS_FCT_CNT
//@return: err error

func CreateTS_FCT_CNT(TFC model.TS_FCT_CNT) (err error) {
	err = global.GVA_DB.Create(&TFC).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteTS_FCT_CNT
//@description: 删除TS_FCT_CNT记录
//@param: TFC model.TS_FCT_CNT
//@return: err error

func DeleteTS_FCT_CNT(TFC model.TS_FCT_CNT) (err error) {
	err = global.GVA_DB.Delete(&TFC).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteTS_FCT_CNTByIds
//@description: 批量删除TS_FCT_CNT记录
//@param: ids request.IdsReq
//@return: err error

func DeleteTS_FCT_CNTByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.TS_FCT_CNT{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateTS_FCT_CNT
//@description: 更新TS_FCT_CNT记录
//@param: TFC *model.TS_FCT_CNT
//@return: err error

func UpdateTS_FCT_CNT(TFC model.TS_FCT_CNT) (err error) {
	err = global.GVA_DB.Save(&TFC).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetTS_FCT_CNT
//@description: 根据id获取TS_FCT_CNT记录
//@param: id uint
//@return: err error, TFC model.TS_FCT_CNT

func GetTS_FCT_CNT(id uint) (err error, TFC model.TS_FCT_CNT) {
	err = global.GVA_DB.Where("id = ?", id).First(&TFC).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetTS_FCT_CNTInfoList
//@description: 分页获取TS_FCT_CNT记录
//@param: info request.TS_FCT_CNTSearch
//@return: err error, list interface{}, total int64

func GetTS_FCT_CNTInfoList(info request.TS_FCT_CNTSearch) (err error, list interface{}, total int64) {
	//limit := info.PageSize
	//offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB_MSSQL.Model(&model.TS_FCT_CNT{})
    var TFCs []model.TS_FCT_CNT
    // 如果有条件搜索 下方会自动创建搜索语句
    //if info.Count != 0 {
    //    db = db.Where("`count` > ?",info.Count)
    //}
    //if info.ErrCount != 0 {
    //    db = db.Where("`ErrCount` > ?",info.ErrCount)
    //}
    //if info.LineID != 0 {
    //    db = db.Where("`LineID` = ?",info.LineID)
    //}
    //if info.OrderNo != "" {
    //    db = db.Where("`OrderNo` <> ?",info.OrderNo)
    //}
    //if info.IssueName != "" {
    //    db = db.Where("`IssueName` = ?",info.IssueName)
    //}
    //if info.Result != 0 {
    //    db = db.Where("`Result` = ?",info.Result)
    //}
    //if info.FCTID != 0 {
    //    db = db.Where("`FCTID` = ?",info.FCTID)
    //}
	//err = db.Count(&total).Error
	//err = db.Limit(limit).Offset(offset).Find(&TFCs).Error
	err = db.Raw("SELECT Count(1) as Count, SUM(CASE Result when 1 then 0 else 1 end) as ErrCount FROM TS_FCT WITH(NOLOCK) where OrderNo <> ''").Scan(&TFCs).Error
	total = int64(len(TFCs))
	return err, TFCs, total
}
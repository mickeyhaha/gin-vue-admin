package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateTS_VI_CNT
//@description: 创建TS_VI_CNT记录
//@param: TVC model.TS_VI_CNT
//@return: err error

func CreateTS_VI_CNT(TVC model.TS_VI_CNT) (err error) {
	err = global.GVA_DB.Create(&TVC).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteTS_VI_CNT
//@description: 删除TS_VI_CNT记录
//@param: TVC model.TS_VI_CNT
//@return: err error

func DeleteTS_VI_CNT(TVC model.TS_VI_CNT) (err error) {
	err = global.GVA_DB.Delete(&TVC).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteTS_VI_CNTByIds
//@description: 批量删除TS_VI_CNT记录
//@param: ids request.IdsReq
//@return: err error

func DeleteTS_VI_CNTByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.TS_VI_CNT{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateTS_VI_CNT
//@description: 更新TS_VI_CNT记录
//@param: TVC *model.TS_VI_CNT
//@return: err error

func UpdateTS_VI_CNT(TVC model.TS_VI_CNT) (err error) {
	err = global.GVA_DB.Save(&TVC).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetTS_VI_CNT
//@description: 根据id获取TS_VI_CNT记录
//@param: id uint
//@return: err error, TVC model.TS_VI_CNT

func GetTS_VI_CNT(id uint) (err error, TVC model.TS_VI_CNT) {
	err = global.GVA_DB.Where("id = ?", id).First(&TVC).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetTS_VI_CNTInfoList
//@description: 分页获取TS_VI_CNT记录
//@param: info request.TS_VI_CNTSearch
//@return: err error, list interface{}, total int64

func GetTS_VI_CNTInfoList(info request.TS_VI_CNTSearch) (err error, list interface{}, total int64) {
	//limit := info.PageSize
	//offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB_MSSQL.Model(&model.TS_VI_CNT{})
    var TVCs []model.TS_VI_CNT
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
    //if info.VIID != 0 {
    //    db = db.Where("`VIID` = ?",info.VIID)
    //}
	//err = db.Count(&total).Error
	//err = db.Limit(limit).Offset(offset).Find(&TVCs).Error
	err = db.Raw("SELECT Count(1) as Count, SUM(CASE Result when 1 then 0 else 1 end) as ErrCount FROM TS_VI WITH(NOLOCK)").Scan(&TVCs).Error
	total = int64(len(TVCs))
	return err, TVCs, total
}
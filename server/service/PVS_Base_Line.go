package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreatePVS_Base_Line
//@description: 创建PVS_Base_Line记录
//@param: PBL model.PVS_Base_Line
//@return: err error

func CreatePVS_Base_Line(PBL model.PVS_Base_Line) (err error) {
	err = global.GVA_DB.Create(&PBL).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeletePVS_Base_Line
//@description: 删除PVS_Base_Line记录
//@param: PBL model.PVS_Base_Line
//@return: err error

func DeletePVS_Base_Line(PBL model.PVS_Base_Line) (err error) {
	err = global.GVA_DB.Delete(&PBL).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeletePVS_Base_LineByIds
//@description: 批量删除PVS_Base_Line记录
//@param: ids request.IdsReq
//@return: err error

func DeletePVS_Base_LineByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.PVS_Base_Line{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdatePVS_Base_Line
//@description: 更新PVS_Base_Line记录
//@param: PBL *model.PVS_Base_Line
//@return: err error

func UpdatePVS_Base_Line(PBL model.PVS_Base_Line) (err error) {
	err = global.GVA_DB.Save(&PBL).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetPVS_Base_Line
//@description: 根据id获取PVS_Base_Line记录
//@param: id uint
//@return: err error, PBL model.PVS_Base_Line

func GetPVS_Base_Line(id uint) (err error, PBL model.PVS_Base_Line) {
	err = global.GVA_DB.Where("id = ?", id).First(&PBL).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetPVS_Base_LineInfoList
//@description: 分页获取PVS_Base_Line记录
//@param: info request.PVS_Base_LineSearch
//@return: err error, list interface{}, total int64

func GetPVS_Base_LineInfoList(info request.PVS_Base_LineSearch) (err error, list interface{}, total int64) {
	//limit := info.PageSize
	//offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB_MSSQL.Model(&model.PVS_Base_Line{})
    var PBLs []model.PVS_Base_Line
    // 如果有条件搜索 下方会自动创建搜索语句
    //if info.LineID != 0 {
    //    db = db.Where("`LineID` = ?",info.LineID)
    //}
    //if info.LineName != "" {
    //    db = db.Where("`LineName` LIKE ?","%"+ info.LineName+"%")
    //}
    //if info.LineType != "" {
    //    db = db.Where("`LineType` = ?",info.LineType)
    //}
    //if info.ReadFilePath != "" {
    //    db = db.Where("`ReadFilePath` = ?",info.ReadFilePath)
    //}
    //if info.LineName2 != "" {
    //    db = db.Where("`LineName2` LIKE ?","%"+ info.LineName2+"%")
    //}
    //if info.MOrderNo != "" {
    //    db = db.Where("`MOrderNo` LIKE ?","%"+ info.MOrderNo+"%")
    //}
    //if info.Remark != "" {
    //    db = db.Where("`Remark` LIKE ?","%"+ info.Remark+"%")
    //}
    //if info.Warnning != nil {
    //    db = db.Where("`Warnning` = ?",info.Warnning)
    //}
    //if info.WarnningStr != "" {
    //    db = db.Where("`WarnningStr` LIKE ?","%"+ info.WarnningStr+"%")
    //}
	//err = db.Count(&total).Error
	//err = db.Limit(limit).Offset(offset).Find(&PBLs).Error
	err = db.Raw("SELECT * FROM PVS_Base_Line WITH(NOLOCK) where MOrderNo != ''").Scan(&PBLs).Error
	total = int64(len(PBLs))
	return err, PBLs, total
}
package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateVM_TimeOutput
//@description: 创建VM_TimeOutput记录
//@param: VTO model.VM_TimeOutput
//@return: err error

func CreateVM_TimeOutput(VTO model.VM_TimeOutput) (err error) {
	err = global.GVA_DB.Create(&VTO).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteVM_TimeOutput
//@description: 删除VM_TimeOutput记录
//@param: VTO model.VM_TimeOutput
//@return: err error

func DeleteVM_TimeOutput(VTO model.VM_TimeOutput) (err error) {
	err = global.GVA_DB.Delete(&VTO).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteVM_TimeOutputByIds
//@description: 批量删除VM_TimeOutput记录
//@param: ids request.IdsReq
//@return: err error

func DeleteVM_TimeOutputByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.VM_TimeOutput{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateVM_TimeOutput
//@description: 更新VM_TimeOutput记录
//@param: VTO *model.VM_TimeOutput
//@return: err error

func UpdateVM_TimeOutput(VTO model.VM_TimeOutput) (err error) {
	err = global.GVA_DB.Save(&VTO).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetVM_TimeOutput
//@description: 根据id获取VM_TimeOutput记录
//@param: id uint
//@return: err error, VTO model.VM_TimeOutput

func GetVM_TimeOutput(id uint) (err error, VTO model.VM_TimeOutput) {
	err = global.GVA_DB.Where("id = ?", id).First(&VTO).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetVM_TimeOutputInfoList
//@description: 分页获取VM_TimeOutput记录
//@param: info request.VM_TimeOutputSearch
//@return: err error, list interface{}, total int64

func GetVM_TimeOutputInfoList(info request.VM_TimeOutputSearch) (err error, list interface{}, total int64) {
	//limit := info.PageSize
	//offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB_MSSQL.Model(&model.VM_TimeOutput{})
    var VTOs []model.VM_TimeOutput
    // 如果有条件搜索 下方会自动创建搜索语句
    //if info.Dest != 0 {
    //    db = db.Where("`Dest` > ?",info.Dest)
    //}
    //if info.Hour != 0 {
    //    db = db.Where("`Hour` = ?",info.Hour)
    //}
    //if !info.BeginTime.IsZero() {
    //     db = db.Where("`BeginTime` > ?",info.BeginTime)
    //}
    //if info.LineID != 0 {
    //    db = db.Where("`LineID` = ?",info.LineID)
    //}
    //if info.ManuDest != nil {
    //    db = db.Where("`ManuDest` = ?",info.ManuDest)
    //}
	//err = db.Count(&total).Error
	//err = db.Limit(limit).Offset(offset).Find(&VTOs).Error
	err = db.Raw("SELECT a.ID, a.Dest, a.Real, a.Hour, a.BeginTime, a.LineID, a.ManuDest FROM VM_TimeOutput a WITH(NOLOCK)").Scan(&VTOs).Error
	total = int64(len(VTOs))
	return err, VTOs, total
}
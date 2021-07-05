package service

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateTBllbShiftManage
//@description: 创建TBllbShiftManage记录
//@param: TSM model.TBllbShiftManage
//@return: err error

func CreateTBllbShiftManage(TSM model.TBllbShiftManage) (err error) {
	err = global.GVA_DB.Create(&TSM).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteTBllbShiftManage
//@description: 删除TBllbShiftManage记录
//@param: TSM model.TBllbShiftManage
//@return: err error

func DeleteTBllbShiftManage(TSM model.TBllbShiftManage) (err error) {
	err = global.GVA_DB.Delete(&TSM).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteTBllbShiftManageByIds
//@description: 批量删除TBllbShiftManage记录
//@param: ids request.IdsReq
//@return: err error

func DeleteTBllbShiftManageByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.TBllbShiftManage{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateTBllbShiftManage
//@description: 更新TBllbShiftManage记录
//@param: TSM *model.TBllbShiftManage
//@return: err error

func UpdateTBllbShiftManage(TSM model.TBllbShiftManage) (err error) {
	err = global.GVA_DB.Save(&TSM).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetTBllbShiftManage
//@description: 根据id获取TBllbShiftManage记录
//@param: id uint
//@return: err error, TSM model.TBllbShiftManage

func GetTBllbShiftManage(id uint) (err error, TSM model.TBllbShiftManage) {
	err = global.GVA_DB.Where("id = ?", id).First(&TSM).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetTBllbShiftManageInfoList
//@description: 分页获取TBllbShiftManage记录
//@param: info request.TBllbShiftManageSearch
//@return: err error, list interface{}, total int64

func GetTBllbShiftManageInfoList(info request.TBllbShiftManageSearch) (err error, list interface{}, total int64) {
    // 创建db
	db := global.GVA_DB.Model(&model.TBllbShiftManage{})
    var TSMs []model.TBllbShiftManage
    // 如果有条件搜索 下方会自动创建搜索语句

	sql := fmt.Sprintf(`
		select  * from T_Bllb_ShiftManage
	`)

	if info.ShiftManageName != "" && info.ShiftManageCode != "" {
		sql = fmt.Sprintf(`
		select  * from T_Bllb_ShiftManage where ShiftManageCode = '%s' and ShiftManageName = '%s'
	`, info.ShiftManageCode, info.ShiftManageName)
	}

	err = db.Raw(sql).Scan(&TSMs).Error
	total = int64(len(TSMs))
	return err, TSMs, total
}


func GetTBllbShiftManageInfoListByShift(info request.TBllbShiftManageSearch) (err error, list interface{}, total int64) {
	// 创建db
	db := global.GVA_DB.Model(&model.TBllbShiftManage{})
	var TSMs []model.TBllbShiftManage
	// 如果有条件搜索 下方会自动创建搜索语句

	sql := fmt.Sprintf(`
		select  * from T_Bllb_ShiftManage where ShiftManageCode = '%s' and ShiftManageName = '%s'
	`, info.ShiftManageCode, info.ShiftManageName)

	err = db.Raw(sql).Scan(&TSMs).Error
	total = int64(len(TSMs))
	return err, TSMs, total
}
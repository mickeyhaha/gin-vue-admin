package service

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateMoniWholeView
//@description: 创建MoniWholeView记录
//@param: MWV model.MoniWholeView
//@return: err error

func CreateMoniWholeView(MWV model.MoniWholeView) (err error) {
	err = global.GVA_DB.Create(&MWV).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteMoniWholeView
//@description: 删除MoniWholeView记录
//@param: MWV model.MoniWholeView
//@return: err error

func DeleteMoniWholeView(MWV model.MoniWholeView) (err error) {
	err = global.GVA_DB.Delete(&MWV).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteMoniWholeViewByIds
//@description: 批量删除MoniWholeView记录
//@param: ids request.IdsReq
//@return: err error

func DeleteMoniWholeViewByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.MoniWholeView{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateMoniWholeView
//@description: 更新MoniWholeView记录
//@param: MWV *model.MoniWholeView
//@return: err error

func UpdateMoniWholeView(MWV model.MoniWholeView) (err error) {
	err = global.GVA_DB.Save(&MWV).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetMoniWholeView
//@description: 根据id获取MoniWholeView记录
//@param: id uint
//@return: err error, MWV model.MoniWholeView

func GetMoniWholeView(id uint) (err error, MWV model.MoniWholeView) {
	err = global.GVA_DB.Where("id = ?", id).First(&MWV).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetMoniWholeViewInfoList
//@description: 分页获取MoniWholeView记录
//@param: info request.MoniWholeViewSearch
//@return: err error, list interface{}, total int64
func GetMoniWholeViewInfoList(info request.MoniWholeViewSearch) (err error, list interface{}, total int64) {
	//limit := info.PageSize
	//offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB_MSSQL.Model(&model.MoniWholeView{})
    var MWVs []model.MoniWholeView
    // 如果有条件搜索 下方会自动创建搜索语句
    //if info.LineID != 0 {
    //    db = db.Where("`LineID` = ?",info.LineID)
    //}
    //if info.MachineCode != 0 {
    //    db = db.Where("`MachineCode` = ?",info.MachineCode)
    //}
    //if info.TableNo != "" {
    //    db = db.Where("`TableNo` = ?",info.TableNo)
    //}
    //if info.BCT != "" {
    //    db = db.Where("`BCT` LIKE ?","%"+ info.BCT+"%")
    //}
    //if info.TrackNo != 0 {
    //    db = db.Where("`TrackNo` = ?",info.TrackNo)
    //}
    //if info.Division != 0 {
    //    db = db.Where("`division` = ?",info.Division)
    //}
    //if info.PickPos != 0 {
    //    db = db.Where("`PickPos` = ?",info.PickPos)
    //}
    //if info.FeederType != "" {
    //    db = db.Where("`feeder_type` = ?",info.FeederType)
    //}
    //if info.FeederNo != "" {
    //    db = db.Where("`FeederNo` = ?",info.FeederNo)
    //}
    //if info.MatrCode != "" {
    //    db = db.Where("`MatrCode` LIKE ?","%"+ info.MatrCode+"%")
    //}
    //if info.PartStatus != 0 {
    //    db = db.Where("`PartStatus` = ?",info.PartStatus)
    //}
    //if info.PlacedQty != 0 {
    //    db = db.Where("`PlacedQty` = ?",info.PlacedQty)
    //}
    //if info.PickError != 0 {
    //    db = db.Where("`PickError` = ?",info.PickError)
    //}
    //if info.StopTrail != nil {
    //    db = db.Where("`StopTrail` = ?",info.StopTrail)
    //}
    //if info.Scanner != "" {
    //    db = db.Where("`Scanner` = ?",info.Scanner)
    //}
	//err = db.Count(&total).Error
	//err = db.Limit(limit).Offset(offset).Find(&MWVs).Error
	err = db.Raw("SELECT top 5 ID, MachineCode, TableNo, PickPos, lefttime, MatrCode FROM MoniWholeView WITH(NOLOCK) " +
		"WHERE LineID = ? and LeftTime is not null and uselevel <> 0 Order by lefttime", "43").Scan(&MWVs).Error
	total = int64(len(MWVs))
	return err, MWVs, total
}

func GetRejectRateList(info request.MoniWholeViewSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//创建db
	db := global.GVA_DB_MSSQL.Model(&model.MoniWholeView{})
	var MWVs []model.MoniWholeView

	//sql := `
	//		SELECT top 5 ID, MachineCode 机器, TableNo 台车, PickPos 吸取位置, TotalError 总错误数, MatrCode 料号, RejectRate 抛料率
	//		  FROM MoniWholeView WITH(NOLOCK)
	//		  WHERE 1=1
	//	`
	//if info.LineID != 0 {
	//	sql += " and LineID = " + fmt.Sprintf("%d", info.LineID)
	//}
	//sql += ` Order by RejectRate desc`

	// 如果有条件搜索 下方会自动创建搜索语句
	if info.LineID != 0 {
	   db = db.Where("LineID = ?",info.LineID)
	}
	if info.MachineCode != 0 {
	   db = db.Where("MachineCode = ?",info.MachineCode)
	}
	if info.TableNo != "" {
	   db = db.Where("TableNo = ?",info.TableNo)
	}
	if info.BCT != "" {
	   db = db.Where("BCT LIKE ?","%"+ info.BCT+"%")
	}
	if info.TrackNo != 0 {
	   db = db.Where("TrackNo = ?",info.TrackNo)
	}
	if info.Division != 0 {
	   db = db.Where("division = ?",info.Division)
	}
	if info.PickPos != 0 {
	   db = db.Where("PickPos = ?",info.PickPos)
	}
	if info.FeederType != "" {
	   db = db.Where("feeder_type = ?",info.FeederType)
	}
	if info.FeederNo != "" {
	   db = db.Where("FeederNo = ?",info.FeederNo)
	}
	if info.MatrCode != "" {
	   db = db.Where("MatrCode LIKE ?","%"+ info.MatrCode+"%")
	}
	if info.PartStatus != 0 {
	   db = db.Where("PartStatus = ?",info.PartStatus)
	}
	if info.PlacedQty != 0 {
	   db = db.Where("PlacedQty = ?",info.PlacedQty)
	}
	if info.PickError != 0 {
	   db = db.Where("PickError = ?",info.PickError)
	}
	if info.StopTrail != nil {
	   db = db.Where("StopTrail = ?",info.StopTrail)
	}
	if info.Scanner != "" {
	   db = db.Where("Scanner = ?",info.Scanner)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&MWVs).Error
	//err = db.Raw(sql, "43").Scan(&MWVs).Error
	total = int64(len(MWVs))
	fmt.Printf("total: %d", total)
	return err, MWVs, total
}
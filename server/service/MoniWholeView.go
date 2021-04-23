package service

import (
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
	return err, MWVs, total
}

func GetRejectRateListByLineName(info request.MoniWholeViewSearch) (err error, list interface{}, total int64) {
	//创建db
	db := global.GVA_DB_MSSQL.Model(&model.MoniWholeView{})
	var MWVs []model.MoniWholeView

	sql := `
			SELECT ID, A.LineID, B.LineName,
				MachineCode = LTRIM(MachineCode) + '-' + LTRIM(TableNo) + '-' + LTRIM(PickPos),
				TotalError,
				MatrCode,
				RejectRate
			FROM CMES3.dbo.MoniWholeView A WITH (NOLOCK)
					 LEFT JOIN CMES3.dbo.PVS_Base_Line B
							   ON A.LineID = B.LineID
			where B.LineName = 'Line04'
			ORDER BY RejectRate DESC;
		`
	err = db.Raw(sql, info.LineName).Scan(&MWVs).Error
	total = int64(len(MWVs))
	return err, MWVs, total
}
//
//func GetRejectRateList4Chart(info request.MoniWholeViewSearch) (err error, list interface{}, total int64) {
//	if info.LineName == "" {
//		err, list, total = GetRejectRateList(info)
//	} else {
//		err, list, total = GetRejectRateListByLineName(info)
//	}
//	TACs := list.([]model.MoniWholeView)
//	var i int64
//	lines := make(map[string]struct{}, 0)
//	dateMap := make(map[string]struct{}, 0)
//	issueNames := make(map[string]struct{}, 0)
//	// line - issueName - errCount
//	lineSeries := make(map[string]map[string]int, 0)
//
//	lineArr := make([]string, 0)
//	dateArr := make([]string, 0)
//	issueNameArr := make([]string, 0)
//
//	for i=0; i<total; i++ {
//		if  _, ok := lines[TACs[i].LineName]; !ok {
//			lines[TACs[i].LineName] = struct{}{}
//			lineArr = append(lineArr, TACs[i].LineName)
//		}
//
//		dateStr :=  TACs[i].CreateTime.Format(global.DateBaseFmt)
//		if  _, ok := lines[dateStr]; !ok {
//			dateMap[dateStr] = struct{}{}
//			dateArr = append(dateArr, dateStr)
//		}
//
//		if  _, ok := issueNames[TACs[i].IssueName]; !ok {
//			issueNames[TACs[i].IssueName] = struct{}{}
//			issueNameArr = append(issueNameArr, TACs[i].IssueName)
//		}
//
//		if  _, ok := lineSeries[TACs[i].LineName]; !ok {
//			lineSeries[TACs[i].LineName] = make(map[string]int, 0)
//		}
//		lineSeries[TACs[i].LineName][TACs[i].IssueName] = TACs[i].ErrCount
//	}
//
//	series := make([]smt.Series, 0)
//	for j:=0; j < len(issueNameArr); j++ {
//		var data []float64
//		for k:=0; k < len(lineArr); k++ {
//			data = append(data, float64(lineSeries[lineArr[k]][issueNameArr[j]]) / (float64(totalCount)))
//		}
//		seri := smt.Series{
//			Name: issueNameArr[j],
//			Data: data,
//		}
//		series = append(series, seri)
//	}
//
//	chartDatas := make([]smt.ChartData, 0)
//	chartData := smt.ChartData{
//		Categories: dateArr,
//		Series: series,
//	}
//	chartDatas = append(chartDatas, chartData)
//	return err, chartDatas, total
//}

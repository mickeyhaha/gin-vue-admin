package service

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/smt"
	"strconv"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateDCSSMTConsumeAndReject
//@description: 创建DCSSMTConsumeAndReject记录
//@param: DSCR model.DCSSMTConsumeAndReject
//@return: err error

func CreateDCSSMTConsumeAndReject(DSCR model.DCSSMTConsumeAndReject) (err error) {
	err = global.GVA_DB_MSSQL.Create(&DSCR).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteDCSSMTConsumeAndReject
//@description: 删除DCSSMTConsumeAndReject记录
//@param: DSCR model.DCSSMTConsumeAndReject
//@return: err error

func DeleteDCSSMTConsumeAndReject(DSCR model.DCSSMTConsumeAndReject) (err error) {
	err = global.GVA_DB_MSSQL.Delete(&DSCR).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteDCSSMTConsumeAndRejectByIds
//@description: 批量删除DCSSMTConsumeAndReject记录
//@param: ids request.IdsReq
//@return: err error

func DeleteDCSSMTConsumeAndRejectByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB_MSSQL.Delete(&[]model.DCSSMTConsumeAndReject{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateDCSSMTConsumeAndReject
//@description: 更新DCSSMTConsumeAndReject记录
//@param: DSCR *model.DCSSMTConsumeAndReject
//@return: err error

func UpdateDCSSMTConsumeAndReject(DSCR model.DCSSMTConsumeAndReject) (err error) {
	err = global.GVA_DB_MSSQL.Save(&DSCR).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetDCSSMTConsumeAndReject
//@description: 根据id获取DCSSMTConsumeAndReject记录
//@param: id uint
//@return: err error, DSCR model.DCSSMTConsumeAndReject

func GetDCSSMTConsumeAndReject(id uint) (err error, DSCR model.DCSSMTConsumeAndReject) {
	err = global.GVA_DB_MSSQL.Where("id = ?", id).First(&DSCR).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetDCSSMTConsumeAndRejectInfoList
//@description: 分页获取DCSSMTConsumeAndReject记录
//@param: info request.DCSSMTConsumeAndRejectSearch
//@return: err error, list interface{}, total int64

func GetDCSSMTConsumeAndRejectInfoList(info request.DCSSMTConsumeAndRejectSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB_MSSQL.Model(&model.DCSSMTConsumeAndReject{})
    var DSCRs []model.DCSSMTConsumeAndReject
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.LineName != "" {
        db = db.Where("`LineName` = ?",info.LineName)
    }
    if info.MachineCode != 0 {
        db = db.Where("`MachineCode` = ?",info.MachineCode)
    }
    if info.MachineName != "" {
        db = db.Where("`MachineName` = ?",info.MachineName)
    }
    if info.Lane != 0 {
        db = db.Where("`Lane` <> ?",info.Lane)
    }
    if info.TableNo != "" {
        db = db.Where("`TableNo` = ?",info.TableNo)
    }
    if info.PickPos != 0 {
        db = db.Where("`PickPos` = ?",info.PickPos)
    }
    if info.FeederNo != "" {
        db = db.Where("`FeederNo` = ?",info.FeederNo)
    }
    if info.PickError != 0 {
        db = db.Where("`PickError` = ?",info.PickError)
    }
    if info.IdentError != 0 {
        db = db.Where("`IdentError` = ?",info.IdentError)
    }
    if info.OtherError != 0 {
        db = db.Where("`OtherError` = ?",info.OtherError)
    }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&DSCRs).Error
	return err, DSCRs, total
}


func GetDCSSMTConsumeAndRejectRateByLine(info request.DCSSMTConsumeAndRejectSearch) (err error, list interface{}, total int64) {
	// 创建db
	db := global.GVA_DB_MSSQL.Model(&model.DCSSMTConsumeAndReject{})
	var DSRs []model.DCSSMTConsumeAndReject

	sql := fmt.Sprintf(`
	     select o.LineName, cast(o.CreateTime as date) CreateTime, 
			sum(o.PickError) PickError, sum(o.IdentError) IdentError, sum(o.OtherError) OtherError, sum(o.PlacedQty) PlacedQty
			from CMES3.dbo.DCS_SMT_ConsumeAndReject o WITH(NOLOCK)
			where o.CreateTime >='%s' AND o.CreateTime <='%s'
		 	and o.LineName = '%s' group by o.LineName, cast(o.CreateTime as date)
		`, info.StartDate, info.EndDate, info.LineName)

	if info.Shift == 1 {
		sql = fmt.Sprintf(`
	     select o.LineName, cast(o.CreateTime as date) CreateTime, 
			sum(o.PickError) PickError, sum(o.IdentError) IdentError, sum(o.OtherError) OtherError, sum(o.PlacedQty) PlacedQty
			from CMES3.dbo.DCS_SMT_ConsumeAndReject o WITH(NOLOCK)
			where o.CreateTime >='%s' AND o.CreateTime <='%s' and DATENAME(hh, o.CreateTime) BETWEEN %d AND %d
		 	and o.LineName = '%s' group by o.LineName, cast(o.CreateTime as date)
		`, info.StartDate, info.EndDate, global.Shift_Day_Begin_Hour, global.Shift_Day_End_Hour, info.LineName)
	} else if info.Shift == 2 {
		sql = fmt.Sprintf(`
	     select o.LineName, cast(o.CreateTime as date) CreateTime, 
			sum(o.PickError) PickError, sum(o.IdentError) IdentError, sum(o.OtherError) OtherError, sum(o.PlacedQty) PlacedQty
			from CMES3.dbo.DCS_SMT_ConsumeAndReject o WITH(NOLOCK)
			where o.CreateTime >='%s' AND o.CreateTime <='%s' and DATENAME(hh, o.CreateTime) BETWEEN %d AND %d
		 	and o.LineName = '%s' group by o.LineName, cast(o.CreateTime as date)
		`, info.StartDate, info.EndDate, global.Shift_Night_Begin_Hour, global.Shift_Night_End_Hour, info.LineName)
	}

	err = db.Raw(sql).Scan(&DSRs).Error
	total = int64(len(DSRs))
	return err, DSRs, total
}

func GetDCSSMTConsumeAndRejectRate4Chart(info request.DCSSMTConsumeAndRejectSearch) (err error, list interface{}, total int64) {
	err, list, total = GetDCSSMTConsumeAndRejectRateByLine(info)
	entities := list.([]model.DCSSMTConsumeAndReject)
	var i int64
	lines := make(map[string]struct{}, 0)
	dateMap := make(map[string]struct{}, 0)
	seriesNameArr := make([]string, 0)
	seriesNameArr = append(seriesNameArr, "抛料率")
	// line - issueName - errCount
	lineSeries := make(map[string]map[string]float64, 0)

	lineArr := make([]string, 0)
	dateArr := make([]string, 0)

	for i=0; i<total; i++ {
		if  _, ok := lines[entities[i].LineName]; !ok {
			lines[entities[i].LineName] = struct{}{}
			lineArr = append(lineArr, entities[i].LineName)
		}

		dateStr :=  entities[i].CreateTime.Format(global.DateBaseFmt)
		if  _, ok := lines[dateStr]; !ok {
			dateMap[dateStr] = struct{}{}
			dateArr = append(dateArr, dateStr)
		}

		if  _, ok := lineSeries[entities[i].LineName]; !ok {
			lineSeries[entities[i].LineName] = make(map[string]float64, 0)
		}

		if entities[i].PlacedQty != 0 {
			rate := float64(entities[i].IdentError+entities[i].PickError+entities[i].OtherError) / float64(entities[i].PlacedQty)
			rate, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", rate), 64)

			lineSeries[entities[i].LineName]["抛料率"] = rate
		} else {
			continue
		}
	}

	series := make([]smt.Series, 0)
	for j:=0; j < len(seriesNameArr) && len(dateArr) > 0; j++ {
		var data []float64
		for k:=0; k < len(lineArr); k++ {
			data = append(data, float64(lineSeries[lineArr[k]][seriesNameArr[j]]))
		}
		seri := smt.Series{
			Name: seriesNameArr[j],
			Data: data,
		}
		series = append(series, seri)
	}

	chartDatas := make([]smt.ChartData, 0)
	chartData := smt.ChartData{
		Categories: dateArr,
		Series: series,
	}
	chartDatas = append(chartDatas, chartData)
	return err, chartDatas, total
}

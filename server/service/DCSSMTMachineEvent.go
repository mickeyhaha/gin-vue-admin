package service

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/smt"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateDCSSMTMachineEvent
//@description: 创建DCSSMTMachineEvent记录
//@param: DSME model.DCSSMTMachineEvent
//@return: err error

func CreateDCSSMTMachineEvent(DSME model.DCSSMTMachineEvent) (err error) {
	err = global.GVA_DB_MSSQL.Create(&DSME).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteDCSSMTMachineEvent
//@description: 删除DCSSMTMachineEvent记录
//@param: DSME model.DCSSMTMachineEvent
//@return: err error

func DeleteDCSSMTMachineEvent(DSME model.DCSSMTMachineEvent) (err error) {
	err = global.GVA_DB_MSSQL.Delete(&DSME).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteDCSSMTMachineEventByIds
//@description: 批量删除DCSSMTMachineEvent记录
//@param: ids request.IdsReq
//@return: err error

func DeleteDCSSMTMachineEventByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB_MSSQL.Delete(&[]model.DCSSMTMachineEvent{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateDCSSMTMachineEvent
//@description: 更新DCSSMTMachineEvent记录
//@param: DSME *model.DCSSMTMachineEvent
//@return: err error

func UpdateDCSSMTMachineEvent(DSME model.DCSSMTMachineEvent) (err error) {
	err = global.GVA_DB_MSSQL.Save(&DSME).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetDCSSMTMachineEvent
//@description: 根据id获取DCSSMTMachineEvent记录
//@param: id uint
//@return: err error, DSME model.DCSSMTMachineEvent

func GetDCSSMTMachineEvent(id uint) (err error, DSME model.DCSSMTMachineEvent) {
	err = global.GVA_DB_MSSQL.Where("id = ?", id).First(&DSME).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetDCSSMTMachineEventInfoList
//@description: 分页获取DCSSMTMachineEvent记录
//@param: info request.DCSSMTMachineEventSearch
//@return: err error, list interface{}, total int64

func GetDCSSMTMachineEventInfoList(info request.DCSSMTMachineEventSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB_MSSQL.Model(&model.DCSSMTMachineEvent{})
    var DSMEs []model.DCSSMTMachineEvent
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.LineName != "" {
        db = db.Where("`LineName` = ?",info.LineName)
    }
    if info.EventName != "" {
        db = db.Where("`EventName` = ?",info.EventName)
    }
    if info.EventRemark != "" {
        db = db.Where("`EventRemark` = ?",info.EventRemark)
    }
    if info.TableNo != "" {
        db = db.Where("`TableNo` = ?",info.TableNo)
    }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&DSMEs).Error
	return err, DSMEs, total
}

func GetDCSSMTMachineEventByRange(info request.DCSSMTMachineEventSearch) (err error, list interface{}, total int64) {
	// 创建db
	db := global.GVA_DB_MSSQL.Model(&model.DCSSMTMachineEvent{})
	var DSMEs []model.DCSSMTMachineEvent

	sql := fmt.Sprintf(`
	     select o.LineName, cast(o.CreateTime as date) CreateTime, EventName, EventRemark,
				   count(1) as Count
			from DCS_SMT_MachineEvent o WITH(NOLOCK)
			where o.CreateTime >='%s' AND o.CreateTime <='%s'
			  and o.LineName = '%s' group by  o.LineName, cast(o.CreateTime as date), EventName, EventRemark
		`, info.StartDate, info.EndDate, info.LineName)

	if info.LineName == "" {
		sql = fmt.Sprintf(`
	     select o.LineName, cast(o.CreateTime as date) CreateTime, EventName, EventRemark,
				   count(1) as Count
			from DCS_SMT_MachineEvent o WITH(NOLOCK)
			where o.CreateTime >='%s' AND o.CreateTime <='%s'
			group by  o.LineName, cast(o.CreateTime as date), EventName, EventRemark
		`, info.StartDate, info.EndDate)
	}

	if info.Shift == 1 {
		sql = fmt.Sprintf(`
	     select o.LineName, cast(o.CreateTime as date) CreateTime, EventName, EventRemark,
				   count(1) as Count
			from DCS_SMT_MachineEvent o WITH(NOLOCK)
			where o.CreateTime >='%s' AND o.CreateTime <='%s' and DATENAME(hh, o.CreateTime) BETWEEN %d AND %d
			  and o.LineName = '%s' group by  o.LineName, cast(o.CreateTime as date), EventName, EventRemark
		`, info.StartDate, info.EndDate, global.Shift_Day_Begin_Hour, global.Shift_Day_End_Hour, info.LineName)
	} else if info.Shift == 2 {
		sql = fmt.Sprintf(`
	     select o.LineName, cast(o.CreateTime as date) CreateTime, EventName, EventRemark,
				   count(1) as Count
			from DCS_SMT_MachineEvent o WITH(NOLOCK)
			where o.CreateTime >='%s' AND o.CreateTime <='%s' and DATENAME(hh, o.CreateTime) BETWEEN %d AND %d
			  and o.LineName = '%s' group by  o.LineName, cast(o.CreateTime as date), EventName, EventRemark
		`, info.StartDate, info.EndDate, global.Shift_Night_Begin_Hour, global.Shift_Night_End_Hour, info.LineName)
	}

	err = db.Raw(sql).Scan(&DSMEs).Error
	total = int64(len(DSMEs))
	return err, DSMEs, total
}

func GetDCSSMTMachineEvent4Chart(info request.DCSSMTMachineEventSearch) (err error, list interface{}, total int64) {
	err, list, total = GetDCSSMTMachineEventByRange(info)
	DSMEs := list.([]model.DCSSMTMachineEvent)
	var i int64
	lines := make(map[string]struct{}, 0)
	dateMap := make(map[string]struct{}, 0)
	issueNames := make(map[string]struct{}, 0)
	// line - issueName - errCount
	lineSeries := make(map[string]map[string]int, 0)

	lineArr := make([]string, 0)
	dateArr := make([]string, 0)
	issueNameArr := make([]string, 0)

	totalCount := 0
	for i=0; i<total; i++ {
		totalCount += DSMEs[i].Count
	}

	for i=0; i<total; i++ {
		if  _, ok := lines[DSMEs[i].LineName]; !ok {
			lines[DSMEs[i].LineName] = struct{}{}
			lineArr = append(lineArr, DSMEs[i].LineName)
		}

		dateStr :=  DSMEs[i].CreateTime.Format(global.DateBaseFmt)
		if  _, ok := lines[dateStr]; !ok {
			dateMap[dateStr] = struct{}{}
			dateArr = append(dateArr, dateStr)
		}

		if  _, ok := issueNames[DSMEs[i].EventRemark]; !ok {
			issueNames[DSMEs[i].EventRemark] = struct{}{}
			issueNameArr = append(issueNameArr, DSMEs[i].EventRemark)
		}

		if  _, ok := lineSeries[DSMEs[i].LineName]; !ok {
			lineSeries[DSMEs[i].LineName] = make(map[string]int, 0)
		}
		lineSeries[DSMEs[i].LineName][DSMEs[i].EventRemark] = DSMEs[i].Count
	}

	series := make([]smt.Series, 0)
	for j:=0; j < len(issueNameArr) && len(dateArr) > 0; j++ {
		var data []float64
		for k:=0; k < len(lineArr); k++ {
			data = append(data, float64(lineSeries[lineArr[k]][issueNameArr[j]]) / (float64(totalCount)))
		}
		seri := smt.Series{
			Name: issueNameArr[j],
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

	chartData2 := smt.ChartData{
		Categories: lineArr,
		Series: series,
	}
	chartDatas = append(chartDatas, chartData2)
	return err, chartDatas, total
}
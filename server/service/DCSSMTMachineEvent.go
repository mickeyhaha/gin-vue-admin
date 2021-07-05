package service

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/smt"
	"strings"
	"time"
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


func GetCurrShiftMachineEvents(lineName string) (err error, list []model.DCSSMTMachineEvent, total int64) {
	db := global.GVA_DB_MSSQL.Model(&model.DCSSMTMachineEvent{})

	shiftStart := time.Now().String()
	dateArr := strings.Split(shiftStart, " ")
	shiftStart = dateArr[0] + " 00:00:00"
	//TODO
	shiftStart = "2021-07-01 00:00:00"

	var DSMEs []model.DCSSMTMachineEvent
	sql := fmt.Sprintf(`
		select  * from DCS_SMT_MachineEvent where CreateTime > '%s' and EventName in('生产开始', '等待进板')
                            and LineName = '%s' order by TableNo, CreateTime asc;
		`, shiftStart, lineName)

	err = db.Raw(sql).Scan(&DSMEs).Error
	total = int64(len(DSMEs))
	return err, DSMEs, total
}


func getAvailLineMachineEvent(info request.DCSSMTMachineEventSearch) (err error, lineAvailMachineEvent map[string]model.DCSSMTMachineEvent) {
	db := global.GVA_DB_MSSQL.Model(&model.DCSSMTMachineEvent{})
	var DSMEs []model.DCSSMTMachineEvent
	sql := fmt.Sprintf(`
	     select o.LineName, max(MachineCode) MachineCode
			from DCS_SMT_MachineEvent o WITH(NOLOCK)
			where o.CreateTime >='%s' AND o.CreateTime <='%s' and o.EventName not in('等待进板','等待出板') group by o.LineName
		`, info.StartDate, info.EndDate)

	err = db.Raw(sql).Scan(&DSMEs).Error

	lineAvailMachineEvent = make(map[string]model.DCSSMTMachineEvent, 0)
	for i:=0; i<len(DSMEs); i++ {
		sql = fmt.Sprintf(`
	     select o.LineName, max(TableNo) TableNo, o.MachineCode
			from DCS_SMT_MachineEvent o WITH(NOLOCK)
			where o.CreateTime >='%s' AND o.CreateTime <='%s' and o.MachineCode = %d and o.LineName = '%s' and o.EventName not in('等待进板','等待出板') group by o.LineName, o.MachineCode
		`, info.StartDate, info.EndDate, DSMEs[i].MachineCode, DSMEs[i].LineName)
		var tmpDSMEs []model.DCSSMTMachineEvent
		err = db.Raw(sql).Scan(&tmpDSMEs).Error
		if err==nil && len(tmpDSMEs) > 0 {
			lineAvailMachineEvent[DSMEs[i].LineName] = tmpDSMEs[0]
		}
	}
	return err, lineAvailMachineEvent
}

func GetDCSSMTMachineEventByRange(info request.DCSSMTMachineEventSearch) (err error, list interface{}, total int64) {
	// 创建db
	db := global.GVA_DB_MSSQL.Model(&model.DCSSMTMachineEvent{})
	var DSMEs []model.DCSSMTMachineEvent

	sql := fmt.Sprintf(`
	     select o.LineName, cast(o.CreateTime as date) CreateTime, EventName, EventRemark,
				   count(1) as Count, MachineCode, TableNo
			from DCS_SMT_MachineEvent o WITH(NOLOCK)
			where o.CreateTime >='%s' AND o.CreateTime <='%s' and EventRemark != '生产'
			  and o.LineName = '%s' group by  o.LineName, cast(o.CreateTime as date), EventName, EventRemark, MachineCode, TableNo
		`, info.StartDate, info.EndDate, info.LineName)

	if info.LineName == "" {
		sql = fmt.Sprintf(`
	     select o.LineName, cast(o.CreateTime as date) CreateTime, EventName, EventRemark,
				   count(1) as Count, MachineCode, TableNo
			from DCS_SMT_MachineEvent o WITH(NOLOCK)
			where o.CreateTime >='%s' AND o.CreateTime <='%s' and EventRemark != '生产'
			group by  o.LineName, cast(o.CreateTime as date), EventName, EventRemark, MachineCode, TableNo
		`, info.StartDate, info.EndDate)
	}

	if info.Shift == 1 {
		sql = fmt.Sprintf(`
	     select o.LineName, cast(o.CreateTime as date) CreateTime, EventName, EventRemark,
				   count(1) as Count, MachineCode, TableNo
			from DCS_SMT_MachineEvent o WITH(NOLOCK)
			where o.CreateTime >='%s' AND o.CreateTime <='%s' and DATENAME(hh, o.CreateTime) BETWEEN %d AND %d  and EventRemark != '生产'
			  and o.LineName = '%s' group by  o.LineName, cast(o.CreateTime as date), EventName, EventRemark, MachineCode, TableNo
		`, info.StartDate, info.EndDate, global.Shift_Day_Begin_Hour, global.Shift_Day_End_Hour, info.LineName)
	} else if info.Shift == 2 {
		sql = fmt.Sprintf(`
	     select o.LineName, cast(o.CreateTime as date) CreateTime, EventName, EventRemark,
				   count(1) as Count, MachineCode, TableNo
			from DCS_SMT_MachineEvent o WITH(NOLOCK)
			where o.CreateTime >='%s' AND o.CreateTime <='%s' and DATENAME(hh, o.CreateTime) BETWEEN %d AND %d  and EventRemark != '生产'
			  and o.LineName = '%s' group by  o.LineName, cast(o.CreateTime as date), EventName, EventRemark, MachineCode, TableNo
		`, info.StartDate, info.EndDate, global.Shift_Night_Begin_Hour, global.Shift_Night_End_Hour, info.LineName)
	}

	err = db.Raw(sql).Scan(&DSMEs).Error
	total = int64(len(DSMEs))
	return err, DSMEs, total
}


func GetRuntimeByRangeLine(info request.DCSSMTMachineEventSearch) (err error, list []model.DCSSMTMachineEvent, total int64) {
	// 创建db
	db := global.GVA_DB_MSSQL.Model(&model.DCSSMTMachineEvent{})
	var DSMEs []model.DCSSMTMachineEvent

	sql := fmt.Sprintf(`
			select o.LineName, MachineCode, MachineName, EventName, EventRemark, CreateTime, TableNo
			from DCS_SMT_MachineEvent o WITH(NOLOCK)
			where o.CreateTime >='%s' AND o.CreateTime <='%s'
			  and o.LineName = '%s' and o.MachineCode = %d and o.TableNo = '%s'  order by o.CreateTime asc
		`, info.StartDate, info.EndDate, info.LineName, info.MachineCode, info.TableNo)

	if info.Shift == 1 {
		sql = fmt.Sprintf(`
			select o.LineName, MachineCode, MachineName, EventName, EventRemark, CreateTime, TableNo
			from DCS_SMT_MachineEvent o WITH(NOLOCK)
			where o.CreateTime >='%s' AND o.CreateTime <='%s' and DATENAME(hh, o.CreateTime) BETWEEN %d AND %d
			  and o.LineName = '%s' and o.MachineCode = %d and o.TableNo = '%s'  order by o.CreateTime asc
		`, info.StartDate, info.EndDate, global.Shift_Day_Begin_Hour, global.Shift_Day_End_Hour, info.LineName, info.MachineCode, info.TableNo)
	} else if info.Shift == 2 {
		sql = fmt.Sprintf(`
			select o.LineName, MachineCode, MachineName, EventName, EventRemark, CreateTime, TableNo
			from DCS_SMT_MachineEvent o WITH(NOLOCK)
			where o.CreateTime >='%s' AND o.CreateTime <='%s' and DATENAME(hh, o.CreateTime) BETWEEN %d AND %d
			  and o.LineName = '%s' and o.MachineCode = %d and o.TableNo = '%s'  order by o.CreateTime asc
		`, info.StartDate, info.EndDate, global.Shift_Night_Begin_Hour, global.Shift_Night_End_Hour, info.LineName, info.MachineCode, info.TableNo)
	}

	err = db.Raw(sql).Scan(&DSMEs).Error
	total = int64(len(DSMEs))
	return err, DSMEs, total
}

// 停机时间
func GetRuntimeByEvent4ChartDash(info request.DCSSMTMachineEventSearch) (err error, lineStops []model.DCSSMTRunTime, total int64) {
	err, lineMachineEvent := getAvailLineMachineEvent(info)
	if err != nil {
		return err, nil, 0
	}

	lineStops = make([]model.DCSSMTRunTime, 0)

	for lineName, machineEvent := range lineMachineEvent {
		info.LineName = lineName
		info.TableNo = machineEvent.TableNo
		info.MachineCode = machineEvent.MachineCode
		err, DSMEs, _ := GetRuntimeByRangeLine(info)
		if err==nil {
			lastStop := model.DCSSMTMachineEvent{
				EventRemark: "生产",
			}
			for i:=0; i<len(DSMEs); i++ {
				if DSMEs[i].EventRemark == "待机" {
					lastStop = DSMEs[i]
				} else if DSMEs[i].EventRemark == "生产" && lastStop.EventRemark == "待机" {
					sub := DSMEs[i].CreateTime.Sub(lastStop.CreateTime).Seconds()
					if sub >= 3 {
						oneStop := model.DCSSMTRunTime {
							CreateTime: lastStop.CreateTime,
							TimeCode: lastStop.EventName,
							TimeValue: sub/60,
							LineName: DSMEs[i].LineName,
						}
						lineStops = append(lineStops, oneStop)
					}
					lastStop.EventRemark = ""
				}
			}
		}
	}
	return nil, lineStops, int64(len(lineStops))
}

func GetDCSSMTMachineEvent4Chart(info request.DCSSMTMachineEventSearch) (err error, list interface{}, total int64) {
	err, lineMachineEvent := getAvailLineMachineEvent(info)
	if err != nil {
		return err, list, 0
	}

	err, list, total = GetDCSSMTMachineEventByRange(info)
	if err != nil {
		return err, list, 0
	}

	DSMEs := list.([]model.DCSSMTMachineEvent)
	var i int64
	lines := make(map[string]struct{}, 0)
	dateMap := make(map[string]struct{}, 0)
	issueNames := make(map[string]struct{}, 0)
	// line - issueName - date - errCount
	lineSeries := make(map[string]map[string]map[string]int, 0)

	lineArr := make([]string, 0)
	dateArr := make([]string, 0)
	issueNameArr := make([]string, 0)

	//totalCount := 0
	//for i=0; i<total; i++ {
	//	totalCount += DSMEs[i].Count
	//}

	for i=0; i<total; i++ {
		if DSMEs[i].Count == 0 {
			continue
		}

		if DSME, ok := lineMachineEvent[DSMEs[i].LineName]; !ok || DSME.MachineCode != DSMEs[i].MachineCode || DSME.TableNo != DSMEs[i].TableNo {
			continue
		}

		if  _, ok := lines[DSMEs[i].LineName]; !ok {
			lines[DSMEs[i].LineName] = struct{}{}
			lineArr = append(lineArr, DSMEs[i].LineName)
		}

		dateStr :=  DSMEs[i].CreateTime.Format(global.DateBaseFmt)
		if  _, ok := dateMap[dateStr]; !ok {
			dateMap[dateStr] = struct{}{}
			dateArr = append(dateArr, dateStr)
		}

		if  _, ok := issueNames[DSMEs[i].EventName]; !ok {
			issueNames[DSMEs[i].EventName] = struct{}{}
			issueNameArr = append(issueNameArr, DSMEs[i].EventName)
		}

		if  _, ok := lineSeries[DSMEs[i].LineName]; !ok {
			lineSeries[DSMEs[i].LineName] = make(map[string]map[string]int, 0)
		}

		if  _, ok := lineSeries[DSMEs[i].LineName][DSMEs[i].EventName]; !ok {
			lineSeries[DSMEs[i].LineName][DSMEs[i].EventName] = make(map[string]int, 0)
		}
		lineSeries[DSMEs[i].LineName][DSMEs[i].EventName][dateStr] = DSMEs[i].Count
	}

	series := make([]smt.Series, 0)
	for j:=0; j < len(issueNameArr) && len(dateArr) > 0 && len(lineArr) > 0; j++ {
		var data []float64
		for k:=0; k < len(dateArr); k++ {
			data = append(data, float64(lineSeries[lineArr[0]][issueNameArr[j]][dateArr[k]]))
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

	return err, chartDatas, total
}


func GetDCSSMTMachineEvent4ChartDash(info request.DCSSMTMachineEventSearch) (err error, list interface{}, total int64) {
	err, lineMachineEvent := getAvailLineMachineEvent(info)
	if err != nil {
		return err, list, 0
	}

	err, list, total = GetDCSSMTMachineEventByRange(info)
	if err != nil {
		return err, list, 0
	}

	DSMEs := list.([]model.DCSSMTMachineEvent)
	var i int64
	lines := make(map[string]struct{}, 0)
	dateMap := make(map[string]struct{}, 0)
	issueNames := make(map[string]struct{}, 0)
	// line - issueName - date - errCount
	lineSeries := make(map[string]map[string]map[string]int, 0)

	lineArr := make([]string, 0)
	dateArr := make([]string, 0)
	issueNameArr := make([]string, 0)

	for i=0; i<total; i++ {
		if DSMEs[i].Count == 0 {
			continue
		}

		if DSME, ok := lineMachineEvent[DSMEs[i].LineName]; !ok || DSME.MachineCode != DSMEs[i].MachineCode || DSME.TableNo != DSMEs[i].TableNo {
			continue
		}

		if  _, ok := lines[DSMEs[i].LineName]; !ok {
			lines[DSMEs[i].LineName] = struct{}{}
			lineArr = append(lineArr, DSMEs[i].LineName)
		}

		dateStr :=  DSMEs[i].CreateTime.Format(global.DateBaseFmt)
		if  _, ok := dateMap[dateStr]; !ok {
			dateMap[dateStr] = struct{}{}
			dateArr = append(dateArr, dateStr)
		}

		if  _, ok := issueNames[DSMEs[i].EventName]; !ok {
			issueNames[DSMEs[i].EventName] = struct{}{}
			issueNameArr = append(issueNameArr, DSMEs[i].EventName)
		}

		if  _, ok := lineSeries[DSMEs[i].LineName]; !ok {
			lineSeries[DSMEs[i].LineName] = make(map[string]map[string]int, 0)
		}

		if  _, ok := lineSeries[DSMEs[i].LineName][DSMEs[i].EventName]; !ok {
			lineSeries[DSMEs[i].LineName][DSMEs[i].EventName] = make(map[string]int, 0)
		}
		lineSeries[DSMEs[i].LineName][DSMEs[i].EventName][dateStr] = DSMEs[i].Count
	}

	series := make([]smt.Series, 0)
	for j:=0; j < len(issueNameArr) && len(dateArr) > 0 && len(lineArr) > 0; j++ {
		var data []float64
		for l:=0; l < len(lineArr); l++  {
			subTotal := 0
			for k:=0; k < len(dateArr); k++ {
				subTotal += lineSeries[lineArr[l]][issueNameArr[j]][dateArr[k]]
			}
			data = append(data, float64(subTotal))
		}
		seri := smt.Series{
			Name: issueNameArr[j],
			Data: data,
		}
		series = append(series, seri)
	}

	chartDatas := make([]smt.ChartData, 0)
	chartData := smt.ChartData{
		Categories: lineArr,
		Series: series,
	}
	chartDatas = append(chartDatas, chartData)

	return err, chartDatas, total
}
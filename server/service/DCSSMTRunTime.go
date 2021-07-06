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
//@function: CreateDCSSMTRunTime
//@description: 创建DCSSMTRunTime记录
//@param: DSRT model.DCSSMTRunTime
//@return: err error

func CreateDCSSMTRunTime(DSRT model.DCSSMTRunTime) (err error) {
	err = global.GVA_DB_MSSQL.Create(&DSRT).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteDCSSMTRunTime
//@description: 删除DCSSMTRunTime记录
//@param: DSRT model.DCSSMTRunTime
//@return: err error

func DeleteDCSSMTRunTime(DSRT model.DCSSMTRunTime) (err error) {
	err = global.GVA_DB_MSSQL.Delete(&DSRT).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteDCSSMTRunTimeByIds
//@description: 批量删除DCSSMTRunTime记录
//@param: ids request.IdsReq
//@return: err error

func DeleteDCSSMTRunTimeByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB_MSSQL.Delete(&[]model.DCSSMTRunTime{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateDCSSMTRunTime
//@description: 更新DCSSMTRunTime记录
//@param: DSRT *model.DCSSMTRunTime
//@return: err error

func UpdateDCSSMTRunTime(DSRT model.DCSSMTRunTime) (err error) {
	err = global.GVA_DB_MSSQL.Save(&DSRT).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetDCSSMTRunTime
//@description: 根据id获取DCSSMTRunTime记录
//@param: id uint
//@return: err error, DSRT model.DCSSMTRunTime

func GetDCSSMTRunTime(id uint) (err error, DSRT model.DCSSMTRunTime) {
	err = global.GVA_DB_MSSQL.Where("id = ?", id).First(&DSRT).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetDCSSMTRunTimeInfoList
//@description: 分页获取DCSSMTRunTime记录
//@param: info request.DCSSMTRunTimeSearch
//@return: err error, list interface{}, total int64

func GetDCSSMTRunTimeInfoList(info request.DCSSMTRunTimeSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB_MSSQL.Model(&model.DCSSMTRunTime{})
    var DSRTs []model.DCSSMTRunTime
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.LineName != "" {
        db = db.Where("`LineName` = ?",info.LineName)
    }
    if info.TimeCode != "" {
        db = db.Where("`TimeCode` LIKE ?","%"+ info.TimeCode+"%")
    }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&DSRTs).Error
	return err, DSRTs, total
}


func GetDCSSMTRunTimeByRange(info request.DCSSMTRunTimeSearch) (err error, list interface{}, total int64) {
	// 创建db
	db := global.GVA_DB_MSSQL.Model(&model.DCSSMTRunTime{})
	var DSRTs []model.DCSSMTRunTime

	sql := fmt.Sprintf(`
	     select o.LineName, cast(o.CreateTime as date) CreateTime, TimeCode,
				   sum(TimeValue) as TimeValue
			from DCS_SMT_Runtime o WITH(NOLOCK)
			where o.CreateTime >='%s' AND o.CreateTime <='%s'
			  and o.LineName = '%s' group by  o.LineName, cast(o.CreateTime as date), TimeCode
		`, info.StartDate, info.EndDate, info.LineName)

	if info.LineName == "" {
		sql = fmt.Sprintf(`
	     select o.LineName, cast(o.CreateTime as date) CreateTime, TimeCode,
				   sum(TimeValue) as TimeValue
			from DCS_SMT_Runtime o WITH(NOLOCK)
			where o.CreateTime >='%s' AND o.CreateTime <='%s'
			 group by  o.LineName, cast(o.CreateTime as date), TimeCode
		`, info.StartDate, info.EndDate)
	}

	if info.Shift == 1 {
		sql = fmt.Sprintf(`
	     select o.LineName, cast(o.CreateTime as date) CreateTime, TimeCode,
				   sum(TimeValue) as TimeValue
			from DCS_SMT_Runtime o WITH(NOLOCK)
			where o.CreateTime >='%s' AND o.CreateTime <='%s' and DATENAME(hh, o.CreateTime) BETWEEN %d AND %d
			  and o.LineName = '%s' group by  o.LineName, cast(o.CreateTime as date), TimeCode
		`, info.StartDate, info.EndDate, global.Shift_Day_Begin_Hour, global.Shift_Day_End_Hour, info.LineName)
	} else if info.Shift == 2 {
		sql = fmt.Sprintf(`
	     select o.LineName, cast(o.CreateTime as date) CreateTime, TimeCode,
				   sum(TimeValue) as TimeValue
			from DCS_SMT_Runtime o WITH(NOLOCK)
			where o.CreateTime >='%s' AND o.CreateTime <='%s' and DATENAME(hh, o.CreateTime) BETWEEN %d AND %d
			  and o.LineName = '%s' group by  o.LineName, cast(o.CreateTime as date), TimeCode
		`, info.StartDate, info.EndDate, global.Shift_Night_Begin_Hour, global.Shift_Night_End_Hour, info.LineName)
	}

	err = db.Raw(sql).Scan(&DSRTs).Error
	total = int64(len(DSRTs))
	return err, DSRTs, total
}

func GetDCSSMTRunTime4Chart(info request.DCSSMTRunTimeSearch) (err error, list interface{}, total int64) {
	//err, list, total = GetDCSSMTRunTimeByRange(info)
	// Get runtime by machine event
	err, DSMEs, total := GetRuntimeByEvent4ChartDash(request.DCSSMTMachineEventSearch{
		DCSSMTMachineEvent: model.DCSSMTMachineEvent {
			LineName: info.LineName,
			DeptFilter: info.DeptFilter,
		},
		PageInfo: info.PageInfo,
	})

	var i int64
	lines := make(map[string]struct{}, 0)
	dateMap := make(map[string]struct{}, 0)
	issueNames := make(map[string]struct{}, 0)
	// line - issueName - date - errCount
	lineSeries := make(map[string]map[string]map[string]float64, 0)

	lineArr := make([]string, 0)
	dateArr := make([]string, 0)
	issueNameArr := make([]string, 0)

	for i=0; i<total; i++ {
		if DSMEs[i].TimeValue == 0 {
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

		if  _, ok := issueNames[DSMEs[i].TimeCode]; !ok {
			issueNames[DSMEs[i].TimeCode] = struct{}{}
			issueNameArr = append(issueNameArr, DSMEs[i].TimeCode)
		}

		if  _, ok := lineSeries[DSMEs[i].LineName]; !ok {
			lineSeries[DSMEs[i].LineName] = make(map[string]map[string]float64, 0)
		}

		if  _, ok := lineSeries[DSMEs[i].LineName][DSMEs[i].TimeCode]; !ok {
			lineSeries[DSMEs[i].LineName][DSMEs[i].TimeCode] = make(map[string]float64, 0)
		}
		lineSeries[DSMEs[i].LineName][DSMEs[i].TimeCode][dateStr] = DSMEs[i].TimeValue
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

	/// 停机分布
	cateArr := make([]string, 0)
	cateArr = append(cateArr, "停机分布")

	series2 := make([]smt.PieSeries, 0)
	for j:=0; j < len(issueNameArr) && len(dateArr) > 0; j++ {
		subTotal := 0.0
		for k:=0; k < len(dateArr); k++ {
			subTotal += float64(lineSeries[lineArr[0]][issueNameArr[j]][dateArr[k]])
		}
		seri := smt.PieSeries{
			Name: issueNameArr[j],
			Data: subTotal,
		}
		series2 = append(series2, seri)
	}

	chartData2 := smt.ChartData {
		Categories: cateArr,
		Series: series2,
	}
	chartDatas = append(chartDatas, chartData2)

	return err, chartDatas, total
}

func GetDCSSMTRunTime4ChartDash(info request.DCSSMTRunTimeSearch) (err error, list interface{}, total int64) {
	//err, list, total = GetDCSSMTRunTimeByRange(info)
	//DSMEs := list.([]model.DCSSMTRunTime)
	// Get runtime by machine event
	err, DSMEs, total := GetRuntimeByEvent4ChartDash(request.DCSSMTMachineEventSearch{
		DCSSMTMachineEvent: model.DCSSMTMachineEvent {
			LineName: info.LineName,
			DeptFilter: info.DeptFilter,
		},
		PageInfo: info.PageInfo,
	})

	var i int64
	lines := make(map[string]struct{}, 0)
	dateMap := make(map[string]struct{}, 0)
	issueNames := make(map[string]struct{}, 0)
	// line - issueName - date - errCount
	lineSeries := make(map[string]map[string]map[string]float64, 0)

	lineArr := make([]string, 0)
	dateArr := make([]string, 0)
	issueNameArr := make([]string, 0)

	//totalCount := 0
	//for i=0; i<total; i++ {
	//	totalCount += DSMEs[i].Count
	//}

	stopTime := 0.0
	for i=0; i<total; i++ {
		if DSMEs[i].TimeValue == 0 {
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

		if  _, ok := issueNames[DSMEs[i].TimeCode]; !ok {
			issueNames[DSMEs[i].TimeCode] = struct{}{}
			issueNameArr = append(issueNameArr, DSMEs[i].TimeCode)
		}

		if  _, ok := lineSeries[DSMEs[i].LineName]; !ok {
			lineSeries[DSMEs[i].LineName] = make(map[string]map[string]float64, 0)
		}

		if  _, ok := lineSeries[DSMEs[i].LineName][DSMEs[i].TimeCode]; !ok {
			lineSeries[DSMEs[i].LineName][DSMEs[i].TimeCode] = make(map[string]float64, 0)
		}
		lineSeries[DSMEs[i].LineName][DSMEs[i].TimeCode][dateStr] = DSMEs[i].TimeValue
		stopTime += DSMEs[i].TimeValue
	}

	series := make([]smt.Series, 0)
	for j:=0; j < len(issueNameArr) && len(dateArr) > 0 && len(lineArr) > 0; j++ {
		var data []float64
		for l:=0; l < len(lineArr); l++  {
			subTotal := 0.0
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

	/// 停机分布
	cateArr := make([]string, 0)
	cateArr = append(cateArr, "停机分布")

	series2 := make([]smt.PieSeries, 0)
	for j:=0; j < len(issueNameArr) && len(dateArr) > 0; j++ {
		subTotal := 0.0
		for k:=0; k < len(dateArr); k++ {
			subTotal += float64(lineSeries[lineArr[0]][issueNameArr[j]][dateArr[k]])
		}
		seri := smt.PieSeries{
			Name: issueNameArr[j],
			Data: subTotal,
		}
		series2 = append(series2, seri)
	}

	chartData2 := smt.ChartData {
		Categories: cateArr,
		Series: series2,
	}
	chartDatas = append(chartDatas, chartData2)


	/// OEE = runRate * okRate
	startTime, err2 := time.Parse(global.TimeBaseFmt, info.StartDate)
	endTime, err3 := time.Parse(global.TimeBaseFmt, info.EndDate)
	if err2 != nil || err3 != nil {
		chartData3 := smt.ChartData{
			Series: smt.Series{
				Name: "OEE",
				Data: []float64{0.0},
			},
		}
		chartDatas = append(chartDatas, chartData3)
		return err, chartDatas, total
	}

	totalTime := endTime.Sub(startTime).Minutes()
	runTime := totalTime - stopTime

	runRate := (runTime - stopTime)/runTime

	// 良率
	totalErrCount := 0
	totalCount := 0
	if err, aoiInfoList, _ := GetTS_AOI_CNTInfoListByLine(GetNowShiftStartEndTime()); err == nil {
		for j := 0; j < len(aoiInfoList); j++  {
			totalErrCount += aoiInfoList[j].ErrCount
			totalCount += aoiInfoList[j].Count
		}
	}
	okRate := 1.0
	if totalCount != 0 {
		okRate = float64(totalCount - totalErrCount) / float64(totalCount)
	}

	oee := runRate * okRate * 100

	series3 := make([]smt.Series, 0)
	var datas []float64
	datas = append(datas, oee)
	serie3 := smt.Series{
		Name: "OEE",
		Data: datas,
	}
	series3 = append(series3, serie3)
	chartData3 := smt.ChartData{
		Series: series3,
	}
	chartDatas = append(chartDatas, chartData3)

	return err, chartDatas, total
}

func GetNowShiftStartEndTime() (start string, end string) {
	now := time.Now()
	nowStr := now.String()[0:strings.Index(now.String(), ".")]
	arrToday := strings.Split(nowStr, " ")

	d, _ := time.ParseDuration("-24h")
	d1 := time.Now().Add(d)
	arrYesterday := strings.Split(d1.String(), " ")

	dayStart := ""
	dayEnd := ""
	yesterDayNightStart := ""

	shift := request.TBllbShiftManageSearch{}
	shift.ShiftManageCode = "白班"
	shift.ShiftManageName = "白班"
	_, days, _ := GetTBllbShiftManageInfoListByShift(shift)
	daysEntities := days.([]model.TBllbShiftManage)
	if len(daysEntities) == 0 {
		fmt.Errorf("未配置白班信息")
		//return err, chartDatas, 0
	} else {
		dayStart = arrToday[0] + " " + daysEntities[0].StartTime + ":00"
		dayEnd = arrToday[0] + " " + daysEntities[0].EndTime + ":00"
	}

	shift.ShiftManageCode = "夜班"
	shift.ShiftManageName = "夜班"
	_, nights, _ := GetTBllbShiftManageInfoListByShift(shift)
	nightsEntities := nights.([]model.TBllbShiftManage)
	if len(nightsEntities) == 0 {
		fmt.Errorf("未配置夜班信息")
		//return err, chartDatas, 0
	} else {
		yesterDayNightStart = arrYesterday[0] + " " + nightsEntities[0].StartTime + ":00"
	}

	t1, _ := time.Parse(global.TimeBaseFmt, dayStart)
	t2, _ := time.Parse(global.TimeBaseFmt, dayEnd)
	if now.After(t2) {
		start = dayEnd
	} else if now.After(t1) {
		start = dayStart
	} else {
		start = yesterDayNightStart
	}

	end = nowStr

	return start, end
}
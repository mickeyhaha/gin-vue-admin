package service

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/smt"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateTS_AOI_CNT
//@description: 创建TS_AOI_CNT记录
//@param: TAC model.TS_AOI_CNT
//@return: err error

func CreateTS_AOI_CNT(TAC model.TS_AOI_CNT) (err error) {
	err = global.GVA_DB.Create(&TAC).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteTS_AOI_CNT
//@description: 删除TS_AOI_CNT记录
//@param: TAC model.TS_AOI_CNT
//@return: err error

func DeleteTS_AOI_CNT(TAC model.TS_AOI_CNT) (err error) {
	err = global.GVA_DB.Delete(&TAC).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteTS_AOI_CNTByIds
//@description: 批量删除TS_AOI_CNT记录
//@param: ids request.IdsReq
//@return: err error

func DeleteTS_AOI_CNTByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.TS_AOI_CNT{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateTS_AOI_CNT
//@description: 更新TS_AOI_CNT记录
//@param: TAC *model.TS_AOI_CNT
//@return: err error

func UpdateTS_AOI_CNT(TAC model.TS_AOI_CNT) (err error) {
	err = global.GVA_DB.Save(&TAC).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetTS_AOI_CNT
//@description: 根据id获取TS_AOI_CNT记录
//@param: id uint
//@return: err error, TAC model.TS_AOI_CNT

func GetTS_AOI_CNT(id uint) (err error, TAC model.TS_AOI_CNT) {
	err = global.GVA_DB.Where("id = ?", id).First(&TAC).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetTS_AOI_CNTInfoList
//@description: 分页获取TS_AOI_CNT记录
//@param: info request.TS_AOI_CNTSearch
//@return: err error, list interface{}, total int64

func GetTS_AOI_CNTInfoList(info request.TS_AOI_CNTSearch) (err error, list interface{}, total int64) {
	//limit := info.PageSize
	//offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB_MSSQL.Model(&model.TS_AOI_CNT{})
    var TACs []model.TS_AOI_CNT
	sql := `
		SELECT a.*, COUNT(1) AS ErrCount
			FROM (
					 SELECT a.IssueName, b.LineID, b.OrderNo, a.AOIID, line.LineName
					 FROM TS_AOI_Repair a WITH(NOLOCK)
							  JOIN  TS_AOI b WITH(NOLOCK)
								   ON b.ID =a.AOIID
							  JOIN  PVS_Base_Line line WITH(NOLOCK)
								   ON b.LineID = line.LineID
					 WHERE b.Result =0 AND b.OrderNo <>''
					 GROUP BY a.IssueName, a.AOIID, b.LineID, b.OrderNo, line.LineName
				 ) a
			GROUP BY a.IssueName, a.LineID, a.OrderNo, a.AOIID, a.LineName
			ORDER BY IssueName, LineName;
		`
	if info.LineName != "" {
		sql = fmt.Sprintf(`
		SELECT a.*, COUNT(1) AS ErrCount
				FROM (
						 SELECT a.IssueName, b.LineID, b.OrderNo, a.AOIID, line.LineName, cast(a.CreateTime as date) CreateTime
						 FROM TS_AOI_Repair a WITH(NOLOCK)
								  JOIN  TS_AOI b WITH(NOLOCK)
										ON b.ID =a.AOIID
								  JOIN  PVS_Base_Line line WITH(NOLOCK)
										ON b.LineID = line.LineID
						 WHERE b.Result =0 AND b.OrderNo <>'' AND line.LineName = '%s'
						   AND b.CreateTime >='%s'  AND b.CreateTime <= '%s'
						   AND a.CreateTime >='%s'  AND a.CreateTime <= '%s'
						 GROUP BY a.IssueName, a.AOIID, b.LineID, b.OrderNo, line.LineName, cast(a.CreateTime as date)
					 ) a
				GROUP BY a.IssueName, a.LineID, a.OrderNo, a.AOIID, a.LineName, a.CreateTime
				ORDER BY IssueName, LineName;
			`, info.LineName, info.StartDate, info.EndDate, info.StartDate, info.EndDate)

		if info.Shift == 1 {
			sql = fmt.Sprintf(`
				SELECT a.*, COUNT(1) AS ErrCount
				FROM (
						 SELECT a.IssueName, b.LineID, b.OrderNo, a.AOIID, line.LineName, cast(a.CreateTime as date) CreateTime
						 FROM TS_AOI_Repair a WITH(NOLOCK)
								  JOIN  TS_AOI b WITH(NOLOCK)
										ON b.ID =a.AOIID
								  JOIN  PVS_Base_Line line WITH(NOLOCK)
										ON b.LineID = line.LineID
						 WHERE b.Result =0 AND b.OrderNo <>'' AND line.LineName = '%s'
						   AND b.CreateTime >='%s'  AND b.CreateTime <= '%s'
						   AND a.CreateTime >='%s'  AND a.CreateTime <= '%s' and DATENAME(hh, a.CreateTime) BETWEEN %d AND %d
						 GROUP BY a.IssueName, a.AOIID, b.LineID, b.OrderNo, line.LineName, cast(a.CreateTime as date)
					 ) a
				GROUP BY a.IssueName, a.LineID, a.OrderNo, a.AOIID, a.LineName, a.CreateTime
				ORDER BY IssueName, LineName;
			`, info.LineName, info.StartDate, info.EndDate, info.StartDate, info.EndDate, global.Shift_Day_Begin_Hour, global.Shift_Day_End_Hour)
		} else if info.Shift == 2 {
			sql = fmt.Sprintf(`
				SELECT a.*, COUNT(1) AS ErrCount
				FROM (
						 SELECT a.IssueName, b.LineID, b.OrderNo, a.AOIID, line.LineName, cast(a.CreateTime as date) CreateTime
						 FROM TS_AOI_Repair a WITH(NOLOCK)
								  JOIN  TS_AOI b WITH(NOLOCK)
										ON b.ID =a.AOIID
								  JOIN  PVS_Base_Line line WITH(NOLOCK)
										ON b.LineID = line.LineID
						 WHERE b.Result =0 AND b.OrderNo <>'' AND line.LineName = '%s'
						   AND b.CreateTime >='%s'  AND b.CreateTime <= '%s'
						   AND a.CreateTime >='%s'  AND a.CreateTime <= '%s' and DATENAME(hh, a.CreateTime) BETWEEN %d AND %d
						 GROUP BY a.IssueName, a.AOIID, b.LineID, b.OrderNo, line.LineName, cast(a.CreateTime as date)
					 ) a
				GROUP BY a.IssueName, a.LineID, a.OrderNo, a.AOIID, a.LineName, a.CreateTime
				ORDER BY IssueName, LineName;
			`, info.LineName, info.StartDate, info.EndDate, info.StartDate, info.EndDate, global.Shift_Night_Begin_Hour, global.Shift_Night_End_Hour)
		}
	}
	err = db.Raw(sql).Scan(&TACs).Error
	total = int64(len(TACs))
	return err, TACs, total
}
//
//func GetCountErrCount(info request.TS_AOI_CNTSearch) (err error, count int, errCount int) {
//	db := global.GVA_DB_MSSQL.Model(&model.TS_AOI_CNT{})
//	var TACs []model.TS_AOI_CNT
//	err = db.Raw(`
//		SELECT Count(1) as Count, SUM(CASE Result when 1 then 0 else 1 end) as ErrCount FROM TS_AOI WITH(NOLOCK)
//			WHERE LineID =? AND OrderNo = 'para_order' and OrderNO <> '';
//		`).Scan(&TACs).Error
//	total = int64(len(TACs))
//	return err, TACs, total
//}

func GetTS_AOI_CNTInfoList4Chart(info request.TS_AOI_CNTSearch) (err error, list interface{}, total int64) {
	// TODO
	//if true {
	//	return nil, make([]smt.ChartData, 0), 0
	//}
	err, list, total = GetTS_AOI_CNTInfoList(info)
	TACs := list.([]model.TS_AOI_CNT)
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
		totalCount += TACs[i].ErrCount
	}

	for i=0; i<total; i++ {
		if  _, ok := lines[TACs[i].LineName]; !ok {
			lines[TACs[i].LineName] = struct{}{}
			lineArr = append(lineArr, TACs[i].LineName)
		}

		dateStr :=  TACs[i].CreateTime.Format(global.DateBaseFmt)
		if  _, ok := lines[dateStr]; !ok {
			dateMap[dateStr] = struct{}{}
			dateArr = append(dateArr, dateStr)
		}

		if  _, ok := issueNames[TACs[i].IssueName]; !ok {
			issueNames[TACs[i].IssueName] = struct{}{}
			issueNameArr = append(issueNameArr, TACs[i].IssueName)
		}

		if  _, ok := lineSeries[TACs[i].LineName]; !ok {
			lineSeries[TACs[i].LineName] = make(map[string]int, 0)
		}
		lineSeries[TACs[i].LineName][TACs[i].IssueName] = TACs[i].ErrCount
	}

	series := make([]smt.Series, 0)
	for j:=0; j < len(issueNameArr); j++ {
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

func GetTS_AOI_CNTInfoListByLineName(lineID string) (err error, list []model.TS_AOI_CNT, total int64) {
	db := global.GVA_DB_MSSQL.Model(&model.TS_AOI_CNT{})
	var TACs []model.TS_AOI_CNT
	err = db.Raw(`
			SELECT Count(1) as Count, SUM(CASE Result when 1 then 0 else 1 end) as ErrCount FROM TS_AOI WITH(NOLOCK) 
			WHERE LineID = ? and OrderNO <> ''
		`, lineID).Scan(&TACs).Error
	total = int64(len(TACs))
	return err, TACs, total
}
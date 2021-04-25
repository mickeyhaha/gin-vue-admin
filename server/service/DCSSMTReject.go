package service

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/smt"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateDCSSMTReject
//@description: 创建DCSSMTReject记录
//@param: DSR model.DCSSMTReject
//@return: err error

func CreateDCSSMTReject(DSR model.DCSSMTReject) (err error) {
	err = global.GVA_DB_MSSQL.Create(&DSR).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteDCSSMTReject
//@description: 删除DCSSMTReject记录
//@param: DSR model.DCSSMTReject
//@return: err error

func DeleteDCSSMTReject(DSR model.DCSSMTReject) (err error) {
	err = global.GVA_DB_MSSQL.Delete(&DSR).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteDCSSMTRejectByIds
//@description: 批量删除DCSSMTReject记录
//@param: ids request.IdsReq
//@return: err error

func DeleteDCSSMTRejectByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB_MSSQL.Delete(&[]model.DCSSMTReject{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateDCSSMTReject
//@description: 更新DCSSMTReject记录
//@param: DSR *model.DCSSMTReject
//@return: err error

func UpdateDCSSMTReject(DSR model.DCSSMTReject) (err error) {
	err = global.GVA_DB_MSSQL.Save(&DSR).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetDCSSMTReject
//@description: 根据id获取DCSSMTReject记录
//@param: id uint
//@return: err error, DSR model.DCSSMTReject

func GetDCSSMTReject(id uint) (err error, DSR model.DCSSMTReject) {
	err = global.GVA_DB_MSSQL.Where("id = ?", id).First(&DSR).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetDCSSMTRejectInfoList
//@description: 分页获取DCSSMTReject记录
//@param: info request.DCSSMTRejectSearch
//@return: err error, list interface{}, total int64

func GetDCSSMTRejectInfoList(info request.DCSSMTRejectSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB_MSSQL.Model(&model.DCSSMTReject{})
    var DSRs []model.DCSSMTReject
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
        db = db.Where("`Lane` = ?",info.Lane)
    }
    if info.TableNo != "" {
        db = db.Where("`TableNo` = ?",info.TableNo)
    }
    if info.PickPos != 0 {
        db = db.Where("`PickPos` = ?",info.PickPos)
    }
    if info.SlotNo != "" {
        db = db.Where("`SlotNo` = ?",info.SlotNo)
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
    if !info.CreateTime.IsZero() {
         db = db.Where("`CreateTime` = ?",info.CreateTime)
    }
    if info.Remark != "" {
        db = db.Where("`Remark` = ?",info.Remark)
    }
    if info.FeederNo != "" {
        db = db.Where("`FeederNo` = ?",info.FeederNo)
    }
    if info.MatrCode != "" {
        db = db.Where("`MatrCode` = ?",info.MatrCode)
    }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&DSRs).Error
	return err, DSRs, total
}


func GetDCSSMTRejectInfoListByLine(info request.DCSSMTRejectSearch) (err error, list interface{}, total int64) {
	// 创建db
	db := global.GVA_DB_MSSQL.Model(&model.DCSSMTReject{})
	var DSRs []model.DCSSMTReject

	sql := fmt.Sprintf(`
	     select o.LineName, cast(o.AddTime as date) AddTime, 
			sum(o.PickError) PickError, sum(o.IdentError) IdentError, sum(o.OtherError) OtherError
			from CMES3.dbo.DCS_SMT_Reject o WITH(NOLOCK)
			where o.AddTime >='%s' AND o.AddTime <='%s'
		 	and o.LineName = '%s' group by o.LineName, cast(o.AddTime as date)
		`, info.StartDate, info.EndDate, info.LineName)

	err = db.Raw(sql).Scan(&DSRs).Error
	total = int64(len(DSRs))
	return err, DSRs, total
}

func GetDCSSMTRejectRate4Chart(info request.DCSSMTRejectSearch) (err error, list interface{}, total int64) {
	err, list, total = GetDCSSMTRejectInfoListByLine(info)
	entities := list.([]model.DCSSMTReject)
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
		// TODO: 获取产量PlacedQty
		lineSeries[entities[i].LineName]["抛料率"] = float64(entities[i].IdentError+entities[i].PickError+entities[i].OtherError) / 1000
	}

	series := make([]smt.Series, 0)
	for j:=0; j < len(seriesNameArr); j++ {
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

package service

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/smt"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateDCSSMTOutPut
//@description: 创建DCSSMTOutPut记录
//@param: DSO model.DCSSMTOutPut
//@return: err error

func CreateDCSSMTOutPut(DSO model.DCSSMTOutPut) (err error) {
	err = global.GVA_DB_MSSQL.Create(&DSO).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteDCSSMTOutPut
//@description: 删除DCSSMTOutPut记录
//@param: DSO model.DCSSMTOutPut
//@return: err error

func DeleteDCSSMTOutPut(DSO model.DCSSMTOutPut) (err error) {
	err = global.GVA_DB_MSSQL.Delete(&DSO).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteDCSSMTOutPutByIds
//@description: 批量删除DCSSMTOutPut记录
//@param: ids request.IdsReq
//@return: err error

func DeleteDCSSMTOutPutByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB_MSSQL.Delete(&[]model.DCSSMTOutPut{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateDCSSMTOutPut
//@description: 更新DCSSMTOutPut记录
//@param: DSO *model.DCSSMTOutPut
//@return: err error

func UpdateDCSSMTOutPut(DSO model.DCSSMTOutPut) (err error) {
	err = global.GVA_DB_MSSQL.Save(&DSO).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetDCSSMTOutPut
//@description: 根据id获取DCSSMTOutPut记录
//@param: id uint
//@return: err error, DSO model.DCSSMTOutPut

func GetDCSSMTOutPut(id uint) (err error, DSO model.DCSSMTOutPut) {
	err = global.GVA_DB_MSSQL.Where("id = ?", id).First(&DSO).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetDCSSMTOutPutInfoList
//@description: 分页获取DCSSMTOutPut记录
//@param: info request.DCSSMTOutPutSearch
//@return: err error, list interface{}, total int64

func GetDCSSMTOutPutInfoList(info request.DCSSMTOutPutSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB_MSSQL.Model(&model.DCSSMTOutPut{})
    var DSOs []model.DCSSMTOutPut
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
    if info.Qty != 0 {
        db = db.Where("`Qty` = ?",info.Qty)
    }
    if !info.CreateTime.IsZero() {
         db = db.Where("`CreateTime` = ?",info.CreateTime)
    }
    if info.Remark != "" {
        db = db.Where("`Remark` = ?",info.Remark)
    }
    if info.Lane != 0 {
        db = db.Where("`Lane` = ?",info.Lane)
    }
    if !info.AddTime.IsZero() {
         db = db.Where("`AddTime` = ?",info.AddTime)
    }
    if info.TableNo != "" {
        db = db.Where("`TableNo` = ?",info.TableNo)
    }
    if info.BoardQty != 0 {
        db = db.Where("`BoardQty` = ?",info.BoardQty)
    }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&DSOs).Error
	return err, DSOs, total
}


func GetDCSSMTOutPutInfoListByLine(info request.DCSSMTOutPutSearch) (err error, list interface{}, total int64) {
	// 创建db
	db := global.GVA_DB_MSSQL.Model(&model.DCSSMTOutPut{})
	var DSOs []model.DCSSMTOutPut
	sql := "select * from DCSSMTOutPut WITH(NOLOCK) where 1=1 "
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.LineName != "" {
		sql += fmt.Sprintf(" and LineName = '%s'", info.LineName)
	}
	if !info.CreateTime.IsZero() {
		sql += fmt.Sprintf(" and CreateTime = '%s'", info.LineName)
	}
	sql += " group by cast(CreateTime as date)"
	err = db.Raw(sql).Scan(&DSOs).Error
	total = int64(len(DSOs))
	return err, DSOs, total
}


func GetDCSSMTOutPutInfoList4Chart(info request.DCSSMTOutPutSearch) (err error, list interface{}, total int64) {
	err, list, total = GetDCSSMTOutPutInfoListByLine(info)
	entities := list.([]model.DCSSMTOutPut)
	var i int64
	lines := make(map[string]struct{}, 0)
	dateMap := make(map[string]struct{}, 0)
	seriesNameArr := make([]string, 0)
	seriesNameArr = append(seriesNameArr, "标准产量")
	seriesNameArr = append(seriesNameArr, "实际产量")
	// line - issueName - errCount
	lineSeries := make(map[string]map[string]int, 0)

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
			lineSeries[entities[i].LineName] = make(map[string]int, 0)
		}
		lineSeries[entities[i].LineName]["实际产量"] = entities[i].Qty
		lineSeries[entities[i].LineName]["标准产量"] = entities[i].BoardQty
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
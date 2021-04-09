package service

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateTS_SPI_CNT
//@description: 创建TS_SPI_CNT记录
//@param: TSC model.TS_SPI_CNT
//@return: err error

func CreateTS_SPI_CNT(TSC model.TS_SPI_CNT) (err error) {
	err = global.GVA_DB.Create(&TSC).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteTS_SPI_CNT
//@description: 删除TS_SPI_CNT记录
//@param: TSC model.TS_SPI_CNT
//@return: err error

func DeleteTS_SPI_CNT(TSC model.TS_SPI_CNT) (err error) {
	err = global.GVA_DB.Delete(&TSC).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteTS_SPI_CNTByIds
//@description: 批量删除TS_SPI_CNT记录
//@param: ids request.IdsReq
//@return: err error

func DeleteTS_SPI_CNTByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.TS_SPI_CNT{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateTS_SPI_CNT
//@description: 更新TS_SPI_CNT记录
//@param: TSC *model.TS_SPI_CNT
//@return: err error

func UpdateTS_SPI_CNT(TSC model.TS_SPI_CNT) (err error) {
	err = global.GVA_DB.Save(&TSC).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetTS_SPI_CNT
//@description: 根据id获取TS_SPI_CNT记录
//@param: id uint
//@return: err error, TSC model.TS_SPI_CNT

func GetTS_SPI_CNT(id uint) (err error, TSC model.TS_SPI_CNT) {
	err = global.GVA_DB.Where("id = ?", id).First(&TSC).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetTS_SPI_CNTInfoList
//@description: 分页获取TS_SPI_CNT记录
//@param: info request.TS_SPI_CNTSearch
//@return: err error, list interface{}, total int64

func GetTS_SPI_CNTInfoList(info request.TS_SPI_CNTSearch) (err error, list interface{}, total int64) {
	//limit := info.PageSize
	//offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB_MSSQL.Model(&model.TS_SPI_CNT{})
    var TSCs []model.TS_SPI_CNT
	sql := `
			SELECT TOP 7 a.IssueName, COUNT(1) AS Num 
        FROM ( 
            SELECT a.IssueName, a.SPIID, b.LineID, b.OrderNo 
            FROM TS_SPI_Repair a WITH(NOLOCK) 
            JOIN TS_SPI b WITH(NOLOCK) 
            ON b.ID =a.SPIID 
            WHERE b.Result =0 AND b.OrderNo <>'' 
            GROUP BY a.IssueName, a.SPIID , b.LineID, b.OrderNo
          ) a 
        GROUP BY a.IssueName 
        ORDER BY Num DESC;
			`	// AND b.CreateTime >=? AND b.CreateTime <? AND a.CreateTime >=?  AND a.CreateTime <? AND b.LineID = ? AND b.OrderNo = ?

	fmt.Printf(sql)
    // 如果有条件搜索 下方会自动创建搜索语句
    //if info.Count != 0 {
    //    db = db.Where("`count` > ?",info.Count)
    //}
    //if info.ErrCount != 0 {
    //    db = db.Where("`ErrCount` > ?",info.ErrCount)
    //}
    //if info.LineID != 0 {
    //    db = db.Where("`LineID` = ?",info.LineID)
    //}
    //if info.OrderNo != "" {
    //    db = db.Where("`OrderNo` <> ?",info.OrderNo)
    //}
    //if info.IssueName != "" {
    //    db = db.Where("`IssueName` = ?",info.IssueName)
    //}
    //if info.Result != 0 {
    //    db = db.Where("`Result` = ?",info.Result)
    //}
    //if info.SPIID != 0 {
    //    db = db.Where("`SPIID` = ?",info.SPIID)
    //}
	//err = db.Count(&total).Error
	//err = db.Limit(limit).Offset(offset).Find(&TSCs).Error
	err = db.Raw("SELECT Count(1) as Count, SUM(CASE Result when 1 then 0 else 1 end) as ErrCount FROM TS_SPI WITH(NOLOCK)").Scan(&TSCs).Error
	total = int64(len(TSCs))
	return err, TSCs, total
}
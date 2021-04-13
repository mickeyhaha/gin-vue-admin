package service

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreatePUBMOrderProduce2
//@description: 创建PUBMOrderProduce2记录
//@param: PUBMOrderProduce model.PUBMOrderProduce2
//@return: err error

func CreatePUBMOrderProduce2(PUBMOrderProduce model.PUBMOrderProduce2) (err error) {
	err = global.GVA_DB.Create(&PUBMOrderProduce).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeletePUBMOrderProduce2
//@description: 删除PUBMOrderProduce2记录
//@param: PUBMOrderProduce model.PUBMOrderProduce2
//@return: err error

func DeletePUBMOrderProduce2(PUBMOrderProduce model.PUBMOrderProduce2) (err error) {
	err = global.GVA_DB.Delete(&PUBMOrderProduce).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeletePUBMOrderProduce2ByIds
//@description: 批量删除PUBMOrderProduce2记录
//@param: ids request.IdsReq
//@return: err error

func DeletePUBMOrderProduce2ByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.PUBMOrderProduce2{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdatePUBMOrderProduce2
//@description: 更新PUBMOrderProduce2记录
//@param: PUBMOrderProduce *model.PUBMOrderProduce2
//@return: err error

func UpdatePUBMOrderProduce2(PUBMOrderProduce model.PUBMOrderProduce2) (err error) {
	err = global.GVA_DB.Save(&PUBMOrderProduce).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetPUBMOrderProduce2
//@description: 根据id获取PUBMOrderProduce2记录
//@param: id uint
//@return: err error, PUBMOrderProduce model.PUBMOrderProduce2

func GetPUBMOrderProduce2(id uint) (err error, PUBMOrderProduce model.PUBMOrderProduce2) {
	err = global.GVA_DB.Where("id = ?", id).First(&PUBMOrderProduce).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetPUBMOrderProduce2InfoList
//@description: 分页获取PUBMOrderProduce2记录
//@param: info request.PUBMOrderProduce2Search
//@return: err error, list interface{}, total int64

func GetPUBMOrderProduce2InfoList(info request.PUBMOrderProduce2Search) (err error, list interface{}, total int64) {
	//limit := info.PageSize
	//offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB_MSSQL.Model(&model.PUBMOrderProduce2{})
    var PUBMOrderProduces []model.PUBMOrderProduce2
	sql := "select distinct LineName, * from PUB_MOrderProduce where 1=1 "
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Status != 0 {
       sql += " and Status = " + fmt.Sprintf("%d", info.Status)
    }
    if info.PanelType != 0 {
		sql += " and PanelType = " + fmt.Sprintf("%d", info.PanelType)
    }
	//err = db.Count(&total).Error
	//err = db.Limit(limit).Offset(offset).Find(&PUBMOrderProduces).Error

	err = db.Raw(sql).Scan(&PUBMOrderProduces).Error

	total = int64(len(PUBMOrderProduces))
	return err, PUBMOrderProduces, total
}

func GetOrderListByLineName(lineName string, status int) (err error, list interface{}, total int64)  {
	db := global.GVA_DB_MSSQL.Model(&model.PUBMOrderProduce2{})
	var PUBMOrderProduces []model.PUBMOrderProduce2

	sql := "select * from PUB_MOrderProduce WITH(NOLOCK) where 1=1 "
	if lineName != "" {
		sql += " and LineName = " + fmt.Sprintf("%s", lineName)
	}
	if status != 0 {
		sql += " and Status = " + fmt.Sprintf("%d", status)
	}
	err = db.Raw(sql).Scan(&PUBMOrderProduces).Error
	total = int64(len(PUBMOrderProduces))
	return err, PUBMOrderProduces, total
}

func GetCurrentOrderByLineName(lineName string) (err error, PMP model.PUBMOrderProduce2, total int64)  {
	err, list, total := GetOrderListByLineName(lineName, 2)
	if total > 0 {
		return err, list.([]model.PUBMOrderProduce2)[0], 1
	} else {
		return nil, model.PUBMOrderProduce2{}, 0
	}
}

func GetOrderListByLineId(lineID int, status int) (err error, list interface{}, total int64)  {
	db := global.GVA_DB_MSSQL.Model(&model.PUBMOrderProduce2{})
	var PUBMOrderProduces []model.PUBMOrderProduce2

	sql := fmt.Sprintf("select o.* from PUB_MOrderProduce o WITH(NOLOCK) join PVS_Base_line l WITH(NOLOCK) on o.LineName = l.LineName where l.LineID = %d",
		lineID)
	if status != 0 {
		sql += fmt.Sprintf(" and l.status = %d", status)
	}

	err = db.Raw(sql).Scan(&PUBMOrderProduces).Error
	total = int64(len(PUBMOrderProduces))
	return err, PUBMOrderProduces, total
}

func GetCurrentOrderByLineId(lineID int) (err error, PMP model.PUBMOrderProduce2, total int64)  {
	err, list, total := GetOrderListByLineId(lineID, 2)
	if total > 0 {
		return err, list.([]model.PUBMOrderProduce2)[0], 1
	} else {
		return nil, model.PUBMOrderProduce2{}, 0
	}
}
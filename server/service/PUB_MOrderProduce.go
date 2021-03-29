package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreatePUB_MOrderProduce
//@description: 创建PUB_MOrderProduce记录
//@param: PMOP model.PUB_MOrderProduce
//@return: err error

func CreatePUB_MOrderProduce(PMOP model.PUB_MOrderProduce) (err error) {
	err = global.GVA_DB.Create(&PMOP).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeletePUB_MOrderProduce
//@description: 删除PUB_MOrderProduce记录
//@param: PMOP model.PUB_MOrderProduce
//@return: err error

func DeletePUB_MOrderProduce(PMOP model.PUB_MOrderProduce) (err error) {
	err = global.GVA_DB.Delete(&PMOP).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeletePUB_MOrderProduceByIds
//@description: 批量删除PUB_MOrderProduce记录
//@param: ids request.IdsReq
//@return: err error

func DeletePUB_MOrderProduceByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.PUB_MOrderProduce{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdatePUB_MOrderProduce
//@description: 更新PUB_MOrderProduce记录
//@param: PMOP *model.PUB_MOrderProduce
//@return: err error

func UpdatePUB_MOrderProduce(PMOP model.PUB_MOrderProduce) (err error) {
	err = global.GVA_DB.Save(&PMOP).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetPUB_MOrderProduce
//@description: 根据id获取PUB_MOrderProduce记录
//@param: id uint
//@return: err error, PMOP model.PUB_MOrderProduce

func GetPUB_MOrderProduce(id uint) (err error, PMOP model.PUB_MOrderProduce) {
	err = global.GVA_DB.Where("id = ?", id).First(&PMOP).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetPUB_MOrderProduceInfoList
//@description: 分页获取PUB_MOrderProduce记录
//@param: info request.PUB_MOrderProduceSearch
//@return: err error, list interface{}, total int64

func GetPUB_MOrderProduceInfoList(info request.PUB_MOrderProduceSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.PUB_MOrderProduce{})
    var PMOPs []model.PUB_MOrderProduce
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.LineName != "" {
        db = db.Where("`LineName` LIKE ?","%"+ info.LineName+"%")
    }
    if info.MOrderNo != "" {
        db = db.Where("`MOrderNo` LIKE ?","%"+ info.MOrderNo+"%")
    }
    if info.WorkOrderNo != "" {
        db = db.Where("`WorkOrderNo` LIKE ?","%"+ info.WorkOrderNo+"%")
    }
    if info.Customer != "" {
        db = db.Where("`Customer` LIKE ?","%"+ info.Customer+"%")
    }
    if info.CustOrderNo != "" {
        db = db.Where("`CustOrderNo` LIKE ?","%"+ info.CustOrderNo+"%")
    }
    if info.MachineType != "" {
        db = db.Where("`MachineType` LIKE ?","%"+ info.MachineType+"%")
    }
    if info.BOMNo != "" {
        db = db.Where("`BOMNo` LIKE ?","%"+ info.BOMNo+"%")
    }
    if info.BOMVersion != "" {
        db = db.Where("`BOMVersion` LIKE ?","%"+ info.BOMVersion+"%")
    }
    if info.Product != "" {
        db = db.Where("`Product` LIKE ?","%"+ info.Product+"%")
    }
    if info.PasteSide != "" {
        db = db.Where("`PasteSide` LIKE ?","%"+ info.PasteSide+"%")
    }
    if !info.BeginTime.IsZero() {
         db = db.Where("`BeginTime` <> ?",info.BeginTime)
    }
    if info.Remark != "" {
        db = db.Where("`Remark` LIKE ?","%"+ info.Remark+"%")
    }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&PMOPs).Error
	return err, PMOPs, total
}
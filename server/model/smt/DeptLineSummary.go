package smt

import (
      "gin-vue-admin/model"
)

type DeptLineSummary struct {
      model.PUBMOrderProduce2
      LineID  int `json:"lineID" form:"lineID" gorm:"column:LineID;comment:;type:int;"`
      Count  int `json:"count" form:"count" gorm:"column:count;comment:Count;type:bigint;"`
      ErrCount  int `json:"errCount" form:"errCount" gorm:"column:ErrCount;comment:ErrCount;type:bigint;"`
      aoiErrRate  float64 `json:"aoiErrRate" form:"aoiErrRate"`
}
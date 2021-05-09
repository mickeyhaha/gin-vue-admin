package global

import (
	"time"
)

type GVA_MODEL struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	//DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type DeptFilter struct {
	StartDate  string `json:"startDate" form:"startDate"`
	EndDate  string `json:"endDate" form:"endDate"`
	Shift  int `json:"shift" form:"shift"`
	WorkOrderNo  int `json:"workOrderNo" form:"workOrderNo"`
}

package model

import (
	"gorm.io/gorm"
	"time"
)

// 公共字段
type CommonModel struct {
	ID        uint           `json:"id" gorm:"primaryKey;autoIncrement"` // 主键ID
	CreatedAt time.Time      `json:"created_at"`                         // 创建时间
	UpdatedAt time.Time      `json:"updated_at"`                         // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index;unique" json:"-"`              // 删除时间
}

func NewCommonModel() *CommonModel {
	return &CommonModel{
		CreatedAt: time.Now(),
	}
}

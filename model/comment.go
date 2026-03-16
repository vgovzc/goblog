package model

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID        uint           `json:"id" gorm:"primaryKey;autoIncrement"` // 加autoIncrement确保自增
	Content   string         `json:"content" gorm:"size:2000;not null"`  // 加not null约束
	PostID    uint           `json:"post_id" gorm:"not null"`            // 驼峰命名+not null
	UserID    uint           `json:"user_id" gorm:"not null"`            // 补充缺失的user_id字段
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`   // 自动填充创建时间
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`   // 自动填充更新时间
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	// 可选：关联关系（便于GORM关联查询）
	Post Post `gorm:"foreignKey:PostID;references:ID" json:"post,omitempty"`
	User User `gorm:"foreignKey:UserID;references:ID" json:"user,omitempty"`
}

// TableName 显式指定表名，避免歧义
func (c *Comment) TableName() string {
	return "comments"
}

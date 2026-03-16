package model

import (
	"time"

	"gorm.io/gorm"
)

// Post 对应posts表（结构体名改为单数）
type Post struct {
	ID        uint           `json:"id" gorm:"primaryKey;autoIncrement"` // 加autoIncrement
	Title     string         `json:"title" gorm:"uniqueIndex;not null;size:100"`
	Content   string         `json:"content" gorm:"size:2000"`
	UserID    uint           `json:"user_id" gorm:"not null"`          // 加not null约束
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"` // 自动填充创建时间
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"` // 自动填充更新时间
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	// 可选：关联用户（便于GORM关联查询）
	User User `gorm:"foreignKey:UserID;references:ID" json:"user"`
}

// TableName 显式指定表名
func (p *Post) TableName() string {
	return "posts"
}

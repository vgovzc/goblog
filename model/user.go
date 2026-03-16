package model

import (
	"time"

	"gorm.io/gorm"
)

// User 对应users表（结构体名改为单数）
type User struct {
	ID        uint           `json:"id" gorm:"primaryKey;autoIncrement"` // 加autoIncrement
	Username  string         `json:"username" gorm:"uniqueIndex;not null;size:50"`
	Code      string         `gorm:"column:usercode;not null" json:"code"`
	Email     string         `json:"email" gorm:"uniqueIndex;not null;size:100"`
	Password  string         `json:"-" gorm:"not null"`
	Age       uint           `json:"age,omitempty"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"` // 自动填充创建时间
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"` // 自动填充更新时间
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	// 可选：关联用户的帖子（一对多）
	// Posts []Post `gorm:"foreignKey:UserID;references:ID" json:"posts,omitempty"`
}

// TableName 显式指定表名
func (u *User) TableName() string {
	return "users"
}

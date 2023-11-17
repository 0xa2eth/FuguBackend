package secrets

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// Secret ...  一对多 has many
type Secret struct {
	gorm.Model
	SecretID int `json:"secretId,omitempty" gorm:"type:bigint"`
	AuthorID int `json:"authorId,omitempty" gorm:"type:bigint"`
	// ViewLevel 1 仅广场 2 仅洞穴 3 广场和洞穴
	ViewLevel int            `json:"viewLevel,omitempty" gorm:"type:int"`
	Timestamp int64          `json:"timestamp,omitempty" gorm:"type:bigint"`
	Views     int64          `json:"views,omitempty" gorm:"type:bigint"`
	Content   string         `json:"content,omitempty" gorm:"type:varchar(255)"`
	Images    datatypes.JSON `json:"images,omitempty" gorm:"column:images"`
	// Status 秘闻状态 平台可以将非法的帖子设置为不可见
	Status bool `json:"status,omitempty" gorm:"type:bool"`
}

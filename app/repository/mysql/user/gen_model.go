package admin

import "time"

// User 管理员表
//
//go:generate gormgen -structs User -input .
type User struct {
	Id          int32  // 主键
	Nickname    string // 昵称
	TwitterID   string
	TwitterName string
	Avatar      string    // 头像
	IsUsed      int32     // 是否启用 1:是  -1:否
	IsDeleted   int32     // 是否删除 1:是  -1:否
	CreatedAt   time.Time `gorm:"time"` // 创建时间
	UpdatedAt   time.Time `gorm:"time"` // 更新时间
	UpdatedUser string    // 更新人
}

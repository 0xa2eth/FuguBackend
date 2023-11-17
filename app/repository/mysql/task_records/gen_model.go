package task_records

import "time"

// TaskRecords
//
//go:generate gormgen -structs TaskRecords -input .
type TaskRecords struct {
	Id        int64     //
	CreatedAt time.Time `gorm:"time"` //
	UpdatedAt time.Time `gorm:"time"` //
	DeletedAt time.Time `gorm:"time"` //
	CaveId    int64     //
	UserId    int64     //
}

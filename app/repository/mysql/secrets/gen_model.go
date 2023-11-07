package secrets

import "time"

// Secrets
//
//go:generate gormgen -structs Secrets -input .
type Secrets struct {
	Id        int64     //
	CreatedAt time.Time `gorm:"time"` //
	UpdatedAt time.Time `gorm:"time"` //
	DeletedAt time.Time `gorm:"time"` //
	SecretId  int64     //
	AuthorId  int64     //
	ViewLevel int64     //
	Timestamp int64     //
	Views     int64     //
	Content   string    //
	Status    int32     //
	Secretid  int64     //
	Authorid  int64     //
	Viewlevel int64     //
}

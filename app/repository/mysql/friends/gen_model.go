package friends

import "time"

// Friends
//
//go:generate gormgen -structs Friends -input .
type Friends struct {
	Id        int64     //
	CreatedAt time.Time `gorm:"time"` //
	UpdatedAt time.Time `gorm:"time"` //
	DeletedAt time.Time `gorm:"time"` //
	BaseId    int64     //
	FriendId  int64     //
	Status    int32     //
}

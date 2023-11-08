package invite_codes

import "time"

// InviteCodes
//
//go:generate gormgen -structs InviteCodes -input .
type InviteCodes struct {
	Id        int64     //
	CreatedAt time.Time `gorm:"time"` //
	UpdatedAt time.Time `gorm:"time"` //
	DeletedAt time.Time `gorm:"time"` //
	Userid    int64     //
	Code      string    //
}

package users

import "time"

// Users
//
//go:generate gormgen -structs Users -input .
type Users struct {
	Id            int64     //
	CreatedAt     time.Time `gorm:"time"` //
	UpdatedAt     time.Time `gorm:"time"` //
	DeletedAt     time.Time `gorm:"time"` //
	UserId        int64     //
	TicketNum     int64     //
	CaveFans      int64     //
	TwitterFans   int64     //
	LastLogin     int64     //
	RegisTime     int64     //
	EarnedPoint   int64     //
	CavePoint     int64     //
	Views         int64     //
	VipLevel      int64     //
	InvitedBy     int64     //
	NickName      string    //
	Bios          string    //
	Avatar        string    //
	Address       string    //
	TwitterId     string    //
	TwitterAvatar string    //
	TwitterName   string    //
	RetweetUrl    string    //
	MintCave      int32     //
	Posts         int64     //
}

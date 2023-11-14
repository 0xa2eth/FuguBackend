package users

import "time"

// Users
//
//go:generate gormgen -structs Users -input .
type Users struct {
	Id             int64     //
	CreatedAt      time.Time `gorm:"time"` //
	UpdatedAt      time.Time `gorm:"time"` //
	DeletedAt      time.Time `gorm:"time"` //
	Userid         int64     //
	Ticketnum      int64     //
	Cavefans       int64     //
	Twitterfans    int64     //
	Lastlogin      int64     //
	Registime      int64     //
	Earnedpoint    int64     //
	Cavepoint      int64     //
	Views          int64     //
	NickName       string    //
	Bios           string    //
	Avatar         string    //
	Address        string    //
	TwitterId      string    //
	TwitterAvatar  string    //
	TwitterName    string    //
	Caveretweeturl string    //
	Mintcave       int32     //
	Nickname       string    //
	Twitterid      string    //
	Twitteravatar  string    //
	Twittername    string    //
	Tag            int64     //
	EarnedPoint    int64     //
	CavePoint      int64     //
	Invitedbycode  string    //
	Numofposts     int64
}

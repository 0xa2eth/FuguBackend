package friends

import "time"

// Friends
//
//go:generate gormgen -structs Friends -input .
type Friends struct {
	BaseId   int64 //
	FriendId int64 //
	Status   int32 //
	Baseid   int64 //
	Friendid int64 //
}

package user

// User
//
//go:generate gormgen -structs User -input .
type User struct {
	Id            int32  //
	Nftnum        int32  //
	Ftnum         int32  //
	Ticketnum     int32  //
	Nickname      string //
	Bios          string //
	Avatar        string //
	Address       string //
	Twitterid     int64  //
	Twitteravatar string //
	Twittername   string //
	Lastlogin     int64  //
	Registertime  int64  //
	Enableroom    int32  //
}

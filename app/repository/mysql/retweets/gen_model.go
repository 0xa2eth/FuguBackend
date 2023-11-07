package retweets

import "time"

// Retweets
//
//go:generate gormgen -structs Retweets -input .
type Retweets struct {
	Id             int64     //
	CreatedAt      time.Time `gorm:"time"` //
	UpdatedAt      time.Time `gorm:"time"` //
	DeletedAt      time.Time `gorm:"time"` //
	CaveId         string    //
	CaveReTweetUrl string    //
	Caveid         string    //
	Caveretweeturl string    //
}

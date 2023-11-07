package fts

import "time"

// Fts
//
//go:generate gormgen -structs Fts -input .
type Fts struct {
	Id        int64     //
	CreatedAt time.Time `gorm:"time"` //
	UpdatedAt time.Time `gorm:"time"` //
	DeletedAt time.Time `gorm:"time"` //
	UserId    int64     //
	ChainId   int64     //
	TokenNum  int64     //
	Address   string    //
	Userid    int64     //
	Chainid   int64     //
	Tokennum  int64     //
}

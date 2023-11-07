package nfts

import "time"

// Nfts
//
//go:generate gormgen -structs Nfts -input .
type Nfts struct {
	Id        int64     //
	CreatedAt time.Time `gorm:"time"` //
	UpdatedAt time.Time `gorm:"time"` //
	DeletedAt time.Time `gorm:"time"` //
	UserId    int64     //
	ChainId   int64     //
	TokenId   int64     //
	Address   string    //
	Userid    int64     //
	Chainid   int64     //
	Tokenid   int64     //
}

package secret_images

import "time"

// SecretImages
//
//go:generate gormgen -structs SecretImages -input .
type SecretImages struct {
	Id        int64     //
	CreatedAt time.Time `gorm:"time"`                       //
	UpdatedAt time.Time `gorm:"time"`                       //
	DeletedAt time.Time `gorm:"type:datetime;default:null"` //
	SecretId  int64     //
	ImageUrl  string    //
	Secretid  int64     //
	Imageurl  string    //
}

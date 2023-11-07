package pkg

import (
	"FuguBackend/config"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ScanStructs ...
// todo
func ScanStructs() []interface{} {
	return nil
}

var cfg config.Config

func Migrate() {
	Db := InitMysql()
	//structs := ScanStructs()
	Db.AutoMigrate(&User{}, &Secret{}, &SecretImage{}, &Retweet{}, &TaskRecord{}, &Friend{}, &Nft{}, &Ft{})
	//images := []SecretImage{
	//	{
	//		Model:    gorm.Model{},
	//		SecretID: 0,
	//		ImageUrl: "",
	//	},
	//}
	//Db.Create(&Secret{
	//	Model:     gorm.Model{},
	//	SecretID:  0,
	//	AuthorID:  0,
	//	ViewLevel: 0,
	//	Timestamp: 0,
	//	Views:     0,
	//	Content:   "",
	//	Images:    nil,
	//	Status:    false,
	//})
	//
	//fmt.Println(u)
	fmt.Println(" migrate success! ")
}
func InitMysql() *gorm.DB {
	//dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local",
	//	cfg.MySQL.Read.User, cfg.MySQL.Read.Password,
	//	cfg.MySQL.Read.Host, cfg.MySQL.Read.Port,
	//	cfg.MySQL.Read.Database,
	//)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local",
		"root", "123456",
		"127.0.0.1", "13306",
		"fugu",
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: false})
	if err != nil {
		panic(fmt.Sprintf("mysql connect failed...,err:%v", zap.Error(err)))

	}
	return db
}

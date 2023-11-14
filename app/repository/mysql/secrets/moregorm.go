package secrets

import (
	"FuguBackend/app/repository/mysql"
	"go.uber.org/zap"
)

type Secret struct {
	SecretID  string    `json:"secretId"`
	Publisher Publisher `json:"publisher" gorm:"embedded"`
	Images    []string  `json:"images"`
	Views     int       `json:"views"`
	Timestamp int64     `json:"timestamp"`
	Content   string    `json:"content"`
	Viewable  string    `json:"viewable"`
}
type Publisher struct {
	CaveID     string `json:"caveID"`
	CaveName   string `json:"caveName"`
	CaveBio    string `json:"caveBio"`
	CaveAvatar string `json:"caveAvatar"`
}

// func FindViewableSecrets(pageNum, pageSize int, ids []int) []Secrets {

// FindViewableSecrets ...
func FindViewableSecrets(ids []int) []Secrets {

	//pageInfo := pagination.PageHelper(pageNum, pageSize, "DESC", int(total))
	// 1,广场 2,洞穴 3,广场&洞穴
	var find []Secrets
	err := mysql.DB.Table("secrets s").
		Where("s.author_id IN ？AND ( s.view_level = ? OR s.view_level = ?)", ids, 1, 3).Order("id DESC").Find(&find).Error
	if err != nil {
		zap.L().Error("FindViewableSecrets failed ... ", zap.Error(err))
	}

	return find
}

func FindExtroVipSecrets() []Secrets {
	var vips []int
	var find []Secrets
	err := mysql.DB.Table("users").Select("id").Where("users.tag > ?", 0).Find(&vips).Error
	if err != nil {
		zap.L().Error("FindExtroVipSecrets failed ... ", zap.Error(err))
	}

	err = mysql.DB.Table("secrets s").
		Where("s.author_id IN ？", vips).Order("RAND()").Limit(1).Find(&find).Error
	if err != nil {
		zap.L().Error("FindExtroVipSecrets failed ... ", zap.Error(err))
	}

	return find
}

func FindExtroRecommendCaveSecrets(ids []int) []Secrets {

	var find []Secrets
	err := mysql.DB.Table("secrets s").
		Not("s.author_id  IN ？", ids).
		Where("s.view_level = ?", 2). // 1,广场 2,洞穴 3,广场&洞穴
		Order("RAND()").Limit(1).Find(&find).Error
	if err != nil {
		zap.L().Error("FindExtroRecommendSecrets failed ... ", zap.Error(err))
	}

	return find
}

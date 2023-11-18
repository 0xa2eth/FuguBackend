package secrets

import (
	"FuguBackend/app/repository/mysql"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Publisher struct {
	CaveID     string `json:"caveID"`
	CaveName   string `json:"caveName"`
	CaveBio    string `json:"caveBio"`
	CaveAvatar string `json:"caveAvatar"`
}

// func FindViewableSecrets(pageNum, pageSize int, ids []int) []Secrets {

// FindViewableSecrets ...
func FindViewableSecrets(ids []int) []Secret {

	//pageInfo := pagination.PageHelper(pageNum, pageSize, "DESC", int(total))
	// 1,广场 2,洞穴 3,广场&洞穴
	var find []Secret
	err := mysql.DB.Table("secrets s").
		Where("s.author_id IN ？AND ( s.view_level = ? OR s.view_level = ?)", ids, 1, 3).Order("id DESC").Find(&find).Error
	if err != nil {
		zap.L().Error("FindViewableSecrets failed ... ", zap.Error(err))
	}

	return find
}

func FindExtroVipSecrets() []Secret {
	var vips []int
	var find []Secret
	err := mysql.DB.Table("users").Select("id").Where("users.vip_level > ?", 0).Find(&vips).Error
	if err != nil {
		zap.L().Error("FindExtroVipSecrets failed ... ", zap.Error(err))

	}
	err = mysql.DB.Table("secrets s").
		Where("s.author_id IN (?)", vips).Order("s.id").Limit(1).Find(&find).Error
	if err != nil {
		zap.L().Error("FindExtroVipSecrets failed ... ", zap.Error(err))
	}

	return find
}

func FindExtroRecommendCaveSecrets(ids []int) []Secret {

	var find []Secret
	err := mysql.DB.Table("secrets s").
		Where("s.view_level = ? ", 2). // 1,广场 2,洞穴 3,广场&洞穴
		Not("s.author_id  IN (?)", ids).
		Order("s.id").Limit(1).Find(&find).Error
	if err != nil {
		zap.L().Error("FindExtroRecommendSecrets failed ... ", zap.Error(err))
	}

	return find
}
func FindSiteFriends(InnerID int) ([]int, error) {
	var siteFriendIds []int
	dberr := mysql.DB.Table("friends").
		Select("friend_id").
		Where("base_id = ? AND status = ?", InnerID, true).
		Find(&siteFriendIds).Error
	if dberr != nil {
		return nil, dberr
	}
	return siteFriendIds, nil
}
func FindSiteFriendsSecrets(friendids []int, pageNum, pageSize int) ([]Secret, error) {
	var dbsecrets []Secret
	offset := (pageNum - 1) * pageSize
	dberr := mysql.DB.Table("secrets").
		Where("author_id IN (?) ", friendids).
		Offset(offset).
		Limit(pageSize).
		Order("id DESC").Find(&dbsecrets).Error
	if dberr != nil {
		return nil, dberr
	}
	return dbsecrets, nil
}

func CreatSecret(in Secret) (int, error) {
	err := mysql.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Table("secrets").Create(&in).Error
		if err != nil {
			return err
		}
		//affected := s.db.GetDbW().Table("users").
		//	Where("id = ?", int64(data.AuthorID)).
		//	Update("posts", gorm.Expr("posts + ?", 1)).RowsAffected
		// 自增字段加1

		err = tx.Exec("UPDATE users SET posts = posts + 1 WHERE id = ?", in.AuthorID).Error
		if err != nil {
			return err
		}
		fmt.Println("affected:", tx.RowsAffected)
		//fmt.Println("affected:", affected)
		return err
	})

	return int(in.ID), err
}

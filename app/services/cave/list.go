package cave

import (
	"FuguBackend/app/pkg/core"
	"FuguBackend/pkg/hash"
)

type Secretcave struct {
	SecretID string `json:"secretId"`
	//Publisher Publisher `json:"publisher"`
	Images    []string `json:"images"`
	Views     int      `json:"views"`
	Timestamp int64    `json:"timestamp"`
	Content   string   `json:"content"`
	Viewable  bool     `json:"viewable"`
}

func (s *service) ListMySecrets(c core.Context, InnerID int, pageNum, pageSize int, hashFunc hash.Hash) (secrets []Secretcave, err error) {

	offset := (pageNum - 1) * pageSize
	err = s.db.GetDbW().Table("secrets").
		Where("secrets.author_id = ?", InnerID).
		Limit(offset).
		Offset(offset).
		Find(&secrets).Error
	if err != nil {
		return nil, err
	}

	for i, v := range secrets {
		var images []string
		err = s.db.GetDbW().Table("secret_images").
			Select("image_url").
			Where("secret_id = ?", v.SecretID).
			Find(&images).Error
		if err != nil {
			return nil, err
		}
		secrets[i].Images = images

	}

	return
}
func (s *service) ListCaveSecrets(c core.Context, InnerID,
	caveID int, pageNum, pageSize int, hashFunc hash.Hash) ([]Secretcave, error) {
	//err := s.db.GetDbW().Table("secret_images").
	//	Select("image_url").
	//	Where("secret_id = ?", v.SecretID).
	//	Find(&images).Error
	return nil, nil
}

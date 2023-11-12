package secret

import (
	"FuguBackend/app/pkg/core"
	"FuguBackend/app/repository/mysql/secrets"
)

// Secret ...  一对多 has many
type CreateSecretData struct {
	SecretID int `json:"secretId,omitempty" gorm:"column:secretid;type:bigint"`
	AuthorID int `json:"authorId,omitempty" gorm:"column:authorid;type:bigint"`
	// ViewLevel 1 仅广场 2 仅洞穴 3 广场和洞穴
	ViewLevel int      `json:"viewLevel,omitempty" gorm:"column:viewlevel;type:int"`
	Timestamp int64    `json:"timestamp,omitempty" gorm:"column:timestamp;type:bigint"`
	Views     int64    `json:"views,omitempty" gorm:"column:views;type:bigint"`
	Content   string   `json:"content,omitempty" gorm:"column:content;type:varchar(255)"`
	Images    []string `json:"images,omitempty" gorm:"foreignKey:SecretID"`
	// Status 秘闻状态 平台可以将非法的帖子设置为不可见
	Status bool `json:"status,omitempty" gorm:"column:status;type:tinyint"`
}
type SecretImage struct {
	SecretID int    `json:"SecretID,omitempty" gorm:"column:secretid;type:bigint"`
	ImageUrl string `json:"ImageUrl,omitempty" gorm:"column:imageurl;type:varchar(255)"`
}

func (s *service) Create(c core.Context, data CreateSecretData) error {
	model := secrets.NewModel()

	id, err = model.Create(s.db.GetDbW().WithContext(ctx.RequestContext()))
	if err != nil {
		return 0, err
	}
	return

}

// parseContent 找内容中的被@的推特用户
func parseContent() {

}

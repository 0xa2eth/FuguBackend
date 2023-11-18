package secret

import (
	"FuguBackend/app/pkg/core"
	"FuguBackend/app/pkg/pagination"
	repo "FuguBackend/app/repository/mysql/secrets"
	"FuguBackend/app/services/user"
	"FuguBackend/pkg/hash"
	"encoding/json"
	"github.com/spf13/cast"
)

type SearchOneData struct {
	Id        int // 用户ID
	TwitterID string
}

type ResData struct {
	SecretsList []SecretRes `json:"secretsList"`
	Extro       []SecretRes `json:"extro"`
}
type SecretRes struct {
	SecretID  string    `json:"secretId"`
	Publisher Publisher `json:"publisher"`
	Images    []string  `json:"images"`
	Views     int       `json:"views"`
	Timestamp int64     `json:"timestamp"`
	Content   string    `json:"content"`
	Viewable  bool      `json:"viewable"`
}
type Publisher struct {
	CaveID     string `json:"caveID"`
	CaveName   string `json:"caveName"`
	CaveBio    string `json:"caveBio"`
	CaveAvatar string `json:"caveAvatar"`
}

//// Secret ...  一对多 has many
//type Secret struct {
//	gorm.Model
//	SecretID int `json:"secretId,omitempty" gorm:"type:bigint"`
//	AuthorID int `json:"authorId,omitempty" gorm:"type:bigint"`
//	// ViewLevel 1 仅广场 2 仅洞穴 3 广场和洞穴
//	ViewLevel int            `json:"viewLevel,omitempty" gorm:"type:int"`
//	Timestamp int64          `json:"timestamp,omitempty" gorm:"type:bigint"`
//	Views     int64          `json:"views,omitempty" gorm:"type:bigint"`
//	Content   string         `json:"content,omitempty" gorm:"type:varchar(255)"`
//	Images    datatypes.JSON `json:"images,omitempty" gorm:"column:images"`
//	// Status 秘闻状态 平台可以将非法的帖子设置为不可见
//	Status bool `json:"status,omitempty" gorm:"type:bool"`
//}

// List  先拿id  再刷新下朋友圈 于上次存cache的对比 找到新增的和去处的再根据这些去拿可见的秘密
func (s *service) List(c core.Context, InnerID int, pageNum, pageSize int, hashFunc hash.Hash) (pagination.PageInfo, error) {

	friendids, err := repo.FindSiteFriends(InnerID)
	if err != nil {
		return pagination.PageInfo{}, err
	}
	secrets, err := repo.FindSiteFriendsSecrets(friendids, pageNum, pageSize)
	if err != nil {
		return pagination.PageInfo{}, err
	}
	NormalSecretsList := s.BuildNormalSecretsRes(c, secrets, hashFunc, true)
	extra := s.GetExtro(c, friendids, hashFunc, InnerID)
	var Response ResData
	Response.SecretsList = NormalSecretsList
	Response.Extro = extra

	helper := pagination.PageHelper(pageNum, pageSize, "DESC", len(secrets))
	helper.Data = Response

	return helper, err

}
func (s *service) BuildNormalSecretsRes(c core.Context, raw []repo.Secret, hashFunc hash.Hash, viewAble bool) []SecretRes {
	userSvc := user.New(s.db, s.cache, s.logger, s.twSvc)
	var SecretResList []SecretRes
	for _, secret := range raw {
		var images []string
		_ = json.Unmarshal(secret.Images, &images)
		encode, _ := hashFunc.HashidsEncode([]int{secret.SecretID})
		res := SecretRes{
			SecretID:  encode,
			Publisher: Publisher{},
			Images:    images,
			Views:     int(secret.Views),
			Timestamp: secret.Timestamp,
			Content:   secret.Content,
			Viewable:  viewAble,
		}
		person, _ := userSvc.Detail(c, &user.SearchOneData{
			Id: int(secret.AuthorID),
		})
		hashId, _ := hashFunc.HashidsEncode([]int{cast.ToInt(person.Id)})
		res.Publisher.CaveID = hashId
		res.Publisher.CaveBio = person.Bios
		res.Publisher.CaveAvatar = person.Avatar
		res.Publisher.CaveName = person.NickName
		SecretResList = append(SecretResList, res)
	}

	return SecretResList
}

func (s *service) GetExtro(c core.Context, ids []int, hashFunc hash.Hash, InnerID int) []SecretRes {

	//viewableSecretsraw := secrets.FindViewableSecrets(ids)
	//viewableSecrets := s.BuildNormalSecretsRes(c, viewableSecretsraw, hashFunc, true)
	var res []SecretRes
	recommendCaveSecretsraw := repo.FindExtroRecommendCaveSecrets(ids)

	recommendCaveSecrets := s.BuildNormalSecretsRes(c, recommendCaveSecretsraw, hashFunc, false)

	vipSecretsraw := repo.FindExtroVipSecrets()
	vipSecrets := s.BuildNormalSecretsRes(c, vipSecretsraw, hashFunc, true)
	res = append(res, recommendCaveSecrets...)
	res = append(res, vipSecrets...)
	//var res ResData
	//if viewableSecrets != nil {
	//	pageinfo = pagination.PageHelper(pageNum, pageSize, "DESC", len(viewableSecrets))
	//
	//	if len(viewableSecrets)-pageinfo.Offset < pageSize {
	//		res.SecretsList = viewableSecrets[pageinfo.Offset : pageinfo.Offset+(len(viewableSecrets)-pageinfo.Offset)]
	//		res.Extro = append(res.Extro, recommendCaveSecrets...)
	//		res.Extro = append(res.Extro, vipSecrets...)
	//		pageinfo.Data = res
	//	} else {
	//		res.SecretsList = viewableSecrets[pageinfo.Offset : pageinfo.Offset+pageSize]
	//		res.Extro = append(res.Extro, recommendCaveSecrets...)
	//		res.Extro = append(res.Extro, vipSecrets...)
	//		pageinfo.Data = res
	//	}
	//
	//} else {
	//	pageinfo = pagination.PageHelper(pageNum, pageSize, "DESC", 0)
	//
	//}
	return res
}

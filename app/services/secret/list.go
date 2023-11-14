package secret

import (
	"FuguBackend/app/pkg/core"
	"FuguBackend/app/pkg/pagination"
	"FuguBackend/app/repository/mysql"
	"FuguBackend/app/repository/mysql/secrets"
	"FuguBackend/app/services/user"
	"FuguBackend/pkg/hash"
	"github.com/spf13/cast"
)

type SearchOneData struct {
	Id        int // 用户ID
	TwitterID string
}
type People struct {
	Id        int // 用户ID
	TwitterID string
}
type ResData struct {
	SecretsList []Secret `json:"secretsList"`
	Extro       []Secret `json:"extro"`
}
type Secret struct {
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

// List  先拿id  再刷新下朋友圈 于上次存cache的对比 找到新增的和去处的再根据这些去拿可见的秘密
func (s *service) List(c core.Context, InnerID int, pageNum, pageSize int, hashFunc hash.Hash) (pagination.PageInfo, error) {
	userSvc := user.New(s.db, s.cache, s.logger, s.twSvc)

	person, _ := userSvc.Detail(c, &user.SearchOneData{
		Id: InnerID,
	})
	all, err := userSvc.FindAll()
	if err != nil {
		return pagination.PageInfo{}, err
	}
	var people []People
	for i := range all {
		//ids = append(ids, all[i].TwitterId)
		people = append(people, People{
			Id:        int(all[i].Id),
			TwitterID: all[i].TwitterId,
		})
	}
	//diff, err2 := s.twSvc.DiffFriendCircle(c, InnerID, person.TwitterName, ids)
	//if err2 != nil {
	//	return pagination.PageInfo{}, err2
	//}
	var siteFriends []People
	circle, _ := s.twSvc.RefreshFriendCircle(c, person.TwitterName)
	for i := range circle {
		for j := range people {

			if circle[i].IDStr == people[j].TwitterID {
				siteFriends = append(siteFriends, people[j])
			}

		}
	}
	var ids []int
	for i := range siteFriends {
		ids = append(ids, siteFriends[i].Id)
	}

	var pageinfo pagination.PageInfo
	viewableSecretsraw := secrets.FindViewableSecrets(ids)
	viewableSecrets := s.BuildRes(c, viewableSecretsraw, hashFunc, true)

	recommendCaveSecretsraw := secrets.FindExtroRecommendCaveSecrets(ids)
	recommendCaveSecrets := s.BuildRes(c, recommendCaveSecretsraw, hashFunc, false)

	vipSecretsraw := secrets.FindExtroVipSecrets()
	vipSecrets := s.BuildRes(c, vipSecretsraw, hashFunc, true)

	var res ResData
	if viewableSecrets != nil {
		pageinfo = pagination.PageHelper(pageNum, pageSize, "DESC", len(viewableSecrets))

		if len(viewableSecrets)-pageinfo.Offset < pageSize {
			res.SecretsList = viewableSecrets[pageinfo.Offset : pageinfo.Offset+(len(viewableSecrets)-pageinfo.Offset)]
			res.Extro = append(res.Extro, recommendCaveSecrets...)
			res.Extro = append(res.Extro, vipSecrets...)
			pageinfo.Data = res
		} else {
			res.SecretsList = viewableSecrets[pageinfo.Offset : pageinfo.Offset+pageSize]
			res.Extro = append(res.Extro, recommendCaveSecrets...)
			res.Extro = append(res.Extro, vipSecrets...)
			pageinfo.Data = res
		}

	} else {
		pageinfo = pagination.PageHelper(pageNum, pageSize, "DESC", 0)

	}
	return pageinfo, err
}
func (s *service) BuildRes(c core.Context, raw []secrets.Secrets, hashFunc hash.Hash, viewAble bool) (list []Secret) {
	userSvc := user.New(s.db, s.cache, s.logger, s.twSvc)

	for _, data := range raw {
		var p Secret
		person, _ := userSvc.Detail(c, &user.SearchOneData{
			Id: int(data.AuthorId),
		})
		hashId, _ := hashFunc.HashidsEncode([]int{cast.ToInt(person.Id)})
		p.Publisher.CaveID = hashId
		p.Publisher.CaveBio = person.Bios
		p.Publisher.CaveAvatar = person.Avatar
		p.Publisher.CaveName = person.NickName
		var images []string
		_ = mysql.DB.Table("secret_images").
			Select("image_url").
			Where("secret_id = ?", data.Id).
			Find(&images).Error
		p.Images = images
		p.Timestamp = data.Timestamp
		hashId, _ = hashFunc.HashidsEncode([]int{cast.ToInt(data.Id)})
		//p.SecretID = strconv.FormatInt(data.SecretId, 10)
		p.SecretID = hashId
		p.Content = data.Content
		p.Viewable = viewAble
		list = append(list, p)
	}
	return
}

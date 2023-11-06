package pkg

import "gorm.io/gorm"

type Nft struct {
	gorm.Model
	UserID  int
	ChainID int
	TokenId int
	Address string
}
type Ft struct {
	gorm.Model
	UserID   int
	ChainID  int
	TokenNum int
	Address  string
}

// User ...内存对齐 省空间
type User struct {
	gorm.Model
	UserID         int `json:"userId,omitempty" gorm:"column:userid;bigint"`
	TicketNum      int `json:"ticketNum,omitempty" gorm:"column:ticketnum;bigint"`
	CaveFans       int `json:"caveFans,omitempty" gorm:"column:cavefans;bigint"`
	TwitterFans    int `json:"twitterFans,omitempty" gorm:"column:twitterfans;bigint"`
	LastLogin      int `json:"lastLogin,omitempty" gorm:"column:lastlogin;bigint"`
	RegisterTime   int `json:"registerTime,omitempty"`
	EarnedPoint    int
	CavePoint      int
	Views          int      `json:"views"`
	NickName       string   `json:"nickName,omitempty"`
	Bios           string   `json:"bios,omitempty"`
	Avatar         string   `json:"avatar,omitempty"`
	Address        string   `json:"address,omitempty"`
	TwitterID      string   `json:"twitterID,omitempty"`
	TwitterAvatar  string   `json:"twitterAvatar,omitempty"`
	TwitterName    string   `json:"twitterName,omitempty"`
	CaveReTweetUrl string   `json:"CaveReTweetUrl" gorm:"column:caveretweeturl;varchar(255)"`
	MintCave       bool     `json:"mintCave,omitempty" gorm:"column:mintcave;tinyint"`
	MyNft          []Nft    `json:"MyNft,omitempty" gorm:"foreignKey:UserID"`
	MyFt           []Ft     `json:"MyFt,omitempty" gorm:"foreignKey:UserID"`
	Secrets        []Secret `json:"secrets,omitempty" gorm:"foreignKey:SecretID"`
	MyFriends      []Friend `json:"MyFriends,omitempty" gorm:"foreignKey:BaseID"`
}
type Cave struct {
}

// Secret ...  一对多 has many
type Secret struct {
	gorm.Model
	SecretID int `json:"secretId"`
	AuthorID int `json:"authorId"`
	// ViewLevel 1 仅广场 2 仅洞穴 3 广场和洞穴
	ViewLevel int           `json:"viewLevel"`
	Timestamp int64         `json:"timestamp"`
	Views     int64         `json:"views"`
	Content   string        `json:"content"`
	Images    []SecretImage `json:"images" gorm:"foreignKey:SecretID"`
	// Status 秘闻状态 可以设置为不可见
	Status bool `json:"status"`
}
type SecretImage struct {
	gorm.Model
	SecretID int    `json:"secretID"`
	ImageUrl string `json:"imageUrl"`
}

type Retweet struct {
	gorm.Model
	CaveID         string
	CaveReTweetUrl string
}

// TaskRecord ... 等价于洞穴权限表 有记录 即可见
type TaskRecord struct {
	gorm.Model
	CaveID string
	UserID int `json:"userId,omitempty"`
}

// Friend ... 推特好友（following + follower）一对多 has many
// eg: select * from secrets s where s.author in basefriend()  AND viewlevel = 3 order by timestamp desc.
type Friend struct {
	BaseID   int
	FriendID int
	Status   bool
}

type Viewable struct {
	SecretID int
	Users    []int
}

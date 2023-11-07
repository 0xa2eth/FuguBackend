package pkg

import "gorm.io/gorm"

type Nft struct {
	gorm.Model
	UserID  int    `json:"UserId,omitempty" gorm:"column:userid;bigint"`
	ChainID int    `json:"ChainId,omitempty" gorm:"column:chainid;bigint"`
	TokenId int    `json:"TokenId,omitempty" gorm:"column:tokenid;bigint"`
	Address string `json:"Address,omitempty" gorm:"column:address;varchar(255)"`
}
type Ft struct {
	gorm.Model
	UserID   int    `json:"UserId,omitempty" gorm:"column:userid;bigint"`
	ChainID  int    `json:"ChainId,omitempty" gorm:"column:chainid;bigint"`
	TokenNum int    `json:"TokenNum,omitempty" gorm:"column:tokennum;bigint"`
	Address  string `json:"Address,omitempty" gorm:"column:address;varchar(255)"`
}

// User ...内存对齐 省空间
type User struct {
	gorm.Model
	UserID         int      `json:"userId,omitempty" gorm:"column:userid;bigint"`
	TicketNum      int      `json:"ticketNum,omitempty" gorm:"column:ticketnum;bigint"`
	CaveFans       int      `json:"caveFans,omitempty" gorm:"column:cavefans;bigint"`
	TwitterFans    int      `json:"twitterFans,omitempty" gorm:"column:twitterfans;bigint"`
	LastLogin      int      `json:"lastLogin,omitempty" gorm:"column:lastlogin;bigint"`
	RegisTime      int      `json:"regisTime,omitempty" gorm:"column:registime;bigint"`
	EarnedPoint    int      `json:"earnedPoint,omitempty" gorm:"column:earnedpoint;bigint"`
	CavePoint      int      `json:"CavePoint,omitempty" gorm:"column:cavepoint;bigint"`
	Views          int      `json:"views" gorm:"column:views;bigint"`
	NickName       string   `json:"nickName,omitempty" gorm:"column:nickname;varchar(255)"`
	Bios           string   `json:"bios,omitempty" gorm:"column:bios;varchar(255)"`
	Avatar         string   `json:"avatar,omitempty" gorm:"column:avatar;varchar(255)"`
	Address        string   `json:"address,omitempty" gorm:"column:address;varchar(255)"`
	TwitterID      string   `json:"twitterID,omitempty" gorm:"column:twitterid;varchar(255)"`
	TwitterAvatar  string   `json:"twitterAvatar,omitempty" gorm:"column:twitteravatar;varchar(255)"`
	TwitterName    string   `json:"twitterName,omitempty" gorm:"column:twittername;varchar(255)"`
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
	SecretID int `json:"secretId,omitempty" gorm:"column:secretid;bigint"`
	AuthorID int `json:"authorId,omitempty" gorm:"column:authorid;bigint"`
	// ViewLevel 1 仅广场 2 仅洞穴 3 广场和洞穴
	ViewLevel int           `json:"viewLevel,omitempty" gorm:"column:viewlevel;bigint"`
	Timestamp int64         `json:"timestamp,omitempty" gorm:"column:timestamp;bigint"`
	Views     int64         `json:"views,omitempty" gorm:"column:views;bigint"`
	Content   string        `json:"content,omitempty" gorm:"column:content;varchar(255)"`
	Images    []SecretImage `json:"images,omitempty" gorm:"foreignKey:SecretID"`
	// Status 秘闻状态 可以设置为不可见
	Status bool `json:"status,omitempty" gorm:"column:status;tinyint"`
}
type SecretImage struct {
	gorm.Model
	SecretID int    `json:"SecretID,omitempty" gorm:"column:secretid;bigint"`
	ImageUrl string `json:"ImageUrl,omitempty" gorm:"column:imageurl;varchar(255)"`
}

type Retweet struct {
	gorm.Model
	CaveID         string `json:"CaveID,omitempty" gorm:"column:caveid;varchar(255)"`
	CaveReTweetUrl string `json:"CaveReTweetUrl,omitempty" gorm:"column:caveretweeturl;varchar(255)"`
}

// TaskRecord ... 等价于洞穴权限表 有记录 即可见
type TaskRecord struct {
	gorm.Model
	CaveID int `json:"CaveID,omitempty" gorm:"column:caveid;bigint"`
	UserID int `json:"userID,omitempty" gorm:"column:userid;bigint"`
}

// Friend ... 推特好友（following + follower）一对多 has many
// eg: select * from secrets s where s.author in basefriend()  AND viewlevel = 3 order by timestamp desc.
type Friend struct {
	BaseID   int  `json:"BaseID,omitempty" gorm:"column:baseid;bigint"`
	FriendID int  `json:"FriendID,omitempty" gorm:"column:friendid;bigint"`
	Status   bool `json:"Status,omitempty" gorm:"column:status;tinyint"`
}

type Viewable struct {
	SecretID int
	Users    []int
}

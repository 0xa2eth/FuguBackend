package pkg

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// User ...内存对齐 省空间
type User struct {
	gorm.Model
	UserID        int    `json:"userId,omitempty" gorm:"type:bigint"`
	TicketNum     int    `json:"ticketNum,omitempty" gorm:"type:bigint"`
	CaveFans      int    `json:"caveFans,omitempty" gorm:"type:bigint"`
	TwitterFans   int    `json:"twitterFans,omitempty" gorm:"type:bigint"`
	LastLogin     int    `json:"lastLogin,omitempty" gorm:"type:bigint"`
	RegisTime     int    `json:"regisTime,omitempty" gorm:"type:bigint"`
	EarnedPoint   int    `json:"earnedPoint,omitempty" gorm:"type:bigint"`
	CavePoint     int    `json:"cavePoint,omitempty" gorm:"type:bigint"`
	Views         int    `json:"views,omitempty" gorm:"type:bigint"`
	VipLevel      int    `json:"vipLevel,omitempty" gorm:"type:bigint"`
	InvitedBy     int    `json:"invitedBy" gorm:"type:bigint"`
	Posts         int    `json:"posts,omitempty" gorm:"type:bigint"`
	NickName      string `json:"nickName,omitempty" gorm:"type:varchar(255)"`
	Bios          string `json:"bios,omitempty" gorm:"type:varchar(255)"`
	Avatar        string `json:"avatar,omitempty" gorm:"type:varchar(255)"`
	Address       string `json:"address,omitempty" gorm:"type:varchar(255)"`
	TwitterID     string `json:"twitterID,omitempty" gorm:"type:varchar(255)"`
	TwitterAvatar string `json:"twitterAvatar,omitempty" gorm:"type:varchar(255)"`
	TwitterName   string `json:"twitterName,omitempty" gorm:"type:varchar(255)"`
	RetweetUrl    string `json:"retweetUrl" gorm:"type:varchar(255)"`
	MintCave      bool   `json:"mintCave,omitempty" gorm:"type:bool"`
}

// Secret ...  一对多 has many
type Secret struct {
	gorm.Model
	SecretID int `json:"secretId,omitempty" gorm:"type:bigint"`
	AuthorID int `json:"authorId,omitempty" gorm:"type:bigint"`
	// ViewLevel 1 仅广场 2 仅洞穴 3 广场和洞穴
	ViewLevel int            `json:"viewLevel,omitempty" gorm:"type:int"`
	Timestamp int64          `json:"timestamp,omitempty" gorm:"type:bigint"`
	Views     int64          `json:"views,omitempty" gorm:"type:bigint"`
	Content   string         `json:"content,omitempty" gorm:"type:varchar(255)"`
	Images    datatypes.JSON `json:"images,omitempty" gorm:"column:images"`
	// Status 秘闻状态 平台可以将非法的帖子设置为不可见
	Status bool `json:"status,omitempty" gorm:"type:bool"`
}

// TaskRecord ... 等价于洞穴权限表 有记录 即可见
type TaskRecord struct {
	gorm.Model
	CaveID int `json:"caveID,omitempty" gorm:"type:int"`
	UserID int `json:"userID,omitempty" gorm:"type:int"`
}

// Friend ... 推特好友（following + follower）一对多 has many
// eg: select * from secrets s where s.author in basefriend()  AND viewlevel = 3 order by timestamp desc.
type Friend struct {
	gorm.Model
	BaseID   int  `json:"baseID,omitempty" gorm:"type:bigint"`
	FriendID int  `json:"friendID,omitempty" gorm:"type:bigint"`
	Status   bool `json:"status,omitempty" gorm:"type:bool"`
}

//type InviteCode struct {
//	gorm.Model
//	UserID int    `json:"userID,omitempty" gorm:"column:userid;type:bigint"`
//	Code   string `json:"code,omitempty" gorm:"column:code;type:varchar(255)"`
//}

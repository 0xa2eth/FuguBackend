package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	u := userInfoResponse{
		UserID:        "",
		TicketNum:     0,
		CaveFans:      0,
		EarnedPoint:   0,
		CavePoint:     0,
		Views:         0,
		NickName:      "",
		Bios:          "",
		Avatar:        "",
		NumberOfPosts: 0,
		FollowedCaves: nil,
	}
	u.FollowedCaves = make([]FollowedCave, 1)
	aa := SecretEntity{
		Timestamp: 0,
		Views:     0,
		SecretID:  "",
		Content:   "",
		Images:    []string{""},
		Publisher: AuthorInfo{},
	}
	var bb []SecretEntity
	bb = append(bb, aa)

	a := secretListResponse{Secrets: bb}
	var list []interface{}
	list = append(list,
		//topResponse{Caves: make([]caveInfo, 1)},
		//userInfoResponse{FollowedCaves: make([]FollowedCave, 1)},
		//u,
		a,
	)

	StructToJson(list)
}

func StructToJson(collection []interface{}) {
	for _, instance := range collection {
		// 判断是否是结构体
		marshal, err := json.Marshal(instance)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(marshal))
		//return fmt.Sprintf("%v", string(marshal))
	}
}

type topResponse struct {
	Caves []caveInfo `json:"caves"`
}
type caveInfo struct {
	CaveID     string `json:"caveID"`
	CaveName   string `json:"caveName"`
	CaveBio    string `json:"caveBio"`
	CaveAvatar string `json:"caveAvatar"`
}

type userInfoResponse struct {
	UserID        string         `json:"userId" gorm:"column:userid;type:bigint"`
	TicketNum     int            `json:"ticketNum" gorm:"column:ticketnum;type:bigint"`
	CaveFans      int            `json:"caveFans" gorm:"column:cavefans;type:bigint"`
	EarnedPoint   int            `json:"earnedPoint" gorm:"column:earned_point;type:bigint"`
	CavePoint     int            `json:"CavePoint" gorm:"column:cave_point;type:bigint"`
	Views         int            `json:"views" gorm:"column:views;type:bigint"`
	NickName      string         `json:"nickName" gorm:"column:nick_name;type:varchar(255)"`
	Bios          string         `json:"bios" gorm:"column:bios;type:varchar(255)"`
	Avatar        string         `json:"avatar" gorm:"column:avatar;type:varchar(255)"`
	NumberOfPosts int            `json:"numberOfPosts"`
	FollowedCaves []FollowedCave `json:"followedCaves"`
}
type FollowedCave struct {
	CaveID     string `json:"caveID"`
	CaveName   string `json:"caveName"`
	CaveAvatar string `json:"caveAvatar"`
}

type secretListResponse struct {
	Secrets []SecretEntity `json:"secrets"`
}
type SecretEntity struct {
	Timestamp int64      `json:"timestamp"`
	Views     int64      `json:"views" gorm:"column:views;type:bigint"`
	SecretID  string     `json:"secretId"`
	Content   string     `json:"content" gorm:"column:content;type:varchar(255)"`
	Images    []string   `json:"images" gorm:"foreignKey:SecretID"`
	Publisher AuthorInfo `json:"publisher"`
}

type AuthorInfo struct {
	CaveID     string `json:"caveID"`
	CaveName   string `json:"caveName"`
	CaveBio    string `json:"caveBio"`
	CaveAvatar string `json:"caveAvatar"`
}

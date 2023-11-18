package twittersvc

import (
	"fmt"
	"go.uber.org/zap"
	"time"
)

type People struct {
	Id        int // 用户ID
	TwitterID string
}

func (s *TwitterServiceMaster) TimeTask() {
	var count = 0
	ticker := time.NewTicker(time.Hour * 3)

	//defer ticker.Stop()
	var ct time.Time
	for {
		select {
		case ct = <-ticker.C:
			count++
			zap.L().Info(fmt.Sprintf(" 第%v次拉取开始  开始时间:%v \t\t\t", count, ct))

			s.DoRefreshTaskFunc()

			zap.L().Info(fmt.Sprintf(" 第%v次拉取结束  结束时间:%v   耗时：%v 秒\n",
				count, time.Now(), time.Since(ct).Seconds()))

		}

	}
}
func (s *TwitterServiceMaster) DoRefreshTaskFunc() {
	// 1 拿在线人数

	// 2 获取在线人数的follower 和 following

	// 3 和平台用户做对比 找到friends

	// 4 先删 再增（或者）

}

//func aaa() {
//	userSvc := user.New(s.db, s.cache, s.logger, s.twSvc)
//
//	person, _ := userSvc.Detail(c, &user.SearchOneData{
//		Id: InnerID,
//	})
//	all, err := userSvc.FindAll()
//	if err != nil {
//		return nil
//	}
//	var people []People
//	for i := range all {
//		//ids = append(ids, all[i].TwitterId)
//		people = append(people, People{
//			Id:        int(all[i].Id),
//			TwitterID: all[i].TwitterId,
//		})
//	}
//	//diff, err2 := s.twSvc.DiffFriendCircle(c, InnerID, person.TwitterName, ids)
//	//if err2 != nil {
//	//	return pagination.PageInfo{}, err2
//	//}
//	var siteFriends []People
//	circle, _ := s.twSvc.RefreshFriendCircle(c, person.TwitterName)
//	for i := range circle {
//		for j := range people {
//
//			if circle[i].IDStr == people[j].TwitterID {
//				siteFriends = append(siteFriends, people[j])
//			}
//
//		}
//	}
//	var ids []int
//	for i := range siteFriends {
//		ids = append(ids, siteFriends[i].Id)
//	}
//}

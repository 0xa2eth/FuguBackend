package twittersvc

import (
	"FuguBackend/app/code"
	"FuguBackend/app/pkg/core"
	"FuguBackend/app/repository/redis"
	"go.uber.org/zap"
	"strconv"
)

const (
	cacheUser    = "user:"
	refreshCount = ":refreshCount"
	timeFriends  = ":times:friends"
)

type FDiff struct {
	Newer  []string
	Remove []string
}

func (s *TwitterServiceMaster) DiffFriendCircle(ctx core.Context, InnerID int, screenName string, ids []string) (diff FDiff, err error) {

	redisCli := redis.RedisC
	idStr := strconv.Itoa(InnerID)

	//count, cacheErr := s.cache.Get(user + idStr + refreshCount)
	// 1 拿这个人最新的刷新次数 从而找到上次的key
	oldCountString, cacheErr := redisCli.Get(cacheUser + idStr + refreshCount).Result()
	if cacheErr != nil {
		s.logger.Error(code.Text(code.CacheGetError), zap.Error(cacheErr))
	}
	count, _ := strconv.Atoi(oldCountString)
	newCount := count + 1
	defer redisCli.Incr(cacheUser + idStr + refreshCount)

	// 2 取他最新朋友圈(todo 一次请求的量的限制？）
	circle, err := s.RefreshFriendCircle(ctx, screenName)
	if err != nil {
		s.logger.Error(code.Text(code.RefreshFriendCircleError), zap.Error(err))
	}
	oldKey := cacheUser + idStr + oldCountString + timeFriends
	newKey := cacheUser + idStr + strconv.Itoa(newCount) + timeFriends
	for i := range circle {
		redisCli.SAdd(newKey, circle[i].ID)
	}
	//SDIFF 旧 新 = 旧的有 新的没有 = 删除的
	deleteFriend, err := redisCli.SDiff(oldKey, newKey).Result()
	if err != nil {
		s.logger.Error(code.Text(code.RefreshFriendCircleError), zap.Error(err))
	}
	defer redisCli.Del(oldKey)
	//SDIFF 新 旧 = 新的有 旧的没有 = 新增的
	newFriend, err := redisCli.SDiff(newKey, oldKey).Result()
	if err != nil {
		s.logger.Error(code.Text(code.RefreshFriendCircleError), zap.Error(err))
	}
	// 从已经注册的用户中 找他的新增或变化的那些朋友
	for i := range newFriend {
		for j := range ids {
			if newFriend[i] == ids[j] {
				diff.Newer = append(diff.Newer, ids[j])
			}
		}
	}

	for i := range deleteFriend {
		for j := range ids {
			if newFriend[i] == ids[j] {
				diff.Remove = append(diff.Remove, ids[j])
			}
		}
	}

	return diff, nil
}

package twittersvc

import (
	"strconv"
	"strings"
)

func (s *TwitterServiceMaster) GetTweetIDByUrl(url string) (int, error) {
	split := strings.Split(url, "/")
	tweetid := split[len(split)-1]

	Tid, err := strconv.Atoi(tweetid)
	if err != nil {
		return 0, err
	}
	return Tid, nil
}

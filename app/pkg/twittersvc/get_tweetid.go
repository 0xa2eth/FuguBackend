package twittersvc

import (
	"strconv"
	"strings"
)

func (s *TwitterServiceMaster) GetTweetIDByUrl(url string) (int, string, error) {
	split := strings.Split(url, "/")
	tweetid := split[len(split)-1]

	Tid, err := strconv.Atoi(tweetid)
	if err != nil {
		return 0, "", nil
	}
	return Tid, tweetid, nil
}

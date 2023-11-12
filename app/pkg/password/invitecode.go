package password

import (
	"crypto/md5"
	"strconv"
	"time"
)

var AlphanumericSet = []rune{
	'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
}

func GenInviteCode(uid string, l int) string {
	if l > 16 {
		return ""
	}
	nano := time.Now().UnixNano()
	nanostr := strconv.Itoa(int(nano))

	in := uid + nanostr
	sum := md5.Sum([]byte(in))
	var code []rune
	for i := 0; i < l; i++ {
		idx := sum[i] % byte(len(AlphanumericSet))
		code = append(code, AlphanumericSet[idx])
	}
	return string(code)
}

// GenRandomStr 随机字符串 长度小于16个字符
func GenRandomStr(base string, l int) string {
	if l > 16 {
		return ""
	}
	nano := time.Now().UnixNano()
	nanostr := strconv.Itoa(int(nano))
	in := base + nanostr
	sum := md5.Sum([]byte(in))
	var code []rune
	for i := 0; i < l; i++ {
		idx := sum[i] % byte(len(AlphanumericSet))
		code = append(code, AlphanumericSet[idx])
	}
	return string(code)
}

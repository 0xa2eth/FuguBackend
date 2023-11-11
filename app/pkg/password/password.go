package password

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

const (
	saltPassword    = "qkhPAGA13HocW3GAEWwb"
	defaultPassword = "123456"
)

func GeneratePassword(str string) (password string) {
	// md5
	m := md5.New()
	m.Write([]byte(str))
	mByte := m.Sum(nil)

	// hmac
	h := hmac.New(sha256.New, []byte(saltPassword))
	h.Write(mByte)
	password = hex.EncodeToString(h.Sum(nil))

	return
}

func ResetPassword() (password string) {
	m := md5.New()
	m.Write([]byte(defaultPassword))
	mStr := hex.EncodeToString(m.Sum(nil))

	password = GeneratePassword(mStr)

	return
}

func GenerateLoginToken(id int32) (token string) {
	m := md5.New()
	m.Write([]byte(fmt.Sprintf("%d%s", id, saltPassword)))
	token = hex.EncodeToString(m.Sum(nil))

	return
}

const (
	secretKey = "kEf%Wr2SsLke"
	Issuer    = "Fugu.club"
)

// CustomClaims 定义 JWT 载荷结构
type CustomClaims struct {
	UserID   string `json:"userID"`
	LoggedIn bool   `json:"loggedIn"`
	jwt.RegisteredClaims
}

// GenerateJWT ...  生成 JWT
func GenerateJWT(userID string) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour) // 设置过期时间为1小时
	claims := &CustomClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Issuer:    Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

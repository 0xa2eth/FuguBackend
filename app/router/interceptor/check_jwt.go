package interceptor

import (
	"FuguBackend/app/code"
	"FuguBackend/app/pkg/core"
	"FuguBackend/config"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
)

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

// CheckJWT ... JWT鉴权中间件
func (i *interceptor) CheckJWT() core.HandlerFunc {
	return func(c core.Context) {
		// 签名信息
		tokenString := c.GetHeader(config.HeaderSignToken)
		if tokenString == "" {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.AuthorizationError,
				code.Text(code.AuthorizationError)).WithError(errors.New("Header 中缺少 Authorization 参数")),
			)
			return
		}
		i.logger.Info("========jwt:get header success ============")
		token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})

		if err != nil || !token.Valid {

			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.AuthorizationError,
				code.Text(code.AuthorizationError)).WithError(errors.New("invalid token")),
			)

			return
		}
		i.logger.Info("===========jwt:  jwt.ParseWithClaims() success ==============")
		claims, ok := token.Claims.(*CustomClaims)
		if !ok {
			c.AbortWithError(core.Error(
				http.StatusUnauthorized,
				code.AuthorizationError,
				code.Text(code.AuthorizationError)).WithError(errors.New("invalid token")),
			)
			return
		}

		i.logger.Info("===========jwt: token.Claims.(*CustomClaims) success  jwt-claims 类型断言成功 ==============")

		// 将用户信息存储在上下文中，供后续处理函数使用

		c.Set("UserID", claims.UserID)
		//这里 userID和SessionUserInfo() 都能渠取到用户的userid(或proposal.SessionUserInfo)
		//c.setSessionUserInfo(config._SessionUserInfo,claims.UserID)

		//c.Next()
	}
}

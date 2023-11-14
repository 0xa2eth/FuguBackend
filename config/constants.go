package config

import "time"

const (
	// MinGoVersion 最小 Go 版本
	MinGoVersion = 1.19

	// ProjectVersion 项目版本
	ProjectVersion = "v1.2.8"

	// ProjectName 项目名称
	ProjectName = "FuguToxic"

	// ProjectDomain 项目域名
	ProjectDomain = "http://127.0.0.1"

	// ProjectPort 项目端口
	ProjectPort = ":9999"

	// ProjectAccessLogFile 项目访问日志存放文件
	ProjectAccessLogFile = "./logs/" + ProjectName + "-access.log"

	// ProjectCronLogFile 项目后台任务日志存放文件
	ProjectCronLogFile = "./logs/" + ProjectName + "-cron.log"

	// ProjectInstallMark 项目安装完成标识
	ProjectInstallMark = "INSTALL.lock"

	// HeaderLoginToken 登录验证 Token，Header 中传递的参数
	HeaderLoginToken = "Token"

	// HeaderSignToken 签名验证 Authorization，Header 中传递的参数
	HeaderSignToken = "Authorization"

	// HeaderSignTokenDate 签名验证 Date，Header 中传递的参数
	HeaderSignTokenDate = "Authorization-Date"

	// HeaderSignTokenTimeout 签名有效期为 2 分钟
	HeaderSignTokenTimeout = time.Minute * 2

	// RedisKeyPrefixLoginUser Redis Key 前缀 - 登录用户信息
	RedisKeyPrefixLoginUser = ProjectName + ":login-user:"

	// RedisKeyPrefixSignature Redis Key 前缀 - 签名验证信息
	RedisKeyPrefixSignature = ProjectName + ":signature:"

	// ZhCN 简体中文 - 中国
	ZhCN = "zh-cn"

	// EnUS 英文 - 美国
	EnUS = "en-us"

	// MaxRequestsPerSecond 每秒最大请求量
	MaxRequestsPerSecond = 10000

	// LoginSessionTTL 登录有效期为 24 小时
	LoginSessionTTL = time.Hour * 24

	// DefaultInviteCodeLength 默认邀请码长度
	DefaultInviteCodeLength = 8

	// DefaultInviteCodeTTL 默认邀请码有效期
	DefaultInviteCodeTTL = time.Hour * 24

	// DefaultAvatar 默认头像
	DefaultAvatar = "https://fugutoxic.s3.ap-northeast-1.amazonaws.com/siteavatar/defaultavatar.jpg"

	// FuguTwitterName
	FuguTwitterName = "fugu_social"
	FuguTwitterID   = "fugu_social"

	// DefaultBio 平台默认Bio
	DefaultBio = "Fugu Toxic Secret is an anonymous content-sharing platform built on Twitter connections, turning your real-life stories into earnings."

	TwitterPostPreFix = "A New Secret Cave =="
	TwitterPostSuffix = "== Has Been Created! Come To See And Find Your Interests! Transfer Link Here https://metahome.tech"
	TransferLink      = ""

	UserPortraitBaseUrl = "https://fugutoxic.s3.ap-northeast-1.amazonaws.com/siteavatar/"

	RetweetPrefix = "https://twitter.com/intent/retweet?tweet_id="
)

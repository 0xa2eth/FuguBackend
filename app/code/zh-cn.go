package code

var zhCNText = map[int]string{
	ServerError:        "内部服务器错误",
	TooManyRequests:    "请求过多",
	ParamBindError:     "参数信息错误",
	AuthorizationError: "签名信息错误",
	UrlSignError:       "参数签名错误",
	CacheSetError:      "设置缓存失败",
	CacheGetError:      "获取缓存失败",
	CacheDelError:      "删除缓存失败",
	CacheNotExist:      "缓存不存在",
	ResubmitError:      "请勿重复提交",
	HashIdsEncodeError: "HashID 加密失败",
	HashIdsDecodeError: "HashID 解密失败",
	RBACError:          "暂无访问权限",
	RedisConnectError:  "Redis 连接失败",
	MySQLConnectError:  "MySQL 连接失败",
	WriteConfigError:   "写入配置文件失败",
	SendEmailError:     "发送邮件失败",
	MySQLExecError:     "SQL 执行失败",
	GoVersionError:     "Go 版本不满足要求",
	SocketConnectError: "Socket 未连接",
	SocketSendError:    "Socket 消息发送失败",

	AuthorizedCreateError:    "创建调用方失败",
	AuthorizedListError:      "获取调用方列表失败",
	AuthorizedDeleteError:    "删除调用方失败",
	AuthorizedUpdateError:    "更新调用方失败",
	AuthorizedDetailError:    "获取调用方详情失败",
	AuthorizedCreateAPIError: "创建调用方 API 地址失败",
	AuthorizedListAPIError:   "获取调用方 API 地址列表失败",
	AuthorizedDeleteAPIError: "删除调用方 API 地址失败",

	AdminCreateError:             "创建管理员失败",
	AdminListError:               "获取管理员列表失败",
	AdminDeleteError:             "删除管理员失败",
	AdminUpdateError:             "更新管理员失败",
	AdminResetPasswordError:      "重置密码失败",
	UserLoginError:               "登录失败",
	AdminLogOutError:             "退出失败",
	AdminModifyPasswordError:     "修改密码失败",
	AdminModifyPersonalInfoError: "修改个人信息失败",
	AdminMenuListError:           "获取管理员菜单授权列表失败",
	AdminMenuCreateError:         "管理员菜单授权失败",
	AdminOfflineError:            "下线管理员失败",
	AdminDetailError:             "获取个人信息失败",

	MenuCreateError:       "创建菜单失败",
	MenuUpdateError:       "更新菜单失败",
	MenuDeleteError:       "删除菜单失败",
	MenuListError:         "获取菜单列表失败",
	MenuDetailError:       "获取菜单详情失败",
	MenuCreateActionError: "创建菜单栏功能权限失败",
	MenuListActionError:   "获取菜单栏功能权限列表失败",
	MenuDeleteActionError: "删除菜单栏功能权限失败",

	CronCreateError:  "创建后台任务失败",
	CronUpdateError:  "更新后台任务失败",
	CronListError:    "获取定时任务列表失败",
	CronDetailError:  "获取定时任务详情失败",
	CronExecuteError: "手动执行定时任务失败",

	VerifyTaskError:          "验证任务失败",
	SecretCreateError:        "发布秘密失败",
	GetSecretsError:          "获取秘密失败",
	RefreshFriendCircleError: "刷新朋友圈失败",
	InvitationVerifyError:    "邀请码效验失败",
}

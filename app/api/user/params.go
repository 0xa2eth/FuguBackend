package user

type createRequest struct {
	Username string `form:"username" binding:"required"` // 用户名
	Nickname string `form:"nickname" binding:"required"` // 昵称
	Mobile   string `form:"mobile" binding:"required"`   // 手机号
	Password string `form:"password" binding:"required"` // MD5后的密码
}

type createResponse struct {
	Id int32 `json:"id"` // 主键ID
}

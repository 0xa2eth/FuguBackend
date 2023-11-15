package router

import (
	"FuguBackend/app/api/cave"
	"FuguBackend/app/api/secret"
	"FuguBackend/app/api/user"
	"FuguBackend/app/pkg/core"
)

//type Resource app.Resource

func SetApiRouter(r *Resource) {

	test := r.Mux.Group("")
	test.GET("/ping", func(c core.Context) {
		c.ResponseWriter().WriteString("pong")
	})

	// 无需登录验证
	login := r.Mux.Group("/api")
	// 创建用户
	userHandler := user.New((*user.Resource)(r))

	login.POST("/twitterLogin", userHandler.RegisterOrLogin())

	// 需要签名验证、登录验证、RBAC 权限验证
	api := r.Mux.Group("/api", core.WrapAuthHandler(r.Interceptors.CheckLogin), r.Interceptors.CheckJWT())
	{
		{
			// user
			userGroup := api.Group("/user")

			// 取用户信息
			userGroup.GET("/:UserID", userHandler.UserInfo())
			// 创建cave，修改cave信息
			userGroup.POST("/:UserID", userHandler.ModifyInfo())
			// 生成邀请码
			userGroup.GET("/invitecode", userHandler.GenInviteCode())
			// 验证邀请码
			userGroup.POST("/verifyinvitation", userHandler.VerifyInviteCode())

			// 用户登出
			userGroup.GET("/logout", userHandler.Logout())

		}
		{
			// secret
			secretGroup := api.Group("/secret")
			secretHandler := secret.New((*secret.Resource)(r))
			// 发布秘密
			secretGroup.POST("/:UserID", secretHandler.Create())
			// 三类：正常广场的， 特权的， 还有洞穴的
			// 特权的和广场的在一个接口里 洞穴的单独一个接口
			secretGroup.GET("/square", secretHandler.List())
			// 投诉
			secretGroup.GET("/complaint/:SecretID", secretHandler.Complaint())
		}
		{
			// cave
			caveGroup := api.Group("/cave")
			caveHandler := cave.New((*cave.Resource)(r))
			caveGroup.GET("/top", caveHandler.Top())
			// 进入自己的 或者已拥有的
			caveGroup.GET("/:CaveID", caveHandler.SecretList())
			//

			// 三类：正常广场的， 特权的， 还有洞穴的
			// 特权的和广场的在一个接口里 洞穴的单独一个接口 逻辑上只能推荐没买票的   废弃
			caveGroup.GET("/recommend", caveHandler.RecommendCave())

			caveGroup.GET("/:CaveID/verify/:Type", caveHandler.VerifyTask())
		}

	}
}

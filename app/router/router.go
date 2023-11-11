package router

import (
	"FuguBackend/app/api/secret"
	"FuguBackend/app/api/user"
	"FuguBackend/app/pkg/core"
)

//type Resource app.Resource

func SetApiRouter(r *Resource) {
	//// helper
	//helperHandler := helper.New(r.logger, r.db, r.cache)
	//
	//helpers := r.mux.Group("/helper")
	//{
	//	helpers.GET("/md5/:str", helperHandler.Md5())
	//	helpers.POST("/sign", helperHandler.Sign())
	//}

	test := r.Mux.Group("")
	test.GET("/ping", func(c core.Context) {
		c.ResponseWriter().WriteString("pong")
	})
	//
	//// 需要签名验证，无需登录验证，无需 RBAC 权限验证
	//login := r.Mux.Group("/api", r.Interceptors.CheckSignature())
	//{
	//
	//}
	//
	//// 需要签名验证、登录验证，无需 RBAC 权限验证
	//notRBAC := r.Mux.Group("/api", core.WrapAuthHandler(r.Interceptors.CheckLogin), r.Interceptors.CheckSignature())
	//{
	//	notRBAC.GET("/admin/info", adminHandler.Detail())
	//	notRBAC.PATCH("/admin/modify_personal_info", adminHandler.ModifyPersonalInfo())
	//}

	// 需要签名验证、登录验证、RBAC 权限验证
	api := r.Mux.Group("/api")
	{
		{
			// user
			userGroup := api.Group("/user")
			userHandler := user.New((*user.Resource)(r))
			// 创建
			userGroup.POST("/twitterlogin", userHandler.Create())
			//
			userGroup.GET("/:hashid", userHandler.Detail())
			//userGroup.GET("/:hashid/caves", userHandler.Detail())
			userGroup.GET("/caves", userHandler.List())

		}
		{
			// secret
			secretGroup := api.Group("/secrets")
			secretHandler := secret.New((*secret.Resource)(r))
			secretGroup.POST("", secretHandler.Create())
			secretGroup.GET("/:hashid", secretHandler.Detail())
			// 三类：正常广场的， 特权的， 还有洞穴的
			// 特权的和广场的在一个接口里 洞穴的单独一个接口
			api.GET("", secretHandler.List())
		}

		//api.PATCH("/authorized/used", authorizedHandler.UpdateUsed())
		//api.DELETE("/authorized/:id", core.AliasForRecordMetrics("/api/authorized/info"), authorizedHandler.Delete())
		//
		//api.POST("/authorized_api", authorizedHandler.CreateAPI())
		//api.GET("/authorized_api", authorizedHandler.ListAPI())
		//api.DELETE("/authorized_api/:id", core.AliasForRecordMetrics("/api/authorized_api/info"), authorizedHandler.DeleteAPI())

		//api.POST("/admin", adminHandler.Create())
		//api.GET("/admin", adminHandler.List())
		//api.PATCH("/admin/used", adminHandler.UpdateUsed())
		//api.PATCH("/admin/offline", adminHandler.Offline())
		//api.PATCH("/admin/reset_password/:id", core.AliasForRecordMetrics("/api/admin/reset_password"), adminHandler.ResetPassword())
		//api.DELETE("/admin/:id", core.AliasForRecordMetrics("/api/admin"), adminHandler.Delete())
		//
		//api.POST("/admin/menu", adminHandler.CreateAdminMenu())
		//api.GET("/admin/menu/:id", core.AliasForRecordMetrics("/api/admin/menu"), adminHandler.ListAdminMenu())
		//
		//// menu
		//menuHandler := menu.New(r.logger, r.db, r.cache)
		//api.POST("/menu", menuHandler.Create())
		//api.GET("/menu", menuHandler.List())
		//api.GET("/menu/:id", core.AliasForRecordMetrics("/api/menu"), menuHandler.Detail())
		//api.PATCH("/menu/used", menuHandler.UpdateUsed())
		//api.PATCH("/menu/sort", menuHandler.UpdateSort())
		//api.DELETE("/menu/:id", core.AliasForRecordMetrics("/api/menu"), menuHandler.Delete())
		//api.POST("/menu_action", menuHandler.CreateAction())
		//api.GET("/menu_action", menuHandler.ListAction())
		//api.DELETE("/menu_action/:id", core.AliasForRecordMetrics("/api/menu_action"), menuHandler.DeleteAction())
		//
		//// tool
		//toolHandler := tool.New(r.logger, r.db, r.cache)
		//api.GET("/tool/hashids/encode/:id", core.AliasForRecordMetrics("/api/tool/hashids/encode"), toolHandler.HashIdsEncode())
		//api.GET("/tool/hashids/decode/:id", core.AliasForRecordMetrics("/api/tool/hashids/decode"), toolHandler.HashIdsDecode())
		//api.POST("/tool/cache/search", toolHandler.SearchCache())
		//api.PATCH("/tool/cache/clear", toolHandler.ClearCache())
		//api.GET("/tool/data/dbs", toolHandler.Dbs())
		//api.POST("/tool/data/tables", toolHandler.Tables())
		//api.POST("/tool/data/mysql", toolHandler.SearchMySQL())
		//api.POST("/tool/send_message", toolHandler.SendMessage())
		//
		//// config
		//configHandler := config.New(r.logger, r.db, r.cache)
		//api.PATCH("/config/email", configHandler.Email())
		//
		//// cron
		//cronHandler := cron.New(r.logger, r.db, r.cache, r.cronServer)
		//api.POST("/cron", cronHandler.Create())
		//api.GET("/cron", cronHandler.List())
		//api.GET("/cron/:id", core.AliasForRecordMetrics("/api/cron/detail"), cronHandler.Detail())
		//api.POST("/cron/:id", core.AliasForRecordMetrics("/api/cron/modify"), cronHandler.Modify())
		//api.PATCH("/cron/used", cronHandler.UpdateUsed())
		//api.PATCH("/cron/exec/:id", core.AliasForRecordMetrics("/api/cron/exec"), cronHandler.Execute())

	}
}

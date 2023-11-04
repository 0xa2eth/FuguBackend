package router

import (
	app "FuguBackend/app"
	"FuguBackend/app/services/user"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetApiRouter(svcCtx *app.ServiceCtx) *gin.Engine {
	gin.ForceConsoleColor()
	gin.SetMode(gin.DebugMode)
	r := gin.New()
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "X-CSRF-Token", "Authorization", "AccessToken", "Token"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "X-GW-Error-Code", "X-GW-Error-Message"},
		AllowCredentials: true,
		MaxAge:           1 * time.Hour,
	}))
	loadv1(r, svcCtx)
	LoadAdmin(r, svcCtx)

	// setup swagger ui server side
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return r
}

func LoadAdmin(r *gin.Engine, svcCtx *app.ServiceCtx) {
	{
		//潜在后台接口
		//设置项目方
		adminGroup := r.Group("/admin")
		adminGroup.GET("/setProject", api.SetProjectHandler())
	}
}

func loadv1(r *gin.Engine, svcCtx *app.ServiceCtx) {
	apiv1 := r.Group("/api")
	d := apiv1.Group("/demo")
	{
		d.POST("/d", func(ctx *gin.Context) {
			xhttp.OkJson(ctx, "hi")
		})
		d.POST("/login", func(ctx *gin.Context) {
			login.Login(ctx, svcCtx)
		})
	}
	//node 通讯
	{
		apiv1.POST("/eventMonitor", api.EventMonitorHandler())
	}

	//公开接口
	{
		apiv1.POST("/walletlogin", api.WalletLoginHandler())
		apiv1.GET("/homeCampaigns", api.HomeCampaignsHandler())
		apiv1.POST("/requestToSettle", api.RequestToSettleHandler())
		apiv1.GET("/recommendation", api.RecommendationHandler())
		apiv1.GET("/gate3dic", api.Gate3DicHandler1())
	}

	//user && space
	user.New()
	{
		apiv1.GET("/userinfo", api.GetUserinfoHandler())
		apiv1.GET("/messages", api.GetMessagesHandler())
		apiv1.GET("/myFollowing", api.MyFollowingHandler())
		apiv1.GET("/readMessage", api.ReadMessageHandler())
		apiv1.POST("/editSpace", api.EditSpaceHandler())
		apiv1.POST("/bindWeb2", api.BindingWeb2Handler())
		apiv1.GET("/spaceProfile", api.GetSpaceprofileHandler())
		apiv1.GET("/spaceCampaigns", api.SpaceCampaignsHandler())
	}
	//project
	{
		apiv1.GET("/getProjectProfile", api.GetProjectProfileHandler())
		apiv1.GET("/getProjectCampaigns", api.ProjectCampaignsHandler())
		//新ui设计项目方也可以follow和unfollow
		apiv1.POST("/projectFollow", api.FollowProjectControllerHandler())
		apiv1.POST("/projectisfollow", api.IsProjectFollowedHandler())
		//apiv1.POST("/subscribeProject", api.ProjectSubscribeHandler())
		//apiv1.POST("/projectSubVerify", api.IsProjectSubscribedHandler())
		apiv1.POST("/launch", api.LaunchHandler())
		//apiv1.GET("/launchCallback", api.LaunchCallbackHanlder())
	}
	//favorite
	{
		//apiv1.GET("/myfavorites", api.MyFavoritesHandler())
		apiv1.GET("/favorite", api.FavoriteHandler())
	}
	//Dashboard
	{
		apiv1.POST("/createCampaign", api.CreateCampaignHandler())
		apiv1.GET("/myCampaigns", api.AllMyCampaignHandler())
		apiv1.POST("/editCampaign", api.EditCampaignHandler())
		////任务玩法
		apiv1.POST("/setTasks", api.SetTaskHandler())
		//
		apiv1.POST("/createNftContract", api.CreateNftContractHandler())
		apiv1.GET("/myNftContracts", api.MyNFTContractHandler())
		//
		apiv1.POST("/setReward", api.SetRewardInfoHandler())
		//
		////startUp
		apiv1.GET("/startUp", api.StartUpHandler())
		////winnerList
		apiv1.GET("/winnerList", api.WinnerListHandler())
	}

	//campaign
	{
		//CampaignInfo
		apiv1.GET("/campaignInfo", api.CampaignInfoHandler())
		apiv1.GET("/campaignActivity", api.CampaignActivityHandler())
		apiv1.GET("/participants", api.ParticipantsHandler())

		//参与项目 和 任务验证
		apiv1.GET("/verify", api.TaskVerifyHandler())
	}

	//mining
	{
		apiv1.GET("/config", api.GetConfigHandler())
		apiv1.GET("/rewardsPool", api.GetRewardPoolHandler())
	}

}

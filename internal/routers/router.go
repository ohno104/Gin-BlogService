package routers

import (
	"time"

	"felix.bs.com/felix/BeStrongerInGO/Gin-BlogService/global"
	"felix.bs.com/felix/BeStrongerInGO/Gin-BlogService/internal/middleware"
	"felix.bs.com/felix/BeStrongerInGO/Gin-BlogService/internal/routers/api"
	v1 "felix.bs.com/felix/BeStrongerInGO/Gin-BlogService/internal/routers/api/v1"
	"felix.bs.com/felix/BeStrongerInGO/Gin-BlogService/pkg/limiter"

	_ "felix.bs.com/felix/BeStrongerInGO/Gin-BlogService/docs"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/gin-swagger/swaggerFiles"
)

var methodLimiters = limiter.NewMethodLimiter().AddBuckets(
	limiter.LimiterBucketRule{
		Key:          "/auth",
		FillInterval: time.Second,
		Capacity:     10,
		Quantm:       10,
	},
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	//r.Use(gin.Recovery())
	r.Use(middleware.AccessLog())
	r.Use(middleware.Recovery())
	r.Use(middleware.Tracing())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/auth", api.GetAuth)
	r.Use(middleware.RateLimiter(methodLimiters))
	r.Use(middleware.ContextTimeout(time.Duration(global.AppSetting.DefaultContextTimeout) * time.Second))
	r.Use(middleware.Translations())

	article := v1.NewArticle()
	tag := v1.NewTag()
	upload := api.NewUpload()
	r.POST("/upload/file", upload.UploadFile)
	r.Static("/static", "./"+global.AppSetting.UploadSavePath)

	apiv1 := r.Group("/api/v1")
	//apiv1.Use(middleware.JWT())
	apiv1.Use()
	{
		apiv1.POST("/tags", tag.Create)
		apiv1.DELETE("/tags/:id", tag.Delete)
		apiv1.PUT("/tags/:id", tag.Update)
		apiv1.PATCH("/tags/:id/state", tag.Update)
		apiv1.GET("/tags", tag.List)

		apiv1.POST("/articles", article.Create)
		apiv1.DELETE("/articles/:id", article.Delete)
		apiv1.PUT("/articles/:id", article.Update)
		apiv1.PATCH("/articles/:id/state", article.Update)
		apiv1.GET("/articles/:id", article.Get)
		apiv1.GET("/articles", article.List)
	}

	return r
}

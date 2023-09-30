package routers

import (
	"blog/global"
	"blog/internal/api"
	v1 "blog/internal/api/v1"
	"blog/internal/middleware"
	"blog/pkg/limiter"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	_ "blog/docs"

	sf "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
)

func limiters() limiter.Limiter {
	return limiter.NewMethodLimit().AddBuckets(
		[]limiter.LimitBucketRule{
			{
				Key:          "/auth",
				FillInterval: time.Second,
				Capacity:     10,
				Quantum:      10,
			},
		}...,
	)
}

func NewRouter() *gin.Engine {
	r := gin.New()
	if global.ServerSetting.RunMode == "debug" {
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	} else {
		r.Use(middleware.AccessLog())
		r.Use(middleware.Recovery())
	}
	r.Use(middleware.RateLimter(limiters()))
	r.Use(middleware.ContextTimeout(global.AppSetting.DefaultContextTimeout * time.Second))
	r.Use(middleware.Translations())
	r.Use(middleware.AccessLog())
	r.Use(middleware.AppInfo())
	r.Use(middleware.Tracing())

	r.GET("/swagger/*any", gs.WrapHandler(sf.Handler))
	upload := api.NewUpload()
	r.POST("/upload", upload.UploadFile)
	r.POST("/auth", api.GetAuth)
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))

	apiV1 := r.Group("/api/v1")
	apiV1.Use(middleware.JWT())

	{

		blog := v1.NewBlogApi()

		apiV1.POST("/blog", blog.Create)
		apiV1.GET("/blog", blog.List)
		apiV1.GET("/blog/:id", blog.Get)
		apiV1.PUT("/blog/:id", blog.Update)
		apiV1.DELETE("/blog/:id", blog.Delete)
	}
	return r
}

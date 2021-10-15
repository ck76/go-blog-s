package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "go-blog-s/docs"
	"go-blog-s/global"
	"go-blog-s/internal/middleware"
	"go-blog-s/internal/routers/api"
	v1 "go-blog-s/internal/routers/api/v1"
	"net/http"
)

func NewRouter() *gin.Engine {
	engine := gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	engine.Use(middleware.Translation())
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	article := v1.NewArticle()
	tag := v1.NewTag()
	upload := api.NewUpload()
	engine.POST("/upload/file", upload.UploadFile)
	//访问 $HOST/static 时，应用程序会读取到 blog-service/storage/uploads 下的文件
	engine.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))
	apiv1 := engine.Group("/api/v1")
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
	return engine
}

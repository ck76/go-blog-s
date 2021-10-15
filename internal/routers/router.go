package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "go-blog-s/docs"
	"go-blog-s/global"
	"go-blog-s/internal/middleware"
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
	upload := v1.NewUpload()
	engine.POST("/upload/file", upload.UploadFile)
	engine.POST("/auth", v1.GetAuth)
	//访问 $HOST/static 时，应用程序会读取到 blog-service/storage/uploads 下的文件
	engine.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))
	group := engine.Group("/api/v1")
	group.Use(middleware.JWT())
	{
		group.POST("/tags", tag.Create)
		group.DELETE("/tags/:id", tag.Delete)
		group.PUT("/tags/:id", tag.Update)
		group.PATCH("/tags/:id/state", tag.Update)
		group.GET("/tags", tag.List)

		group.POST("/articles", article.Create)
		group.DELETE("/articles/:id", article.Delete)
		group.PUT("/articles/:id", article.Update)
		group.PATCH("/articles/:id/state", article.Update)
		group.GET("/articles/:id", article.Get)
		group.GET("/articles", article.List)
	}
	return engine
}

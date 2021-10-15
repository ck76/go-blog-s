package v1

import (
	"github.com/gin-gonic/gin"
	"go-blog-s/internal/service"
	"go-blog-s/pkg/app"
	"go-blog-s/pkg/convert"
	"go-blog-s/pkg/errcode"
	"go-blog-s/pkg/upload"
)

type Upload struct{}

func NewUpload() Upload {
	return Upload{}
}

//curl -X POST http://127.0.0.1:8000/upload/file -F file=@"/Users/chengkun02/Downloads/a.txt" -F type=1
//
//{"file_access_url":"http://127.0.0.1:8000/static/0cc175b9c0f1b6a831c399e269772661.txt"}%
func (u Upload) UploadFile(c *gin.Context) {
	response := app.NewResponse(c)
	file, fileHeader, err := c.Request.FormFile("file")
	//https://www.cnblogs.com/yh2924/p/12377309.html 多文件上传
	//form,err:=c.MultipartForm()
	//files:=form.File["f1s"]
	if err != nil {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}

	fileType := convert.StrTo(c.PostForm("type")).MustInt()
	if fileHeader == nil || fileType <= 0 {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}

	svc := service.New(c.Request.Context())
	fileInfo, err := svc.UploadFile(upload.FileType(fileType), file, fileHeader)
	if err != nil {
		//global.Logger.Errorf(c, "svc.UploadFile err: %v", err)
		response.ToErrorResponse(errcode.ErrorUploadFileFail.WithDetails(err.Error()))
		return
	}

	response.ToResponse(gin.H{
		"file_access_url": fileInfo.AccessUrl,
	})
}

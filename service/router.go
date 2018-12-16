package service

import (
	"github.com/gin-gonic/gin"
	"github.com/rakyll/statik/fs"
	"log"
	"net/http"
)

// Router ...
func Router(engine *gin.Engine) error {
	//api document
	//engine.Static("/doc", "./doc")
	st, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}
	engine.StaticFS("/doc", st)

	ver := "/v1"
	group := engine.Group(ver)
	group.Use()
	group.Static("/transfer", "./transfer")

	group.Any("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "%s", "hello world")
	})

	//上传文件
	group.POST("/upload", UploadPost(ver))

	//视频转换
	group.POST("/transfer", TransferPost(ver))

	//下载并转换
	group.POST("/downloadTransform", func(ctx *gin.Context) {

	})

	//服务器视频列表
	group.GET("/list", ListGet(ver))

	//查看状态
	group.GET("/info/:id", InfoGet(ver))

	return nil
}
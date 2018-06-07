package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/devfeel/dotweb"
)

func stock(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("filepond.html")
	t.Execute(w, nil)
}

func initServer() {
	app := dotweb.New()
	app.SetDevelopmentMode()
	//启用访问日志
	app.SetEnabledLog(true)
	app.UseRequestLog()

	//设置路由
	InitRoute(app.HttpServer)

	// 开始服务
	port := 8080
	fmt.Println("dotweb.StartServer => " + strconv.Itoa(port))
	err := app.StartServer(port)
	fmt.Println("dotweb.StartServer error => ", err)
}

func InitRoute(server *dotweb.HttpServer) {
	server.Router().POST("/file", FileUpload)
}

func FileUpload(ctx dotweb.Context) error {
	upload, err := ctx.Request().FormFile("file")
	if err != nil {
		err := ctx.WriteString("FormFile error " + err.Error())
		return err
	} else {
		_, err = upload.SaveFile("/Users/hgz/Desktop/Projects/" + upload.FileName())
		updateOrder("/Users/hgz/Desktop/Projects/" + upload.FileName())
		if err != nil {
			err := ctx.WriteString("SaveFile error => " + err.Error())
			return err
		} else {
			err := ctx.WriteString("SaveFile success || " + upload.FileName() + " || " + upload.GetFileExt() + " || " + fmt.Sprint(upload.Size()))

			return err
		}
	}

}

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"goproject/service"
	"io"
	"net/http"
	"os"
)
import "goproject/greet"

func init() {
	fmt.Println("config log")
	//logPath, _ := os.Getwd() // 后期可以配置
	//logName := fmt.Sprintf("%s/access_log.", logPath)
	//r, _ := rotatelogs.New(logName + "%Y%m%d")
	//mw := io.MultiWriter(os.Stdout, r)
	//log.SetOutput(mw)
	//log.Info("something ....")
}

func init() {
	fmt.Println("call init() b...")

	log.SetLevel(log.DebugLevel)
	logFile, _ := os.OpenFile("log/info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
		TimestampFormat: "2006-01-02.15:04:05.000000",
	})

}




func main() {

	runWeb()
}

func runWeb()  {
	// 创建一个默认的路由引擎
	engine := gin.Default()
	engine.MaxMultipartMemory = 8 << 20 // 8MB

	// GET：请求方式；/hello：请求的路径
	// 当客户端以GET方法请求/hello路径时，会执行后面的匿名函数
	engine.GET("/hello", func(c *gin.Context) {
		// c.JSON：返回JSON格式的数据
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world!",
		})
	})


	engine.GET("/user/:name", service.ParameterInPath1)
	engine.GET("/user/:name/*action", service.ParameterInPath2)
	engine.GET("/welcome", service.QueryStringParam)


	engine.POST("/upload", service.UploadSingleFile)
	engine.POST("/loginJSON", service.BindWithJson)


	// 启动HTTP服务，默认在0.0.0.0:8080启动服务
	if err := engine.Run(); err != nil {
		fmt.Printf("startup service failed, err: %v\n\n", err)
	}
}



func learnMath() {
	fmt.Println(greet.Morning)
}








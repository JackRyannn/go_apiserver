package main

import (
	"apiserver/config"
	"apiserver/model"
	"apiserver/router"
	"apiserver/router/middleware"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"time"

	//配置读取库
	"github.com/spf13/pflag"
	//日志包
	"github.com/lexkong/log"
)

var (
	cfg = pflag.StringP("config", "c", "", "apiserver config file path.")
)

func main() {
	//var s string
	//for {
	//
	//	fmt.Scanln(&s)
	//	println(service.Receive(s))
	//}

	pflag.Parse()

	// init config
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}
	//初始化数据库

	model.DB.Init()
	defer model.DB.Close()

	fmt.Println("chaoren:" + viper.GetString("runmode"))
	// Set gin mode.
	gin.SetMode(viper.GetString("runmode"))

	// Create the Gin engine.
	g := gin.New()
	g.Static("/templates", "./templates")
	//添加和加载中间件，但是logging本身就实现了，其实去掉也没差
	middlewares := []gin.HandlerFunc{}
	middlewares = append(middlewares, middleware.Logging())
	// Routes.
	router.Load(
		// Cores.
		g,
		//middleware.Logging(),
		// Middlwares.
		middlewares...,
	)
	//加载templates下的web页面
	g.LoadHTMLGlob("templates/*")

	// Ping the server to make sure the router is working.
	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Info("The router has been deployed successfully.")
	}()

	log.Infof("Start to listening the incoming requests on http address: %s", ":80")
	log.Infof(http.ListenAndServe(":8080", g).Error())
}

// pingServer pings the http server to make sure the router is working.
func pingServer() error {
	for i := 0; i < 2; i++ {
		// Ping the server by sending a GET request to `/health`.
		resp, err := http.Get("http://127.0.0.1:8080" + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		// Sleep for a second to continue the next ping.
		log.Info("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the router.")
}

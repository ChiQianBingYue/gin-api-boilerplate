package main

import (
	"fmt"

	"gin-api-boilerplate/middleware"
	"gin-api-boilerplate/routes"

	"github.com/ChiQianBingYue/go-api-lib/config"
	"github.com/ChiQianBingYue/go-api-lib/log"
	"github.com/ChiQianBingYue/go-api-lib/mongo"
	"github.com/gin-gonic/gin"
)

func main() {
	name := config.GetString("name")
	port := config.GetString("port")

	// 初始化默认redis db, 后面在使用的时候import "github.com/ihahoo/go-api-lib/redis" 通过redis.DB(0)调用实例
	// 如果要初始化多个redis的db，则在这里添加，比如redis.Init(1)就建立了一个db 1的连接
	// 如果不使用redis，则删除这里以及其它和redis相关的包引入
	// redis.Init(0)

	// 初始化数据库
	mongo.Init()

	r := gin.Default()
	r.HandleMethodNotAllowed = true
	r.Use(middleware.NoCache)

	r.Use(middleware.Logger(log.GetLog()), gin.Recovery())

	routes.Routes(r)

	r.NoMethod(func(c *gin.Context) {
		c.JSON(405, gin.H{"errcode": 405, "errmsg": "Method Not Allowed"})
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"errcode": 404, "errmsg": "Not Found"})
	})

	log.GetLog().Info(name + " start listening at http://loaclhost:" + port)
	fmt.Printf("==> 🚀 %s listening at %s\n", name, port)
	r.Run(":" + port)
}

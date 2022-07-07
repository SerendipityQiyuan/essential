package main

import (
	"awesomeProject1/common"
	"awesomeProject1/routes"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
)

func main() {
	//初始化配置
	InitConfig()
	//运行数据库初始化
	common.InitDB()
	// 创建路由
	r := gin.Default()
	routes.CollectRoute(r)
	port := viper.GetString("server.port")
	err := r.Run(port)
	if err != nil {
		return
	}
}

// InitConfig 获取配置文件参数
func InitConfig() {
	workdir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workdir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic("config read failed,err:" + err.Error())

	}
}

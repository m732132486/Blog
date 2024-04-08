package main

import (
	"fmt"
	"practice/dao/mysql"
	"practice/dao/redis"
	"practice/logger"
	snowflake "practice/pkg"
	"practice/routes"
	"practice/settings"

	"go.uber.org/zap"
)

func main() {
	//初始化配置文件
	if err := settings.Init(); err != nil {
		fmt.Println("配置文件初始化失败")
		return
	}
	fmt.Println("配置文件初始化成功")
	//初始化日志
	if err := logger.InitLogger(settings.Conf.LogConfig); err != nil {
		fmt.Println("日志初始化失败")
		return
	} else {
		fmt.Println("日志初始化成功")
	}

	defer zap.L().Sync()

	if err := mysql.Init(settings.Conf.MySQLConfig); err != nil {
		fmt.Println("mysql连接失败")
		return
	} else {
		fmt.Println("MySQL连接成功......")
		//defer Close.Close()
	}

	if err := redis.Init(settings.Conf.RedisConfig); err != nil {
		fmt.Println("redis连接失败")
		return
	} else {
		fmt.Println("Redis连接成功......")
		defer redis.Close()
	}
	if err := snowflake.Init(settings.Conf.StartTime, settings.Conf.MachineID); err != nil {
		fmt.Println("雪花算法初始化失败", err)
		return
	}
	//注册路由
	routes.Setup()
	//redis.V1()
}

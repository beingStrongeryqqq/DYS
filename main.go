package main

import (
	"DYS/controllers"
	"DYS/dao/mysql"
	"DYS/dao/redis"
	"DYS/logger"
	"DYS/pkg/snowflake"
	"DYS/routes"
	"DYS/settings"
	"fmt"
	"go.uber.org/zap"
)

func main() {
	//1.加载配置
	if err := settings.Init("D:\\DYS\\config.yaml"); err != nil {
		fmt.Println("failed, err:%v\n", err)
		return
	}
	//2.初始化日志
	if err1 := logger.Init(settings.Conf.LogConfig, settings.Conf.Mode); err1 != nil {
		fmt.Println("failed, err1:%v\n", err1)
		return
	}
	defer zap.L().Sync()
	zap.L().Debug("logger init success...")
	//3.初始化MySQL连接
	if err := mysql.Init(settings.Conf.MySQLConfig); err != nil {
		fmt.Println("failed, err:%v\n", err)
		return
	}
	defer mysql.Close()
	//4.初始化Redis连接
	if err := redis.Init(settings.Conf.RedisConfig); err != nil {
		fmt.Println("failed, err:%v\n", err)
		return
	}
	defer redis.Close()
	if err := snowflake.Init(settings.Conf.StartTime, settings.Conf.MachineID); err != nil {
		fmt.Println("failed, err:%v\n", err)
		return
	}

	//初始化gin框架内置的校验器使用的翻译器
	if err := controllers.InitTrans("zh"); err != nil {
		fmt.Printf("init validator trans failed, err:%v\n", err)
		return
	}
	//5.注册路由
	r := routes.SetUp(settings.Conf.Mode)
	err := r.Run(fmt.Sprintf(":%d", settings.Conf.Port))
	if err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}
}

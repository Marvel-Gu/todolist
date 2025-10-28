package main

import (
	"github.com/Marvel-Gu/todolist/conf"
	"github.com/Marvel-Gu/todolist/interfaces/initilaize"
)

func main() {
	// 初始化配置文件
	conf.InitConfig()

	// 初始化路由
	r := initilaize.NewRouter()
	// 启动服务
	_ = r.Run(conf.Conf.Server.Port)
}

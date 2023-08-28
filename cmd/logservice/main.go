package main

import (
	"context"
	"distributed/log"
	"distributed/registry"
	"distributed/service"
	"fmt"
	Stlog "log" //引入标准库log
)

func main() {
	log.Run("./distributed.log")      //文件地址
	host, port := "localhost", "4000" //地址和端口
	serviceAddress := fmt.Sprintf("http://%s:%s", host, port)
	r := registry.Registration{
		ServiceName: registry.LogService,
		ServiceURL:  serviceAddress,
	}
	ctx, err := service.Start(
		context.Background(),
		//"log serviceName",
		host,
		port,
		r,
		log.RegisterHandlers, //注册函数
	)
	if err != nil {
		Stlog.Fatalln(err) //标准库的log 把错误传进去
	}
	<-ctx.Done()                              //等ctx的信号  ->service ->go func()->cancel()
	fmt.Println("Shutting down log service.") //打印 关闭日志服务
}

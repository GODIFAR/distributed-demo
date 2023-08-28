package service

import (
	"context"
	"distributed/registry"
	"fmt"
	"log"
	"net/http"
)

func Start(ctx context.Context, host, port string, reg registry.Registration, registerHanlersFunc func()) (context.Context, error) {
	registerHanlersFunc()

	ctx = startService(ctx, reg.ServiceName, host, port)

	err := registry.RegisterService(reg)
	if err != nil {
		return ctx, err
	}

	return ctx, nil
}

func startService(ctx context.Context, serviceName registry.ServiceName, host, port string) context.Context {
	ctx, cancel := context.WithCancel(ctx) //取消

	var srv http.Server
	srv.Addr = ":" + port //设置地址

	go func() {
		log.Println(srv.ListenAndServe()) //打印出错误
		err := registry.ShutdownService(fmt.Sprintf("http://%s:%s", host, port))
		if err != nil {
			log.Println(err)
		}
		cancel() //发生错误就取消context取消上下文
	}()

	go func() {
		fmt.Printf("%v started. Press any key to stop.\n", serviceName) //按任意键停止
		var s string
		fmt.Scanln(&s) //如果按任何键，那么就继续往下走，否则就停在那等待用户的输入
		err := registry.ShutdownService(fmt.Sprintf("http://%s:%s", host, port))
		if err != nil {
			log.Println(err)
		}
		srv.Shutdown(ctx) //继续往下走就要关闭服务
		cancel()
	}()
	return ctx
}

//  func Start(ctx context.Context, serviceName, host, port string, registerHanlersFunc func()) (context.Context, error) {
// 	registerHanlersFunc()

// 	ctx = startService(ctx, serviceName, host, port)

// 	return ctx, nil
// }

// func startService(ctx context.Context, serviceName, host, port string) context.Context {
// 	ctx, cancel := context.WithCancel(ctx) //取消

// 	var srv http.Server
// 	srv.Addr = ":" + port //设置地址

// 	go func() {
// 		log.Println(srv.ListenAndServe()) //打印出错误
// 		cancel()                          //发生错误就取消context取消上下文
// 	}()

// 	go func() {
// 		fmt.Printf("%v started. Press any key to stop.\n", serviceName) //按任意键停止
// 		var s string
// 		fmt.Scanln(&s)    //如果按任何键，那么就继续往下走，否则就停在那等待用户的输入
// 		srv.Shutdown(ctx) //继续往下走就要关闭服务
// 		cancel()
// 	}()
// 	return ctx
//  }

package main

import (
	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	"github.com/ZengQingAn/go-server/service"
)

func main() {
	config.SetProviderService(new(service.HelloService))
	if err := config.Load(config.WithPath("conf/dubbogo.yml")); err != nil {
		panic(err)
	}
	select {}
}

package main

import (
	"fmt"
	"go_templates/proxy"
)

func proxyWorker() {
	fmt.Println("---------- Proxy -----------")

	warechouse := proxy.ProxyWarehouse{}
	if res, ok := warechouse.GetData(1); ok {
		fmt.Println(res)
	}

	fmt.Println("-------------------")
}

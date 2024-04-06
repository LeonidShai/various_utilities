package proxy

import (
	"fmt"
)

func ProxyWorker() {
	fmt.Println("---------- Proxy -----------")

	warechouse := ProxyWarehouse{}
	if res, ok := warechouse.GetData(1); ok {
		fmt.Println(res)
	}

	fmt.Println("-------------------")
}

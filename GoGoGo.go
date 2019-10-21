package main

import (
	"github.com/wongyinlong/walletBackend/start"
	"log"
)

func main() {
	// 程序入口
	// start wallet
	err := start.OnStart()
	if err != nil {
		log.Fatal(err)
	}

}

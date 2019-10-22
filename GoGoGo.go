package main

import (
	"github.com/wongyinlong/walletBackend/conf"
	"github.com/wongyinlong/walletBackend/start"
	"log"
)

func init() {
	// check Settings
	config := conf.NewConfig()
	if config.Db.DbString == "" {
		log.Fatal("dbString is empty! Please Check 'conf/config.go'")
		return
	}

	if config.Service.IP == "" || config.Service.Port == "" {
		log.Fatal("IP and Port can not be empty! Please Check 'conf/config.go'")
		return
	}

	if config.JiG.AppKey == "" || config.JiG.AuthString == "" {
		if config.JiG.MasterSecret == "" {
			log.Fatal("AppKey and AuthString OR MasterSecret Can not be empty!  Please Check 'conf/config.go'")
			return
		}
	}
}

func main() {
	// start main process
	err := start.OnStart()
	if err != nil {
		log.Fatal(err)
	}

}

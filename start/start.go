package start

import (
	"github.com/wongyinlong/walletBackend/conf"
	"github.com/wongyinlong/walletBackend/middleware"
	"github.com/wongyinlong/walletBackend/rest"
	"github.com/wongyinlong/walletBackend/rest/wallet"
	"net/http"
	"time"
)

var IPAddr string
var ReadTimeout time.Duration
var WriteTimeout time.Duration

func init(){

	config := conf.NewConfig()
	IPAddr = config.Service.IP+":"+config.Service.Port
	ReadTimeout = config.Service.ReadTimeout * time.Second
	WriteTimeout = config.Service.WriteTimeout * time.Second
}

func OnStart() error {
	apiServer := registeredFunc()
	srv := http.Server{
		Addr:         IPAddr,
		Handler:      apiServer.Mux,
		ReadTimeout:  ReadTimeout,
		WriteTimeout: WriteTimeout,
	}
	middleware.SetMiddleWare(apiServer)
	err := srv.ListenAndServe()
	return err
}
func registeredFunc() *rest.ApiServer {
	restServer := rest.NewApiServer()

	// 注册 手机app的url
	front := restServer.Mux.PathPrefix("/v1/wallet").Subrouter()
	wallet.RegisterFrontApi(front)

	// 注册管理后台的url
	backend := restServer.Mux.PathPrefix("/v1/backend").Subrouter()
	wallet.RegisterBackendApi(backend)

	// other urls
	//jPushURl := restServer.Mux.PathPrefix("/api/v1/jPush").Subrouter()

	return restServer
}

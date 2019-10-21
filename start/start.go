package start

import (
	"github.com/wongyinlong/walletBackend/rest"
	"github.com/wongyinlong/walletBackend/rest/wallet"
	"net/http"
	"time"
)

func OnStart() error {
	apiServer := registeredFunc()
	srv := http.Server{
		Addr:         "127.0.0.1:7789",
		Handler:      apiServer.Mux,
		ReadTimeout:  time.Second * 15,
		WriteTimeout: time.Second * 15,
	}

	err := srv.ListenAndServe()
	return err
}
func registeredFunc() *rest.ApiServer {
	restServer := rest.NewApiServer()
	walletUrl := restServer.Mux.PathPrefix("/api/v1/wallet").Subrouter()
	wallet.RegisterPublishInfoApi(walletUrl)
	wallet.RegisterGetInfoApi(walletUrl)
	// other urls
	//jPushURl := restServer.Mux.PathPrefix("/api/v1/jPush").Subrouter()


	return restServer
}

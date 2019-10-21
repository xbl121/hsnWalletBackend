package wallet

import (
	"github.com/gorilla/mux"
	"github.com/wongyinlong/walletBackend/handler/wallet"
	"net/http"
)

func RegisterPublishInfoApi(r *mux.Router) {
	// 发布通告
	r.HandleFunc("/teat",wallet.TestFunc).Methods(http.MethodGet)
	//发布活动
	//server.Mux.HandleFunc("/publishNotice", handle.TestFunc).Methods("POST")
}

package wallet

import (
	"github.com/gorilla/mux"
	"github.com/wongyinlong/walletBackend/handler/wallet"
	"net/http"
)

func RegisterPublishInfoApi(r *mux.Router) {
	// 发布通告
	r.HandleFunc("/notice", wallet.PostNotice).Methods(http.MethodPost)
	// 发布活动
	r.HandleFunc("/events", wallet.PostEvents).Methods(http.MethodPost)
	// 提交交易查询
	r.HandleFunc("/pushTx", wallet.PostTx).Methods(http.MethodPost)
}

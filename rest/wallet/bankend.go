package wallet

import (
	"github.com/gorilla/mux"
	"github.com/wongyinlong/walletBackend/handler/wallet"
	"net/http"
)

func RegisterBackendApi(r *mux.Router) {
	// 获取添加/更新/删除/获取 活动信息
	r.HandleFunc("/event", wallet.HelpWord).Methods(http.MethodPost)
	r.HandleFunc("/event", wallet.HelpWord).Methods(http.MethodGet)
	r.HandleFunc("/event", wallet.HelpWord).Methods(http.MethodPut)
	r.HandleFunc("/event", wallet.HelpWord).Methods(http.MethodDelete)
	// 获取/删除历史活动
	r.HandleFunc("/historyEvents", wallet.PostEvents).Methods(http.MethodGet)
	r.HandleFunc("/historyEvents", wallet.PostEvents).Methods(http.MethodDelete)

	// 发布/或许通告信息
	r.HandleFunc("/notice", wallet.PostNotice).Methods(http.MethodPost)
	r.HandleFunc("/notice", wallet.PostNotice).Methods(http.MethodGet)

	// 登陆/推出
	r.HandleFunc("/login", wallet.PostTx).Methods(http.MethodPost)
	r.HandleFunc("/logout", wallet.APMRInsert).Methods(http.MethodPost)

}

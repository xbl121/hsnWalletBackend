package wallet

import (
	"github.com/gorilla/mux"
	"github.com/wongyinlong/walletBackend/handler/wallet"
	"net/http"
)

func RegisterGetInfoApi(r *mux.Router) {
	// 获取助记词
	r.HandleFunc("/words", wallet.HelpWord).Methods(http.MethodPost)
	// 获取首页banner信息
	r.HandleFunc("/banner", wallet.Banner).Methods(http.MethodGet)
	// 获取历史活动信息
	r.HandleFunc("/historyNotice", wallet.Notice).Methods(http.MethodGet)
	// 获取历史提示信息
	r.HandleFunc("/historyEvents", wallet.Events).Methods(http.MethodGet)

}

package wallet

import (
	"github.com/gorilla/mux"
	"github.com/wongyinlong/walletBackend/handler/wallet"
	"net/http"
)

func RegisterFrontApi(r *mux.Router) {
	// 获取助记词
	r.HandleFunc("/words", wallet.HelpWord).Methods(http.MethodPost)
	// 获取首页banner信息
	r.HandleFunc("/banner", wallet.Banner).Methods(http.MethodGet)
	// 获取活动信息
	r.HandleFunc("/event", wallet.Events).Methods(http.MethodGet)
	// 获取历史活动信息
	r.HandleFunc("/historyEvents", wallet.Events).Methods(http.MethodGet)
	// 获取通告信息
	r.HandleFunc("/notice", wallet.Events).Methods(http.MethodGet)
	// 提交交易查询
	r.HandleFunc("/tx", wallet.PostTx).Methods(http.MethodPost)
	// 处理账户和手机的对应关系
	r.HandleFunc("/jPushId", wallet.APMRInsert).Methods(http.MethodPost)
	r.HandleFunc("/jPushId", wallet.APMRRemove).Methods(http.MethodDelete)
	//r.HandleFunc("/jPushId", wallet.APMRUpdate).Methods(http.MethodPut)

}

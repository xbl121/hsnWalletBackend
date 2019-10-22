package wallet

import (
	"github.com/gorilla/mux"
	"github.com/wongyinlong/walletBackend/handler/backend"
	"net/http"
)

func RegisterBackendApi(r *mux.Router) {
	// 获取添加/更新/删除/获取 活动信息
	r.HandleFunc("/event", backend.PostEvent).Methods(http.MethodPost)
	r.HandleFunc("/event", backend.GetEvent).Methods(http.MethodGet)
	r.HandleFunc("/event", backend.PutEvent).Methods(http.MethodPut)
	r.HandleFunc("/event", backend.DeleteEvent).Methods(http.MethodDelete)
	// 获取/删除历史活动
	r.HandleFunc("/historyEvents", backend.GetEvents).Methods(http.MethodGet)
	r.HandleFunc("/historyEvents", backend.DeleteEvents).Methods(http.MethodDelete)

	// 发布/或许通告信息
	r.HandleFunc("/notice", backend.PostNotice).Methods(http.MethodPost)
	r.HandleFunc("/notice", backend.GetNotice).Methods(http.MethodGet)

	// 登陆/推出
	r.HandleFunc("/login", backend.LogIn).Methods(http.MethodPost)
	r.HandleFunc("/logout", backend.LogOut).Methods(http.MethodPost)

}

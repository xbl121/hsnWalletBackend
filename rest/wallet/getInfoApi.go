package wallet

import "github.com/gorilla/mux"

func RegisterGetInfoApi(r *mux.Router) {
	// 获取助记词
	r.HandleFunc("/publishNotice", handle.TestFunc()).Methods("POST")
	//// 获取首页banner信息
	//r.HandleFunc("/publishNotice", handle.TestFunc()).Methods("POST")
	//// 获取历史活动信息
	//r.HandleFunc("/publishNotice", handle.TestFunc()).Methods("POST")
	//// 获取历史提示信息
	//r.HandleFunc("/publishNotice", handle.TestFunc()).Methods("POST")

}

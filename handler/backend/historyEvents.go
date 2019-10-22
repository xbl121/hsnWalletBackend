package backend

import (
	"github.com/wongyinlong/walletBackend/models/moves"
	"net/http"
)

func GetEvents(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err !=nil{
		// 返回错误信息，
	}
	word := r.PostFormValue("word")
	var helpWordsList moves.Words
	//get data form db
	helpWordsList.GetHelpWords(word)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type","application/Json")
	_, _ = w.Write([]byte("hello"))

	//_, _ = w.Write([]byte("hello"))
}
func DeleteEvents(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err !=nil{
		// 返回错误信息，
	}
	word := r.PostFormValue("word")
	var helpWordsList moves.Words
	//get data form db
	helpWordsList.GetHelpWords(word)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type","application/Json")
	_, _ = w.Write([]byte("hello"))

	//_, _ = w.Write([]byte("hello"))
}
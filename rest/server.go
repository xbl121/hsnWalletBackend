package rest

import (
	"github.com/gorilla/mux"
	"github.com/wongyinlong/walletBackend/logger"
	"go.uber.org/zap"
	localLog "log"
	"net/http"
)

type ApiServer struct {
	Mux *mux.Router
	Log zap.Logger
	//listener net.Listener
}

func NewApiServer() *ApiServer {
	r := mux.NewRouter()
	r.Use(mux.CORSMethodMiddleware(r))

	r.Use(printURl)
	r.Use(setHttpHeader)
	log := logger.NewLogger()
	return &ApiServer{
		Mux: r,
		Log: log,
	}
}

func printURl(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		localLog.Println(r.URL)
		next.ServeHTTP(w, r)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
	})
}

func setHttpHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		w.Header().Set("Access-Control-Allow-Origin", "*")
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
		// after handler
		w.Header().Set("Content-Type", "application/json")
	})
}

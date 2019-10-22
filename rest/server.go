package rest

import (
	"github.com/gorilla/mux"
	"github.com/wongyinlong/walletBackend/logger"
	"go.uber.org/zap"
)

type ApiServer struct {
	Mux *mux.Router
	Log zap.Logger
	//listener net.Listener
}

func NewApiServer() *ApiServer {
	r := mux.NewRouter()
	log := logger.NewLogger()
	return &ApiServer{
		Mux: r,
		Log: log,
	}
}


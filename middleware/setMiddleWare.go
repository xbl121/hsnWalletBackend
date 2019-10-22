package middleware

import (
	"github.com/gorilla/mux"
	"github.com/wongyinlong/walletBackend/rest"
)

func SetMiddleWare(server *rest.ApiServer) *rest.ApiServer{
	server.Mux.Use(mux.CORSMethodMiddleware(server.Mux))
	server.Mux.Use(printURl)
	server.Mux.Use(setHttpHeader)
	server.Mux.Use(JWT)
	return server
}


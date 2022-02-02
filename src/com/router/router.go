package router

import (
	//"example.com/http_demo/handler/user"
	//"example.com/http_demo/handler/ws"
	//"example.com/http_demo/middleware"
	"github.com/gorilla/mux"
	"go-reco-service/src/com/handler"
	"net/http"
	//"example.com/http_demo/handler"
)

func RegisterRoutes(r *mux.Router) {
	// serve static file request
	fs := http.FileServer(http.Dir("assets/"))
	serveFileHandler := http.StripPrefix("/static/", fs)
	r.PathPrefix("/static/").Handler(serveFileHandler)

	// apply Logging middleware
	//r.Use(middleware.Logging(), middleware.AccessLogging)

	indexRouter := r.PathPrefix("/index").Subrouter()
	//indexRouter.Handle("/", &handler.HelloHandler{})
	//indexRouter.HandleFunc("/password_hashing", handler.PassWordHashingHandler)
	//indexRouter.HandleFunc("/display_headers", handler.DisplayHeadersHandler)
	//indexRouter.HandleFunc("/display_url_params", handler.DisplayUrlParamsHandler)
	//indexRouter.HandleFunc("/display_form_data", handler.DisplayFormDataHandler).Methods("POST")
	//indexRouter.HandleFunc("/read_cookie", handler.ReadCookieHandler)
	indexRouter.HandleFunc("/parse_json_request", handler.ParseJsonRequestHandler)
	indexRouter.HandleFunc("/parse_json_request2", handler.Test)
	indexRouter.HandleFunc("/get_json_response", handler.WriteJsonResponseHandler)

	//userRouter := r.PathPrefix("/user").Subrouter()
	//userRouter.HandleFunc("/names/{name}/countries/{country}", handler.ShowVisitorInfo)
	//userRouter.HandleFunc("/login", user.Login).Methods("POST")
	//userRouter.HandleFunc("/secret", user.Secret)
	//userRouter.HandleFunc("/logout", user.Logout)
	//
	//viewRouter := r.PathPrefix("/view").Subrouter()
	//viewRouter.HandleFunc("/index", handler.ShowIndexView)
	//
	//wsRouter := r.PathPrefix("/ws").Subrouter()
	//wsRouter.HandleFunc("/echo", ws.EchoMessage)
	//wsRouter.HandleFunc("/echo_display", ws.DisplayEcho)
}

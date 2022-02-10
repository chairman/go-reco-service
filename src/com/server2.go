package main

import (
	"context"
	"github.com/gorilla/mux"
	_ "go-reco-service/src/com/drivers"
	"go-reco-service/src/com/router"
	"go-reco-service/src/com/utils/vlog"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	muxRouter := mux.NewRouter()

	// register route handlers
	router.RegisterRoutes(muxRouter)

	// set error log writer
	errorWriter := vlog.ErrorLog.Writer()
	defer errorWriter.Close()

	server := &http.Server{
		Addr:     ":8080",
		Handler:  muxRouter,
		ErrorLog: log.New(vlog.ErrorLog.Writer(), "", 0),
	}

	// 创建系统信号接收器
	done := make(chan os.Signal)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-done

		if err := server.Shutdown(context.Background()); err != nil {
			log.Fatal("Shutdown server:", err)
		}
	}()

	log.Println("Starting HTTP server...")
	err := server.ListenAndServe()
	if err != nil {
		if err == http.ErrServerClosed {
			log.Print("Server closed under request")
		} else {
			log.Fatal("Server closed unexpected")
		}
	}
}

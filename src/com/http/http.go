package http

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-reco-service/src/com/handler"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func Lanuch(serverAddr string) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	auditcbHandler := r.Group("/reco/config/v1")
	auditcbHandler.POST("/:app_name/:res_type", handler.AddHandler)
	auditcbHandler.PUT("/:app_name/:res_type", handler.UpdateHandler)
	auditcbHandler.DELETE("/:app_name/:res_type/:rule_id", handler.DeleteHandler)
	auditcbHandler.GET("/:app_name/:res_type/:rule_id", handler.GetHandler)

	r.GET("/", func(c *gin.Context) {
		c.Redirect(302, "https://www.baidu.com/")
	})

	srv := &http.Server{
		Addr:    serverAddr,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Listen: %s", err)
		}
	}()

	quit := make(chan os.Signal)

	signal.Notify(quit, os.Interrupt)

	fmt.Printf("----------------- auditcb startup ok!-----------------")
	fmt.Println()
	fmt.Println()

	<-quit

	log.Println("Stutdown Server ... ")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")
}

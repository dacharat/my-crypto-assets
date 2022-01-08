package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dacharat/my-crypto-assets/cmd/api/route"
	"github.com/dacharat/my-crypto-assets/pkg/config"
)

func main() {
	config.NewConfig()

	router := route.NewRouter()

	httpSrv := &http.Server{
		Addr:    fmt.Sprintf(":%s", config.Cfg.Port),
		Handler: router,
	}

	// 1 set up handler to go routine
	go func() {
		if err := httpSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	// 2 make channel for listen os.Siganal and setup notify Signal
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	// 3 setup withTimeout to preserve connection before close
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := httpSrv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}
	log.Println("Server exiting")
}

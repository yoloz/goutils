package main

import (
	"birthdayreminder/control"
	"birthdayreminder/server"
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	initDb := flag.Bool("init", false, "init database")
	flag.Parse()
	if *initDb {
		log.Println("init database......")
		server.InitDb()
	}
	log.Println("start email notify......")
	go server.EmailNotify()

	srv := http.Server{
		Addr:         ":8004",
		WriteTimeout: time.Second * 10,
	}

	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		// We received an interrupt signal, shut down.
		if err := srv.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("HTTP server Shutdown: %v", err)
		}
		close(idleConnsClosed)
	}()

	log.Println("start http server......")
	if err := control.StartHttpServer(srv); err != http.ErrServerClosed {
		// Error starting or closing listener:
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}

	<-idleConnsClosed

}

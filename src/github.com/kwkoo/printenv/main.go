package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Print("Request for URI: ", r.URL.Path)

	w.Header().Set("Content-Type", "text/plain")
	dumpEnv(w)
}

func main() {
	dumpEnv(os.Stdout)

	var port int

	flag.IntVar(&port, "port", 8080, "HTTP listener port")
	flag.Parse()

	env := getPortEnv()
	if env > 0 {
		port = env
	}

	// Setup signal handling.
	shutdown := make(chan os.Signal)
	signal.Notify(shutdown, os.Interrupt)

	var wg sync.WaitGroup
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: http.HandlerFunc(handler),
	}
	go func() {
		log.Printf("listening on port %v", port)
		http.HandleFunc("/", handler)
		wg.Add(1)
		defer wg.Done()
		if err := server.ListenAndServe(); err != nil {
			if err == http.ErrServerClosed {
				log.Print("web server graceful shutdown")
				return
			}
			log.Fatal(err)
		}
	}()

	// Wait for SIGINT
	<-shutdown
	log.Print("interrupt signal received, initiating web server shutdown...")
	signal.Reset(os.Interrupt)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	server.Shutdown(ctx)

	wg.Wait()
	log.Print("Shutdown successful")
}

func getPortEnv() int {
	s := os.Getenv("PORT")
	if len(s) == 0 {
		return 0
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}

func dumpEnv(w io.Writer) {
	for _, env := range os.Environ() {
		fmt.Fprintln(w, env)
	}
}

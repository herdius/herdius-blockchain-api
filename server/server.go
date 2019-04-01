package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"github.com/herdius/herdius-blockchain-api/handler"
)

func main() {
	log.Println("Opening API")
	//test()
	LaunchServer()
}
func test() {
	fmt.Println("test it")

	ctx := make(chan *TX)
	go func() {
		i := 1
		for {
			time.Sleep(time.Millisecond * 1000)
			txpool := GetMemPool()

			txpool.data = append(txpool.data, i)
			fmt.Println(len(txpool.data))
			ctx <- txpool
			i++
		}

	}()

	for {
		select {
		case tx := <-ctx:
			fmt.Println(tx.data)
		}
	}

}

type TX struct {
	data []int
}

var memPool *TX
var once sync.Once

func GetMemPool() *TX {
	once.Do(func() {
		d := make([]int, 0)
		memPool = &TX{data: d}
	})
	return memPool
}

// LaunchServer ...
func LaunchServer() {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	router := mux.NewRouter()

	srv := &http.Server{
		Addr: "0.0.0.0:80",
		// Timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}

	router.HandleFunc("/account/{address}", handler.GetAccount).Methods("GET")
	router.HandleFunc("/block/{height}", handler.GetBlockByHeight).Methods("GET")
	router.HandleFunc("/tx", handler.SendTx).Methods("POST")

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()
	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}

package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/herdius/herdius-blockchain-api/config"
	"github.com/herdius/herdius-blockchain-api/handler"
	"github.com/herdius/herdius-blockchain-api/network"
)

func main() {
	log.Println("Opening API")
	LaunchServer()
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

	env := os.Getenv("ENV")
	configuration := config.GetConfiguration(env)
	builder := network.GetNetworkBuilder(env)
	net, err := builder.Build()
	if err != nil {
		log.Fatalf("Failed to build network:%v", err)
	}
	go net.Listen()
	defer net.Close()
	supervisorAdds := make([]string, 1)
	supervisorAdds = append(supervisorAdds, configuration.GetSupervisorAddress())
	handler.BootStrap(net, supervisorAdds)

	log.Println("supervsioraddress:", configuration.GetSupervisorAddress())
	if !net.ConnectionStateExists(configuration.GetSupervisorAddress()) {
		log.Println("No peers discovered in network")
	}

	router.HandleFunc("/account/{address}",
		func(w http.ResponseWriter, r *http.Request) {
			handler.GetAccount(w, r, *net, env)
		}).Methods("GET")

	router.HandleFunc("/block/{height}", handler.GetBlockByHeight).Methods("GET")
	router.HandleFunc("/tx", handler.SendTx).Methods("POST")
	router.HandleFunc("/tx/{id}", handler.GetTx).Methods("GET")

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

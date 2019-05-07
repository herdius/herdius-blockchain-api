package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/herdius/herdius-blockchain-api/config"
	"github.com/herdius/herdius-blockchain-api/handler"
	"github.com/herdius/herdius-blockchain-api/middleware"
	"github.com/herdius/herdius-blockchain-api/network"
	coreNet "github.com/herdius/herdius-core/p2p/network"
)

func main() {
	log.Println("Opening API")
	LaunchServer()
}

// LaunchServer ...
func LaunchServer() {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*3, "the duration for which the server gracefully waits for existing connections to finish - e.g. 15s or 1m")
	envFlag := flag.String("env", "dev", "environment to build network and run process for")
	flag.Parse()

	env := *envFlag
	builder := network.GetNetworkBuilder(env)
	net, err := builder.Build()
	if err != nil {
		log.Fatalf("Failed to build network:%v", err)
	}

	confg := config.GetConfiguration(env)
	supervisorAddr := confg.GetSupervisorAddress()

	go net.Listen()
	defer net.Close()
	supervisorAdds := make([]string, 0)
	supervisorAdds = append(supervisorAdds, supervisorAddr)

	connTest := new(middleware.Connected)
	router := *mux.NewRouter()
	addRoutes(net, env, &router)
	router.Use(connTest.IsConnected)

	srv := &http.Server{
		Addr:         "0.0.0.0:80",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      &router,
	}
	go func() {
		connPinging(net, supervisorAdds, connTest)
	}()
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()
	log.Println("Supervisor discovered at:", supervisorAddr)

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

func connPinging(net *coreNet.Network, supervisorAdds []string, connTest *middleware.Connected) {
	for {
		for _, supervisorAddr := range supervisorAdds {
			if !net.ConnectionStateExists(supervisorAddr) {
				net.Bootstrap(supervisorAdds...)
				log.Println("No peers discovered in network, retrying")
				*connTest = false
				continue
			}
			*connTest = true
			break
		}
		time.Sleep(time.Second * 3)
	}
}

func addRoutes(net *coreNet.Network, env string, router *mux.Router) {
	router.HandleFunc("/account/{address}",
		func(w http.ResponseWriter, r *http.Request) {
			handler.GetAccount(w, r, net, env)
		}).Methods("GET")
	router.HandleFunc("/block/{height}",
		func(w http.ResponseWriter, r *http.Request) {
			handler.GetBlockByHeight(w, r, net, env)
		}).Methods("GET")
	router.HandleFunc("/tx",
		func(w http.ResponseWriter, r *http.Request) {
			handler.PostTx(w, r, net, env)
		}).Methods("POST")
	router.HandleFunc("/tx/{id}",
		func(w http.ResponseWriter, r *http.Request) {
			handler.GetTx(w, r, net, env)
		}).Methods("GET")
	router.HandleFunc("/txs/{address}",
		func(w http.ResponseWriter, r *http.Request) {
			handler.GetTxsByAddress(w, r, net, env)
		}).Methods("GET")
	router.HandleFunc("/cancel/{address}",
		func(w http.ResponseWriter, r *http.Request) {
			handler.GetTxsByAddress(w, r, net, env)
		}).Methods("GET")
	router.HandleFunc("/cancel",
		func(w http.ResponseWriter, r *http.Request) {
			handler.PutCancelTxByTxID(w, r, net, env)
		}).Methods("PUT")
	router.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			json.NewEncoder(w).Encode("That path does not exist")
		})
}

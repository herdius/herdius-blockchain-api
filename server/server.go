package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/gorilla/mux"

	"github.com/herdius/herdius-blockchain-api/config"
	"github.com/herdius/herdius-blockchain-api/handler"
	"github.com/herdius/herdius-blockchain-api/middleware"
	"github.com/herdius/herdius-blockchain-api/network"
	"github.com/herdius/herdius-blockchain-api/store"
	"github.com/herdius/herdius-blockchain-api/store/postgres"
	coreNet "github.com/herdius/herdius-core/p2p/network"
	"github.com/herdius/herdius-core/p2p/network/discovery"
)

func main() {
	log.Println("Opening API")
	LaunchServer()
}

var (

	db *postgres.Store
	net *coreNet.Network
	err error
	connTest *middleware.Connected
)

// LaunchServer ...
func LaunchServer() {
	var (
		wait         time.Duration
		syncInterval time.Duration
	)
	flag.DurationVar(&wait, "graceful-timeout", time.Second*3, "the duration for which the server gracefully waits for existing connections to finish - e.g. 15s or 1m")
	flag.DurationVar(&syncInterval, "sync-interval", time.Second*5, "time to wait between db sync")
	envFlag := flag.String("env", "dev", "environment to build network and run process for")
	flag.Parse()

	env := *envFlag
	builder := network.GetNetworkBuilder(env)

	// Register peer discovery plugin.
	builder.AddPlugin(new(discovery.Plugin))

	net, err = builder.Build()
	if err != nil {
		log.Fatalf("Failed to build network:%v", err)
	}

	confg := config.GetConfiguration(env)
	supervisorAddr := confg.GetSupervisorAddress()

	go net.Listen()
	defer net.Close()
	supervisorAdds := make([]string, 0)
	supervisorAdds = append(supervisorAdds, supervisorAddr)

	connTest = new(middleware.Connected)
	router := mux.NewRouter()
	addRoutes(net, env, router)
	router.Use(connTest.IsConnected)

	srv := &http.Server{
		Addr:         "0.0.0.0:80",
		WriteTimeout: time.Second * 60,
		ReadTimeout:  time.Second * 60,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}
	go func() {
		connPinging(net, supervisorAdds, connTest)
	}()
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	var wg sync.WaitGroup
	stopDBSyncCh := make(chan struct{}, 1)

	wg.Add(1)
	go func() {
		defer wg.Done()

		configuration := config.GetConfiguration(env)
		db, err = postgres.NewStore(configuration.DBConnString())
		if err != nil {
			log.Printf("Failed to create store: %v", err)
			return
		}

		for {
			select {
			case <-stopDBSyncCh:
				return
			default:
				if *connTest {
					if err := store.SyncPendingTxs(db, net, env); err != nil {
						log.Printf("failed to sync pending tx: %v", err)
					}
				} else {
					log.Printf("not connected to core, skip sync pending txs")
				}
			}
			time.Sleep(syncInterval)
		}

	}()
	log.Println("Supervisor discovered at:", supervisorAddr)

	c := make(chan os.Signal, 2)

	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	srv.Shutdown(ctx)
	log.Println("shutting down")

	log.Println("Notify sync db goroutine")
	stopDBSyncCh <- struct{}{}
	wg.Wait()
	log.Println("sync db goroutine stopped, exiting...")
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
	router.HandleFunc("/tx/{id}",
		func(w http.ResponseWriter, r *http.Request) {
			handler.PutUpdateTxByTxID(w, r, net, env)
		}).Methods("PUT")
	router.HandleFunc("/tx/{id}",
		func(w http.ResponseWriter, r *http.Request) {
			handler.DeleteTx(w, r, net, env)
		}).Methods("DELETE")
	router.HandleFunc("/txs/{address}",
		func(w http.ResponseWriter, r *http.Request) {
			handler.GetTxsByAddress(w, r, net, env)
		}).Methods("GET")
	router.HandleFunc("/txs/{asset}/{address}",
		func(w http.ResponseWriter, r *http.Request) {
			handler.GetTxsByAssetAndAddress(w, r, net, env)
		}).Methods("GET")
	router.HandleFunc("/lock/txs/{block_number}",
		func(w http.ResponseWriter, r *http.Request) {
			handler.GetLockedTxsByBlockNumber(w, r, net, env)
		}).Methods("GET")
	router.HandleFunc("/redeem/txs/{block_number}",
		func(w http.ResponseWriter, r *http.Request) {
			handler.GetRedeemTxsByBlockNumber(w, r, net, env)
		}).Methods("GET")
	router.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			json.NewEncoder(w).Encode("That path does not exist")
		})
}

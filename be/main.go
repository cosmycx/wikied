package main

import (
	"context"
	"log"
	"net/http"
	"time"

	config "github.com/cosmycx/wikied/config"
	elastic "github.com/cosmycx/wikied/elastic"
	seed "github.com/cosmycx/wikied/seed"
	server "github.com/cosmycx/wikied/server"
)

// main
func main() {
	// wait for elastic docker connection
	time.Sleep(time.Second * 30)

	// configure log
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// get port from config
	port := ":" + config.AppPort

	// server connections
	router := http.NewServeMux()

	ctx := context.Background()

	elasticClient, err := elastic.NewClient(ctx, config.ElasticHost, false)
	if err != nil {
		// Can't connect to Elastic
		log.Printf("Can't connect to Elastic, err: %v", err)

	} else {
		// Connected to Elastic
		log.Printf("Elastic Client connected, client: %v\n", elasticClient)

		// server
		s := server.Server{
			Router:        router,
			ElasticClient: elasticClient,
		} // .server

		_ = s.ElasticClient.CreateIndexIfNotExists(ctx, elastic.IndexName)

		// ---------------------------
		// 			seed the DB
		// ---------------------------
		s.SeedInfoPost(seed.Seed1, seed.Seed1ID)
		s.SeedInfoPost(seed.Seed2, seed.Seed2ID)
		s.SeedInfoPost(seed.Seed3, seed.Seed3ID)


		// initiate routes
		s.RoutesInit()

		// starting server
		log.Printf("Starting server on port %s", port)
		log.Fatalln(http.ListenAndServe(port, s.Router))
	} // .else

} // .main

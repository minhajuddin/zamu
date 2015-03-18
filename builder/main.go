package main

import (
	"log"
	"net/http"

	cl "github.com/minhajuddin/config"
)

var (
	config struct {
		Host             string `yaml:"host"`
		Port             string `yaml:"port"`
		ENV              string `yaml:"env"`
		DB               string `yaml:"db"`
		MaxDBConnections int    `yaml:"max_db_connections"`
	}

	err error
)

func main() {
	//wireup routes
	wireupRoutes()
	//start server
	log.Println("starting server at http://localhost:3000/")
	log.Fatal(http.ListenAndServe(config.Port, nil))
}

func init() {
	//set log format
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)

	//load config
	cl.LoadFromFile("./config.yml", &config, log.Println)

	log.Printf("loaded config: %+v\n", config)

	//connectDB()
}

func wireupRoutes() {
}

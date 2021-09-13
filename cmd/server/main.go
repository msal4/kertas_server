package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Netflix/go-env"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/msal4/hassah_school_server/server"
	"gopkg.in/yaml.v2"
)

var debg *bool

func init() {
	debg = flag.Bool("debug", false, "Run server in debug mode")
	flag.Parse()
}

func loadConfig() server.Config {
	var cfgPath = os.Getenv("SCHOOL_CONFIG")
	if cfgPath == "" {
		cfgPath = "./config.yml"
	}

	f, err := os.Open(cfgPath)
	if err != nil {
		log.Fatal(err)
	}

	var cfg server.Config
	err = yaml.NewDecoder(f).Decode(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	if err := godotenv.Load(); err != nil {
		log.Println(err)
	}

	if _, err := env.UnmarshalFromEnviron(&cfg); err != nil {
		log.Fatal(err)
	}

	if cfg.Port == 0 {
		cfg.Port = 3000
	}

	if *debg {
		cfg.Debug = true
	}

	return cfg
}

func main() {
	cfg := loadConfig()
	srv, err := server.NewDefaultServer(cfg)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("listening on :%d", cfg.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), srv))
}

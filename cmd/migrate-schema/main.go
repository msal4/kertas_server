package main

import (
	"context"
	"flag"
	"log"
	"os"
	"strings"

	"github.com/Netflix/go-env"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/msal4/hassah_school_server/ent"
	"github.com/msal4/hassah_school_server/ent/migrate"
	"github.com/msal4/hassah_school_server/server"
	"gopkg.in/yaml.v2"
)

var debug *bool

func init() {
	debug = flag.Bool("debug", false, "Run migration in debug mode")
	flag.Parse()
}

func main() {
	f, err := os.Open("./config.yml")
	if err != nil {
		log.Fatal(err)
	}

	var cfg server.Config
	err = yaml.NewDecoder(f).Decode(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	godotenv.Load()

	env.UnmarshalFromEnviron(&cfg)

	if cfg.Port == 0 {
		cfg.Port = 3000
	}

	if *debug {
		cfg.Debug = true
	}

	if cfg.DatabaseDialect == "" {
		cfg.DatabaseDialect = cfg.DatabaseURL[:strings.Index(cfg.DatabaseURL, ":")]
	}

	client, err := ent.Open(cfg.DatabaseDialect, cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("failed establishing connection: %v", err)
	}
	defer client.Close()

	if *debug {
		err = client.Debug().Schema.Create(context.Background(), migrate.WithDropIndex(true), migrate.WithDropColumn(true))
	} else {
		err = client.Schema.Create(context.Background())
	}
	if err != nil {
		log.Fatalf("failed creating schema: %v", err)
	}
}

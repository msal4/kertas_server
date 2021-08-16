package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"entgo.io/ent/dialect"
	"github.com/Netflix/go-env"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/msal4/hassah_school_server/ent"
	"github.com/msal4/hassah_school_server/graph"
	"github.com/msal4/hassah_school_server/service"
	"gopkg.in/yaml.v2"
)

var debg *bool

func init() {
	debg = flag.Bool("debug", false, "Run server in debug mode")
	flag.Parse()
}

type Config struct {
	DatabaseURL string `yaml:"database_url" env:"DATABASE_URL"`
	Port        int    `yaml:"port" env:"PORT"`

	Minio struct {
		Endpoint  string `yaml:"endpoint" env:"MINIO_ENDPOINT"`
		AccessKey string `yaml:"access_key" env:"MINIO_ACCESS_KEY"`
		Token     string `yaml:"token" env:"MINIO_TOKEN"`
	} `yaml:"minio"`

	service.Config `yaml:"service"`
}

func main() {
	godotenv.Load()
	f, err := os.Open("./config.yml")
	if err != nil {
		log.Fatal(err)
	}

	var cfg Config
	err = yaml.NewDecoder(f).Decode(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", cfg)
	if cfg.Port == 0 {
		cfg.Port = 3000
	}
	env.UnmarshalFromEnviron(&cfg)

	ec, err := ent.Open(dialect.Postgres, cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("establishing db connection: %v", err)
	}
	mc, err := minio.New(cfg.Minio.Endpoint, &minio.Options{
		Creds: credentials.NewStaticV4(cfg.Minio.AccessKey, cfg.Minio.Token, ""),
	})
	if err != nil {
		log.Fatalf("instantiating minio client: %v", err)
	}

	s, err := service.New(ec, mc, nil)
	if err != nil {
		log.Fatalf("initializing service: %v", err)
	}

	srv := graph.NewServer(s, *debg)

	log.Printf("listening on :%d", cfg.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), srv))
}

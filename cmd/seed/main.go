package main

import (
	"context"
	"flag"
	"log"
	"os"
	"strings"

	"github.com/99designs/gqlgen/graphql"
	"github.com/Netflix/go-env"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/msal4/hassah_school_server/ent"
	"github.com/msal4/hassah_school_server/ent/user"
	"github.com/msal4/hassah_school_server/server"
	"github.com/msal4/hassah_school_server/server/model"
	"github.com/msal4/hassah_school_server/service"
	"gopkg.in/yaml.v2"
)

func init() {
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

	if cfg.DatabaseDialect == "" {
		cfg.DatabaseDialect = cfg.DatabaseURL[:strings.Index(cfg.DatabaseURL, ":")]
	}

	ec, err := ent.Open(cfg.DatabaseDialect, cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("failed establishing connection: %v", err)
	}
	defer ec.Close()

	mc, err := minio.New(cfg.Minio.Endpoint, &minio.Options{
		Creds: credentials.NewStaticV4(cfg.Minio.AccessKey, cfg.Minio.Token, ""),
	})
	if err != nil {
		log.Fatalf("instantiating minio client: %v", err)
	}

	s, err := service.New(ec, mc, &cfg.Config)
	if err != nil {
		log.Fatalf("creating service: %v", err)
	}

	if err = seed(context.Background(), s); err != nil {
		log.Fatalf("seeding: %v", err)
	}

	log.Println("Finished seeding 🎉.")
}

func seed(ctx context.Context, s *service.Service) error {
	f, err := os.Open("./testfiles/harvard.jpg")
	if err != nil {
		return err
	}

	stat, err := f.Stat()
	if err != nil {
		return err
	}

	sch, err := s.AddSchool(ctx, model.AddSchoolInput{
		Name:   "Palestine School",
		Active: true,
		Image: graphql.Upload{
			File:        f,
			Filename:    f.Name(),
			Size:        stat.Size(),
			ContentType: "image/jpg",
		},
	})
	if err != nil {
		return err
	}
	log.Printf("Created school: %v\n\n", sch)

	stg, err := s.AddStage(ctx, model.AddStageInput{
		Name:          "1st Elementry Stage",
		Active:        true,
		TuitionAmount: 15000000,
		SchoolID:      sch.ID,
	})
	log.Printf("Created stage: %v\n\n", stg)

	stdt, err := s.AddUser(ctx, model.AddUserInput{
		Name:     "John Doe",
		Username: "student01",
		Password: "student01pass",
		Phone:    "07712345678",
		Role:     user.RoleStudent,
		StageID:  &stg.ID,
		Active:   true,
	})
	if err != nil {
		return err
	}
	log.Printf("Created student: %v\n\n", stdt)

	return nil
}

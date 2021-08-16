package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Netflix/go-env"
	"github.com/msal4/hassah_school_server/service"
	"gopkg.in/yaml.v2"
)

func main() {
	f, err := os.Open("./config.yml")
	if err != nil {
		log.Fatal(err)
	}

	var cfg service.Config
	yaml.NewDecoder(f).Decode(&cfg)

	fmt.Printf("just yaml: %+v\n", cfg)

	env.UnmarshalFromEnviron(&cfg)

	fmt.Println()
	fmt.Printf("yaml+env: %+v\n", cfg)
}

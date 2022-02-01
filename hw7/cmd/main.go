package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"

	config "github.com/AbnormalReality/backend1_homeworks/hw7/configs"
)

const (
	usageConfig = "use this flag for set path to configuration file"
)

var (
	USECONFIG = "../configs/config.yml"
)

func main() {

	// Load config file
	cfg, err := config.InitConfig(&USECONFIG)
	if err != nil {
		log.Fatalf("Can't load configuration file: %s", err)
	}

	fmt.Printf("Host: %s:%s\n", cfg.App.ServerAddress, strconv.Itoa(cfg.App.ServerPort))
}

// init func, parse flags
func init() {
	flag.StringVar(&USECONFIG, "path", "../configs/config.yml", usageConfig)
	flag.Parse()
}
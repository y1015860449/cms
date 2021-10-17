package main

import (
	"cms/advert/initial"
	"flag"
	_ "go.uber.org/automaxprocs"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	defaultConfig := "./config.yaml"
	//dmEnv := os.Getenv("DM_MODE")
	//if dmEnv != "" {
	//	defaultConfig = fmt.Sprintf("./config-%s.yaml", dmEnv)
	//}
	confPath := flag.String("conf", defaultConfig, "config file path")
	flag.Parse()

	initial.Initial(*confPath)
}

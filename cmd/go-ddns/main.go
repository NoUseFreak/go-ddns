package main

import (
	"flag"
	"os"

	"github.com/NoUseFreak/go-ddns/internal/pkg/config"
	"github.com/NoUseFreak/go-ddns/internal/pkg/controller"
	log "github.com/sirupsen/logrus"
)

func init() {
	// log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		log.Error("Path argument required")
		return
	}

	handler := controller.HandlerController{}
	config := config.Config{}
	config.ParseFile(args[0])

	handler.Handle(config)
}

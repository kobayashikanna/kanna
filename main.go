package main

import (
	"flag"
	"fmt"
	"os"
	log "github.com/sirupsen/logrus"
	"strings"
)

type progArguments struct {
	getVersion bool
	logFormat string
	logLevel string
}

const VERSION = "0.0.0-alpha"

var arguments = progArguments{}

func main() {
	flag.BoolVar(&arguments.getVersion, "version", false, "Print kanna version.")
	flag.StringVar(&arguments.logLevel, "level", "debug","Log level")
	flag.Parse()

	if arguments.getVersion {
		fmt.Printf("kanna, a tunnel for lolicon. ver: %s\n", VERSION)
		os.Exit(0)
	}

	// set log level for logrus
	switch (strings.ToLower(arguments.logLevel)) {
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "panic":
		log.SetLevel(log.PanicLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "fatal":
		log.SetLevel(log.FatalLevel)
	default:
		fmt.Println("log level only can be set as debug/error/info/panic/warn/fatal")
		os.Exit(1)
	}

}

package main

import (
	"flag"
	"os"
	"time"

	"github.com/robertlestak/preflight-netpath/pkg/preflightnetpath"
	log "github.com/sirupsen/logrus"
)

func init() {
	ll, err := log.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		ll = log.InfoLevel
	}
	log.SetLevel(ll)
}

func main() {
	l := log.WithFields(log.Fields{
		"app": "preflight-netpath",
	})
	l.Debug("starting preflight-netpath")
	preflightFlags := flag.NewFlagSet("preflight-netpath", flag.ExitOnError)
	logLevel := preflightFlags.String("log-level", log.GetLevel().String(), "log level")
	endpoint := preflightFlags.String("endpoint", "", "endpoint to test in the form of <host>:<port>")
	timeout := preflightFlags.Duration("timeout", time.Second*5, "timeout in seconds")
	configFile := preflightFlags.String("config", "", "path to config file")
	preflightFlags.Parse(os.Args[1:])
	ll, err := log.ParseLevel(*logLevel)
	if err != nil {
		ll = log.InfoLevel
	}
	log.SetLevel(ll)
	preflightnetpath.Logger = l.Logger
	pf := &preflightnetpath.PreflightNetPath{
		Endpoint: *endpoint,
		Timeout:  *timeout,
	}
	if *configFile != "" {
		if pf, err = preflightnetpath.LoadConfig(*configFile); err != nil {
			l.WithError(err).Error("error loading config")
			os.Exit(1)
		}
	}
	if err := pf.Run(); err != nil {
		l.WithError(err).Error("error running preflight-netpath")
		os.Exit(1)
	}
}

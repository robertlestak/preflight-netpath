package preflightnetpath

import (
	"net"
	"time"

	log "github.com/sirupsen/logrus"
)

type PreflightNetPath struct {
	// Endpoint to test in the form of <host>:<port>
	Endpoint string        `json:"endpoint,omitempty" yaml:"endpoint,omitempty"`
	Timeout  time.Duration `json:"timeout,omitempty" yaml:"timeout,omitempty"`
}

func (pf *PreflightNetPath) Init() error {
	l := log.WithFields(log.Fields{
		"fn": "Init",
	})
	l.Debug("starting preflight-netpath")
	if pf.Timeout == 0 {
		pf.Timeout = 5 * time.Second
	}
	return nil
}

func (pf *PreflightNetPath) Run() error {
	l := log.WithFields(log.Fields{
		"fn": "Run",
	})
	l.Debug("starting preflight-netpath")
	if pf.Endpoint == "" {
		l.Error("endpoint is required")
		return nil
	}
	// create a tcp connection to the endpoint
	// if successful, return nil
	// if unsuccessful, return error
	c, err := net.DialTimeout("tcp", pf.Endpoint, pf.Timeout)
	if err != nil {
		l.WithError(err).Error("error dialing endpoint")
		return err
	}
	defer c.Close()
	l.Debug("successfully dialed endpoint")
	l.Info("preflight successful")
	return nil
}

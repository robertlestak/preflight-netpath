package preflightnetpath

import (
	"encoding/json"
	"net"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

type PreflightNetPath struct {
	// Endpoint to test in the form of <host>:<port>
	Endpoint string        `json:"endpoint" yaml:"endpoint"`
	Timeout  time.Duration `json:"timeout" yaml:"timeout"`
}

func LoadConfig(filepath string) (*PreflightNetPath, error) {
	l := log.WithFields(log.Fields{
		"fn": "LoadConfig",
	})
	l.Debug("loading config")
	var err error
	pf := &PreflightNetPath{}
	bd, err := os.ReadFile(filepath)
	if err != nil {
		l.WithError(err).Error("error reading file")
		return pf, err
	}
	if err := yaml.Unmarshal(bd, pf); err != nil {
		// try with json
		if err := json.Unmarshal(bd, pf); err != nil {
			l.WithError(err).Error("error unmarshalling config")
			return pf, err
		}
	}
	return pf, err
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

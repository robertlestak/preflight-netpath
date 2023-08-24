package preflightnetpath

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

var (
	Logger *log.Logger
)

func init() {
	if Logger == nil {
		Logger = log.New()
		Logger.SetOutput(os.Stdout)
		Logger.SetLevel(log.InfoLevel)
	}
}

type PreflightNetPath struct {
	// Endpoint to test in the form of <host>:<port>
	Endpoint string        `json:"endpoint" yaml:"endpoint"`
	Timeout  time.Duration `json:"timeout" yaml:"timeout"`
}

func LoadConfig(filepath string) (*PreflightNetPath, error) {
	l := Logger.WithFields(log.Fields{
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
	l := Logger.WithFields(log.Fields{
		"fn": "Init",
	})
	l.Debug("starting preflight-netpath")
	if pf.Timeout == 0 {
		pf.Timeout = 5 * time.Second
	}
	return nil
}

func (pf *PreflightNetPath) Equivalent() {
	l := Logger
	l.Debug("printing equivalent command")
	timeoutSeconds := int(pf.Timeout.Seconds())
	cmd := fmt.Sprintf(`HOST="$(echo %s | cut -d: -f1)" && PORT="$(echo %s | cut -d: -f2)" && nc -z -w%d $HOST $PORT`, pf.Endpoint, pf.Endpoint, timeoutSeconds)
	cmd = fmt.Sprintf(`sh -c '%s'`, cmd)
	l.Infof("equivalent command: %s", cmd)
}

func (pf *PreflightNetPath) Run() error {
	l := Logger.WithFields(log.Fields{
		"preflight": "netpath",
	})
	l.Debug("starting preflight-netpath")
	if pf.Endpoint == "" {
		l.Error("endpoint is required")
		return nil
	}
	if err := pf.Init(); err != nil {
		l.WithError(err).Error("error initializing preflight-netpath")
		return err
	}
	pf.Equivalent()
	// create a tcp connection to the endpoint
	// if successful, return nil
	// if unsuccessful, return error
	c, err := net.DialTimeout("tcp", pf.Endpoint, pf.Timeout)
	if err != nil {
		l.WithError(err).Error("failed")
		return err
	}
	defer c.Close()
	l.Debug("successfully dialed endpoint")
	l.Info("passed")
	return nil
}

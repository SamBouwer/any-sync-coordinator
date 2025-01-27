package config

import (
	"github.com/anyproto/any-sync-coordinator/cafeapi"
	"github.com/anyproto/any-sync-coordinator/db"
	"github.com/anyproto/any-sync-coordinator/filelimit"
	"github.com/anyproto/any-sync-coordinator/spacestatus"
	commonaccount "github.com/anyproto/any-sync/accountservice"
	"github.com/anyproto/any-sync/app"
	"github.com/anyproto/any-sync/metric"
	"github.com/anyproto/any-sync/net/rpc"
	"github.com/anyproto/any-sync/net/transport/yamux"
	"github.com/anyproto/any-sync/nodeconf"
	"gopkg.in/yaml.v3"
	"os"
)

const CName = "config"

func NewFromFile(path string) (c *Config, err error) {
	c = &Config{}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	if err = yaml.Unmarshal(data, c); err != nil {
		return nil, err
	}
	return
}

type Config struct {
	Account                  commonaccount.Config   `yaml:"account"`
	Drpc                     rpc.Config             `yaml:"drpc"`
	Metric                   metric.Config          `yaml:"metric"`
	Network                  nodeconf.Configuration `yaml:"network"`
	NetworkStorePath         string                 `yaml:"networkStorePath"`
	NetworkUpdateIntervalSec int                    `yaml:"networkUpdateIntervalSec"`
	Mongo                    db.Mongo               `yaml:"mongo"`
	SpaceStatus              spacestatus.Config     `yaml:"spaceStatus"`
	Yamux                    yamux.Config           `yaml:"yamux"`
	CafeApi                  cafeapi.Config         `yaml:"cafeApi"`
	FileLimit                filelimit.Config       `yaml:"fileLimit"`
}

func (c *Config) Init(a *app.App) (err error) {
	return
}

func (c Config) Name() (name string) {
	return CName
}

func (c Config) GetAccount() commonaccount.Config {
	return c.Account
}

func (c Config) GetDrpc() rpc.Config {
	return c.Drpc
}

func (c Config) GetMetric() metric.Config {
	return c.Metric
}

func (c Config) GetNodeConf() nodeconf.Configuration {
	return c.Network
}

func (c Config) GetMongo() db.Mongo {
	return c.Mongo
}

func (c Config) GetSpaceStatus() spacestatus.Config {
	return c.SpaceStatus
}

func (c Config) GetNodeConfStorePath() string {
	return c.NetworkStorePath
}

func (c Config) GetNodeConfUpdateInterval() int {
	return c.NetworkUpdateIntervalSec
}

func (c Config) GetYamux() yamux.Config {
	return c.Yamux
}

func (c Config) GetCafeApi() cafeapi.Config {
	return c.CafeApi
}

func (c Config) GetFileLimit() filelimit.Config {
	return c.FileLimit
}

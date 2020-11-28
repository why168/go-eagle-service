package conf

import (
	"encoding/json"
	"io/ioutil"

	"git.inke.cn/BackendPlatform/golang/logging"
)

type Config struct {
	Server struct {
		ServiceName int    `json:"service_name"`
		Port        string `json:"port"`
	} `json:"server"`

	Redis struct {
		Host     string `json:"host"`
		Password string `json:"password"`
		DbNum    int    `json:"db_num"`
	} `json:"redis"`

	Database struct {
		Schema   string `json:"schema"`
		Host     string `json:"host"`
		User     string `json:"user"`
		Password string `json:"password"`
		Database string `json:"database"`
	} `json:"database"`
}

func Init() (*Config, error) {
	cfg := &Config{}
	err := LoadConf("config2.toml", cfg)
	return cfg, err
}

func LoadConf(confFile string, v interface{}) error {
	cbytes, err := ioutil.ReadFile(confFile)
	if err != nil {
		logging.Errorf("match|MatchDispatcher|LoadConf|read conf file error. file:%s, err:%+v", confFile, err)
		return err
	}

	if err := json.Unmarshal(cbytes, v); err != nil {
		logging.Errorf("match|MatchDispatcher|LoadConf|json Unmarshal conf error. confFile:%s, cbytes:%s, err:%+v",
			confFile, cbytes, err)
		return err
	}

	return nil
}

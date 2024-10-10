package config

import (
	"fmt"
	"github.com/pkg/errors"
	"sync"
)

var (
	once   sync.Once
	config *Config
	err    error
)

func Load() (*Config, error) {
	once.Do(func() {
		config, err = loadConfig()
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to load config")
	}
	return config, err
}

func loadConfig() (*Config, error) {

	conf := &Config{}
	// loads config from file
	k, _, err := LoadConfigUsingKoanf("static")
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	err = k.UnmarshalWithConf("", conf, DefaultUnmarshallingConfig(conf))
	if err != nil {
		return nil, fmt.Errorf("failed to refresh config: %w", err)
	}
	return conf, nil
}

type Config struct {
	Flags     *Flags
	MapOfFile map[string]*FileStruct
	Test1     *Test1
}

type Flags struct {
	EnablePreCallLocationCheck bool `dynamic:"true"`
}

type FileStruct struct {
	Key1 string
}

type Test1 struct {
	Test2 *Test2
}

type Test2 struct {
	Test3 *Test3
}

type Test3 struct {
	MapOfFile map[string]*FileStruct
}

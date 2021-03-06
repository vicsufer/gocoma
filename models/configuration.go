package models

import (
	"github.com/spf13/viper"
	"sort"
)

var Conf *Configuration

type Configuration struct {
	CurrentEnv string `mapstructure: "currentenv"`
	Environments map[string] Environment `mapstructure:"environments"`
}

func (conf *Configuration) GetEnvsNames() []string {
	envNames := make([]string, 0, len(conf.Environments))
	for name := range conf.Environments {
		envNames = append(envNames, name )
	}
	sort.Strings(envNames)
	return envNames
}

func (conf *Configuration) GetEnv(e string) Environment {
	return conf.Environments[e]
}

func (conf *Configuration) SetCurrentEnv(e string) bool {
	if env, ok := conf.Environments[e]; ok{
		conf.CurrentEnv = e
		env.Activate()
		return save("currentenv", conf.CurrentEnv)
	}
	return false
}

func save (key string, value interface{}) bool {
	viper.Set(key, value)
	if err := viper.WriteConfig(); err != nil {
		return false
	}
	return true
}

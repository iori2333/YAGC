package config

import (
	"gopkg.in/yaml.v2"
	"os"
	"path"
)

const Version = "0.0.1"

type UserConfig struct {
	Name  string `yaml:"name"`
	EMail string `yaml:"e_mail"`
}

type AppConfig struct {
	User UserConfig `yaml:"user,omitempty"`
}

var conf = &AppConfig{}

func init() {
	_ = conf.Load()
}

func App() *AppConfig {
	return conf
}

func (config *AppConfig) Load() error {
	configDir, _ := os.UserConfigDir()
	workDir, _ := os.Getwd()

	var err error
	var content []byte

	if content, err = os.ReadFile(path.Join(configDir, ".yagcconfig")); err == nil {
		err = yaml.Unmarshal(content, config)
	}

	if content, err = os.ReadFile(path.Join(workDir, ".yagcconfig")); err == nil {
		err = yaml.Unmarshal(content, config)
	}

	return err
}

func (config *AppConfig) Save(global bool) error {
	var savePath string
	if global {
		savePath, _ = os.UserConfigDir()
	} else {
		savePath, _ = os.Getwd()
	}

	var err error
	var content []byte
	if content, err = yaml.Marshal(config); err == nil {
		err = os.WriteFile(path.Join(savePath, ".yagcconfig"), content, os.ModeType)
	}

	return err
}

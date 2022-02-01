package config

import (
	"errors"
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"path"
	"sync"
	"yagc/util"
)

type RepoCore struct {
	Bare bool `yaml:"bare,omitempty"`
}

type RepoRemote struct {
	Name string `yaml:"name"`
	URL  string `yaml:"url,omitempty"`
}

type RepoBranch struct {
	Name   string `yaml:"name"`
	Remote string `yaml:"remote,omitempty"`
	Merge  string `yaml:"merge,omitempty"`
}

type RepoConfig struct {
	Core     RepoCore     `yaml:"core,omitempty"`
	Remotes  []RepoRemote `yaml:"remotes,omitempty"`
	Branches []RepoBranch `yaml:"branches,omitempty"`
}

var repoConf = &RepoConfig{}

var repo sync.Once

func Repo() (*RepoConfig, bool) {
	_, ok := util.GetRepoRoot()
	if !ok {
		return nil, false
	}

	repo.Do(func() {
		repoConf = &RepoConfig{}

		if err := repoConf.Load(); err != nil {
			log.Fatalf("%s", err.Error())
		}
	})
	return repoConf, true
}

func (config *RepoConfig) Load() error {
	root, ok := util.GetRepoRoot()
	if !ok {
		return errors.New("repo root not found")
	}
	configFile := path.Join(root, ".yagc", "config")

	if content, err := os.ReadFile(configFile); err != nil {
		return err
	} else if err := yaml.Unmarshal(content, config); err != nil {
		return err
	} else {
		return nil
	}
}

func (config *RepoConfig) Save() error {
	root, ok := util.GetRepoRoot()
	if !ok {
		return errors.New("repo root not found")
	}
	configFile := path.Join(root, ".yagc", "config")

	if content, err := yaml.Marshal(config); err != nil {
		return err
	} else if err := os.WriteFile(configFile, content, os.ModeType); err != nil {
		return err
	} else {
		return nil
	}
}

// This file is a part of git.thekondor.net/zvuchno.git (mirror: github.com/thekondor/zvuchno)

package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type notificationConfigSection struct {
	Timeout uint32
}

type appearanceConfigSection struct {
	Width  byte                `yaml:"width,omitempty"`
	Format formatConfigSection `yaml:"format,omitempty"`
	Text   textConfigSection   `yaml:"text,omitempty"`
}

type formatConfigSection struct {
	Full string `yaml:"full,omitempty"`
	Bar  string `yaml:"bar,omitempty"`
}

type textConfigSection struct {
	Title    string `yaml:"title"`
	OnMute   string `yaml:"on_mute"`
	OnUnmute string `yaml:"on_unmute"`
}

type Config struct {
	Notification notificationConfigSection `yaml:"notification,omitempty"`
	Appearance   appearanceConfigSection   `yaml:"appearance,omitempty"`
}

func NewConfig() *Config {
	configPath := locateConfigPath()
	log.Printf(`Config = %s`, configPath)

	self, err := newConfig(configPath)
	if nil != err {
		log.Printf("W: Failed to load config: %s, default values to be used", err)
	}

	return self
}

func locateConfigPath() string {
	configPath := "${HOME}/.zvuchno.yml"
	if _, ok := os.LookupEnv("XDG_CONFIG_HOME"); ok {
		configPath = "${XDG_CONFIG_HOME}/zvuchno.yml"
	}

	return os.ExpandEnv(configPath)
}

func newConfig(path string) (*Config, error) {
	self := &Config{
		Notification: notificationConfigSection{
			Timeout: 1000,
		},
		Appearance: appearanceConfigSection{
			Width: 20,
			Format: formatConfigSection{
				Full: "{{ .Percent }}% {{ .Bar }}",
				Bar:  "[=> ]",
			},
			Text: textConfigSection{
				Title:    "Volume",
				OnMute:   "🔇 muted",
				OnUnmute: "🔈 unmuted",
			},
		},
	}

	file, err := ioutil.ReadFile(path)
	if nil != err {
		return self, err
	}

	err = yaml.Unmarshal(file, self)
	if nil != err {
		return self, err
	}

	return self, nil
}

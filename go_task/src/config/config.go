package config

import (
	"log"

	"gopkg.in/go-ini/ini.v1"
)

type ConfigList struct {
	MaxSell      int
	MaxBuy       int
	ApiKey       string
	ApiSecret    string
	LineSecret   string
	LineToken    string
	SlackToken   string
	SlackChannel string
}

var BaseURL string
var TestBaseURL string
var Config ConfigList

func NewConfig() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Println("iniFileLoadError")
	}
	Config = ConfigList{
		ApiKey:       cfg.Section("bybit").Key("api_key").String(),
		ApiSecret:    cfg.Section("bybit").Key("api_secret").String(),
		LineSecret:   cfg.Section("line").Key("secret").String(),
		LineToken:    cfg.Section("line").Key("token").String(),
		SlackToken:   cfg.Section("slack").Key("token").String(),
		SlackChannel: cfg.Section("slack").Key("channel").String(),
	}
	BaseURL = cfg.Section("bybit").Key("base_url").String()
	TestBaseURL = cfg.Section("bybit").Key("test_base_url").String()
}

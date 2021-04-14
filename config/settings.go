package config

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

type server struct {
	RunMode string
	HttpPort string
	ReadTimeout time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &server{}

type database struct {
	Type string
	User string
	Password string
	Host string
	Port string
	Dbname string
}

var DatabaseSetting = &database{}

var cfg *ini.File

func init() {
	var err error
	cfg, err = ini.Load("config/config.ini")
	if err != nil {
		log.Fatalln("fail to load settings")
	}
	mapTo("server", ServerSetting)
	mapTo("database", DatabaseSetting)
}

func mapTo(s string, i interface{}){
	err := cfg.Section(s).MapTo(i)
	if err != nil {
		log.Fatalln("Cfg.MapTo", s, "err", err)
	}
}


package server

import (
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/BurntSushi/toml"
	log "github.com/Sirupsen/logrus"
	"github.com/julienschmidt/httprouter"
)

const (
	CONFIG_FILE = "config.toml"
)

type Config struct {
	dev      bool
	static   string
	js       string
	style    string
	template string
	address  string
	api      string
}

type Server struct {
	router *httprouter.Router
	config Config
}

func (s *Server) Bootstrap() {
	log.Infoln("bootstrapping server")
	router := httprouter.New()
	s.router = router

	log.Infoln("loading routes")
	s.mapRoutes()
	log.Infoln("loading endpoints")
	s.DefineEndpoints()

	log.Infoln("serving")
	http.ListenAndServe(s.config.address, s.router)
}

func (c *Server) ParseConfig() {
	log.Infoln("parsing config")
	config := struct {
		Common struct {
			Env    string
			Js     string
			Style  string
			Api    string
			Static string
		}
		Server struct {
			Template string
			Address  string
		}
	}{}
	f, err := os.Open(CONFIG_FILE)
	if err != nil {
		log.Errorln("Config open", err)
		os.Exit(1)
	}

	data, err := ioutil.ReadAll(f)
	if err != nil {
		log.Errorln("Config read", err)
		os.Exit(1)
	}

	_, err = toml.Decode(string(data), &config)
	if err != nil {
		log.Errorln("Config parse", err)
		os.Exit(1)
	}

	var dev bool
	if config.Common.Env == "dev" {
		dev = true
	} else if config.Common.Env == "prod" {
		dev = false
	} else {
		log.Errorln("Invalid env", config.Common.Env)
		os.Exit(1)
	}

	c.config = Config{
		dev:      dev,
		static:   config.Common.Static,
		js:       config.Common.Js,
		style:    config.Common.Style,
		api:      config.Common.Api,
		template: config.Server.Template,
		address:  config.Server.Address,
	}
}

func init() {
	log.SetFormatter(&log.TextFormatter{
		TimestampFormat: time.RFC822,
		FullTimestamp:   true,
	})
}

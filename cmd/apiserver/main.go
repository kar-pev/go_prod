package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/iekar-pov/go_prod/internal/app/apiserver"
)

var (
	configPath         string
	passwordConfigPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config")
	flag.StringVar(&passwordConfigPath, "password-config-path",
		"configs/apiserver_passwords.toml", "path to password config")
}

func main() {
	flag.Parse()
	config := apiserver.NewConfig()
	{
		//adding config file
		_, err := toml.DecodeFile(configPath, config)
		if err != nil {
			log.Println(err)
		}
	}

	{
		//adding password config file
		_, err := toml.DecodeFile(passwordConfigPath, config)
		if err != nil {
			log.Println(err)
		}
	}

	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}

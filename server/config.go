package server

import (
	"encoding/json"
	"log"
	"os"
)

type AppConfig struct {
	Db     DbConfig
	Smtp   SmtpConfig
	Rabbit RabbitConfig
}

type DbConfig struct {
	Host   string
	Port   string
	DbName string
	User   string
	Pass   string
}

type RabbitConfig struct {
	Host string
	Port string
}

type SmtpConfig struct {
	Host string
	Port string
}

func LoadConfigFromFile(configAbsPath string) (c AppConfig) {
	log.Printf("Using config file: %v\n", configAbsPath)
	file, err := os.Open(configAbsPath)

	if err != nil {
		log.Fatalln("Error openning config file: ", err)
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&c)
	if err != nil {
		log.Fatalln("Error when decoding config file: ", err)
	}

	return c
}

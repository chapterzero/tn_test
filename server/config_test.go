package server

import (
	"path/filepath"
	"testing"
)

var configAbsPath string = ""

func init() {
	absPath, err := filepath.Abs("../app.config.json")
	if err != nil {
		panic("Could not find app.conf.json")
	}
	configAbsPath = absPath
}

func TestLoadConfigFromFile(t *testing.T) {
	config := LoadConfigFromFile(configAbsPath)
	if config.Db.Host == "" || config.Db.Port == "" || config.Db.DbName == "" || config.Db.User == "" {
		t.Error("Database host, port, dbname and user config must not empty")
	}
}

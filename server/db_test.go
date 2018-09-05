package server

import (
	"testing"
)

func TestDbGetterSetter(t *testing.T) {
	dbConfig := DbConfig{
		Host:   "127.0.0.1",
		Port:   "3306",
		DbName: "test",
		User:   "root",
		Pass:   "mypass",
	}

	SetDbConfig(dbConfig)
	actualConfig := GetDbConfig()

	if actualConfig.Host != "127.0.0.1" {
		t.Errorf("Wrong host value, expected %v, got %v", "127.0.0.1", actualConfig.Host)
	}

	if actualConfig.Port != "3306" {
		t.Errorf("Wrong port value, expected %v, got %v", "3306", actualConfig.Port)
	}

	if actualConfig.DbName != "test" {
		t.Errorf("Wrong Db name value, expected %v, got %v", "test", actualConfig.DbName)
	}

	if actualConfig.User != "root" {
		t.Errorf("Wrong user value, expected %v, got %v", "root", actualConfig.User)
	}

	if actualConfig.Pass != "mypass" {
		t.Errorf("Wrong pass value, expected %v, got %v", "mypass", actualConfig.Pass)
	}
}

func TestGetDsnString(t *testing.T) {
	dbConfig := DbConfig{
		Host:   "127.0.0.1",
		Port:   "3306",
		DbName: "test",
		User:   "root",
		Pass:   "mypass",
	}

	actual := GetDsnString(dbConfig)
	expectedDsnString := "root:mypass@tcp(127.0.0.1:3306)/test"

	if actual != expectedDsnString {
		t.Errorf("Expected %v, got %v", expectedDsnString, actual)
	}
}

package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
)

// DBConfig interface
type DBConfig interface {
	FormatDSN() string
}

type Config struct {
	Databases struct {
		NbaApp DB
	}
}

// DB config
type DB struct {
	Name string
	Host string
	Port string
	USER string
	Pass string
}

var Conf Config

// Load Config
func Load(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	if err := toml.Unmarshal(b, &Conf); err != nil {
		return err
	}

	if err := f.Close(); err != nil {
		return err
	}

	return nil
}

func (d DB) FormatDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Asia%%2FTokyo")
}

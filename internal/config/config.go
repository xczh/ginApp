package config

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/pelletier/go-toml"
)

type config struct {
	Address   string `validate:"required,tcp_addr"`
	URLPrefix string `toml:"url_prefix" validate:"required,startswith=/,endswith=/"`
	MySQLDSN  string `toml:"mysql_dsn" validate:"omitempty"`
	RedisDSN  string `toml:"redis_dsn" validate:"omitempty"`
}

func newConfig() *config {
	return &config{
		Address:   "0.0.0.0:3000",
		URLPrefix: "/api/v1/",
	}
}

func (c *config) String() string {
	return fmt.Sprintf("%#v", c)
}

func (c *config) applyEnvVar(envPrefix string) {
	refValC := reflect.ValueOf(c).Elem()

	for i := 0; i < refValC.NumField(); i++ {
		sf := refValC.Type().Field(i)
		envName := strings.ToUpper(envPrefix + sf.Name)
		envVar, ok := os.LookupEnv(envName)
		if !ok {
			continue
		}
		_ = fillVar(refValC.Field(i), envVar)
	}
}

func (c *config) applyConfigFile(cfgFilePath string) error {
	ext := filepath.Ext(cfgFilePath)
	switch ext {
	case "":
		return errors.New("Config file path must end with an extension, eg. '.toml'")
	case ".toml":
		cfgBytes, err := ioutil.ReadFile(cfgFilePath)
		if err != nil {
			return err
		}
		return toml.Unmarshal(cfgBytes, c)
	default:
		return errors.New("Unsupported config file format: " + ext)
	}
}

func (c *config) validate() error {
	return validator.New().Struct(c)
}

package config

import (
	"io/ioutil"

	"github.com/BurntSushi/toml"
)

type serverConfig struct {
	Name        string              `toml:"name"`
	Description string              `toml:"description"`
	Domain      string              `toml:"domain"`
	Host        string              `toml:"host"`
	TLS         serverTLSConfig     `toml:"tls"`
	App         serverAppConfig     `toml:"app"`
	API         serverAPIConfig     `toml:"api"`
	Logging     serverLoggingConfig `toml:"logging"`
}

type serverAppConfig struct {
	Port int `toml:"port"`
}

type serverAPIConfig struct {
	Port int `toml:"port"`
}

type serverLoggingConfig struct {
	Level    string `toml:"level"`
	Location string `toml:"location"`
}

type serverTLSConfig struct {
	Enabled        bool   `toml:"enabled"`
	TruststorePath string `toml:"truststore"`
	KeystorePath   string `toml:"keystore"`
	Truststore     []byte
	Keystore       []byte
}

// Config describes the otaks configuration
type Config struct {
	Server serverConfig `toml:"server"`
}

// NewConfig parses arguments to create a Config
func NewConfig(configPath string) (*Config, error) {
	return LoadConfigFile(configPath)
}

// LoadConfigFile ...
func LoadConfigFile(path string) (*Config, error) {
	var c Config
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	_, err = toml.Decode(string(data), &c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

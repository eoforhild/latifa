package config

import (
	"io/ioutil"
	"sync"

	"gopkg.in/yaml.v2"
)

const DefaultLocation = "/etc/latifa/config.yml"

var (
	mu      sync.RWMutex
	_config *Configuration
)

type Configuration struct {
	path string
	//Key string `yaml:"key"`
	LogDirectory  string                   `default:"/var/log/latifa" yaml:"log_directory"`
	RootDirectory string                   `default:"/var/lib/latfia" yaml:"root_directory"`
	MongoDBUri    string                   `yaml:"mongodb_uri"`
	Debug         bool                     `default:"true" yaml:"debug"`
	Api           ApiConfiguration         `yaml:"api"`
	GoogleCloud   GoogleCloudConfiguration `yaml:"google_cloud"`
}

type GoogleCloudConfiguration struct {
	Storage struct {
		CredentialsFile string `default:"" yaml:"credentials_file"`
	} `yaml:"storage"`
}

type ApiConfiguration struct {
	Host string `default:"0.0.0.0" yaml:"host"`
	Port int    `default:"8080" yaml:"port"`

	Ssl struct {
		Enabled         bool   `default:"true" yaml:"enabled"`
		CertificateFile string `default:"" yaml:"certifcate_file"`
		KeyFile         string `default:"" yaml:"key_file"`
	}
}

func FromFile(path string) error {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	c := new(Configuration)
	c.path = path

	if err := yaml.Unmarshal(f, c); err != nil {
		return err
	}

	Set(c)
	return nil
}

func Set(c *Configuration) {
	mu.Lock()
	_config = c
	mu.Unlock()
}

func Get() *Configuration {
	mu.RLock()
	c := *_config
	mu.RUnlock()
	return &c
}

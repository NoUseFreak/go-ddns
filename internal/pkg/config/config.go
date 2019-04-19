package config

import (
	"io/ioutil"

	log "github.com/sirupsen/logrus"
	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	sets []ConfigSet
}

func (c *Config) GetSets() []ConfigSet {
	return c.sets
}

type ConfigFile struct {
	Timeout int `yaml:"timeout"`
	Spec    []struct {
		Type    string            `yaml:"type"`
		Domain  string            `yaml:"domain"`
		Options map[string]string `yaml:"options"`
	} `yaml:"spec"`
}

func (c *Config) ParseFile(file string) {
	yamlFile, err := ioutil.ReadFile(file)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	var sets ConfigFile
	err = yaml.Unmarshal(yamlFile, &sets)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
		return
	}
	for _, spec := range sets.Spec {
		cs := ConfigSet{
			record:  spec.Domain,
			setType: spec.Type,
			options: spec.Options,
			timeout: sets.Timeout,
		}
		c.sets = append(c.sets, cs)
		log.Infof("Found config for %s", spec.Domain)
	}
}

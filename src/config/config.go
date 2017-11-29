package config

import (
	"github.com/go-chinese-site/cfg"
)

var YamlConfig *cfg.YamlConfig

func YmlFileRead(configflie string) {
	var err error

	YamlConfig, err = cfg.ParseYaml(configflie)

	if err != nil {
		panic(err)
	}
}

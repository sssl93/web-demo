package main

import (
	"flag"
	"fmt"
	"github.com/ghodss/yaml"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
)

const VERSION = "v1.0.0"

type Config struct {
	Path string      `yaml:"path"`
	Data interface{} `yaml:"data"`
}

func getConfigs(configPath string) map[string]interface{} {
	var sourceConfigs []Config
	yamlFile, err := ioutil.ReadFile(configPath)

	if err != nil {
		log.Fatalf("yamlFile.Get err #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &sourceConfigs)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	configs := make(map[string]interface{})
	for _, config := range sourceConfigs {
		configs[config.Path] = config.Data
	}
	return configs
}

func registerURL(configPath string, r *gin.Engine) {
	configs := getConfigs(configPath)
	for path := range configs {
		r.GET(path, func(c *gin.Context) {
			currentConfigs := getConfigs(configPath)
			c.PureJSON(200, currentConfigs[path])
		})
	}
}

func main() {
	var (
		listenAddress = flag.String("web.listen-address", ":8089", "Address to listen on for web interface and telemetry.")
		configPath    = flag.String("config", "config.yaml", "the path of config yml file.")
		showVersion   = flag.Bool("version", false, "Show the version.")
	)

	flag.Parse()
	if *showVersion {
		fmt.Printf("generic-exporter %s", VERSION)
		return
	}

	r := gin.Default()
	registerURL(*configPath, r)

	// listen and serve on 0.0.0.0:8089
	if err := r.Run(*listenAddress); err != nil {
		fmt.Println(err)
	}
}

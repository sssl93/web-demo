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

var configs []Config
var pathMap = make(map[string]interface{})

type Config struct {
	Path string      `yaml:"path"`
	Data interface{} `yaml:"data"`
}

func reloadConfigs(configPath string) {
	yamlFile, err := ioutil.ReadFile(configPath)

	log.Println("yamlFile:", yamlFile)
	if err != nil {
		log.Printf("yamlFile.Get err #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &configs)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	for _, config := range configs {
		pathMap[config.Path] = config.Data
	}
}

func registerURL(configPath string, r *gin.Engine) {
	reloadConfigs(configPath)
	for path := range pathMap {
		r.GET(path, func(c *gin.Context) {
			reloadConfigs(configPath)
			c.PureJSON(200, pathMap[path])
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

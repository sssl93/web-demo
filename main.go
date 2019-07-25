package main

import (
	"fmt"
	"github.com/ghodss/yaml"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
)

type Config struct {
	Path string      `yaml:"path"`
	Data interface{} `yaml:"data"`
}

func main() {
	var configs []Config
	yamlFile, err := ioutil.ReadFile("config.yaml")

	log.Println("yamlFile:", yamlFile)
	if err != nil {
		log.Printf("yamlFile.Get err #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &configs)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	r := gin.Default()

	for _, config := range configs {
		// Serves unicode entities
		r.GET(config.Path, func(c *gin.Context) {
			c.PureJSON(200, config.Data)
		})
	}

	// listen and serve on 0.0.0.0:8089
	if err := r.Run(":8089"); err != nil {
		fmt.Println(err)
	}
}

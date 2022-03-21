package main

import (
	"fmt"
	yamlConfig "github.com/rowdyroad/go-yaml-config"
	"github.com/rowdyroad/grpc-hw/internal/client"
	"os"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error:",err)
			os.Exit(-1)
		}
	}()
	var config client.Config
	yamlConfig.LoadConfig(&config, "configs/client.yaml", nil)
	app, err := client.NewClient(config)
	if err != nil {
		panic(err)
	}
	defer app.Close()
	if err = app.Run(); err != nil {
		panic(err)
	}
}




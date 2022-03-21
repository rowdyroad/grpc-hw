package main

import (
	"fmt"
	"github.com/rowdyroad/grpc-hw/internal/server"
	yamlConfig "github.com/rowdyroad/go-yaml-config"
	"os"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error:",err)
			os.Exit(-1)
		}
	}()
	var config server.Config
	yamlConfig.LoadConfig(&config, "configs/server.yaml", nil)
	app, err := server.NewServer(config)
	if err != nil {
		panic(err)
	}
	defer app.Close()
	if err = app.Run(); err != nil {
		panic(err)
	}
}




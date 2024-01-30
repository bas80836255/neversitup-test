package main

import (
	"fmt"
	"my-project/app"
	"my-project/config"
	"os"
)

func init() {
	err := config.LoadStandardVersion()
	if err != nil {
		println(err.Error())
		os.Exit(-1)
	}

	err = config.LoadEnv()
	if err != nil {
		println(err.Error())
		os.Exit(-1)
	}

	err = config.LoadFile()
	if err != nil {
		println(err.Error())
		os.Exit(-1)
	}

	err = config.LoadOverrideENV()
	if err != nil {
		println(err.Error())
		os.Exit(-1)
	}
}

func main() {

	fmt.Printf("starting application %s, version: %s, env: %s\n", config.App.Application.Name, config.Env.AppVersion, config.Env.AppEnv)
	app.Init()
	app.Start()
}

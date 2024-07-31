package main

import "github.com/waldrey/go-expert/apis/apis/configs"

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBHost)

}

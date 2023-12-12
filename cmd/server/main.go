package main

import "github.com/kleytonsolinho/pos-go-expert/08_API/configs"

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBDriver)
}

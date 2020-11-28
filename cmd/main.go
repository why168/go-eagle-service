package main

import "go-eagle-service/conf"

func init() {

}

func main() {
	_, _ = conf.Init()
	//toml.DecodeReader(bytes.NewBuffer(tm.ch.Data), v)
}

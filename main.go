package main

import Server "github.com/vrianta/Server"

func main() {
	Server.New("", "1080", Routes, &ServerConfig).Start()
}

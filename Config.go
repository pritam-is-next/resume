package main

import Config "github.com/vrianta/Server/Config"

var ServerConfig = Config.New(false, "Views", []string{
	"Static",
}, []string{
	"Css",
}, []string{
	"Js",
})

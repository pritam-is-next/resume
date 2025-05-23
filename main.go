package main

import (
	server "github.com/vrianta/Server"
)

func main() {
	server.New("", "1080", Routes, &server.Config{
		JS_Folders: []string{
			"Js",
		},
		CSS_Folders: []string{
			"Css",
		},
		Static_folders: []string{
			"Static",
		},
		Views_folder: "Views",
	}).Start()
}

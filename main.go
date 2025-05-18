package main

import (
	Controllers "github.com/pritam-is-next/resume/Controllers"

	server "github.com/pritam-is-next/resume/Server"
)

func main() {
	server.New("localHost", "1080", server.Routes{
		"/": &Controllers.Home{},
	}, &server.Config{
		JS_Folders: []string{
			"Js",
		},
		CSS_Folders: []string{
			"Css",
		},
		Static_folders: []string{
			"Static",
		},
		Views_folder: "./Views",
	}).Start()
}

package main

import "{{ .ModulePath }}/app"

func main() {
	app.New().Run()
}

package server

import "fmt"

func buildLoaderSource(importPath string) string {
	return fmt.Sprintf(`
package main

import (
	"encoding/json"
	"os"

	"%s"
)

func main() {
	var p Routes
	routes := p.Routes()
	_ = json.NewEncoder(os.Stdout).Encode(routes)
}
`, importPath)
}

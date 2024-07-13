package main

import "github.com/dencoseca/toolkit"

func main() {
	var tools toolkit.Tools

	tools.CreateDirIfNotExist("./test-dir")
}

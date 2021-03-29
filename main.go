// main serves one purpose: to initialize Cobra. find true src in /cmd
package main

import (
	"duck/cmd"
)

func main() {
	cmd.Execute()
}

/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"mittere/cmd"
	"mittere/requester"
)

func main() {
	cmd.Execute()

	requester.Request()
}

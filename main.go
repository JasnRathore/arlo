package main

import (
	"fmt"
	cmd "github.com/JasnRathore/arlo/commands"
	"os"
)

const version = "v0.0.4"

func main() {
	size := len(os.Args)
	if size < 2 {
		fmt.Println("No Commands Passed")
		fmt.Println("use \narlo help or arlo -h\nTo get a list of commands")
		return
	}
	args := os.Args[1:]
	switch args[0] {
	case "init", "-i":
		cmd.InitProject()
	case "dev", "-d":
		cmd.RunDevBuild()
	case "build", "-b":
		cmd.RunProdBuild()
	case "help", "-h":
		cmd.PrintHelp()
	case "version", "-v":
		fmt.Println(version)
	case "upgrade", "-u":
		cmd.UpgradeArlo()
	default:
		fmt.Println("Invalid Commands")
		fmt.Println("use \narlo help or arlo -h\nTo get a list of commands")
	}
}

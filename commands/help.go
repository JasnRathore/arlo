package commands

import "fmt"

func PrintHelp() {
	fmt.Println("Usage:")
	fmt.Println()
	fmt.Println("    arlo <command> [shorthand]")
	fmt.Println()
	fmt.Println("The commands are:")
	fmt.Println()
	fmt.Println("    init     (-i)    initialize a new arlo project")
	fmt.Println("    dev      (-d)    starts your development environment")
	fmt.Println("    build    (-b)    builds the final binary for distribution")
	fmt.Println("    upgrade  (-u)    upgrades arlo to the latest version")
	fmt.Println("    version  (-v)    prints app version")
	fmt.Println("    help     (-h)    prints all the available commands")
	fmt.Println()
}

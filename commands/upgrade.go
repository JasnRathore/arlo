package commands

import (
	"fmt"
	utils "github.com/JasnRathore/arlo/utils"
)

func UpgradeArlo() {
	fmt.Println("Running Upgrade")
	utils.RunCommand("go", "install", "github.com/JasnRathore/arlo@latest")
	fmt.Println("Arlo Upgraded")
}

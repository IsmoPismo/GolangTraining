package main

import (
	"fmt"

	"github.com/IsmoPismo/GolangTraining/05_Scope/01_Package/02_Visibility/vis"
)

func main() {
	fmt.Printf(vis.MyName)
	vis.PrintVar()
}

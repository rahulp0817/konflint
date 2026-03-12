package main

import (
	"fmt"

	"github.com/rahulp0817/konflint/cmd"
)

func printBanner() {
	banner := `
  _  __            __ _         _  
 | |/ /___  _ __  / _| (_)_ __ | |_ 
 | ' // _ \| '_ \| |_| | | '_ \| __|
 | . \ (_) | | | |  _| | | | | | |_ 
 |_|\_\___/|_| |_|_| |_|_|_| |_|\__|

Konflint - Deployment Intelligence CLI Tool
`
	fmt.Println(banner)
}

func main() {
	printBanner()
	cmd.Execute()
}

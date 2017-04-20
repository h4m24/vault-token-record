/*
This will generate orphan tokens according to provided policy, store it in vault and spit out the token itself

first read json config file
  read vault address
  read token store path
  read script token

CLI flags

create an orphan token

*/
package main

import (
	"fmt"
	"os"
)

// const EnvVaultAddress = "VAULT_ADDR"
var policyName string

func main() {
	app := cliconfig()
	if len(os.Args) > 1 {
		app.Run(os.Args)
	} else {
		fmt.Println("enter  --help to show available options")
		os.Exit(2)
	}
	// fmt.Println(policyName)
	// config := getconfigs()
	//  how to get 1 config
	// fmt.Println(config.TokensPath)
	// fmt.Println(toJSON(config))
}

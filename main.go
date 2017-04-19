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

import "os"

func main() {
	app := cliconfig()
	app.Run(os.Args)

	// config := getconfigs()
	//  how to get 1 config
	// fmt.Println(config.TokensPath)
	// fmt.Println(toJSON(config))
}

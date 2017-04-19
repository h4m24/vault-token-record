/*
This will generate orphan tokens according to provided policy, store it in vault and spit out the token itself

first read json config file
  read vault address
  read token store path
  read script token
*/
package main

import "fmt"

func main() {

	pages := getPages()

	fmt.Println(toJSON(pages))
}

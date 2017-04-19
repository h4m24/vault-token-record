/*
This will generate orphan tokens according to provided policy, store it in vault and spit out the token itself

first read json config file
  read vault address
  read token store path
  read script token
*/
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type config struct {
	VaultAddress string `json:"vaultaddress"`
	TokensPath   string `json:"tokenspath"`
	VaultToken   string `json:"vaulttoken"`
}

func toJson(p interface{}) string {
	bytes, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return string(bytes)
}
func (p config) toString() string {
	return toJson(p)
}

func getPages() config {
	raw, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var c config
	json.Unmarshal(raw, &c)
	return c
}

func main() {

	pages := getPages()

	fmt.Println(toJson(pages))
}

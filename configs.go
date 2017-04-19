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

func toJSON(p interface{}) string {
	bytes, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return string(bytes)
}
func (p config) toString() string {
	return toJSON(p)
}

func getconfigs() config {
	raw, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var c config
	json.Unmarshal(raw, &c)
	return c
}

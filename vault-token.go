package main

import (
	"fmt"

	"github.com/hashicorp/vault/api"
)

func genToken() {
	// vclient := api.NewClient(c * Config)
	// vclient.SetToken(config.VaultToken)
	// vclient := api.Client(addr == config.VaultAddress, token == config.VaultToken)
	config := getconfigs()
	clientConfig := &api.Config{
		Address: config.VaultAddress,
	}

	client, err := api.NewClient(clientConfig)
	if err != nil {
		panic(err)
	}

	secrets, err := client.Logical().Read("tokens")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v+", secrets.Data)
}

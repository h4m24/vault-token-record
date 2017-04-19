package main

import (
	"fmt"
	"strings"

	"github.com/hashicorp/vault/api"
)

func genToken() {
	// get configs from json
	config := getconfigs()

	//  pass config from json to vault client struct
	clientConfig := &api.Config{
		Address: config.VaultAddress,
	}

	// pass policy name from CLI to vault client struct
	orphanRequest := &api.TokenCreateRequest{
		Policies: strings.Fields(policyName),
	}

	// create vault client
	client, err := api.NewClient(clientConfig)
	if err != nil {
		panic(err)
	}
	// setting vault token
	client.SetToken(config.VaultToken)

	vaultResponse, err := client.Auth().Token().CreateOrphan(orphanRequest)

	if err != nil {
		println("ERROR Could not create Orphan token")
		panic(err)
	}

	// printing property of the token returned by vault
	fmt.Println(vaultResponse.Renewable)

	// fmt.Printf("%v", vaultResponse)

	// secrets, err := client.Logical().Read("tokens")
	// fmt.Printf("%v+", secrets.Data)

	// generate the token with properties above
	// vaultAnswer := client.Auth().Token().CreateOrphan(orphanRequest)
	// client.Token()

	// fmt.Printf("%v+", vaultAnswer.LeaseID)

	//
	// secrets, err := client.Logical().Read("tokens")
	// if err != nil {
	// 	panic(err)
	// }
	//
	// fmt.Printf("%v+", secrets.Data)
}

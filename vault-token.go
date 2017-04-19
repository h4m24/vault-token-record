package main

import (
	"fmt"
	"strings"
	"time"

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
	type vaultRepliedToken struct {
		Token         string
		Accessor      string
		LeaseDuration string
		Policies      string
		Renewable     string
		Metadata      string
	}

	tokenExpireTime := time.Now().Local().Add(time.Second * time.Duration(vaultResponse.Auth.LeaseDuration)).Format("2006-01-02")

	println("accessor:  " + vaultResponse.Auth.Accessor)
	println("token:  " + vaultResponse.Auth.ClientToken)
	println("expires: " + tokenExpireTime)
	println("policies:  ")
	fmt.Printf("%v", vaultResponse.Auth.Policies)
	println("\nis renewable:")
	fmt.Println(vaultResponse.Auth.Renewable)

	// var f interface{}
	// err := json.Unmarshal(vaultResponse.Data, &f)
	// fmt.Println(vaultResponse.Data)

	// tokenValues = vaultResponse.Data
	// fmt.Println(tokenValues)
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

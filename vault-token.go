package main

import (
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

	tokenExpireTime := time.Now().Local().Add(time.Second * time.Duration(vaultResponse.Auth.LeaseDuration)).Format("2006-01-02")
	println("writing to " + config.VaultAddress)

	client.Logical().Write(config.TokensPath, map[string]interface{}{
		"accessor":       vaultResponse.Auth.Accessor,
		"token":          vaultResponse.Auth.ClientToken,
		"lease duration": vaultResponse.Auth.LeaseDuration,
		"expires":        tokenExpireTime,
		"policies":       vaultResponse.Auth.Policies,
		"renewable":      vaultResponse.Auth.Renewable,
	})

}

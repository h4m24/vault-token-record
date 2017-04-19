package main

import (
	"github.com/urfave/cli"
)

func cliconfig() *cli.App {

	app := cli.NewApp()
	app.Name = "token-log"
	//  fix usage later
	// app.Usage = "--policy-name default "
	app.Description = "this program will generate a token according to a provided policy\n   calculate the expiration time for that token, and store all that info\n   in vault under a provided path."
	app.Version = "0.1"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "policy-name, p",
			Usage:       "vault policy name",
			Destination: &policyName,
		},
	}

	app.Action = func(c *cli.Context) {

		genToken()

	}

	// app.Commands = []cli.Command{
	//   {
	//     Name:    "request",
	//     Usage:   "request certificate",
	//     Action:  func(c *cli.Context) error {
	//       url, token := getURLToken()
	//       vaults := askConsul(url)
	//       activeVault := getActVault(vaults)
	//       getCerts(dir, cname, ttl, format, activeVault, path, token)
	//       return nil
	//     },
	//   },
	// }
	return app

}

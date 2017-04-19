package main

import "github.com/urfave/cli"

//
// import (
// 	"github.com/urfave/cli"
// )
//
func cliconfig() *cli.App {

	app := cli.NewApp()
	app.Name = "token-log"
	app.Usage = "this program will generate a token according to a provided policy\n   calculate the expiration time for that token, and store all that info\n   in vault under a provided path."

	// app.Flags = []cli.Flag{
	// 	cli.StringFlag{
	// 		Name:        "consuladdr, a",
	// 		Usage:       "consul address string",
	// 		Destination: &addr,
	// 		EnvVar:      "PKICONSULADDR",
	// 	}
	// }
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

package main

import (
"os"
"github.com/urfave/cli"
rancher "github.com/rancher/go-rancher/client"
"errors"
"fmt"
)

func main() {
	app := cli.NewApp()
	app.Flags = []cli.Flag {
		cli.StringFlag{
			Name: "api-url",
			Usage: "URL to your rancher API",
			EnvVar: "RANCHER_API_URL",
		},
		cli.StringFlag{
			Name: "access-key",
			Usage: "The access key for your API",
			EnvVar: "RANCHER_ACCESS_KEY",
		},
		cli.StringFlag{
			Name: "secret-key",
			Usage: "The secret key for your API",
			EnvVar: "RANCHER_SECRET_KEY",
		},
		cli.StringFlag{
			Name: "environment-id",
			Usage: "The Environment Id to get the token for",
			EnvVar: "RANCHER_ENVIRONMENT_ID",
		},
	}
	app.Name = "ret - Rancher Environment Token"
	app.Usage = "Receive the current Registration Token for a given Rancher Environment"
	app.Action = func(c *cli.Context) error {

		var rancherApiUrl 		= c.String("api-url")
		var rancherAccessKey		= c.String("access-key")
		var rancherSecretKey		= c.String("secret-key")
		var rancherEnvironmentId	= c.String("environment-id")

		var registrationToken 		= ""

		if rancherApiUrl == "" {
			return errors.New("You have to provide an url by using --api-url")
		}

		if rancherAccessKey == "" {
			return errors.New("You have to provide an access key by using --access-key")
		}

		if rancherSecretKey == "" {
			return errors.New("You have to provide an url by using --secret-key")
		}

		if rancherEnvironmentId == "" {
			return errors.New("You have to provide an url by using --environment-id")
		}

		client, err := rancher.NewRancherClient(&rancher.ClientOpts{
			Url:       rancherApiUrl,
			AccessKey: rancherAccessKey,
			SecretKey: rancherSecretKey,
		})

		if err != nil {
			fmt.Printf("Error creating Rancher Client %s", err)
		}

		registrationTokens, err := client.RegistrationToken.List(nil)

		if err != nil {
			fmt.Printf("Error fetching Tokens")
		}

		for _, token := range registrationTokens.Data {
			if token.AccountId == rancherEnvironmentId {
				registrationToken = token.Token
			}
		}

		if registrationToken != "" {
			fmt.Printf(registrationToken)
		} else {
			return errors.New("No token found for the given Envirionment " + rancherEnvironmentId)
		}

		return nil
	}

	app.Run(os.Args)
}

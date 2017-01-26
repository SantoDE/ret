package main

import (
"os"

"github.com/urfave/cli"
"errors"
)

func main() {
	app := cli.NewApp()
	app.Flags = []cli.Flag {
		cli.StringFlag{
			Name: "api-url",
			Usage: "URL to your rancher API",
		},
		cli.StringFlag{
			Name: "access-key",
			Usage: "The access key for your API",
		},
		cli.StringFlag{
			Name: "secret-key",
			Usage: "The secret key for your API",
		},
		cli.StringFlag{
			Name: "environment-id",
			Usage: "The Environment Id to get the token for",
		},
	}
	app.Name = "ret - Rancher Environment Token"
	app.Usage = "Receive the current Registration Token for a given Rancher Environment"
	app.Action = func(c *cli.Context) error {

		var rancherApiUrl 		= c.String("api-url")
		var rancherAccessKey		= c.String("access-key")
		var rancherSecretKey		= c.String("secret-key")
		var rancherEnvironmentId	= c.String("environment-id")


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

		return nil
	}

	app.Run(os.Args)
}

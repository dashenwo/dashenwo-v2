package agent

import "github.com/urfave/cli/v2"

func Command() *cli.Command {
	command := &cli.Command{
		Name:  "agent",
		Usage: "Run the api gateway agent",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "address",
				Usage:   "Set the gateway agent address :10001",
				EnvVars: []string{"OKGATEWAY_ADDRESS"},
			},
		},
		Action: func(ctx *cli.Context) error {
			return run(ctx)
		},
	}
	return command
}

func run(ctx *cli.Context) error {

	return nil
}
package command

import (
	cli "github.com/urfave/cli/v2"

	"github.com/nyanpassu/systemd-oci/common"
)

// Ps .
var Ps = cli.Command{
	Name:      "ps",
	Usage:     "ps displays the processes running inside a container",
	ArgsUsage: `<container-id> [ps options]`,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "format, f",
			Value: "table",
			Usage: `select one of: ` + formatOptions,
		},
	},
	Action: func(context *cli.Context) error {
		return common.ErrNotImplemented
	},
	// SkipArgReorder: true,
}

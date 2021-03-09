package command

import (
	cli "github.com/urfave/cli/v2"

	"github.com/nyanpassu/systemd-oci/container"
)

// Delete .
var Delete = cli.Command{
	Name:  "delete",
	Usage: "delete any resources held by the container often used with detached container",
	ArgsUsage: `<container-id>

Where "<container-id>" is the name for the instance of the container.

EXAMPLE:
For example, if the container id is "ubuntu01" and runc list currently shows the
status of "ubuntu01" as "stopped" the following will delete resources held for
"ubuntu01" removing "ubuntu01" from the runc list of containers:

       # runc delete ubuntu01`,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "force, f",
			Usage: "Forcibly deletes the container if it is still running (uses SIGKILL)",
		},
	},
	Action: func(context *cli.Context) error {
		var (
			err error
			c   container.Container
		)
		if c, err = getContainer(context); err != nil {
			return err
		}
		if err = c.Delete(); err != nil {
			return err
		}
		return nil
	},
}

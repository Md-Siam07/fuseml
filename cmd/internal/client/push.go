package client

import (
	"os"

	"github.com/fuseml/fuseml/cli/paas"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var ()

// CmdPush implements the fuseml orgs command
var CmdPush = &cobra.Command{
	Use:   "push NAME [PATH_TO_APPLICATION_SOURCES]",
	Short: "Push an application from the specified directory, or the current working directory",
	Args:  cobra.RangeArgs(1, 2),
	RunE: func(cmd *cobra.Command, args []string) error {
		client, cleanup, err := paas.NewFusemlClient(cmd.Flags(), nil)
		defer func() {
			if cleanup != nil {
				cleanup()
			}
		}()

		if err != nil {
			return errors.Wrap(err, "error initializing cli")
		}

		var path string
		if len(args) == 1 {
			path, err = os.Getwd()
			if err != nil {
				return errors.Wrap(err, "error pushing app")
			}
		} else {
			path = args[1]
		}

		err = client.Push(args[0], path)
		if err != nil {
			return errors.Wrap(err, "error pushing app")
		}

		return nil
	},
	SilenceErrors: true,
	SilenceUsage:  true,
}

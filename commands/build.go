package commands

import (
	"github.com/avitacco/jig/internal/build"
	"github.com/spf13/cobra"
)

func (a *App) buildCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "build",
		Short: "Build a Puppet module package",
		RunE: func(cmd *cobra.Command, args []string) error {
			return build.DoBuild()
		},
	}
	return cmd
}

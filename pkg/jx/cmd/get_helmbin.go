package cmd

import (
	"github.com/jenkins-x/jx/pkg/jx/cmd/templates"
	"github.com/jenkins-x/jx/pkg/log"
	"github.com/jenkins-x/jx/pkg/util"
	"github.com/spf13/cobra"
)

// GetHelmBinOptions containers the CLI options
type GetHelmBinOptions struct {
	GetOptions
}

var (
	helmBinsAliases = []string{
		"branch pattern",
	}

	getHelmBinLong = templates.LongDesc(`
		Display the Helm binary name used in pipelines.

		This setting lets you switch from the stable release to early access releases (e.g. from Helm 2 <-> 3)
`)

	getHelmBinExample = templates.Examples(`
		# List the git branch patterns for the current team
		jx get helmbin
	`)
)

// NewCmdGetHelmBin creates the new command for: jx get env
func NewCmdGetHelmBin(commonOpts *CommonOptions) *cobra.Command {
	options := &GetHelmBinOptions{
		GetOptions: GetOptions{
			CommonOptions: commonOpts,
		},
	}
	cmd := &cobra.Command{
		Use:     "helmbin",
		Short:   "Display the Helm binary name used in the pipelines",
		Aliases: []string{"helm"},
		Long:    getHelmBinLong,
		Example: getHelmBinExample,
		Run: func(cmd *cobra.Command, args []string) {
			options.Cmd = cmd
			options.Args = args
			err := options.Run()
			CheckErr(err)
		},
	}

	options.addGetFlags(cmd)
	return cmd
}

// Run implements this command
func (o *GetHelmBinOptions) Run() error {
	helm, _, _, err := o.TeamHelmBin()
	if err != nil {
		return err
	}
	log.Infof("Your team uses the helm binary: %s\n", util.ColorInfo(helm))
	log.Infof("To change this value use: %s\n", util.ColorInfo("jx edit helmbin helm3"))
	return nil
}

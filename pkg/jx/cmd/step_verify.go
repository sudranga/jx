package cmd

import (
	"github.com/spf13/cobra"
)

// StepVerifyOptions contains the command line flags
type StepVerifyOptions struct {
	StepOptions
}

func NewCmdStepVerify(commonOpts *CommonOptions) *cobra.Command {
	options := &StepVerifyOptions{
		StepOptions: StepOptions{
			CommonOptions: commonOpts,
		},
	}

	cmd := &cobra.Command{
		Use:   "verify",
		Short: "verify [command]",
		Run: func(cmd *cobra.Command, args []string) {
			options.Cmd = cmd
			options.Args = args
			err := options.Run()
			CheckErr(err)
		},
	}
	cmd.AddCommand(NewCmdStepVerifyPod(commonOpts))
	return cmd
}

// Run implements this command
func (o *StepVerifyOptions) Run() error {
	return o.Cmd.Help()
}

package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wasuken/wcgorep"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "wcgorep",
	Short: "wcgorep is grep with web contents as the target implemented by go.",
	Long: `wcgorep is grep with web contents as the target implemented by go.
now, only supported by text contents`,
	Run: func(cmd *cobra.Command, args []string) {
		wcgorep.Gorep(args)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

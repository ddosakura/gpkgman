package core

import (
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install the packages",
	Long:  `Install the packages.`,

	Run: executeMod(Install),
}

package core

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	// ErrModNotFound Mod_Not_Found
	ErrModNotFound = errors.New("Mod Not Found")
)

var modeListCmd = &cobra.Command{
	Use:   "modelist",
	Short: "Print the list of mod which can be used",
	Long:  `Print the list of mod which can be used.`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Available Mods:")
		for _, m := range modCenter {
			a, b, c := m.Version()
			fmt.Printf("  %-16s %-9s %s\n",
				m.Name(),
				fmt.Sprintf("v%d.%d.%d", a, b, c),
				m.Description(),
			)
		}
	},
}

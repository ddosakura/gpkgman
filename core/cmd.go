package core // import "github.com/ddosakura/gpkgman/core"

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/ddosakura/gklang"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

var (
	// root is the root of the command
	root = &cobra.Command{
		Use:   "gpkgman",
		Short: "Generic Package Manager",
		Long:  `The generic package manage tools.`,
	}

	dev = false
	mod = "generic"
)

func init() {
	cobra.OnInitialize(initCLI)

	root.PersistentFlags().StringVarP(&mod, "mod", "m", "generic", "Choose the manager.")

	root.AddCommand(modeListCmd)
	root.AddCommand(installCmd)
}

// Execute is the entry of the CLI
func Execute() {
	l := log.New(os.Stdout, "[GPKG]: ", log.LstdFlags)
	if os.Getenv("GPKG_DEV") == "true" {
		dev = true
	}
	if dev {
		gklang.LoadLogger(l, gklang.LvDebug)
	} else {
		gklang.LoadLogger(l, gklang.LvInfo)
	}
	fmt.Println(logo)

	root.Version = Version
	root.Execute()
}

func initCLI() {
	home, err := homedir.Dir()
	if err != nil {
		gklang.Er(err)
	}
	HomeDir = path.Join(home, ".ddosakura", ".gpkgman")
	CacheDir = path.Join(HomeDir, "cache")
	err = os.MkdirAll(CacheDir, 0755)
	if err != nil {
		gklang.Er(err)
	}
}

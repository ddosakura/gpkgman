package core

import (
	"fmt"
	"os/exec"

	"github.com/ddosakura/gklang"
	"github.com/spf13/cobra"
)

type executable func(Action, *cobra.Command, []string)

// Mod is the interface for mods
type Mod interface {
	Name() string
	Version() (int, int, int)
	Description() string
	Dependences() [][]string
	// TODO: SubDependences() map[string][][]string
	Execute(Action, *cobra.Command, []string)
}

var (
	// modCenter is the place for mod to register
	modCenter = make(map[string]Mod)
)

// Register is the func for mod to register itself
func Register(mod Mod) {
	modCenter[mod.Name()] = mod
}

// Action is used to choose action
type Action int

const (
	// Actions is the sub command which mod can use
	Actions = iota
	// Install Action
	Install
)

func executeMod(a Action) func(*cobra.Command, []string) {
	return func(cmd *cobra.Command, args []string) {
		m := modCenter[mod]
		if m == nil {
			gklang.Er(ErrModNotFound)
		}

		for _, clis := range m.Dependences() {
			var err error
			for _, cli := range clis {
				_, err = exec.LookPath(cli)
				if err == nil {
					break
				}
			}
			if err != nil {
				gklang.Er(fmt.Errorf("%s %v %s", "one of", clis, "must be installed"))
			}
		}

		m.Execute(a, cmd, args)
	}
}

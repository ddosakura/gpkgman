package generic

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"
	"strings"

	"../../core"
	"github.com/ddosakura/gklang"
	"github.com/spf13/cobra"
)

func init() {
	core.Register(&Mod{})
}

// Mod is
type Mod struct {
}

// Execute is the entry of the mod
func (*Mod) Execute(a core.Action, cmd *cobra.Command, args []string) {

	getDepDir()
	switch a {
	case core.Install:
		for _, pkg := range args {
			install(pkg)
		}
	}
}

func install(pkg string) {
	var (
		part   = strings.Split(pkg, "/")
		remote = "github.com"
		user   = "ddosakura"
		repo   = ""
	)
	switch len(part) {
	case 1:
		repo = part[0]
	case 2:
		user = part[0]
		repo = part[1]
	default:
		remote = part[0]
		user = part[1]
		repo = part[2]
	}
	dir := path.Join(core.CacheDir, remote, user)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		gklang.Er(err)
	}
	cmd := exec.Command("git", "clone", fmt.Sprintf("https://%s/%s/%s.git", remote, user, repo))
	cmd.Dir = dir
	cmd.Env = os.Environ()
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	//cmd.Stdout = os.Stdout
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		gklang.Er(err)
	}
	err = cmd.Start()
	if err != nil {
		gklang.Er(err)
	}
	reader := bufio.NewReader(stdout)
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		fmt.Print(line)
	}
	err = cmd.Wait()
	if err != nil {
		gklang.Er(err)
	}
	if cmd.Process != nil {
		if err = cmd.Process.Kill(); err != nil && err.Error() != "os: process already finished" {
			gklang.Er(err)
		}
	}

	fmt.Println(path.Join(dir, repo))
	err = os.Symlink(path.Join(dir, repo), path.Join(getDepDir(), repo))
	if err != nil {
		gklang.Er(err)
	}
}

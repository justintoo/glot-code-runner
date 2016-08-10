package nim

import (
	"github.com/prasmussen/glot-code-runner/cmd"
	"path/filepath"
)

func Run(files []string, stdin string) (string, string, error) {
	workDir := filepath.Dir(files[0])
	return cmd.RunStdin(workDir, stdin, "nim", "--verbosity:0", "compile", "--run", files[0])
}

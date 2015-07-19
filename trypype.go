package main

import (
	"fmt"
	"os/exec"

	"labix.org/v2/pipe"
)

func main() {
	cmd := exec.Command("find", "/usr/bin", "-maxdepth", "2", "-iname", "*go*")
	line := pipe.Line(
		pipe.Exec(cmd.Path, cmd.Args...),
		pipe.Exec("head", "-10"),
	)
	out, _ := pipe.Output(line)
	fmt.Println(string(out))
}

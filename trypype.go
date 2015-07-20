package main

import (
	"fmt"
	"strings"

	"labix.org/v2/pipe"
)

func findCommandBinaries(loc, value string) []string {
	return []string{loc, "-maxdepth", "2", "-iname", "*" + value + "*"}
}

func findCommandOutput(cmd []string, c chan []string) {
	results := make([]string, 0)
	line := pipe.Line(
		pipe.Exec("find", cmd...),
		pipe.Exec("head", "-10"),
	)
	output, err := pipe.Output(line)
	if err != nil {
		fmt.Println(string(output))
		panic(err)
	}
	split := strings.Split(string(output), "\n")
	for _, i := range split {
		results = append(results, i)
	}
	c <- results
}

func main() {
	results := make([]string, 0)
	c := make(chan []string)
	cmd := findCommandBinaries("/usr/bin", "go")
	go findCommandOutput(cmd, c)
	bin := <-c
	results = append(results, bin...)
	fmt.Println(results)
}

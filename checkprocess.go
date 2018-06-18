package main

import (
	"flag"
	"fmt"
	"os/exec"
	"regexp"
	"time"
)

func errChk(e error) {
	if e != nil {
		panic(e)
	}
}

func processChk(program string) bool {
	re := regexp.MustCompile("(?m)^.*" + regexp.QuoteMeta(program) + "$")
	out, err := exec.Command("/usr/bin/ps", "-A").Output()
	errChk(err)
	if re.MatchString(string(out)) {
		return true
	}
	return false
}

func main() {
	var program string
	flag.StringVar(&program, "program", "kdeinit5", "which program to watch")
	flag.Parse()

	ticker := time.NewTicker(time.Second * 1)

	for range ticker.C {
		if processChk(program) {
			fmt.Println("program is running. quit")
			ticker.Stop()
			break
		}
		fmt.Println("program is not running. wait for another second.")
	}
}

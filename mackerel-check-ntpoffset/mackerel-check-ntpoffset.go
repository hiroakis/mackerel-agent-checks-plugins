package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

const (
	OK       = 0
	WARNING  = 1
	CRITICAL = 2
	UNKNOWN  = 3
)

func main() {
	offsetCrit := flag.Float64("crit", 100.0, "critical if the number of processes is over")
	offsetWarn := flag.Float64("warn", 10.0, "warning if the number of processes is under")
	flag.Parse()

	out, err := exec.Command("ntpq", "-c", "rv 0 offset").Output()
	if err != nil {
		fmt.Println("UNKNOWN: couldn't execute 'ntpq -c 'rv 0 offset''")
		os.Exit(UNKNOWN)
	}
	o := strings.Split(string(out), "=")

	if len(o) != 2 {
		fmt.Println("UNKNOWN: couldn't get ntp offset. ntpd process may be down.")
		os.Exit(UNKNOWN)
	}

	offset, err := strconv.ParseFloat(strings.Trim(o[1], "\n"), 64)
	if err != nil {
		fmt.Println("UNKNOWN: couldn't parse result.")
		os.Exit(UNKNOWN)
	}

	if *offsetCrit < math.Abs(offset) {
		msg := fmt.Sprintf("CRITICAL: ntp offset is over %f(actual) > %f(threshold)", math.Abs(offset), *offsetCrit)
		fmt.Println(msg)
		os.Exit(CRITICAL)
	} else if *offsetWarn < math.Abs(offset) {
		msg := fmt.Sprintf("WARNING: ntp offset is over %f(actual) > %f(threshold)", math.Abs(offset), *offsetWarn)
		fmt.Println(msg)
		os.Exit(WARNING)
	}

	msg := fmt.Sprintf("OK: ntp offset is %f(actual) < %f(warning threshold), %f(critial threshold)", offset, *offsetWarn, *offsetCrit)
	fmt.Println(msg)
	os.Exit(OK)
}

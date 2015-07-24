package main

import (
	"flag"
	"fmt"
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

func GetTheNumberOfProcesses(procName string) int {
	out, _ := exec.Command("ps", "awx").Output()
	procs := strings.Split(string(out), "\n")

	selfPid := strconv.Itoa(os.Getpid())
	selfPPid := strconv.Itoa(os.Getppid())

	num := 0
	for _, proc := range procs {
		if strings.Index(proc, procName) == -1 || strings.Index(proc, selfPid) != -1 || strings.Index(proc, selfPPid) != -1 {
			continue
		}
		num++
	}
	return num
}

func main() {
	procName := flag.String("name", "", "process name")
	procCritOver := flag.Int("critover", 1, "critical if the number of processes is over")
	procCritUnder := flag.Int("critunder", 1, "critical if the number of processes is under")
	procWarnOver := flag.Int("warnover", 1, "warning if the number of processes is over")
	procWarnUnder := flag.Int("warnunder", 1, "warning if the number of processes is under")
	flag.Parse()

	if *procName == "" {
		fmt.Println("-name argument is empty")
		os.Exit(UNKNOWN)
	}

	num := GetTheNumberOfProcesses(*procName)

	if num < *procCritUnder || *procCritOver < num {
		msg := fmt.Sprintf("CRITICAL: %s %d proc found, critover: %d, critunder: %d, warnover: %d, warnunder: %d",
			*procName, num, *procCritOver, *procCritUnder, *procWarnOver, *procWarnUnder)
		fmt.Println(msg)
		os.Exit(CRITICAL)
	} else if num < *procWarnUnder || *procWarnOver < num {
		msg := fmt.Sprintf("WARNING: %s %d proc found, critover: %d, critunder: %d, warnover: %d, warnunder: %d",
			*procName, num, *procCritOver, *procCritUnder, *procWarnOver, *procWarnUnder)
		fmt.Println(msg)
		os.Exit(WARNING)
	}

	msg := fmt.Sprintf("OK: %s %d proc found, critover: %d, critunder: %d, warnover: %d, warnunder: %d",
		*procName, num, *procCritOver, *procCritUnder, *procWarnOver, *procWarnUnder)
	fmt.Println(msg)
	os.Exit(OK)
}

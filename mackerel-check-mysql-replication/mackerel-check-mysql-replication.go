package main

import (
	"flag"
	"fmt"
	"github.com/ziutek/mymysql/mysql"
	_ "github.com/ziutek/mymysql/native"
	"os"
)

const (
	OK       = 0
	CRITICAL = 1
	WARNING  = 2
	UNKNOWN  = 3
)

func main() {
	optHost := flag.String("host", "localhost", "Hostname")
	optPort := flag.String("port", "3306", "Port")
	optUser := flag.String("username", "root", "Username")
	optPass := flag.String("password", "", "Password")
	optCrit := flag.Int("crit", 1, "critical if the second behind master is over")
	optWarn := flag.Int("warn", 1, "warning if the second behind master is over")
	flag.Parse()

	target := fmt.Sprintf("%s:%s", *optHost, *optPort)
	db := mysql.New("tcp", "", target, *optUser, *optPass, "")
	err := db.Connect()
	if err != nil {
		fmt.Println("UNKNOWN: couldn't connect DB")
		os.Exit(UNKNOWN)
	}
	defer db.Close()

	rows, res, err := db.Query("show slave status")
	if err != nil {
		fmt.Println("UNKNOWN: couldn't execute query")
		os.Exit(UNKNOWN)
	}
	if len(rows) == 0 {
		fmt.Println("OK: MySQL is not slave")
		os.Exit(OK)
	}

	idx_io_thread_runninng := res.Map("Slave_IO_Running")
	idx_sql_thread_runninng := res.Map("Slave_SQL_Running")
	idx_seconds_behind_master := res.Map("Seconds_Behind_Master")
	io_thead_status := rows[0].Str(idx_io_thread_runninng)
	sql_thead_status := rows[0].Str(idx_sql_thread_runninng)
	seconds_behind_master := rows[0].Int(idx_seconds_behind_master)

	if io_thead_status == "No" || sql_thead_status == "No" {
		fmt.Println("CRITICAL: MySQL replication has been stopped")
		os.Exit(CRITICAL)
	}

	if seconds_behind_master > *optCrit {
		msg := fmt.Sprintf("CRITICAL: MySQL replication behind master %d seconds", seconds_behind_master)
		fmt.Println(msg)
		os.Exit(CRITICAL)
	} else if seconds_behind_master > *optWarn {
		msg := fmt.Sprintf("WARNING: MySQL replication behind master %d seconds", seconds_behind_master)
		fmt.Println(msg)
		os.Exit(WARNING)
	}

	fmt.Println("OK: MySQL replication works well")
	os.Exit(OK)
}

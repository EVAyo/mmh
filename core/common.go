package core

import (
	"os"
	"os/user"
	"strings"

	"fmt"
)

// print layout func
func listLayout(name string) string {
	if len(name) < 14 {
		return fmt.Sprintf("%-14s", name)
	} else {
		return fmt.Sprintf("%-14s", shortenString(name, 14))
	}
}

// merge tags
func mergeTags(tags []string) string {
	return strings.Join(tags, ",")
}

func checkAndExit(err error) {
	if err != nil {
		fmt.Println("😱 " + err.Error())
		os.Exit(1)
	}
}

func checkErr(err error) bool {
	if err != nil {
		fmt.Println("😱 " + err.Error())
		return false
	}

	return true
}

func shortenString(str string, n int) string {
	if len(str) <= n {
		return str
	} else {
		return str[:n]
	}
}

func isRoot() bool {
	u, err := user.Current()
	checkAndExit(err)
	return u.Uid == "0" || u.Gid == "0"
}

func Exit(message string, code int) {
	if strings.TrimSpace(message) == "" {
		message = "No message"
	}
	fmt.Println("😱 " + message)
	os.Exit(code)
}

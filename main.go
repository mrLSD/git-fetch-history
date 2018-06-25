package main

import (
	"fmt"
	"os/exec"
	"strings"
	"time"
)

// commitData - commit data
type commitData struct {
	hash string
	date time.Time
}

func main() {
	hashes := getGitHashes()

	res, err = exec.Command("git", "log", "--format=%aI").Output()
	if err != nil {
		fmt.Printf("Error: %#v\n", err.Error())
		return
	}
	dates := strings.Split(string(res), "\n")
	if len(dates) < 2 {
		println("Wrong dates length")
		return
	}
	if len(dates) != len(hashes) {
		println("Hashes length and dates length not equal")
		return
	}

	for i := 0; i < len(dates)-1; i++ {
		t, err := time.Parse(time.RFC3339, dates[i])
		if err != nil {
			println("Parse time error: ", err.Error())
			return
		}
		if t.Year() > 2016 {
			t = t.AddDate(-2, 0, 0)
		}
		fmt.Println("Hash: ", hashes[i])
		fmt.Println("Date: ", t.String())
		hash := hashes[i]
		date := t.String()

		filter := fmt.Sprintf("git filter-branch -f --env-filter "+
			"'if [ $GIT_COMMIT = %s ] then "+
			" export GIT_AUTHOR_DATE=\"%s\" fi'", hash, date)
		println(filter)
	}
}

func getFilterString() string {

}

func getGitHashes() (hashes []string) {
	res, err := exec.Command("git", "log", "--format=%h").Output()
	if err != nil {
		fmt.Printf("Error: %#v\n", err.Error())
		return
	}

	hashes = strings.Split(string(res), "\n")
	if len(hashes) < 2 {
		println("Wrong hashes length")
		return
	}
	return
}

// applyChanges - apply git filter changes
func applyChanges(filter string) {
	_, err = exec.Command("git", "filter-branch", "-f", "--env-filter", filter).Output()
	if err != nil {
		fmt.Printf("Error apply changes: %#v\n", err.Error())
		return
	}
}

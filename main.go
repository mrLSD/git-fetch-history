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

		filter := fmt.Sprintf("'if [ $GIT_COMMIT = %s ] then "+
			" export GIT_AUTHOR_DATE=\"%s\" fi'", hash, date)
		println(filter)
	}
}

// getFilterString - generate filter param string
func getFilterString(commit commitData) string {
	return fmt.Sprintf("'if [ $GIT_COMMIT = %s ] then "+
		" export GIT_AUTHOR_DATE=\"%s\" fi'", commit.hash, commit.date.String())
}

// getCommitsHash - get all commits hashes
func getCommitsHash() (hashes []string, resErr error) {
	res, err := exec.Command("git", "log", "--format=%h").Output()
	if err != nil {
		resErr = fmt.Errorf("Error: %#v\n", err.Error())
		return
	}

	hashes = strings.Split(string(res), "\n")
	if len(hashes) < 2 {
		resErr = fmt.Errorf("Wrong hashes length")
		return
	}
	return
}

// applyChanges - apply git filter changes for commit
func applyChanges(filter string) error {
	_, err := exec.Command("git", "filter-branch", "-f", "--env-filter", filter).Output()
	if err != nil {
		return fmt.Errorf("Error apply changes: %#v\n", err.Error())
	}
	return nil
}

// getCommitsDate - get all commits date time
func getCommitsDate() (dates []time.Time, resErr error) {
	res, err := exec.Command("git", "log", "--formaI=%h").Output()
	if err != nil {
		resErr = fmt.Errorf("Error: %#v\n", err.Error())
		return
	}

	dateStr := strings.Split(string(res), "\n")
	if len(dateStr) < 2 {
		resErr = fmt.Errorf("Wrong hashes length")
		return
	}

	for i := 0; i < len(dateStr)-1; i++ {
		t, err := time.Parse(time.RFC3339, dateStr[i])
		if err != nil {
			resErr = fmt.Errorf("Parse time error: ", err.Error())
			return
		}
		dates = append(dates, t)
	}
	return
}

// getCommitsData - get all commits data
func getCommitsData() (data []commitData, resErr error) {
	hashes, err := getCommitsHash()
	if err != nil {
		resErr = err
		return
	}

	dates, err := getCommitsDate()
	if err != nil {
		resErr = err
		return
	}

	if len(dates) != len(hashes) {
		resErr = fmt.Errorf("Hashes length and dates length not equal")
		return
	}
	for i:= 0; i< len(hashes); i++ {
		data = append(data, commitData{
			hash: hashes[i],
			date: dates[i],
		})
	}
	return
}
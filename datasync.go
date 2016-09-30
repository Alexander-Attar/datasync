package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var dataCenters [][]string
	scanner := bufio.NewScanner(os.Stdin)
	_ = scanner.Scan()

	dcId := 0
	master := 0
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)
		}

		line := scanner.Text()
		fmt.Println(line)
		dataCenters = append(dataCenters, strings.Split(line, " "))

		// Find the data center with the most dataset ids
		if len(strings.Split(line, " ")) > len(dataCenters[master]) {
			master = dcId
		}

		dcId++
	}

	for fromDc, dc := range dataCenters {

		var diff = difference(dc, dataCenters[master])

		for _, dsId := range diff {
			fmt.Println(dsId, fromDc+1, master+1)
			dataCenters[master] = append(dataCenters[master], dsId)
		}
	}

	for fromDc, dc := range dataCenters {

		var diff = difference(dataCenters[master], dc)

		for _, dsId := range diff {
			fmt.Println(dsId, fromDc+1, master+1)
			dataCenters[fromDc] = append(dataCenters[fromDc], dsId)
		}
	}

	fmt.Println("done")
	fmt.Println(dataCenters)
}

func difference(slice1 []string, slice2 []string) []string {
	var diff []string

	for _, s1 := range slice1 {
		found := false
		for _, s2 := range slice2 {
			if s1 == s2 {
				found = true
				break
			}
		}
		// String not found. We add it to return slice
		if !found {
			diff = append(diff, s1)
		}
	}

	return diff
}

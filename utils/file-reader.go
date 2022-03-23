package utils

import (
	"bufio"
	"os"
	"strings"
)

func ReadLines(filename string) ([]string, error) {
	data, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	var hosts []string
	sc := bufio.NewScanner(data)
	for sc.Scan() {
		hosts = append(hosts, strings.ToLower(sc.Text()))
	}

	return hosts, data.Close()
}

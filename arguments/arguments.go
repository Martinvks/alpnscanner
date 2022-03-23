package arguments

import (
	"errors"
	"flag"
	"fmt"
)

const (
	ModeIana = iota
	ModeHttp
	ModeH2Draft
)

type Arguments struct {
	Concurrency int
	Port        int
	Mode        int
	HostsFile   string
}

func GetArguments() (Arguments, error) {
	var concurrency int
	flag.IntVar(&concurrency, "c", 10, "concurrency")

	var port int
	flag.IntVar(&port, "p", 443, "port")

	var hostsFile string
	flag.StringVar(&hostsFile, "h", "", "hosts file path")

	var mode string
	flag.StringVar(
		&mode,
		"m",
		"iana",
		"mode specifies which protocols to scan for. Must be one of \"iana\", \"http\" or \"h2draft\"",
	)

	flag.Parse()

	if hostsFile == "" {
		return Arguments{}, errors.New("missing hosts file")
	}

	var intMode int
	switch mode {
	case "iana":
		intMode = ModeIana
	case "http":
		intMode = ModeHttp
	case "h2draft":
		intMode = ModeH2Draft
	default:
		return Arguments{}, errors.New(
			fmt.Sprintf("unknown mode %v", mode),
		)
	}

	return Arguments{
		Concurrency: concurrency,
		Port:        port,
		Mode:        intMode,
		HostsFile:   hostsFile,
	}, nil
}

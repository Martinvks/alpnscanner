package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/Martinvks/alpnscanner/arguments"
	"github.com/Martinvks/alpnscanner/scanner"
	"github.com/Martinvks/alpnscanner/utils"
)

func main() {
	args, err := arguments.GetArguments()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error parsing arguments: %v\n", err)
		os.Exit(1)
	}

	hosts, err := utils.ReadLines(args.HostsFile)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error reading hosts file: %v\n", err)
		os.Exit(1)
	}

	protocols := utils.GetProtocols(args.Mode)

	scannerWG := sync.WaitGroup{}
	outputWG := sync.WaitGroup{}

	scannerChan := make(chan string)
	outputChan := make(chan string)

	for i := 0; i < args.Concurrency; i++ {
		scannerWG.Add(1)
		go func() {
			for host := range scannerChan {
				result := scanner.Scan(host, protocols)
				outputChan <- fmt.Sprintf("%v %v", host, result)
			}
			scannerWG.Done()
		}()
	}

	go func() {
		scannerWG.Wait()
		close(outputChan)
	}()

	outputWG.Add(1)
	go func() {
		for output := range outputChan {
			fmt.Println(output)
		}
		outputWG.Done()
	}()

	for _, host := range hosts {
		scannerChan <- fmt.Sprintf("%v:%d", host, args.Port)
	}
	close(scannerChan)

	outputWG.Wait()
}

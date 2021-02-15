package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"sync"
)

func resolve(isVerbose bool, host string, wg *sync.WaitGroup) {
	defer wg.Done()
	addr, _ := net.LookupIP(host)

	if addr != nil {
		if isVerbose {
			fmt.Printf("%v,%v\n", host, addr[0])
		} else {
			fmt.Printf("%v\n", host)
		}
	}
}

func main() {
	isVerbose := false
	fileName := ""

	flag.BoolVar(&isVerbose, "v", false, "If this flag is set, ip adress is also displayed with comma seperated")
	flag.StringVar(&fileName, "l", "", "Input file")
	flag.Parse()

	var sc *bufio.Scanner = nil

	if fileName != "" {
		file, _ := os.Open(fileName)
		sc = bufio.NewScanner(file)
	} else {
		sc = bufio.NewScanner(os.Stdin)
	}

	var wg sync.WaitGroup
	for sc.Scan() {
		line := sc.Text()
		wg.Add(1)
		//fmt.Printf("%v\n", line)
		go resolve(isVerbose, line, &wg)
	}
	wg.Wait()

}

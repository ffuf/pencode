package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/ffuf/pencode/pkg/pencode"
)

func main() {
	chain := pencode.NewChain()

	flag.Usage = func() {
		fmt.Printf("pencode - complex payload encoder v%s\n\n", pencode.VERSION)
		fmt.Printf("Usage: %s ENCODER1 ENCODER2 ENCODER3...\n\n", os.Args[0])
		fmt.Printf("%s reads input from stdin, which is typically piped from another process.\n\n", os.Args[0])
		fmt.Printf("Available encoders:\n")
		chain.Usage()
	}

	var listMode bool
	flag.BoolVar(&listMode, "list", false, "list available encoders")
	flag.Parse()

	if listMode {
		fmt.Println(strings.Join(chain.GetEncoders(), "\n"))
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(1)
	}

	err := chain.Initialize(os.Args[1:])
	if err != nil {
		flag.Usage()
		fmt.Printf("\n[!] %s\n", err)
		os.Exit(1)
	}

	input := readInput()
	output, err := chain.Encode([]byte(input))
	if err != nil {
		fmt.Printf("  [!] %s\n", err)
	}
	fmt.Println(string(output))
}

func readInput() string {
	input := os.Stdin
	defer input.Close()
	reader := bufio.NewScanner(input)

	data := ""
	for reader.Scan() {
		data = data + reader.Text()
	}
	return data
}

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

	inputWordlist := flag.String("input", "", "A wordlist to encode")

	flag.Usage = func() {
		fmt.Printf("pencode - complex payload encoder v%s\n\n", pencode.VERSION)
		fmt.Printf("Usage: %s FUNC1 FUNC2 FUNC3...\n\n", os.Args[0])
		fmt.Printf("%s reads input from stdin by default, which is typically piped from another process.\n\n", os.Args[0])
		fmt.Printf("OPTIONS:\n-input reads input from a file, line by line.\n\n")
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

	err := chain.Initialize(flag.Args())
	if err != nil {
		flag.Usage()
		fmt.Printf("\n[!] %s\n", err)
		os.Exit(1)
	}

	if *inputWordlist != "" {
		// read the input wordlist and output to stdout
		file, err := os.Open(*inputWordlist)

		if err != nil {
			fmt.Println(err)
		}

		defer file.Close()

		fs := bufio.NewScanner(file)
		fs.Split(bufio.ScanLines)

		for fs.Scan() {
			output, err := chain.Encode(fs.Bytes())
			if err != nil {
				fmt.Printf("  [!] %s\n", err)
			}
			fmt.Println(string(output))
		}
	} else {
		input := readInput()
		output, err := chain.Encode([]byte(input))
		if err != nil {
			fmt.Printf("  [!] %s\n", err)
		}
		fmt.Print(string(output))
	}
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

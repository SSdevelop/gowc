package main

import (
	"fmt"
	"io"
	"os"

	"github.com/SSdevelop/gowc/argparser"
	"github.com/SSdevelop/gowc/counter"
)

func main() {
	args := os.Args[1:]
	parsedArgs, err := argparser.ParseArgs(args)
	if err != nil {
		os.Stderr.WriteString("Error parsing arguments: " + err.Error() + "\n")
		os.Exit(1)
	}
	var currentReader io.Reader
	var display string
	// print(parsedArgs.CountBytes, " ", parsedArgs.CountLines, " ", parsedArgs.CountWords, " ", parsedArgs.CountChars, " ", parsedArgs.FileName, "\n")
	if parsedArgs.FileName == "" {
		currentReader = os.Stdin
	} else {
		file, err := os.Open(parsedArgs.FileName)
		if err != nil {
			os.Stderr.WriteString("Error opening file: " + err.Error() + "\n")
			os.Exit(1)
		}
		defer file.Close()
		currentReader = file
	}
	counts, err := counter.ProcessFile(currentReader)
	if err != nil {
		os.Stderr.WriteString("Error processing file: " + err.Error() + "\n")
		os.Exit(1)
	}
	display = formatOutput(parsedArgs, counts)
	_, err = os.Stdout.WriteString(display)
	if err != nil {
		os.Stderr.WriteString("Error writing output: " + err.Error() + "\n")
		os.Exit(1)
	}
}

func formatOutput(args argparser.Arg, counts counter.Counter) string {
	var output string
	if args.CountLines {
		output += fmt.Sprintf("%d\t", counts.Lines)
	}
	if args.CountWords {
		output += fmt.Sprintf("%d\t", counts.Words)
	}
	if args.CountBytes {
		output += fmt.Sprintf("%d\t", counts.Bytes)
	}
	if args.CountChars {
		output += fmt.Sprintf("%d\t", counts.Chars)
	}
	return output + args.FileName + "\n"
}

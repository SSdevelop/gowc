package counter

import (
	"bufio"
	"io"
	"unicode"
)

type Counter struct {
	Bytes int
	Lines int
	Words int
	Chars int
}

func ProcessFile(r io.Reader) (Counter, error) {
	reader := bufio.NewReader(r)
	var counts Counter
	inWord := false
	currentLineHasContent := false
	for {
		char, size, err := reader.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			return Counter{}, err
		}

		counts.Bytes += size
		counts.Chars += 1
		if char == '\n' {
			if currentLineHasContent {
				counts.Lines += 1
			}
			currentLineHasContent = false
		} else {
			currentLineHasContent = true
		}
		if unicode.IsSpace(char) {
			if inWord {
				inWord = false
			}
		} else if !inWord {
			counts.Words += 1
			inWord = true
		}
	}
	return counts, nil
}

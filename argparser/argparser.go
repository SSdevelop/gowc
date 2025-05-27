package argparser

import (
	"fmt"
	"strings"
)

type Arg struct {
	CountBytes bool
	CountLines bool
	CountWords bool
	CountChars bool
	FileName   string
}

func ParseArgs(rawArgs []string) (Arg, error) {
	parsed := Arg{
		CountLines: false,
		CountBytes: false,
		CountWords: false,
		CountChars: false,
		FileName:   "",
	}
	explicitOptions := false
	parsing := true

	for _, arg := range rawArgs {
		if parsing && strings.HasPrefix(arg, "-") && len(arg) > 1 {
			option := arg[1:]
			for _, char := range option {
				explicitOptions = true
				switch char {
				case 'l':
					parsed.CountLines = true
				case 'c':
					parsed.CountBytes = true
				case 'w':
					parsed.CountWords = true
				case 'm':
					parsed.CountChars = true
				default:
					return Arg{}, fmt.Errorf("unknown option: %s", string(char))
				}
			}
		} else if parsing && !strings.HasPrefix(arg, "-") {
			parsed.FileName = arg
		} else {
			parsing = false
		}
	}
	if !explicitOptions {
		parsed.CountLines = true
		parsed.CountWords = true
		parsed.CountBytes = true
	}
	return parsed, nil
}

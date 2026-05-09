package iteration

import (
	"strings"
)

func Repeat(char string, numOfRepeats int) string {
	var repeat strings.Builder

	if numOfRepeats < 0 {
		numOfRepeats = numOfRepeats * -1
	}

	for i := 0; i < numOfRepeats; i++ {
		repeat.WriteString(char)
	}

	return repeat.String()
}

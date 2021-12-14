package main

import (
	"strings"
)

func main() {
	learnStrconv()
	learnTime()
	learnRegexp()
}

func parseComment(comments string) string {
	type Status int
	const (
		Init Status = iota
		Empty
		NotEnd
		End
	)
	var (
		switchPrintMap = []map[Status]string{
			Init:   {Empty: "", NotEnd: "", End: ""},
			Empty:  {Empty: "", NotEnd: "", End: ""},
			NotEnd: {Empty: "\n\n", NotEnd: " ", End: " "},
			End:    {Empty: "\n\n", NotEnd: "\n", End: "\n"},
		}
		builder = strings.Builder{}
		status  = Init
	)
	for _, line := range strings.Split(comments, "\n") {
		line = strings.TrimSpace(strings.Trim(line, "//"))
		var currStatus Status
		if line == "" {
			currStatus = Empty
		} else if strings.HasSuffix(line, ".") || strings.HasSuffix(line, "。") {
			currStatus = End
		} else {
			currStatus = NotEnd
		}

		builder.WriteString(switchPrintMap[status][currStatus])
		builder.WriteString(line)

		status = currStatus
	}
	return builder.String()
}

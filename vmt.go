package vmt

import (
	"strings"
	"fmt"
	"strconv"
	"bytes"
)

func Sprintq(format string, a ...interface{}) string {
	updatedFormat := strings.Replace(format,"?", "%s", -1)
	return fmt.Sprintf(updatedFormat, a...)
}

func isNum(n rune) bool {
	return  '0' <= n && n <= '9'
}

func Sprints(format string, a ...interface{}) string {

	var i = 0
	formatRunes := []rune(format)
	lenFormat := len(formatRunes)

	lenArguments := len(a)
	var buffer bytes.Buffer

	for i<lenFormat {
		var c = formatRunes[i]

		if c=='$' {
			j:=i+1
			if j<lenFormat && isNum(formatRunes[j]) {
				for j < lenFormat {
					if isNum(formatRunes[j]) {
						j++
					} else {
						break
					}
				}

				strNum,err := strconv.Atoi(string(formatRunes[i+1:j]))
				if err!=nil {
					fmt.Errorf("error in processing number ")
				}

				if(strNum-1>lenArguments) {
					fmt.Errorf("arguments number index not found")
				}

				buffer.WriteString(fmt.Sprintf("%v", a[strNum-1]))
				i = j-1
			} else {
				buffer.WriteRune(c)
			}
		} else {
			buffer.WriteRune(c)
		}

		i++
	}

	return buffer.String()
}

func Printq(format string, a ...interface{}) {
	fmt.Print(Sprintq(format, a...))
}

func Prints(format string, a ...interface{}) {
	fmt.Print(Sprints(format, a...))
}
package vmt

import (
	"strings"
	"fmt"
)

func Sprintq(format string, a ...interface{}) string {
	updatedFormat := strings.Replace(format,"?", "%s", -1)
	return fmt.Sprintf(updatedFormat, a...)
}

func Sprints(format string, a ...interface{}) string {

	var i = 0
	formatRunes := []rune(format)
	lenFormat := len(formatRunes)

	for i<lenFormat {
		var c = formatRunes[i]
		fmt.Printf("character %c starts at byte position %d\n", c, i)

		if c=='$' {
			fmt.Println("dollar sign")

			j:=i+1
			for j<lenFormat {
				cInner:=formatRunes[j]
				if '0' <= cInner && cInner <= '9' {
					j++
				} else {
					break
				}
			}
		}

		i++
	}

	return format
}


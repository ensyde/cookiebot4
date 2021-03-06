package main

import (
	"strings"
	"regexp"
)

func WildcardCompare(s string, pattern string) bool {
        s = strings.ToLower(s)
        pattern = strings.ToLower(pattern)

        pattern = createRegex(pattern)
        regex := regexp.MustCompile(pattern)

        return regex.MatchString(s)
}

func createRegex(s string) string {
        var wc string
        for i := 0; i < len(s); i++ {
                if s[i] == '*' {
                        wc += ".*"
                } else if s[i] == '?' {
                        wc += "."
                } else if s[i] >= '0' && s[i] <= '9' {
                        wc += string(s[i])
                } else if s[i] >= 'a' && s[i] <= 'z' {
                        wc += string(s[i])
                } else {
                        wc += "\\" + string(s[i])
                }
        }

        return "^" + wc + "$"
}

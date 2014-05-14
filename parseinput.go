package main

import "os"

//import "regexp"
import "strings"

func popclip() string {
	//	env := os.Environ()
	//	popclipPattern := regexp.MustCompile("^POPCLIP_TEXT=(.*)$")
	//	for _, kw := range env {
	//		if popclipPattern.MatchString(kw) {
	//			return popclipPattern.FindStringSubmatch(kw)[1]
	//		}
	//	}
	return os.ExpandEnv("$POPCLIP_TEXT")
}
func commandline() string {
	return strings.Join(os.Args[1:], " ")
}

func getText() string {
	if len(os.Args) > 1 {
		return commandline()
	} else {
		return popclip()
	}
	return ""
}

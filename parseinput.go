package main

import "os"

//import "regexp"
import "strings"
import "net/url"

func popclip() string {
	var popcliptext string
	if i := os.ExpandEnv("$POPCLIP_URLENCODED_TEXT"); i != "" {
		popcliptext = i
	} else if i := os.ExpandEnv("$POPCLIP_TEXT"); i != "" {
		popcliptext = url.QueryEscape(i)
	}
	return popcliptext
}
func commandline() string {
	return url.QueryEscape(strings.Join(os.Args[1:], " "))
}

func getText() string {
	if len(os.Args) > 1 {
		return commandline()
	} else {
		return popclip()
	}
	return ""
}

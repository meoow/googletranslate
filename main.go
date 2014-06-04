package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	//	"unicode/utf8"
)

const (
	//	requestAddrPattern string = "http://translate.google.com/translate_a/t?client=x&text=%s&hl=en&sl=%s&tl=%s"
	//requestAddrPattern string = "http://74.125.235.201/translate_a/t?client=x&text=%s&hl=en&sl=%s&tl=%s"
	httpP              string = "http://"
	defaultAddr        string = "translate.google.com"
	requestAddrPattern string = "/translate_a/t?client=x&text=%s&hl=en&sl=%s&tl=%s"
	referer            string = "http://translate.google.com"
	userAgent          string = "Mozilla/5.0 (compatible; Googlebot/2.1)"
	language           string = "en-us,en;q=0.5"
	charset            string = "ISO-8859-1,utf-8;q=0.7,*;q=0.7"
	format             string = "text/xml,application/xml,application/xhtml+xml,text/html;q=0.9,text/plain;q=0.8,image/png,*/*;q=0.5"
	SENTENCES          uint8  = 1
	DICT               uint8  = 2
)

/*
type indivEntry struct {
	word string
	rtrans []string
	score float64
}*/

type indivTerm struct {
	pos   string
	terms []string
	//	entry []*indivEntry
}

type dict []*indivTerm

func main() {
	gtFrom := os.ExpandEnv("$GTF")
	if gtFrom == "" {
		gtFrom = "en"
	}
	gtTo := os.ExpandEnv("$GTT")
	if gtTo == "" {
		gtTo = "zh-CN"
	}
	gtAddr := os.ExpandEnv("$GTADDR")
	if gtAddr == "" {
		gtAddr = defaultAddr
	}
	requestAddrPtn := httpP + gtAddr + requestAddrPattern
	text := getText()
	replWhiteSpace := regexp.MustCompile("[ \n\t]")
	text = replWhiteSpace.ReplaceAllString(text, "%20")
	requestAddr := fmt.Sprintf(requestAddrPtn, text, gtFrom, gtTo)
	reqst, _ := http.NewRequest("GET", requestAddr, nil)
	reqst.Header.Add("Referer", referer)
	reqst.Header.Add("User-Agent", userAgent)
	reqst.Header.Add("Accept-Language", language)
	reqst.Header.Add("Accept-Charset", charset)

	client := &http.Client{}

	resp, err := client.Do(reqst)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	resultJson := make(map[string]interface{})
	err = json.Unmarshal(result, &resultJson)
	if err != nil {
		log.Fatal(err)
	}

	queryType := SENTENCES
	if len(os.Args) == 2 && os.ExpandEnv("$GTDICT") != "" {
		queryType = DICT
	}
	fmt.Printf("%s\n", parseJson(resultJson, queryType))
}

func parseJson(js map[string]interface{}, catalog uint8) string {
	switch catalog {
	case SENTENCES:
		sentences, _ := js["sentences"].([]interface{})
		output := make([]string, 0, 3)
		for _, sentence := range sentences {
			line := sentence.(map[string]interface{})["trans"]
			output = append(output, line.(string))
		}
		return strings.Join(output, " ")
	case DICT:
		if d, ok := js["dict"]; ok {
			output := make(dict, 0, 5)
			dicts, _ := d.([]interface{})
			for _, term := range dicts {
				term_, _ := term.(map[string]interface{})
				pos := term_["pos"].(string)
				terms_ := term_["terms"].([]interface{})
				terms := make([]string, len(terms_))
				for i, _ := range terms_ {
					terms[i] = terms_[i].(string)
				}
				/*entry := make([]*indivEntry, 0, 3)
				for _, e_ := range term_["entry"].([]interface{}){
					e := e_.(map[string]interface{})
					word := e["word"].(string)
					rtrans_ := e["reverse_translation"].([]interface{})
					rtrans := make([]string, len(rtrans_))
					for i,_:= range rtrans_ {
						rtrans[i] = rtrans_[i].(string)
					}
					var score float64
					switch x:=e["score"].(type) {
					case float64: score = x
					case nil: score = 0.0
					}
					entry = append(entry, &indivEntry{word,rtrans,score})
				}*/
				output = append(output, &indivTerm{pos, terms}) //, entry})
			}
			return parseDict(output)
		} else {
			return ""
		}
	}
	return ""
}

func parseDict(d dict) string {
	lines := make([]string, 0, 10)
	for _, v := range d {
		lines = append(lines, fmt.Sprintf("%s: %s", strings.ToUpper(v.pos),
			strings.Join(v.terms, ", ")))
		/*for _, p2entry := range v.entry {
			byteCount := len([]byte(p2entry.word))
			runeCount := utf8.RuneCount([]byte(p2entry.word))
			deltaWidth := byteCount - runeCount
			lines = append(lines, fmt.Sprintf("%4s%-*s: %s","",
							12 - deltaWidth/2,
							p2entry.word,
							strings.Join(p2entry.rtrans,", ")))
		}*/
	}
	return strings.Join(lines, "\n")
}

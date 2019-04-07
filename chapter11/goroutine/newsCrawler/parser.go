package main

import (
	"regexp"
)

func ParseIndex(data []byte) [][]string {
	regexpString :=
		`\<li.*?\>\s*\<a.*?href=\"(.*?)\".*?\>(.+?)\<\/a\>\s*\<\/li\>`
	re := regexp.MustCompile(regexpString)
	content := string(data)
	matches := re.FindAllStringSubmatch(content, -1)
	urlTitlePair := make([][]string, len(matches))
	for i, value := range matches {
		urlTitlePair[i] = []string{value[1], value[2]}
	}
	return urlTitlePair
}

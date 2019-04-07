package main

import (
	"io/ioutil"
	"net/http"
)

var ITNewsURIs = []string{
	"https://www.naver.com/nvnewsstand?953",
	"https://www.naver.com/nvnewsstand?092",
	"https://www.naver.com/nvnewsstand?910",
	"https://www.naver.com/nvnewsstand?977",
	"https://www.naver.com/nvnewsstand?952",
	"https://www.naver.com/nvnewsstand?917",
	"https://www.naver.com/nvnewsstand?029",
	"https://www.naver.com/nvnewsstand?030",
	"https://www.naver.com/nvnewsstand?818",
	"https://www.naver.com/nvnewsstand?138",
	"https://www.naver.com/nvnewsstand?293",
}

type NewsLink struct {
	URI     string
	Title   string
	Content string
}

type NewsIndex struct {
	Link []NewsLink
}

func Indexing() (newsIndex NewsIndex, err error) {
	index := NewsIndex{[]NewsLink{}}
	for _, url := range ITNewsURIs {
		resp, err := http.Get(url)
		if err != nil {
			return NewsIndex{}, err
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		titleUrlPair := ParseIndex(body)
		for _, pair := range titleUrlPair {
			index.Link = append(index.Link, NewsLink{
				URI:   pair[0],
				Title: pair[1],
			})
		}
	}
	return index, nil
}

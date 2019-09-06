package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Sitemapindex struct {
	Sitemaps []string `xml:"sitemap>loc"`
}

type News struct {
	Location []string `xml:"url>loc"`
	Title    []string `xml:"url>news>title"`
	Keywords []string `xml:"url>news>keywords"`
}

type NewsMap struct {
	Location string
	Keywords string
}

func main11() {

	var newsData News
	var sms Sitemapindex
	newsMap := make(map[string]NewsMap)

	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes, &sms)

	for _, biner := range sms.Sitemaps {
		resp, _ := http.Get(strings.TrimSpace(biner))
		bytes, _ := ioutil.ReadAll(resp.Body)

		//unmarshall all bytes/news
		xml.Unmarshal(bytes, &newsData)
		for idx := range newsData.Keywords {
			newsMap[newsData.Title[idx]] = NewsMap{newsData.Keywords[idx], newsData.Location[idx]}
		}
		resp.Body.Close()
		fmt.Println(newsMap)

	}

	resp.Body.Close()
}

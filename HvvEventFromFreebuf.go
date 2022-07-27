package main

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	Spider()
}

type HvvEventData struct {
	Time      string `json:time`
	WeakPoint string `json:weakpoint`
	Level     string `json:level`
	Describe  string `json:describe`
	Effect    string `json:effect`
}

func Spider() {

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://www.freebuf.com/news/340081.html", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Connection", "keep-alive") //设置请求头
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Host", "www.freebuf.com")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2")

	resp, err := client.Do(req) //拿到返回的内容
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	reader := resp.Body

	if resp.Header.Get("Content-Encoding") == "gzip" {
		reader, _ = gzip.NewReader(resp.Body)
	}

	docDetail, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		fmt.Println("err !! don`t read response body.")
		log.Fatal(err)
	}
	eventSet := docDetail.Find("div#tinymce-editor.content-detail > div > blockquote")

	var hvvEventSet []HvvEventData

	eventSet.Each(func(i int, s *goquery.Selection) {
		// HvvEventData;

		var hvvEventData00 HvvEventData
		s.Find("p").Each(func(index int, elem *goquery.Selection) {
			// 第一个 blockquote
			content := elem.Text()
			if i == 0 {
				if index == 0 {
					hvvEventData00.WeakPoint = content
				} else {
					htmlContent, err := elem.Html()
					if err != nil {
						log.Fatal(err)
					}
					contentElems := strings.Split(htmlContent, "<br/>")
					for _, ssss := range contentElems {
						if strings.Contains(ssss, "漏洞等级") {
							hvvEventData00.Level = ssss
						} else if strings.Contains(ssss, "影响范围") {
							hvvEventData00.Effect = ssss
						} else if strings.Contains(ssss, "披露时间") {
							hvvEventData00.Time = ssss
						} else if strings.Contains(ssss, "描述") {
							hvvEventData00.Describe = ssss
						}
					}

				}
			} else {
				if index == 0 {
					hvvEventData00.WeakPoint = content
				} else {
					if strings.Contains(content, "漏洞等级") {
						hvvEventData00.Level = content
					} else if strings.Contains(content, "影响范围") {
						hvvEventData00.Effect = content
					} else if strings.Contains(content, "披露时间") {
						hvvEventData00.Time = content
					} else if strings.Contains(content, "描述") {
						hvvEventData00.Describe = content
					}
				}
			}
		})
		hvvEventSet = append(hvvEventSet, hvvEventData00)
		// fmt.Println(s.Text())
	})

	for _, hvv := range hvvEventSet {
		bytes, _ := json.Marshal(hvv)
		fmt.Println(string(bytes))
	}

}

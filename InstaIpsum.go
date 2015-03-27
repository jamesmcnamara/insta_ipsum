package main 

import (
	"net/http"
	"net/url"
	"fmt"
        "os"
	"regexp"
	"strings"
	)

func pullText(paragraphs string) []byte {
	params := url.Values{}
	params.Set("paras", paragraphs)
	params.Set("type", "hipster-centric")

	client, error := http.Get("http://hipsum.co/?" + params.Encode())
	if error != nil {
		fmt.Println(error)	
	}
	body := make([]byte, client.ContentLength + 1)
	if len(body) == 0 {
		buffer := make([]byte, 1024)
		amt, error := client.Body.Read(buffer)
		for  amt > 0 && error == nil {
			body = append(body, buffer...)
			amt, error = client.Body.Read(buffer) 
		}
	} else {
		client.Body.Read(body)
	}
	return body
}

func cleanText(text []byte, clean_text bool) string {
	block_hunter := regexp.MustCompile("<div class=.hipsum.>.*</div>")
	text_block := string(block_hunter.Find(text))
	if !clean_text {
		return text_block
	}
	plain_text := strings.Split(text_block, "<p>")
	for i, s := range plain_text {
		plain_text[i] = strings.TrimSuffix(s, "</p>")
	}
	return strings.Join(plain_text[1:len(plain_text) - 1], " ")
}

func GetIpsum(paragraphs string, clean_text bool) string {
	raw_text := pullText(paragraphs)
	return cleanText(raw_text, clean_text)
}

func main() {
    fmt.Println(GetIpsum(os.Args[1], true))
}

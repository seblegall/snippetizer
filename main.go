package main

import (
	"github.com/golang-commonmark/markdown"
	"fmt"
	"net/http"
	"io/ioutil"
	"log"
	"flag"
	"net/url"
)

type snippet struct {
	content string
	lang string
}


func getSnippet(tok markdown.Token) snippet {
	switch tok := tok.(type) {
	case *markdown.CodeBlock:
		return snippet{
			tok.Content,
			"code",
		}
	case *markdown.CodeInline:
		return snippet{
			tok.Content,
			"code inline",
		}
	case *markdown.Fence:
		return snippet{
			tok.Content,
			tok.Params,
		}
	}
	return snippet{}
}

func readFromWeb(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}


func main() {

	var urlString string
	flag.StringVar(&urlString, "url", "", `The url of the github repository`)
	flag.Parse()

	if urlString == "" {
		log.Fatalln("Please, provide an url for the readme to parse.")
	}

	//Parse URL
	u, err := url.Parse(urlString)

	if err != nil {
		log.Fatalln("Impossible to parse the URL")
	}

	//read the readme file
	readMe, err := readFromWeb(fmt.Sprintf("https://raw.githubusercontent.com%s/master/README.md", u.Path))
	if err != nil {
		log.Fatalf(err.Error())
	}

	md := markdown.New(markdown.XHTMLOutput(true), markdown.Nofollow(true))
	tokens := md.Parse(readMe)

	for _, t := range tokens {
		snippet := getSnippet(t)

		if snippet.content != "" {
			fmt.Printf("##### Lang : %s ###### \n", snippet.lang)
			fmt.Println(snippet.content)
		}
	}

}
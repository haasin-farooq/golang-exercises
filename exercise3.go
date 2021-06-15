/*
Exercise 3: Parsing files
Goal is to create a package that makes it easy to parse an HTML file and extract all of the links (<a href="">...</a> tags).
For each extracted link you should return a data structure that includes both the href, as well as the text inside the link.
Any HTML inside of the link can be stripped out, along with any extra whitespace including newlines, back-to-back spaces, etc.

See details here: https://github.com/meer-online/link
*/

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	href string
	text string
}

func main() {
	f := readFile("ex4.html")

	var links []Link

	r := strings.NewReader(f)

	tokenizer := html.NewTokenizer(r)
	for {
		tokenType := tokenizer.Next()

		if tokenType == html.ErrorToken {
			err := tokenizer.Err()
			if err == io.EOF {
				break
			}
			log.Fatalf("Error: %v", tokenizer.Err())
		}

		if tokenType == html.StartTagToken {
			token := tokenizer.Token()

			if token.Data == "a" {
				var h string
				for _, a := range token.Attr {
					if a.Key == "href" {
						h = a.Val
					}
				}

				tokenType = tokenizer.Next() // go to the next token which will be of type Text
				token = tokenizer.Token() // get additional information about the token
				data := strings.TrimSpace(token.Data) // remove any leading and trailing whitespace from the token data

				links = append(links, Link{h, data}) // append custom data type variable Link to slice links
			}
		}
	}

	for _, l := range links {
		fmt.Printf("href: %v, text: %v\n", l.href, l.text)
	}
}

func readFile(filename string) string {
	f, err := os.OpenFile(filename, os.O_RDONLY, 0644)

	if err != nil {
		log.Fatal(err)
	}

	content, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	return string(content)
}
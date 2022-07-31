package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

func main() {

	resp, err := http.Get("http://webcode.me")

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {

		log.Fatal(err)
	}

	content := string(body)

	re := regexp.MustCompile("<[^>]*>")
	replaced := re.ReplaceAllString(content, "")

	fmt.Println(strings.TrimSpace(replaced))
}

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// https://google.github.io/styleguide/go/decisions.html

func main() {
	resp, err := http.Get("https://quran-api.santrikoding.com/api/surah")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)
	fmt.Println(sb)
}

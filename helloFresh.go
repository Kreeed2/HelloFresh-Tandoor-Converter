package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
)

type BearerToken struct {
	AccessToken string
}

func new() {
	req, err := http.NewRequest("POST", "https://www.hellofresh.de/gw/auth/token?client_id=senf&grant_type=client_credentials", nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header = http.Header{
		"User-Agent": {"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:97.0) Gecko/20100101 Firefox/97.0"},
		"Referer":    {"https://www.hellofresh.de/recipes"},
		"Origin":     {"https://www.hellofresh.de"},
		"Accept":     {"application/json"},
	}

	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{Jar: jar}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	for name, values := range res.Header {
		// Loop over all values for the name.
		for _, value := range values {
			fmt.Println(name, value)
		}
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	fmt.Println(string(body))
}

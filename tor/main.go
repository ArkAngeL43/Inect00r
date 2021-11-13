package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"time"
)

var (
	flagTarget = flag.String("t", "", `target URL`)
)

// URL to fetch

var Proxy string = "socks5://127.0.0.1:9050"

func main() {
	flag.Parse()
	if *flagTarget == "true" {
		fmt.Println("[-] Please input a url")
		fmt.Println("[-] go run main.go -t http://testphp.vulnweb.com/listproducts.php?cat=1")
		os.Exit(1)
	}
	//
	// tor
	//
	//
	torProxyUrl, err := url.Parse(Proxy)
	//
	if err != nil {
		log.Fatal(err)
	}
	//
	//
	//
	//
	injections := []string{
		"baseline",
		")",
		"(",
		"\"",
		"'",
	}
	errors := []string{
		"SQL",
		"MySQL",
		"ORA-",
		"syntax",
	}

	errRegexes := []*regexp.Regexp{}

	for _, e := range errors {
		re := regexp.MustCompile(fmt.Sprintf(".*%s.*", e))
		errRegexes = append(errRegexes, re)
	}

	for _, payload := range injections {
		torTransport := &http.Transport{Proxy: http.ProxyURL(torProxyUrl)}
		client := &http.Client{Transport: torTransport, Timeout: time.Second * 10}

		client = new(http.Client)
		body := []byte(fmt.Sprintf("username=%s&password=p", payload))

		resp, err := http.NewRequest(
			"POST",
			*flagTarget,
			bytes.NewReader(body),
		)
		//
		//
		//
		//
		if err != nil {
			log.Fatalf("[!] Unable to generate request: %s\n", err)
		}
		//
		//
		//
		//
		resp.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		resp, err = client.Do(resp)
		if err != nil {
			log.Fatalf("[!] Unable to process response: %s\n", err)
		}
		//
		//
		//
		//
		body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("[!] Unable to read response body: %s\n", err)
		}
		//
		//
		//
		resp.Body.Close() // close response
		//
		//
		//
		//
		//
		for idx, re := range errRegexes {
			if re.MatchString(string(body)) {
				fmt.Printf("[+] SQL Error found [Server->%s] for payload: %s\n", errors[idx], payload)
				break
			}
		}
	}
}

package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

type Network struct {
	timeout time.Duration
}

var (
	singleton Network
	once      sync.Once
)

func NewNetwork(timeout time.Duration) Network {
	once.Do(func() {
		singleton = Network{timeout}
	})

	return singleton
}

func (n *Network) Client() *http.Client {
	return &http.Client{
		Timeout: n.timeout * time.Second,
	}
}

func (n *Network) PrepareRequest(url string) (*http.Request, error) {
	return n.prepareGETRequest(url)
}

func (n *Network) prepareGETRequest(url string) (*http.Request, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	n.addHeaders(req)
	return req, nil
}

func (n *Network) addHeaders(req *http.Request) {
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")
	req.Header.Add("Accept-Language", "en-us, en;q=0.50")
}

func (n *Network) DoRequest(client *http.Client, request *http.Request, callback func(res *http.Response)) {
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	//n.CopyResponseToStdout(response)
	callback(response)
}

func (n *Network) CopyResponseToStdout(response *http.Response) {
	_, err := io.Copy(os.Stdout, response.Body)
	if err != nil {
		log.Fatal(err)
	}
}

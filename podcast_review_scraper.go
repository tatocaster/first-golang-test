package main

import (
	"net/http"
)

type App struct {
	baseUrl   string
	podcastId string
}

func NewApp(baseUrl, podcastId string) App {
	return App{
		baseUrl:   baseUrl,
		podcastId: podcastId,
	}
}

func main() {
	app := NewApp("https://podcasts.apple.com/us/podcast/", "id1447502263")
	network := NewNetwork(30)
	req, _ := network.PrepareRequest(app.baseUrl + app.podcastId)
	network.DoRequest(network.Client(), req, func(res *http.Response) {
		//network.CopyResponseToStdout(res)
		parseDOM(createDOM(res))
	})
}

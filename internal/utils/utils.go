package utils

import (
	"io"
	"net/http"
	"strconv"
	"strings"

	config "egeucak.com/kamil/cmd/configservice"
)

func MirrorRequest(w http.ResponseWriter, req *http.Request, url string) (body io.ReadCloser) {
	newReq, _ := http.NewRequest(req.Method, url, req.Body)
	for headerKey, headerValues := range req.Header {
		for _, val := range headerValues {
			newReq.Header.Add(headerKey, val)
		}
	}
	client := http.Client{}
	resp, err := client.Do(newReq)
	if err != nil {
		http.Error(w, err.Error(), 400)
	}
	return resp.Body
}

func BuildURL(routeConfig config.Endpoint, req *http.Request) string {
	var sb strings.Builder
	var port string = strconv.Itoa(routeConfig.Port)
	sb.WriteString("http://")
	sb.WriteString(routeConfig.Host)
	sb.WriteString(":")
	sb.WriteString(port)
	sb.WriteString(req.URL.Path)
	sb.WriteString("?")
	sb.WriteString(req.URL.RawQuery)
	url := sb.String()
	sb.Reset()
	return url
}

package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"

	config "egeucak.com/kamil/cmd/configservice"
	utils "egeucak.com/kamil/internal/utils"
)

var (
	gatewayPort       *int
	fileCheckInterval *int
	configFileName    *string
)

func parseFlags() {
	fileCheckInterval = flag.Int("config-check-interval", 2, "an int")
	configFileName = flag.String("config-file-name", "./config.yaml", "a path")
	gatewayPort = flag.Int("port", 3000, "an int")
	flag.Parse()
}

func isInRequests(a string, targetSlice []string) bool {
	for _, b := range targetSlice {
		if targetSlice != nil && len(targetSlice) > 0 && b == a {
			return true
		}
	}
	return len(targetSlice) == 0
}

func matchRequest(req *http.Request, routeMap config.APIConfigFile) (config.Endpoint, error) {
	reqPath := req.URL.Path
	for i, val := range routeMap.Endpoint {
		route := routeMap.Endpoint[i]
		matchingString := regexp.MustCompile(route.Route).FindString(reqPath)
		if matchingString == reqPath && isInRequests(req.Method, route.RequestTypes) {
			return val, nil
		}
	}
	return config.Endpoint{}, errors.New("Wrong request")
}

func main() {
	parseFlags()
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		routeConfig := *config.GetInstance(configFileName, fileCheckInterval)
		conf, err := matchRequest(req, routeConfig)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		url := utils.BuildURL(conf, req)
		fmt.Println(url)
		respBody := utils.MirrorRequest(w, req, url)
		defer respBody.Close()

		io.Copy(w, respBody)
	})
	port := ":" + strconv.Itoa(*gatewayPort)
	fmt.Println("Listening on ", port)
	http.ListenAndServe(port, nil)
}

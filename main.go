package main

import (
	"encoding/json"
	"fmt"
	"github.com/t-k/fluent-logger-golang/fluent"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	tgif := tgif()
	logger, err := fluent.New(fluent.Config{FluentPort: 24224, FluentHost: "localhost"})
	if err != nil {
		panic("Could not find fluentd.")
	}
	port := "3000"
	fmt.Println("Pipe has found fluentd and starting the server at port: " + port + " ...")
	tag := "com.kaiinui.pipe"

	// GET /e.gif
	http.HandleFunc("/e.gif", func(res http.ResponseWriter, req *http.Request) {
		fmt.Println("GET " + req.URL.String())

		encodedJson := req.URL.Query().Get("data")
		data, err := getJsonFromUnescapedString(encodedJson)
		if err != nil {
			fmt.Println("400 Bad Request")
			res.WriteHeader(http.StatusBadRequest)
			return
		}

		logger.Post(tag, data)

		res.Header().Set("Content-Type", "image/gif")
		res.Header().Set("Cache-Control", "nocache")
		res.Write(tgif)
	})

	// POST /e
	http.HandleFunc("/e", func(res http.ResponseWriter, req *http.Request) {
		fmt.Println("POST " + req.URL.String())
		dat, _ := ioutil.ReadAll(req.Body)
		body := string(dat)
		fmt.Println(body)

		data, err := getJsonFromUnescapedString(body)
		if err != nil {
			fmt.Println("400 Bad Request")
			res.WriteHeader(http.StatusBadRequest)
			return
		}

		logger.Post(tag, data)

		res.Header().Set("Cache-Control", "nocache")
	})

	http.ListenAndServe(":"+port, nil)
}

func getJsonFromUnescapedString(encodedJson string) (interface{}, error) {
	json_string, err := url.QueryUnescape(encodedJson)
	if err != nil {
		return nil, err
	}
	var data interface{}
	e := json.Unmarshal([]byte(json_string), &data)
	if e != nil {
		return nil, e
	}

	return data, nil
}

func tgif() []byte {
	dat, err := ioutil.ReadFile("1x1.gif")
	if err != nil {
		fmt.Println(err)
	}

	return dat
}

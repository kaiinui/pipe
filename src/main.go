package main

import (
	"encoding/json"
	"fmt"
	"github.com/t-k/fluent-logger-golang/fluent"
	"io/ioutil"
	"net/http"
	"net/url"
	"compress/gzip"
	"io"
	"bytes"
)

func main() {
	tgif := tgif()
	logger, err := fluent.New(fluent.Config{FluentPort: 24224, FluentHost: "localhost"})
	if err != nil {
		panic("Could not find fluentd.")
	}
	port := "8080"
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

	http.HandleFunc("/_ah/health", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("ok"))
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

var _tgif = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x72\xf7\x74\xb3\xb0\x4c\x64\x64\x60\x64\x68\x60\x00\x81\xff\xff\xff\x2b\xfe\x64\x61\x04\x31\x75\x40\x04\x48\x86\x81\x89\xd1\x85\xc1\x1a\x10\x00\x00\xff\xff\x78\x13\x95\x27\x2a\x00\x00\x00")

func tgif() []byte {
	data, _ := bindata_read(_tgif, "1x1.gif")
	return data
}

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

package transformer

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

type TestTransformer struct {
}

func NewTestTransformer() Transformer {
	return &TestTransformer{}
}

func (t *TestTransformer) TransformRequestBody(routeID string, req *http.Request) error {
	log.Printf("Transforming request body for route: %s", routeID)

	b, _ := ioutil.ReadAll(req.Body)
	body := string(b)

	if routeID == "rt1" {
		body += " transformed"
	}

	size := int64(len(body))

	req.Body = ioutil.NopCloser(bufio.NewReader(bytes.NewBufferString(body)))
	req.Header.Del("Content-Length")
	req.ContentLength = size

	log.Println("Done")

	return nil
}

func (t *TestTransformer) TransformResponseBody(routeID string, res *http.Response) error {
	if routeID == "rt1" {
		b, _ := ioutil.ReadAll(res.Body)
		body := string(b)

		log.Printf("Transforming response body for route: %s", routeID)

		body += " transformed"

		res.Body = ioutil.NopCloser(bufio.NewReader(bytes.NewBufferString(body)))

		size := int64(len(body))
		res.Header.Del("Content-Length")
		res.ContentLength = size

		log.Println("Done")
	}

	return nil
}

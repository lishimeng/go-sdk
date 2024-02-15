package rest

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Post(url string, resp any, headers ...Header) (code int, err error) {
	method := "POST"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	for _, h := range headers {
		req.Header.Add(h.Name, h.Value)
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		_ = res.Body.Close()
	}()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(body) <= 0 {
		return
	}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return
	}
	code = res.StatusCode
	return
}

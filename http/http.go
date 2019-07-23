package http

import (
	"bytes"
	"errors"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

func PostJson(url string, data []byte) ([]byte, error) {
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func Get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func HttpRebalance(urls []string, function func(url string) (interface{}, error)) (success bool, resp interface{}, errs []error) {
	if len(urls) == 0 {
		return false, nil, []error{errors.New("urls is empty")}
	}

	rand.Seed(time.Now().UnixNano())
	for _, index := range rand.Perm(len(urls)) {
		url := urls[index]
		resp, err := function(url)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		return true, resp, errs
	}
	return false, nil, errs
}

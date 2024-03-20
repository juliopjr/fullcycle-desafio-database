package service

import (
	"context"
	"errors"
	"io"
	"net/http"
	"time"
)

func NewRequester(url string, timeout time.Duration) *requester {
	chReqSuccess := make(chan *[]byte)
	chReqFails := make(chan error)
	return &requester{url, timeout, chReqSuccess, chReqFails}
}

type requester struct {
	url     string
	timeout time.Duration

	chReqSuccess chan *[]byte
	chReqFails   chan error
}

func (e *requester) Execute() (*[]byte, error) {
	ctxReq := context.Background()
	ctxReq, cancelReq := context.WithTimeout(ctxReq, e.timeout)
	defer cancelReq()

	go e.getData(ctxReq)

	select {
	case <-ctxReq.Done():
		return nil, errors.New("request timeout")
	case err := <-e.chReqFails:
		return nil, err
	case data := <-e.chReqSuccess:
		return data, nil
	}
}

func (e *requester) getData(ctxReq context.Context) {
	req, err := http.NewRequestWithContext(ctxReq, "GET", e.url, nil)
	if err != nil {
		e.chReqFails <- err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		e.chReqFails <- err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		e.chReqFails <- err
	}

	e.chReqSuccess <- &data
}

package util

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type ImgInfo struct {
	Data []byte
	Len  int64
	Name string
}

func GetImgReader(url string) (*ImgInfo, error) {
	now := time.Now()
	rsp, err := http.Get(url)
	defer rsp.Body.Close()
	data, err := ioutil.ReadAll(rsp.Body)
	return &ImgInfo{
		Data: data,
		Len:  int64(len(data)),
		Name: fmt.Sprintf("%d%d%d%d.png", now.Year(), now.Month(), now.Day(), now.Unix()),
	}, err
}

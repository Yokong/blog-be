package util

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
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
	if err != nil {
		return nil, err
	}

	defer rsp.Body.Close()
	data, err := ioutil.ReadAll(rsp.Body)
	return &ImgInfo{
		Data: data,
		Len:  int64(len(data)),
		Name: fmt.Sprintf("%d%d%d%d.png", now.Year(), now.Month(), now.Day(), now.Unix()),
	}, err
}

func GetImgsByString(s string) []string {
	reg := regexp.MustCompile(`\((http.*\.(png|jpg))\)`)
	ret := reg.FindAllStringSubmatch(s, len(s))
	l := make([]string, 0)
	for _, v := range ret {
		if len(v) > 1 {
			l = append(l, v[1])
		}
	}
	return l
}

package util

import (
	"testing"
)

func TestGetImgsByString(t *testing.T) {
	s := `
	[](http://img.test.com/1.png)
	[](http://img.test.com/2.png)
	[](http://img.test.com/3.jpg)
	`

	ret := map[string]int{
		"http://img.test.com/1.png": 1,
		"http://img.test.com/2.png": 2,
		"http://img.test.com/3.jpg": 3,
	}

	newRet := GetImgsByString(s)

	for _, v := range newRet {
		if _, ok := ret[v]; !ok {
			t.Errorf("未找到结果: %s", v)
		}
	}
}

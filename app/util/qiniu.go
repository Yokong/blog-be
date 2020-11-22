package util

import (
	"blog-be/app/config"
	"bytes"
	"context"
	"strings"

	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
)

func UploadImg(url string) (string, error) {
	c := config.GetConfig()
	if strings.Contains(url, c.QiNiu.Domain) {
		return url, nil
	}

	imgInfo, err := GetImgReader(url)
	if err != nil {
		return "", err
	}

	putPolicy := storage.PutPolicy{
		Scope: c.QiNiu.Bucket,
	}

	mac := qbox.NewMac(c.QiNiu.Ak, c.QiNiu.Sk)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{
		Zone:          &storage.ZoneHuanan,
		UseHTTPS:      false,
		UseCdnDomains: false,
	}

	uploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "yokoa image",
		},
	}

	err = uploader.Put(context.Background(), &ret, upToken, imgInfo.Name, bytes.NewReader(imgInfo.Data), imgInfo.Len, &putExtra)
	return c.QiNiu.Domain + ret.Key, err
}

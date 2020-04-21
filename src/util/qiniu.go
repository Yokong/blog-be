package util

import (
	"blog-be/src/config"
	"bytes"
	"context"

	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
)

func UploadImg(url string) (string, error) {
	imgInfo, err := GetImgReader(url)
	if err != nil {
		return "", err
	}

	putPolicy := storage.PutPolicy{
		Scope: config.QiNiuBucket,
	}

	mac := qbox.NewMac(config.QiNiuAK, config.QiNiuSK)
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
	return config.QiNiuDomain + ret.Key, err
}

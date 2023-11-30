package util

import (
	"auth-service/config"
	"commons/util"
	"context"
	"github.com/google/uuid"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"log"
	"mime/multipart"
)

// UploadToQiNiu 文件上传操作
// 封装上传图片到七牛云然后返回状态和图片的url
func UploadToQiNiu(file multipart.File, fileSize int64) (string, error) {

	var AccessKey = config.QinuConfig["AccessKey"]
	var SecretKey = config.QinuConfig["SecretKey"]
	var Bucket = config.QinuConfig["Bucket"]
	var ImgUrl = config.QinuConfig["Url"]

	putPolicy := storage.PutPolicy{
		Scope: Bucket,
	}
	mac := qbox.NewMac(AccessKey, SecretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{
		Zone:          &storage.ZoneHuadong,
		UseCdnDomains: false,
		UseHTTPS:      false,
	}

	putExtra := storage.PutExtra{}
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	key := GetRandomImgName()
	err := formUploader.Put(context.Background(), &ret, upToken, key, file, fileSize, &putExtra)
	if err != nil {
		return "上传图片失败", err
	}
	url := ImgUrl + "/" + ret.Key
	return url, nil
}

// GetRandomImgName 获取不会重复的文件名
func GetRandomImgName() string {
	uuid, err := uuid.NewUUID()
	if err != nil {
		log.Printf("uuid生成错误 err => %s", err)
	}
	path := "DataPulse/" + util.GetTodayString() + "-" + uuid.String()
	return path
}

package myziyue

import (
	"github.com/qiniu/api.v7/storage"
	"github.com/qiniu/api.v7/auth/qbox"
	"golang.org/x/net/context"
	"github.com/qiniu/x/log.v7"
	"strings"
	"os"
)

var (
	AccessKey, _      = GetOption("AccessKey", "qiniu")
	SecretKey, _      = GetOption("SecretKey", "qiniu")
	Bucket, _         = GetOption("Bucket", "qiniu")
	StorageZone, _    = GetOption("StorageZone", "qiniu")
	UseHTTPS, _       = GetOption("UseHTTPS", "qiniu")
	UseCdnDomains, _  = GetOption("UseHTTPS", "qiniu")
	ForceOverwrite, _ = GetOption("ForceOverwrite", "qiniu")
)

func UploadFile(filePath string, uploadFileName string) {
	putPolicy := storage.PutPolicy{
		Scope: Bucket,
	}
	mac := qbox.NewMac(AccessKey, SecretKey)
	cfg := storage.Config{}
	// 空间对应的机房
	storageZone, _ := GetOption("StorageZone", "qiniu")
	switch storageZone {
	case "Huadong":
		cfg.Zone = &storage.ZoneHuadong
	case "Huabei":
		cfg.Zone = &storage.ZoneHuabei
	case "Huanan":
		cfg.Zone = &storage.ZoneHuanan
	case "Beimei":
		cfg.Zone = &storage.ZoneBeimei
	}

	// 是否使用https域名
	cfg.UseHTTPS = UseHTTPS == "1"
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = UseCdnDomains == "1"
	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	// 可选配置
	//putExtra := storage.PutExtra{
	//	Params: map[string]string{
	//		"x:name": "github logo",
	//	},
	//}

	upToken := putPolicy.UploadToken(mac)
	err := formUploader.PutFile(context.Background(), &ret, upToken, uploadFileName, filePath, nil)
	if err != nil {
		log.Fatalf("'" + uploadFileName + "' upload file error", filePath, err)
		return
	}

	log.Infof("'" + uploadFileName + "' upload file success! ", ret.Key, ret.Hash)
	return
}

func DeleteFile(fileName string) {
	mac := qbox.NewMac(AccessKey, SecretKey)
	cfg := storage.Config{}
	// 空间对应的机房
	switch StorageZone {
	case "Huadong":
		cfg.Zone = &storage.ZoneHuadong
	case "Huabei":
		cfg.Zone = &storage.ZoneHuabei
	case "Huanan":
		cfg.Zone = &storage.ZoneHuanan
	case "Beimei":
		cfg.Zone = &storage.ZoneBeimei
	}

	// 是否使用https域名
	UseHTTPS, _ := GetOption("UseHTTPS", "qiniu")
	cfg.UseHTTPS = UseHTTPS == "1"
	// 上传是否使用CDN上传加速
	UseCdnDomains, _ := GetOption("UseHTTPS", "qiniu")
	cfg.UseCdnDomains = UseCdnDomains == "1"

	bucketManager := storage.NewBucketManager(mac, &cfg)

	err := bucketManager.Delete(Bucket, fileName)
	if err != nil {
		log.Fatalf("Delete file faild!", fileName, err)
		return
	}
	log.Infof("Delete file success! ", fileName)
	return
}

func MoveFile(srcName string, destName string) {
	mac := qbox.NewMac(AccessKey, SecretKey)
	cfg := storage.Config{}
	// 空间对应的机房
	switch StorageZone {
	case "Huadong":
		cfg.Zone = &storage.ZoneHuadong
	case "Huabei":
		cfg.Zone = &storage.ZoneHuabei
	case "Huanan":
		cfg.Zone = &storage.ZoneHuanan
	case "Beimei":
		cfg.Zone = &storage.ZoneBeimei
	}

	// 是否使用https域名
	UseHTTPS, _ := GetOption("UseHTTPS", "qiniu")
	cfg.UseHTTPS = UseHTTPS == "1"
	// 上传是否使用CDN上传加速
	UseCdnDomains, _ := GetOption("UseHTTPS", "qiniu")
	cfg.UseCdnDomains = UseCdnDomains == "1"

	bucketManager := storage.NewBucketManager(mac, &cfg)

	err := bucketManager.Move(Bucket, srcName, Bucket, destName, ForceOverwrite == "1")
	if err != nil {
		log.Fatalf("Move file faild!", srcName, "=>", destName, err)
		return
	}
	log.Infof("Move file success! ", srcName, "=>", destName, )
	return
}

func CopyFile(srcName string, destName string) {
	mac := qbox.NewMac(AccessKey, SecretKey)
	cfg := storage.Config{}
	// 空间对应的机房
	switch StorageZone {
	case "Huadong":
		cfg.Zone = &storage.ZoneHuadong
	case "Huabei":
		cfg.Zone = &storage.ZoneHuabei
	case "Huanan":
		cfg.Zone = &storage.ZoneHuanan
	case "Beimei":
		cfg.Zone = &storage.ZoneBeimei
	}

	// 是否使用https域名
	UseHTTPS, _ := GetOption("UseHTTPS", "qiniu")
	cfg.UseHTTPS = UseHTTPS == "1"
	// 上传是否使用CDN上传加速
	UseCdnDomains, _ := GetOption("UseHTTPS", "qiniu")
	cfg.UseCdnDomains = UseCdnDomains == "1"

	bucketManager := storage.NewBucketManager(mac, &cfg)

	err := bucketManager.Copy(Bucket, srcName, Bucket, destName, ForceOverwrite == "1")
	if err != nil {
		log.Fatalf("Copy file faild!", srcName, "=>", destName, err)
		return
	}
	log.Infof("Copy file success! ", srcName, "=>", destName)
	return
}
// 获取上传文件名
func GetFilePath(filePath string) string {
	prefixPath,err := GetOption("WatcherPath", "watcher")
	if err != nil {
		return ""
	}

	var prefixPaths []string
	prefixPaths = strings.Split(prefixPath, ";")
	for i:=0;i < len(prefixPaths); i++ {
		filePath = strings.Replace(filePath, prefixPaths[i], "", -1)
	}
	pathSeparator := string(os.PathSeparator)
	return strings.Replace(strings.Replace(filePath, pathSeparator + pathSeparator, pathSeparator, -1), "\\", "/", -1)
}

func Batch()  {
	
}
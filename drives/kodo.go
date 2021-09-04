package drives

import (
	ObjectStorage "github.com/qwxingzhe/go-object-storage"
)

type KodoDrives struct {
	ObjectStorageDrive *ObjectStorage.ObjectStorage
}

//func (receiver KodoDrives) init() *ObjectStorage.ObjectStorage {
//	return &ObjectStorage.ObjectStorage{
//		Drive: ObjectStorageDrives.Kodo{
//			AccessKey: "",
//			SecretKey: "",
//			Bucket:    "",
//		},
//		IsAutomaticProductionPath: true,
//		FilePathPrefix:            "test/",
//		IsAppendExt:               false,
//		BaseUrl:                   "http://qynr9haq9.hd-bkt.clouddn.com/",
//	}
//}

func (receiver KodoDrives) Save(imgUrl string) (string, error) {
	uploadFileInfo, err := receiver.ObjectStorageDrive.PutNetFile(imgUrl)
	return uploadFileInfo.Url, err
}

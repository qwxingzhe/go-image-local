package go_image_local

import (
	"fmt"
	"github.com/qwxingzhe/go-image-local/drives"
	ObjectStorage "github.com/qwxingzhe/go-object-storage"
	ObjectStorageDrives "github.com/qwxingzhe/go-object-storage/drives"
	"testing"
)

func TestUnit_ImageToLocal(t *testing.T) {

	fileDrives := drives.FileDrives{
		BasePath: "./storage/",
	}

	content := "<img src=\"https://img-home.csdnimg.cn/images/20210827090533.png\">XXX<img  src=\"https://img-blog.csdnimg.cn/606d5fec02674021be47ea0358417e46.jpg?x-oss-process=image/resize,m_fixed,h_200\">"
	content = ImageToLocal(fileDrives, content)
	fmt.Println("content |---------------------------------")
	fmt.Println(content)
}

func TestUnit_ImageToKodo(t *testing.T) {

	kodoDrives := drives.KodoDrives{
		ObjectStorageDrive: &ObjectStorage.ObjectStorage{
			Drive: ObjectStorageDrives.Kodo{
				AccessKey: kodoConfig.AccessKey,
				SecretKey: kodoConfig.SecretKey,
				Bucket:    kodoConfig.Bucket,
			},
			IsAutomaticProductionPath: true,
			FilePathPrefix:            "test2/",
			IsAppendExt:               false,
			BaseUrl:                   "http://qynr9haq9.hd-bkt.clouddn.com/",
		},
	}

	content := "<img src=\"https://img-home.csdnimg.cn/images/20210827090533.png\">XXX<img  src=\"https://img-blog.csdnimg.cn/606d5fec02674021be47ea0358417e46.jpg?x-oss-process=image/resize,m_fixed,h_200\">"
	content = ImageToLocal(kodoDrives, content)
	fmt.Println("content |---------------------------------")
	fmt.Println(content)
}

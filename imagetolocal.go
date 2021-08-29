package go_image_local

import (
	"bufio"
	"fmt"
	"github.com/google/uuid"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"regexp"
	"strings"
)

type StorageDrive interface {

}

type ImageLocal struct {

}

// 去除url参数
func clearImageUrl(imgUrl string) (string,error)  {
	u, err := url.Parse(imgUrl)
	if err!=nil {
		return "",err
	}
	return u.Scheme +"://"+ u.Host + u.Path,nil
}

func ImageToLocal(content string) string{

	var imgRE = regexp.MustCompile(`<img[^>]+\bsrc=["']([^"']+)["']`)
	imgs := imgRE.FindAllStringSubmatch(content, -1)
	imgList := make([]string, len(imgs))
	for i := range imgList {
		// 去除url参数
		//clearPath,err := clearImageUrl(imgs[i][1])
		//if err!=nil {
		//	continue
		//}

		// 将这些图片下载到本地
		newPath,err := DownLoadImage(imgs[i][1])
		if err==nil{
			// 替换文本中所有图片地址
			content = strings.Replace(content, imgs[i][1], newPath, 1)
		}
	}
	fmt.Println(imgList)
	return content
}

func UuidPath(imgUrl string) (string,error) {
	imgUrl,err := clearImageUrl(imgUrl)
	if err != nil {
		return "",err
	}
	uuidPath := uuid.New().String() + path.Ext(imgUrl)
	return uuidPath,nil
}

func DownLoadImage(imgUrl string) (string,error){
	imgPath := "./"
	//fileName := path.Base(imgUrl)
	fileName,err := UuidPath(imgUrl)
	if err!=nil{
		return "",err
	}

	res, err := http.Get(imgUrl)
	if err != nil {
		fmt.Println("A error occurred!")
		return imgUrl,nil
	}
	defer res.Body.Close()
	// 获得get请求响应的reader对象
	reader := bufio.NewReaderSize(res.Body, 32 * 1024)

	filePath := imgPath + fileName
	file, err := os.Create(filePath)
	if err != nil {
		return "",err
	}
	// 获得文件的writer对象
	writer := bufio.NewWriter(file)

	written, _ := io.Copy(writer, reader)
	fmt.Printf("Total length: %d", written)

	return filePath,nil
}
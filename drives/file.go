package drives

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
)

type FileDrives struct {
	BasePath string
}

//判断文件或文件夹是否存在
func isExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		if os.IsNotExist(err) {
			return false
		}
		fmt.Println(err)
		return false
	}
	return true
}

func (receiver FileDrives) Save(imgUrl string) (string, error) {

	fileName, err := UuidPath(imgUrl)
	if err != nil {
		return "", err
	}

	res, err := http.Get(imgUrl)
	if err != nil {
		return imgUrl, err
	}
	defer res.Body.Close()
	// 获得get请求响应的reader对象
	reader := bufio.NewReaderSize(res.Body, 32*1024)

	if isExist(receiver.BasePath) == false {
		errFileCreate := os.MkdirAll(receiver.BasePath, os.ModePerm)
		if errFileCreate != nil {
			panic(errFileCreate)
		}
	}

	filePath := receiver.BasePath + fileName
	file, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	// 获得文件的writer对象
	writer := bufio.NewWriter(file)

	written, _ := io.Copy(writer, reader)
	fmt.Printf("Total length: %d", written)

	return filePath, nil
}

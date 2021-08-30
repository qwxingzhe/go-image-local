package go_image_local

import (
	"fmt"
	"github.com/qwxingzhe/go-image-local/drives"
	"regexp"
	"strings"
)

func ImageToLocal(imageSaveDrive drives.ImageSaveDrive, content string) string {

	var imgRE = regexp.MustCompile(`<img[^>]+\bsrc=["']([^"']+)["']`)
	imgs := imgRE.FindAllStringSubmatch(content, -1)
	imgList := make([]string, len(imgs))

	for i := range imgList {
		// 将这些图片下载到本地
		newPath, err := imageSaveDrive.Save(imgs[i][1])
		if err == nil {
			// 替换文本中所有图片地址
			content = strings.Replace(content, imgs[i][1], newPath, 1)
		}
	}
	fmt.Println(imgList)
	return content
}

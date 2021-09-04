package go_image_local

import (
	"fmt"
	"github.com/qwxingzhe/go-image-local/drives"
	"regexp"
	"strings"
)

func ImageToLocal(imageSaveDrive drives.ImageSaveDrive, content string) string {
	// 将文本中的图片地址提取到数组中
	var imgRE = regexp.MustCompile(`<img[^>]+\bsrc=["']([^"']+)["']`)
	imgs := imgRE.FindAllStringSubmatch(content, -1)
	imgList := make([]string, len(imgs))

	// 遍历图片数组分别进行本地化，并用本地化后的地址替换文本中原有的地址
	for i := range imgList {
		// 使用本地化引擎对图片做本地化处理
		newPath, err := imageSaveDrive.Save(imgs[i][1])
		// 成功，则替换文本中对应的图片地址
		if err == nil {
			content = strings.Replace(content, imgs[i][1], newPath, 1)
		}
	}
	fmt.Println(imgList)
	// 将新的文本内容返回
	return content
}

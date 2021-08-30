package go_image_local

import (
	"fmt"
	"github.com/qwxingzhe/go-image-local/drives"
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

package drives

import (
	"github.com/google/uuid"
	"net/url"
	"path"
)

// ImageSaveDrive 图片存储驱动接口
type ImageSaveDrive interface {
	Save(imgUrl string) (string, error)
}

// ClearImageUrl 去除url参数
func ClearImageUrl(imgUrl string) (string, error) {
	u, err := url.Parse(imgUrl)
	if err != nil {
		return "", err
	}
	return u.Scheme + "://" + u.Host + u.Path, nil
}

// UuidPath 获取uuid图片路径
func UuidPath(imgUrl string) (string, error) {
	imgUrl, err := ClearImageUrl(imgUrl)
	if err != nil {
		return "", err
	}
	uuidPath := uuid.New().String() + path.Ext(imgUrl)
	return uuidPath, nil
}

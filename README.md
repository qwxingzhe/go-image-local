# go-image-local

[![Godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/qwxingzhe/go-image-local)
[![Build Status](https://travis-ci.org/qwxingzhe/go-image-local.svg)](https://travis-ci.org/qwxingzhe/go-image-local)
[![GitHub release](http://img.shields.io/github/release/qwxingzhe/go-image-local.svg?style=flat-square)](release)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](license)
[![Commit activity](https://img.shields.io/github/commit-activity/m/qwxingzhe/go-image-local)](https://github.com/qwxingzhe/go-image-local/graphs/commit-activity)
[![Average time to resolve an issue](http://isitmaintained.com/badge/resolution/qwxingzhe/go-image-local.svg)](http://isitmaintained.com/project/qwxingzhe/go-image-local "Average time to resolve an issue")
[![Percentage of issues still open](http://isitmaintained.com/badge/open/qwxingzhe/go-image-local.svg)](http://isitmaintained.com/project/qwxingzhe/go-image-local "Percentage of issues still open")



<p align="center">go-image-local 是一个用go开发的将指定文本中的图片本地化的包。其可配置将图片本地化或者上传到对象存储（对象存储使用的是<a href="https://github.com/qwxingzhe/go-object-storage">qwxingzhe/go-object-storage</a>包），可轻易的集成到go项目中，欢迎使用。</p>

# 安装

```shell
$ go get -u github.com/qwxingzhe/go-image-local
```

# 使用指南

配置本地存储引擎

~~~
localDrives := drives.FileDrives{
	BasePath: "./storage/",
}


~~~

配置对象存储引擎（七牛云kodo）

> 参数配置可参看 [qwxingzhe/go-object-storage](https://github.com/qwxingzhe/go-object-storage)

~~~
localDrives := drives.KodoDrives{
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
~~~

调用

~~~
content := "<img src=\"https://img-home.csdnimg.cn/images/20210827090533.png\">XXX<img  src=\"https://img-blog.csdnimg.cn/606d5fec02674021be47ea0358417e46.jpg?x-oss-process=image/resize,m_fixed,h_200\">"
content = ImageToLocal(localDrives, content)
fmt.Println(content)
~~~

# Enjoy it! :heart:

# License

MIT


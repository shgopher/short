# short
[![Goproxy.cn](https://goproxy.cn/stats/github.com/shgoppher/short/badges/download-count.svg)](https://goproxy.cn)

Use murmur hash function.

URL shortening service.

## Usage
[example](./example/example.go)
```go
package main

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/googege/short"
	"os"
)

var (
	path    = "https://t.cn/"
	longURL = "https://github.com/googege/GOFamily/blob/master/%E5%9F%BA%E7%A1%80%E7%9F%A5%E8%AF%86/%E7%AE%97%E6%B3%95/%E7%AE%97%E6%B3%95%E9%A2%98/leetcode/1.md"
)

func main() {
	db := short.NewMapDB()
	// add db engine to short.
	s := short.NewShort(db)
	//
	shortURL, err := s.ShortAdd(longURL)
	if err != nil {
		glog.Error(err)
	} else {
		fmt.Println(shortURL)
	}
	//
	longURL, err = s.ShortFind(path + shortURL)
	// if http
	//http.Redirect(nil,nil,longURL,302
	//)
	if err != nil {
		glog.Error(err)
	} else {
		fmt.Println("longURL:", longURL)
	}
	//
	shortURL, err = s.ShortFind("a")
	if err != nil {
		glog.Error(err)
	} else {
		fmt.Println("short: ", shortURL)
	}
	//
	file, err := os.Getwd()
	if err != nil {
		glog.Error(err)
	}
	if err = s.SetQR(path, 256, file+"/text.png"); err != nil {
		glog.Error(err)
	}
}


```

## HERE
|items|description|
|:---:|:---:|
|aliyun vps|[aliyun vps](https://www.aliyun.com/minisite/goods?userCode=ol87kpmz)，China's largest cloud computing manufacturer, the world's top 5 cloud computing service providers|
|VPS|[vps](https://app.cloudcone.com/?ref=2525) fast and cheap just 2-3.75$/mo|
|WeChat public account 微信公众号|![p](https://raw.githubusercontent.com/googege/GOFamily/master/joinUsW.jpg)|
|My bilibili b站|[b](https://space.bilibili.com/23170151)|

/*
 * @Author: shgopher shgopher@gmail.com
 * @Date: 2023-05-26 04:15:12
 * @LastEditors: shgopher shgopher@gmail.com
 * @LastEditTime: 2023-05-26 04:33:30
 * @FilePath: /short/example/example.go
 * @Description:
 *
 * Copyright (c) 2023 by shgopher, All Rights Reserved.
 */
package main

import (
	"fmt"
	"os"

	"github.com/golang/glog"
	"github.com/shgopher/short"
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

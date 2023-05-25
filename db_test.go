/*
 * @Author: shgopher shgopher@gmail.com
 * @Date: 2023-05-26 04:15:12
 * @LastEditors: shgopher shgopher@gmail.com
 * @LastEditTime: 2023-05-26 04:20:48
 * @FilePath: /short/db_test.go
 * @Description:
 *
 * Copyright (c) 2023 by shgopher, All Rights Reserved.
 */
package short

import "testing"

func TestMapDB(t *testing.T) {
	n := &Node{
		LongValue:  "https://github.com/googege/12",
		ShortValue: "short",
	}
	//
	db := NewMapDB()
	//ad
	_, err := db.Add(n)
	if err != nil {
		t.Error(err)
	}
	// find
	shortURL, err := db.Find("https://github.com/googege/12")
	if err != nil {
		t.Error(err)
	}
	t.Log(shortURL)
	// change
	err = db.Change(n, "short1")
	if err != nil {
		t.Error(err)
	}
	// find
	shortURL, err = db.Find("https://github.com/googege/12")
	if err != nil {
		t.Error(err)
	} else {
		t.Log(shortURL)
	}
	// delete
	db.Delete("https://github.com/googege/12")
	// find
	shortURL, err = db.Find("https://github.com/googege/12")
	if err != nil {
		t.Error(err)
	} else {
		t.Log(shortURL)
	}
}

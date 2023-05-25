/*
 * @Author: shgopher shgopher@gmail.com
 * @Date: 2023-05-26 04:22:10
 * @LastEditors: shgopher shgopher@gmail.com
 * @LastEditTime: 2023-05-26 04:22:18
 * @FilePath: /short/mem.go
 * @Description:
 *
 * Copyright (c) 2023 by shgopher, All Rights Reserved.
 */
package short

import "fmt"

type MapDB struct {
	Value map[string]string
}

func NewMapDB() *MapDB {
	return &MapDB{
		Value: make(map[string]string),
	}
}
func (m *MapDB) Add(v *Node) (string, error) {
	if _, ok := m.Value[v.ShortValue]; ok {
		return "", fmt.Errorf("this long URL is existed. ")
	}
	m.Value[v.ShortValue] = v.LongValue
	return v.ShortValue, nil
}
func (m *MapDB) Delete(shortURL string) {
	delete(m.Value, shortURL)
}

// change the longurl's shorturl to positionShortURL.
func (m *MapDB) Change(n *Node, positionShortURL string) error {
	shortURL := n.ShortValue

	if _, ok := m.Value[shortURL]; !ok {
		return fmt.Errorf("this is no long ulr")
	} else {
		m.Value[positionShortURL] = n.LongValue
		n.ShortValue = positionShortURL
	}
	return nil
}

// find the short url
func (m *MapDB) Find(shortURL string) (string, error) {
	if val, ok := m.Value[shortURL]; !ok {
		return "", fmt.Errorf("this is no long url")
	} else {
		return val, nil
	}
}

package short

import "testing"

func TestMapDB(t *testing.T) {
	n := &Node{
		LongValue: "https://github.com/googege/12",
		ShortValue:"short",
	}
	//
	db := NewMapDB()
	//ad
	err := db.Add(n)
	if err != nil {
		t.Error(err)
	}
	// find
	shortURL,err := db.Find("https://github.com/googege/12")
	if err != nil {
		t.Error(err)
	}
	t.Log(shortURL)
	// change
	err = db.Change(n,"short1")
	if err != nil {
		t.Error(err)
	}
	// find
	shortURL,err = db.Find("https://github.com/googege/12")
	if err != nil {
		t.Error(err)
	}else {
		t.Log(shortURL)
	}
	// delete
	db.Delete("https://github.com/googege/12")
	// find
	shortURL,err = db.Find("https://github.com/googege/12")
	if err != nil {
		t.Error(err)
	}else {
		t.Log(shortURL)
	}
}
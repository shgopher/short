package short

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/glog"
	"sync"
)



type MySqlDB struct {
	db *sql.DB
}
type MySqlDBValue struct {
	ShortURL string `json:"short_url"`
	LongURL  string `json:"long_url"`
}

func NewMySqlDB(dataSourName string) *MySqlDB {
	once := new(sync.Once)
	md := new(MySqlDB)
	once.Do(func() {
		var err error
		md.db, err = sql.Open("mysql", dataSourName)
		if err != nil {
			glog.Error(err)
		}
		err = md.db.Ping()
		if err != nil {
			glog.Error(err)
		}
	})
	return md
}

// add
func (m *MySqlDB) Add(v *Node) (shortURL string, err error) {
	// short-url long-URL
	st, err := m.db.Prepare("INSERT INTO short (short_url,long_url) VALUES (?,?)")
	if err != nil {
		return "", err
	}
	// into data.
	_, err = st.Exec(v.ShortValue, v.LongValue)
	if err != nil {
		return "", err
	}
	// close the connection
	st.Close()
	return v.ShortValue, nil
}
func (m *MySqlDB) Delete(shortURL string) {
	st, err := m.db.Prepare("DELETE FROM short WHEN short_url=?")
	if err != nil {
		glog.Error(err)
	}
	_, err = st.Exec(shortURL)
	if err != nil {
		glog.Error(err)
	}
	st.Close()
}

func (m *MySqlDB) Change(vi *Node, r string) error {
	st, err := m.db.Prepare("UPDATE short SET short_url=?WHERE long_url=?")
	if err != nil {
		return err
	}
	_, err = st.Exec(r, vi.LongValue)
	if err != nil {
		return err
	}
	st.Close()
	return nil
}

func (m *MySqlDB) Find(shortURL string) (string, error) {
	long := ""
	row:= m.db.QueryRow("SELECT long_url FROM short WHERE short_url=?",shortURL)
	row.Scan(&long)
	return long, nil
}


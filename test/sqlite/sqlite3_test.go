package sqlite

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"testing"
)

func Test_Sqlite3(t *testing.T) {
	t.Log("Start test")

	db, err := sql.Open("sqlite3", "data.db")
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	sqlTable := `
			CREATE TABLE IF NOT EXISTS userinfo(
					id INTEGER PRIMARY KEY AUTOINCREMENT,
					user_id text,
					user_name text
			);`
	_, err = db.Exec(sqlTable)

	t.Error(err)

	// _, err = db.Exec("CREATE TABLE log(id INT, content VARCHAR(1024))")
	// if err == errors.New("table log already exists") {
	// 	t.Error(err)
	// }
}

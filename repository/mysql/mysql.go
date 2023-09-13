package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	_interface "github.com/muhammadarash1997/movies/repository/interface"
)

type dbReadWriter struct {
	db *sql.DB
}

func NewDBReadWriter(url string, schema string, user string, password string) (_interface.ReadWriter, error) {
	schemaURL := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", user, password, url, schema)
	db, err := sql.Open("mysql", schemaURL)
	if err != nil {
		return nil, err
	}

	return &dbReadWriter{db: db}, nil
}

// Close is used for closing the sql connection
func (rw *dbReadWriter) Close() error {
	if rw.db != nil {
		if err := rw.db.Close(); err != nil {
			return err
		}
		rw.db = nil
	}

	return nil
}

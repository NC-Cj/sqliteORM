package session

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteAdapter interface {
	Open(dbFilePath string) error
	Close() error
}

type Session struct {
	db *sql.DB
}

func NewSession(dbFilePath string) *Session {
	s := Session{}
	s.Open(dbFilePath)
	return &s
}

func (s *Session) Open(dbFilePath string) error {
	var err error
	s.db, err = sql.Open("sqlite3", dbFilePath)
	if err != nil {
		panic(err.Error())
	}
}

func (s *Session) Close() error {
	if s.db != nil {
		err := s.db.Close()
		if err != nil {
			s.db = nil
			return err
		}
	}
	fmt.Println("Database connection has been successfully closed.")
	return nil
}

package dbdriver

import (
	"testing"
)

func TestSqlite(t *testing.T) {
	s := NewSqlite("./test.db")
	if err := s.Open(nil); err != nil {
		t.Fatal(err)
	}

	s.Close()
}
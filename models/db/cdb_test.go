package db

import (
	"testing"
)

func TestCDB(t *testing.T) {
	db := CouchDB{}
	err := db.Init()
	if err != nil {
		t.Error(err)
	}

	db.enroll([]byte("test"))
}

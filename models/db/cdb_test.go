package db

import (
	"testing"
)

func TestCDB(t *testing.T) {
	db := CouchDB{}
	err := db.Init("dfs")
	if err != nil {
		t.Error(err)
	}

	//db.set()
}

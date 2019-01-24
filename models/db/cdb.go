package db

import (
	"errors"
	"github.com/leesper/couchdb-golang"
	log "github.com/sirupsen/logrus"
)

type CouchDB struct {
	db *couchdb.Database
}

func (this *CouchDB) Init(urlStr string) error {
	log.WithFields(log.Fields{"urlStr": urlStr}).Info("cdb Init")

	db, err := couchdb.NewDatabase(urlStr)
	if err != nil {
		log.Error("the cdb Init error %v", err)
		return err
	}

	this.db	= db
	return nil
}

func (this *CouchDB) set(value []byte) (key string, err error) {
	log.Info("cdb set")

	if this.db == nil {
		log.WithError(errors.New("the db is nil"))
		return "", errors.New("please call Init first")
	}

	if value == nil {
		err = errors.New("the value can not is nil")
		log.WithError(err)
		return "", err
	}

	doc := make(map[string]interface{})
	// TODO:

	id, rev, err := this.db.Save(doc, nil)
	if err != nil {
		log.Error("couch db save error: %v", err)
		return "", err
	}
	log.WithFields(log.Fields{"id": id, "rev": rev}).Info("couch db save document")

	return id, nil
}

func (this *CouchDB) get(key string) (value []byte, err error) {
	log.WithField("key", key).Info("cdb get")

	if this.db == nil {
		err = errors.New("the db is nil")
		log.WithError(err)
		return nil, err
	}

	if key == "" {
		err = errors.New("the key can not is empty")
		log.WithError(err)
		return nil, err
	}

	return nil, nil
}



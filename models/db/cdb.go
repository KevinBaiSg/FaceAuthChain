package db

import (
	"encoding/base64"
	"errors"
	"github.com/leesper/couchdb-golang"
	log "github.com/sirupsen/logrus"
)

type CouchDB struct {
	db *couchdb.Database
}

func (c *CouchDB) Init() error {
	dbName := "faceAuthDescriptor"
	address := "http://localhost:5984/" + dbName

	var db *couchdb.Database
	var err error

	log.WithFields(log.Fields{"address": address}).Info("cdb Init")

	server := couchdb.Server{}

	if server.Contains(dbName) {
		log.Info("the couchdb server: %s is exist", dbName)
		db, err = couchdb.NewDatabase(address)
		if err != nil {
			log.Error("the cdb Init error", err.Error())
			return err
		}
	} else {
		log.Info("the couchdb server: %s is not exist, and create it", dbName)
		db, err = server.Create(dbName)
		if err != nil {
			log.Error("create couchdb: %s failed, error: ", dbName, err.Error())
			return err
		}
	}

	c.db = db
	return nil
}

func (c *CouchDB) enroll(descriptor []byte) (id string, err error) {
	log.Info("cdb set")

	if c.db == nil {
		log.WithError(errors.New("the db is nil"))
		return "", errors.New("please call Init first")
	}

	if descriptor == nil {
		err = errors.New("the value can not is nil")
		log.WithError(err)
		return "", err
	}

	doc := make(map[string]interface{})
	doc["descriptor"] = base64.StdEncoding.EncodeToString(descriptor)

	id, rev, err := c.db.Save(doc, nil)
	if err != nil {
		log.Error("couch db save error: ", err)
		return "", err
	}
	log.WithFields(log.Fields{"id": id, "rev": rev}).Info("couch db save document")
	return id, nil
}

func (c *CouchDB) enrollBy(id string, descriptor []byte) error {
	panic("implement me")
}

func (c *CouchDB) query(id string) (descriptor []byte, err error) {
	panic("implement me")
	//log.WithField("id", id).Info("cdb query")
	//
	//if c.db == nil {
	//	err = errors.New("the db is nil")
	//	log.WithError(err)
	//	return nil, err
	//}
	//
	//if id == "" {
	//	err = errors.New("the id can not is empty")
	//	log.WithError(err)
	//	return nil, err
	//}
	//
	//return nil, nil
}



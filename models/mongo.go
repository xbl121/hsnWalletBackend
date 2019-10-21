package models

import (
	"go.uber.org/zap"
	"gopkg.in/mgo.v2"
	"log"
	"time"

	"github.com/wongyinlong/walletBackend/conf"
)

var Session *mgo.Session

func init() {
	config := conf.NewConfig()
	if config.Db.DbString == "" {
		log.Fatal("dbString is empty")
		return
	}
	dialInfo, _ := mgo.ParseURL(config.Db.DbString)
	dialInfo.PoolLimit = 50
	dialInfo.Timeout = time.Minute * 1
	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		log.Fatal("creat db connection failed, check db state", zap.String("e", err.Error()), zap.String("url", config.Db.DbString))
		return
	}
	Session = session
}

func connect(db, collections string) (*mgo.Session, *mgo.Collection) {
	s := Session.Copy()
	c := s.DB(db).C(collections)
	return s, c
}

func Insert(db, collection string, docs ...interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Insert(docs...)
}

func FindAll(db string, collection string, query map[string]interface{}, result interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Find(query).One(result)
}
func FindOne(db, collection string, query map[string]interface{}, selector, result interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Find(query).Select(selector).All(result)
}

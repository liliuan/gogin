package lib

import (
	"gogin/config"
	"gopkg.in/mgo.v2"
)

var Mgodb *mgo.Database

type Mgoer interface {
	FindOne(query, selector map[string]interface{}) (map[string]interface{}, error)
	FindAll(query, selector map[string]interface{}) ([]map[string]interface{}, error)
	Upsert(query, update interface{}) (*mgo.ChangeInfo, error)
	Update(query, update interface{}) error
	Remove(query map[string]interface{}) error
	Insert(docs interface{}) error
}

func initMongoDB() {
	cfg := config.Cfg
	session, err := mgo.Dial(cfg.MgoUrl)

	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Eventual, true)

	Mgodb = session.DB(cfg.MgoName)

}

func init() {

	initMongoDB()
}

package model

import (
	"gogin/lib"
	"gopkg.in/mgo.v2"
)

type Model struct {
	CollName string
	Limit    int
	Page     int
}

func (m Model) FindOne(query, selector map[string]interface{}) (map[string]interface{}, error) {

	coll := lib.Database.C(m.CollName)
	result := make([]map[string]interface{}, 0)
	err := coll.Find(query).Select(selector).All(&result)

	if err != nil {
		return nil, err
	}

	if len(result) > 0 {
		return result[0], nil
	}

	return nil, nil
}

func (m Model) FindAll(query, selector map[string]interface{}) ([]map[string]interface{}, error) {

	coll := lib.Database.C(m.CollName)
	result := make([]map[string]interface{}, 0)
	skip := (m.Page - 1) * m.Limit
	err := coll.Find(query).Select(selector).Skip(skip).All(&result)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (m Model) Upsert(query, update interface{}) (*mgo.ChangeInfo, error) {

	coll := lib.Database.C(m.CollName)

	info, err := coll.Upsert(query, update)

	if err != nil {
		return nil, err
	}

	return info, nil
}

func (m Model) Update(query, update interface{}) error {

	coll := lib.Database.C(m.CollName)

	err := coll.Update(query, update)

	return err
}

func (m Model) Remove(query map[string]interface{}) error {

	coll := lib.Database.C(m.CollName)

	err := coll.Remove(query)

	return err
}

func (m Model) Insert(docs interface{}) error {

	coll := lib.Database.C(m.CollName)

	err := coll.Insert(docs)

	return err
}

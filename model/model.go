package model

import "gopkg.in/mgo.v2"

type Model struct {
}

func (m Model) FindOne(query, selector map[string]interface{}) (map[string]interface{}, error) {
	return nil, nil
}

func (m Model) FindAll(query, selector map[string]interface{}) ([]map[string]interface{}, error) {
	return nil, nil
}

func (m Model) Upsert(query, update interface{}) (*mgo.ChangeInfo, error) {
	return nil, nil
}

func (m Model) Update(query, update interface{}) error {
	return nil
}

func (m Model) Remove(query map[string]interface{}) error {
	return nil
}

func (m Model) Insert(docs interface{}) error {
	return nil
}

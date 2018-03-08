package dbservices

import (
    // "log"
	"go-select/entity"
	"go-select/config"
    "gopkg.in/mgo.v2"

)

func UpdateSinglesById(s []entity.SingleRes) error {
	session, err := mgo.Dial(config.DBurl)
    if err != nil {
        return err
    }
	defer session.Close()
	
	c := session.DB("Select").C("Single")
	for i := range s {
		selector := s[i].Id
		err := c.UpdateId(selector,s[i])
		if err != nil {
			return err
		}
	}
	return nil
}


func AddSingleExam(ids []string) error {
	session, err := mgo.Dial(config.DBurl)
    if err != nil {
        return err
    }
	defer session.Close()
	
	c := session.DB("Select").C("Single")
	for i := range ids {
		
	}
}
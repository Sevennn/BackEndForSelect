package dbservices

import (
	// "log"
	"go-select/entity"
	"go-select/config"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func UpdateSingleById(s entity.SingleReq) error {
	session, err := mgo.Dial(config.DBurl)
    if err != nil {
        return err
    }
	defer session.Close()
	c := session.DB("Select").C("Problems")

	err = c.UpdateId(bson.ObjectIdHex(s.Id),s)
	if err != nil {
		return err
	}

	return nil
}


func UpdateMultiById(s entity.MultiReq) error {
	session, err := mgo.Dial(config.DBurl)
    if err != nil {
        return err
    }
	defer session.Close()
	c := session.DB("Select").C("Problems")

	err = c.UpdateId(bson.ObjectIdHex(s.Id),s)
	if err != nil {
		return err
	}

	return nil
}

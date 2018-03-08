package dbservices
import (
    // "log"
	"go-select/entity"
	"go-select/config"
    "gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2/bson"
)


func InsertSingles(s []entity.Single) (error) {
	session, err := mgo.Dial(config.DBurl)
	if err != nil {
		return err;
	}
	defer session.Close()
	var data []interface{}

	for i := range s {
		data = append(data, s[i]);
	}

	c := session.DB("Select").C("Problems")
	err = c.Insert(data...)

	if err != nil {
		return err;
	}

	return nil
}

func InsertMultis(s []entity.Multiple) (error) {
	session, err := mgo.Dial(config.DBurl)
	if err != nil {
		return err;
	}
	defer session.Close()

	var data []interface{}
	
	for i := range s {
		data = append(data, s[i]);
	}

	c := session.DB("Select").C("Multi")
	err = c.Insert(data...)

	if err != nil {
		return err;
	}

	return nil
}


func AddSingleExam(ids []string) (error,string) {
	session, err := mgo.Dial(config.DBurl)
    if err != nil {
        return err
    }
	defer session.Close()
	id := bson.NewObjectId()
	c := session.DB("Select").C("Exam")
	err := c.insert(bson.M{"_id":id, "prolist": ids})
	if err != nil {
		return err, nil
	}
	return nil, id
}
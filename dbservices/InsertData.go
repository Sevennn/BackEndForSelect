package dbservices
import (
    // "log"
	"go-select/entity"
	"go-select/config"
    "gopkg.in/mgo.v2"
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

	c := session.DB("Select").C("Problems")
	err = c.Insert(data...)

	if err != nil {
		return err;
	}

	return nil
}


func AddExam(exam entity.Exam) (error) {
	session, err := mgo.Dial(config.DBurl)
    if err != nil {
		return err
    }
	defer session.Close()
	c := session.DB("Select").C("Exam")
	err = c.Insert(exam)
	if err != nil {
		return err
	}
	return nil
}
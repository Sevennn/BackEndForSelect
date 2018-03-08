package dbservices
import (
    // "log"
	"go-select/entity"
	"go-select/config"
    "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"fmt"

)



func FindAllSingle() (error,[]entity.SingleRes) {
    session, err := mgo.Dial(config.DBurl)
    if err != nil {
        return err,nil
    }
	defer session.Close()
	
	c := session.DB("Select").C("Single")

	var dbresult []entity.SingleRes

	err = c.Find(bson.M{}).All(&dbresult)
	if err != nil {
		return err,nil
	}
	// result,err := json.MarshalIndent(dbresult,""," ")

	// if err != nil {
	// 	return err,nil
	// }

	return nil,dbresult
}


func FindAllMulti() (error,[]entity.MultipleRes) {
    session, err := mgo.Dial(config.DBurl)
    if err != nil {
        return err,nil
    }
	defer session.Close()
	
	c := session.DB("Select").C("Multi")

	var dbresult []entity.MultipleRes

	err = c.Find(bson.M{}).All(&dbresult)

	if err != nil {
		return err,nil
	}
	// result,err := json.MarshalIndent(dbresult,""," ")

	// if err != nil {
	// 	return err,nil
	// }
	return nil,dbresult
}


func FindSinglesByIds(ids []string) (error,[]entity.SingleRes) {
    session, err := mgo.Dial(config.DBurl)
    if err != nil {
        return err,nil
    }
	defer session.Close()
	var s entity.SingleRes
	c := session.DB("Select").C("Single")
	var dbresult []entity.SingleRes
	for i := range ids {
		objId := bson.ObjectIdHex(ids[i])

		err := c.Find(bson.M{`_id`:objId}).One(&s)
		fmt.Println(s)
		if err != nil {
			return err, nil
		}
		dbresult = append(dbresult,s)
	}
	return nil,dbresult
}
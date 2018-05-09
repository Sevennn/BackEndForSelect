package dbservices
import (
    // "log"
	"go-select/entity"
	"go-select/config"
    "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"fmt"

)

func FindUserByName(uname string) (error, *entity.User) {
	session, err := mgo.Dial(config.DBurl)
    if err != nil {
		return err,nil
    }
	defer session.Close()
	
	c := session.DB("Select").C("User")
	
	var dbresult entity.User

	err = c.Find(bson.M{"username":uname}).One(&dbresult)
	if err != nil {
		return err,nil
	}
	return nil,&dbresult
}

func FindAllSingle() (error,[]entity.SingleRes) {
    session, err := mgo.Dial(config.DBurl)
    if err != nil {
        return err,nil
    }
	defer session.Close()
	
	c := session.DB("Select").C("Problems")

	var dbresult []entity.SingleRes

	err = c.Find(bson.M{"answertype":"single"}).All(&dbresult)
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
	
	c := session.DB("Select").C("Problems")

	var dbresult []entity.MultipleRes

	err = c.Find(bson.M{"answertype":"multi"}).All(&dbresult)

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
	c := session.DB("Select").C("Problems")
	var dbresult []entity.SingleRes
	for i := range ids {
		objId := bson.ObjectIdHex(ids[i])

		err := c.Find(bson.M{`_id`:objId,`answertype`:`single`}).One(&s)
		fmt.Println(s)
		if err == nil {
			dbresult = append(dbresult,s)
		}

	}
	return nil,dbresult
}

func FindMultisByIds(ids []string) (error,[]entity.MultipleRes) {
    session, err := mgo.Dial(config.DBurl)
    if err != nil {
        return err,nil
    }
	defer session.Close()
	var s entity.MultipleRes
	c := session.DB("Select").C("Problems")
	var dbresult []entity.MultipleRes
	for i := range ids {
		objId := bson.ObjectIdHex(ids[i])

		err := c.Find(bson.M{`_id`:objId, `answertype`:`multi`}).One(&s)
		fmt.Println(s)
		if err == nil {
			dbresult = append(dbresult,s)
		}


	}
	return nil,dbresult
}


func FindAllExams() (error,[]entity.ExamListRes) {
    session, err := mgo.Dial(config.DBurl)
    if err != nil {
        return err,nil
    }
	defer session.Close()
	c := session.DB("Select").C("Exam")
	var dbresult []entity.ExamListRes
	err = c.Find(bson.M{}).Select(bson.M{"title":true, "_id":true}).All(&dbresult)

	if err != nil {
		return err, nil
	}
	return nil,dbresult
}

func GetExamById(id string) (error, entity.ExamPage) {
	var res entity.ExamPage
    session, err := mgo.Dial(config.DBurl)
    if err != nil {
        return err,res
    }
	defer session.Close()
	c := session.DB("Select").C("Exam")
	var dbresult entity.ExamRes;
	err = c.FindId(bson.ObjectIdHex(id)).One(&dbresult)
	if err != nil {
		return err,res
	}
	err,sresult := FindSinglesByIds(dbresult.Problems)
	if err != nil {
		return err,res
	}
	err,mresult := FindMultisByIds(dbresult.Problems)
	if err != nil {
		return err,res
	}
	res.Single = sresult
	res.Multi = mresult
	res.Title = dbresult.Title;
	return nil, res
}

func GetBelongs(id string) (error,[]entity.Exam) {
	session, err := mgo.Dial(config.DBurl)
    if err != nil {
        return err,nil
    }
	defer session.Close()
	c := session.DB("Select").C("Exam")
	var dbresult []entity.Exam
	err = c.Find(bson.M{`problems`: id}).All(&dbresult)
	if err != nil {
		return err,nil
	}
	return nil, dbresult
}
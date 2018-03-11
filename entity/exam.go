package entity
import "gopkg.in/mgo.v2/bson"


type Exam struct {
	Title string `json:"title"`
	Problems []string `json:"problems"`
}

type ExamRes struct {
	Title string `json:"title"`
	Problems []string `json:"problems"`
	Id bson.ObjectId `json:"id" bson:"_id"`
}

type ExamListRes struct {
	Title string `json:"title"`
	Id bson.ObjectId `json:"id" bson:"_id"`
}

type ExamPage struct {
	Single []SingleRes `json:"single"`
	Multi []MultipleRes `json:"multi"`
	Title string `json:"title"`
}
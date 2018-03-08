package entity

import "gopkg.in/mgo.v2/bson"


type Multiple struct {
	Text string `json:"text"`
	Options []option `json:"options"`
	Answer []string `json:"answer"`
	Brief string `json:"brief"`
	Images []string `json:"images"`
}

type MultipleRes struct {
	Text string `json:"text"`
	Options []option `json:"options"`
	Answer []string `json:"answer"`
	Brief string `json:"brief"`
	Images []string `json:"images"`
	Id bson.ObjectId `json:"id" bson:"_id"`
}
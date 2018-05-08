package entity

import "gopkg.in/mgo.v2/bson"


type Multiple struct {
	Text string `json:"text"`
	Options []option `json:"options"`
	Answer []string `json:"answer"`
	Brief string `json:"brief"`
	Images []image `json:"images"`
	Type string `json:"type"`
	Answertype string `json:"answertype"`
}

type MultipleRes struct {
	Text string `json:"text"`
	Options []option `json:"options"`
	Answer []string `json:"answer"`
	Brief string `json:"brief"`
	Images []image `json:"images"`
	Id bson.ObjectId `json:"id" bson:"_id"`
	Type string `json:"type"`
	Answertype string `json:"answertype"`
}

type MultiReq struct {
	Text string `json:"text"`
	Options []option `json:"options"`
	Answer []string `json:"answer"`
	Brief string `json:"brief"`
	Images []image `json:"images"`
	Id string `json:"id"`
	Type string `json:"type"`
	Answertype string `json:"answertype"`
}
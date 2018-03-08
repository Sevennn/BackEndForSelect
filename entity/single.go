package entity

import (
	"gopkg.in/mgo.v2/bson"
	// "os"
)
type option struct {
	Text string `json:"text"`
	Value string `json:"value"`
}

type image struct {
	Name string `json:"name"`
	Data string `json:"data"`
	Index string `json:"index"`
}

type Single struct {
	Text string `json:"text"`
	Options []option `json:"options"`
	Answer string `json:"answer"`
	Brief string `json:"brief"`
	Images []image `json:"images"`
	Type string `json:"type"`
	Answertype string `json:"answertype"`
}

type SingleRes struct {
	Text string `json:"text"`
	Options []option `json:"options"`
	Answer string `json:"answer"`
	Brief string `json:"brief"`
	Images []image `json:"images"`
	Id bson.ObjectId `json:"id" bson:"_id"`
	Type string `json:"type"`
	Answertype string `json:"answertype"`
}

package models

type Question struct {
	QuestionId int64
	Number     int64
	ParentId   int64
	Question   string
	Answer     string
	Authors    string
	Comments   string
}

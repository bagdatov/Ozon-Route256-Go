package models

// Question - сущность вопроса, парсится с сайта ЧГК.
type Question struct {
	QuestionId   int64  `xml:"QuestionId"`   // Уникальный ID вопроса
	ParentId     int64  `xml:"ParentId"`     // ID тура куда относится вопрос
	Number       int64  `xml:"Number"`       // Порядковый номер вопроса
	Type         string `xml:"Type"`         // Тип сущности
	TextId       string `xml:"TextId"`       // Текстовый ключ вопроса <ключ турнира + номер тура + номер вопроса> например: lidamajor20_u.1-1
	Question     string `xml:"Question"`     // Текст вопроса
	Answer       string `xml:"Answer"`       // Ответ
	Authors      string `xml:"Authors"`      // Авторы вопроса
	Sources      string `xml:"Sources"`      // Источники
	Comments     string `xml:"Comments"`     // Комментарии к вопросу
	ParentTextId string `xml:"ParentTextId"` // Ключ тура
}

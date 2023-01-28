package models

import "encoding/xml"

// Tour - сущность тура. Каждый турнир состоит из туров, которые состоят из вопросов.
// Парсится с сайта ЧГК.
type Tour struct {
	XMLName      xml.Name   `xml:"tournament"`
	ID           int64      `xml:"Id"`           // Уникальный ID тура
	ParentId     int64      `xml:"ParentId"`     // ID турнира к которму относится тур
	QuestionsNum int64      `xml:"QuestionsNum"` // Количество вопросов в туре
	Number       int64      `xml:"Number"`       // Порядковый номер тура в турнире
	Title        string     `xml:"Title"`        // Название тура
	TextId       string     `xml:"TextId"`       // Ключ тура <ключ тура + номер тура> например lidamajor20_u.1
	Type         string     `xml:"Type"`         // Тип сущности
	Info         string     `xml:"Info"`         // Допольнительная информация
	Editors      string     `xml:"Editors"`      // Редакторы тура
	CreatedAt    string     `xml:"CreatedAt"`    // Дата создания
	Questions    []Question `xml:"question"`     // Вопросы тура
}

package models

import "encoding/xml"

// Tournament - сущность турнира. Каждый турнир состоит из туров, которые сами состоят из вопросов.
// Парсится с сайта ЧГК. Турнир можно еще назвать пакетом вопросов.
type Tournament struct {
	XMLName      xml.Name `xml:"tournament"`
	ID           int64    `xml:"Id"`           // Уникальный ID турнира
	QuestionsNum int64    `xml:"QuestionsNum"` // Количество вопросов во всем турнире
	ChildrenNum  int64    `xml:"ChildrenNum"`  // Количество туров в турнире
	Title        string   `xml:"Title"`        // Название турнира
	TextId       string   `xml:"TextId"`       // Ключ турнира
	Type         string   `xml:"Type"`         // Тип сущности
	Info         string   `xml:"Info"`         // Дополнительная информация
	Editors      string   `xml:"Editors"`      // Редакторы
	CreatedAt    string   `xml:"CreatedAt"`    // Дата создания
	Tours        []TourI  `xml:"tour"`         // Туры
}

type TourI struct {
	ID           int64  `xml:"Id"`           // Уникальный ID тура
	ParentId     int64  `xml:"ParentId"`     // ID турнира к которму относится тур
	QuestionsNum int64  `xml:"QuestionsNum"` // Количество вопросов в туре
	ChildrenNum  int64  `xml:"ChildrenNum"`  // Количество связанных сущностей
	Number       int64  `xml:"Number"`       // Порядковый номер тура в турнире
	Title        string `xml:"Title"`        // Название тура
	TextId       string `xml:"TextId"`       // Ключ тура <ключ тура + номер тура> например lidamajor20_u.1
	Type         string `xml:"Type"`         // Тип сущности
	Info         string `xml:"Info"`         // Допольнительная информация
	Editors      string `xml:"Editors"`      // Редакторы тура
	CreatedAt    string `xml:"CreatedAt"`    // Дата создания
	ParentTextId string `xml:"ParentTextId"` // Ключ родительской сущности
}

package models

import "time"

// Session - сущность сессии. Каждая запущенная игра представляется в виде сессии.
type Session struct {
	Tournament string    // Ключ турнира
	Title      string    // Название турнира
	IsActive   bool      // Завершенность игры
	Created    time.Time // Дата создания
}

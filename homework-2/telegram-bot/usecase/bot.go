package usecase

import (
	"context"
	"fmt"
	"gitlab.ozon.dev/bagdatov/homework-2/telegram-bot/models"
	"log"
	"strings"
	"sync"
	"time"
)

type bot struct {
	messenger Messenger
	game      Game
	c         cache
}

type cache struct {
	Sessions map[int64]*models.Session
	*sync.RWMutex
}

func New(g Game, m Messenger) *bot {
	return &bot{
		game:      g,
		messenger: m,
		c: cache{
			make(map[int64]*models.Session),
			new(sync.RWMutex),
		},
	}
}

func (b *bot) Help() string {
	return `Список доступных комманд:
/help - список доступных комманд
/score - счёт
/stop - остановить турнир
/next - следующий вопрос
/random - показать пять случайных пакетов

/begin <ключ турнира> - начать турнир
/desc <ключ турнира> - описание турнира

Для начала предоставьте к боту доступ к сообщениям.
Чтобы ответить сделайте реплай на сообщение с вопросом.
`
}

func (b *bot) Question(chatID, questionID int64) int {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	q, err := b.game.Question(ctx, questionID)
	if err != nil {
		log.Printf("couldn't get question from server: %s", err)
		return 0
	}

	text := fmt.Sprintf("Вопрос %d\n%s", q.Number, q.Question)

	id, err := b.messenger.Send(chatID, text)
	if err != nil {
		log.Printf("couldn't send answer to chat: %s\n", err)
	}
	return id
}

func (b *bot) FinishTournament(chatID int64) {

	b.c.Lock()
	defer b.c.Unlock()

	_, ok := b.c.Sessions[chatID]
	if !ok {
		if _, err := b.messenger.Send(chatID, "Сессия не начата"); err != nil {
			log.Println(err)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := b.game.Finish(ctx, chatID)
	if err != nil {
		if _, err2 := b.messenger.Send(chatID, err.Error()); err2 != nil {
			log.Println(err2)
		}
	}

	delete(b.c.Sessions, chatID)
	if _, err := b.messenger.Send(chatID, "Остановлено"); err != nil {
		log.Println(err)
	}
}

func (b *bot) Answer(chatID, questionID int64) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	q, err := b.game.Answer(ctx, questionID)
	if err != nil {
		log.Printf("couldn't send answer to chat: %s", err)
		return
	}

	text := fmt.Sprintf(`Ответ на вопрос %d
%s
Комментарий: %s
Авторы вопроса: %s`,
		q.Number, q.Answer, q.Comments, q.Authors)

	if _, err := b.messenger.Send(chatID, text); err != nil {
		log.Printf("couldn't send answer to chat: %s\n", err)
	}
}

func (b *bot) Next(chatID int64) string {
	b.c.Lock()
	defer b.c.Unlock()

	s, ok := b.c.Sessions[chatID]
	if !ok {
		return "Сессия не начата"
	}

	go func() {
		s.Next <- struct{}{}
	}()

	return "Скипаем..."
}

func (b *bot) Stop(chatID int64) string {
	b.c.Lock()
	defer b.c.Unlock()

	s, ok := b.c.Sessions[chatID]
	if !ok {
		return "Сессия не начата"
	}

	go func() {
		s.Stop <- struct{}{}
	}()
	return "Останавливаем..."
}

func (b *bot) Begin(ctx context.Context, chatID int64, tournament string) string {
	qIDs, err := b.game.Begin(ctx, chatID, tournament)
	if err != nil {
		return err.Error()
	}
	b.c.Lock()
	defer b.c.Unlock()

	s := &models.Session{
		TournamentKey:   tournament,
		TournamentID:    0,
		Questions:       qIDs,
		Current:         0,
		QuestionMessage: 0,
		Next:            make(chan struct{}),
		Stop:            make(chan struct{}),
		Mu:              new(sync.RWMutex),
	}

	go b.startSession(s, 60*time.Second, chatID)

	b.c.Sessions[chatID] = s

	return "Поехали"
}

func (b *bot) startSession(session *models.Session, dur time.Duration, chID int64) {
	time.Sleep(3 * time.Second)

	session.Mu.Lock()
	session.QuestionMessage = b.Question(chID, session.Questions[0])
	session.Mu.Unlock()

	ticker := time.NewTicker(dur)
	defer ticker.Stop()

	for {
		select {
		case <-session.Stop:
			b.FinishTournament(chID)
			return

		case <-session.Next:
			i := session.Current
			b.Answer(chID, session.Questions[i])

			if len(session.Questions)-1 == i {
				b.FinishTournament(chID)
				return
			}

			i++
			session.Mu.Lock()
			session.QuestionMessage = b.Question(chID, session.Questions[i])
			session.Current = i
			session.Mu.Unlock()
			ticker.Reset(dur)

		case <-ticker.C:
			i := session.Current
			b.Answer(chID, session.Questions[i])

			if len(session.Questions)-1 == i {
				b.FinishTournament(chID)
				return
			}

			i++
			session.Mu.Lock()
			session.QuestionMessage = b.Question(chID, session.Questions[i])
			session.Current = i
			session.Mu.Unlock()
		}
	}
}

func (b *bot) Submit(ctx context.Context, chatID int64, replyTo int, username, guess string) string {

	b.c.RLock()
	defer b.c.RUnlock()

	s, ok := b.c.Sessions[chatID]
	if !ok {
		return "Сессия не началась"
	}

	s.Mu.RLock()
	defer s.Mu.RUnlock()

	if replyTo != s.QuestionMessage {
		return "Чтобы ответить сделайте реплай на сообщение с вопросом"
	}

	questionID := s.Questions[s.Current]

	correct, err := b.game.Submit(ctx, chatID, questionID, username, guess)
	if err != nil {
		return err.Error()
	}

	if !correct {
		return "Неверный ответ"
	}

	return "Ответ верный"
}

func (b *bot) Score(ctx context.Context, chatID int64) string {

	users, err := b.game.Score(ctx, chatID)
	if err != nil {
		return err.Error()
	}

	var res strings.Builder
	res.WriteString("Результаты чата:\n")

	for _, u := range users {
		line := fmt.Sprintf("%s - %d\n", u.Name, u.Score)
		res.WriteString(line)
	}

	return res.String()
}

func (b *bot) Details(ctx context.Context, tournamentKey string) string {

	t, err := b.game.Tournament(ctx, tournamentKey)
	if err != nil {
		return err.Error()
	}

	return fmt.Sprintf(`Название: %s
Ключ: %s,
ID: %d,
Количество туров: %d
Количество вопросов: %d
Дата: %s
`, t.Name, t.Key, t.ID, t.ToursNum, t.QuestionNum, t.Date)
}

func (b *bot) Random(ctx context.Context) string {

	tournaments, err := b.game.Random(ctx)
	if err != nil {
		return err.Error()
	}

	var res strings.Builder

	for _, t := range tournaments {

		line := fmt.Sprintf(`Название: %s
ID: %d
Ключ: %s
Количество туров: %d
Количество вопросов: %d


`, t.Name, t.ID, t.Key, t.ToursNum, t.QuestionNum)

		res.WriteString(line)

	}

	return res.String()
}

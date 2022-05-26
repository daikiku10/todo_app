package models

import (
	"log"
	"time"
)

type Todo struct {
	ID int
	Content string
	UserID int
	CreatedAt time.Time
}

/** タスクの作成 */
func (u *User) CreateTodo(content string) (err error) {
	cmd := `insert into todos (
		content,
		user_id,
		created_at) values (?, ?, ?)`

	_, err = Db.Exec(cmd, content, u.ID, time.Now())

	// エラーハンドリング
	if err != nil {
		log.Fatalln(err)
	}
	
	return err
}
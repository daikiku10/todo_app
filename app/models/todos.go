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

/** タスクの取得 */
func GetTodo(id int) (todo Todo, err error) {
	cmd := `select id, content, user_id, created_at from todos
	where id = ?`

	todo = Todo{}

	err = Db.QueryRow(cmd, id).Scan(
		&todo.ID,
		&todo.Content,
		&todo.UserID,
		&todo.CreatedAt,
	)

	return todo, err
}
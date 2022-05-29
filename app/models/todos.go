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

/** タスクの全取得 */
func GetTodos() (todos []Todo, err error) {
	cmd := `select id, content, user_id, created_at from todos`

	rows, err := Db.Query(cmd)

	// エラーハンドリング
	if err != nil {
		log.Fatalln(err)
	}

	for rows.Next() {
		var todo Todo
		err = rows.Scan(
			&todo.ID,
			&todo.Content,
			&todo.UserID,
			&todo.CreatedAt,
		)

		// エラーハンドリング
		if err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}

	rows.Close()

	return todos, err
	
}

/** 特定のユーザーのタスクを取得する */
func (u *User) GetTodosByUser() (todos []Todo, err error) {
	cmd := `select id, content, user_id, created_at from todos
	where user_id = ?`

	rows, err := Db.Query(cmd, u.ID)

	// エラーハンドリング
	if err != nil {
		log.Fatalln(err)
	}
	
	for rows.Next() {
		var todo Todo
		err = rows.Scan(
			&todo.ID,
			&todo.Content,
			&todo.UserID,
			&todo.CreatedAt,
		)

		// エラーハンドリング
		if err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}

	rows.Close()

	return todos, err
}
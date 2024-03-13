package storage

import (
	"database/sql"
	"fmt"
	"lms_backed_pr/model"
	pkg "lms_backed_pr/pkg/check"

	"github.com/google/uuid"
)

type tasklms struct {
	db *sql.DB
}

func Newtask(db *sql.DB) tasklms {

	return tasklms{

		db: db,
	}

}

func (c *tasklms) Creatask(task model.Tasks) (string, error) {

	id := uuid.New()

	query := `INSERT INTO tasks(
		id,
	lesson_id,
	group_id,
	task,
	score,
	created_at
	updated_at)
	VALUES($1,$2,$3,$4)`

	_, err := c.db.Exec(query,
		id.String(), task.Id,
		task.Lesson_id,
		task.Group_id,
		task.Task,
		task.Score,
	)
	if err != nil {
		fmt.Println("error while creating task:", err)
		return "", err
	}

	return id.String(), nil

}

func (c *tasklms) Updatetask(tasks model.Tasks) (string, error) {

	queryy := `UPDATE task set
	id=$1,
	lesson_id=$2,
	group_id=$3,
	task=$4,
	score=$5,
    updated_at=CURRENT_TIMESTAMP,
	id=$9 WHERE id=$9 AND deleted_at=0
    `

	_, err := c.db.Exec(queryy,
		tasks.Id,
		tasks.Lesson_id,
		tasks.Group_id,
		tasks.Task,
		tasks.Score)
	if err != nil {
		fmt.Println("Error while updating task:", err)
		return "", err
	}

	return tasks.Id, nil
}

func (c *tasklms) Getallstudents(search string) (model.Gettask, error) {

	var (
		resp   = model.Gettask{}
		filter = ""
	)

	if search != "" {
		filter += fmt.Sprintf(` and name ILIKE  '%%%v%%' `, search)
	}

	fmt.Println("filter: ", filter)

	rows, err := c.db.Query(`select 
				count(id) OVER(),
				id,
				lesson_id,
				group_id,
				task,
				score, 
				created_at::date,
				updated_at
	  FROM student WHERE deleted_at = 0 ` + filter + ``)

	if err != nil {
		return resp, err
	}

	for rows.Next() {

		var (
			task   = model.Tasks{}
			update sql.NullString
		)

		if err := rows.Scan(

			&resp.Count,
			&task.Id,
			&task.Lesson_id,
			&task.Group_id,
			&task.Task,
			&task.Score,
			&update); err != nil {
			fmt.Println("error while scaning all infos", err)
			return resp, err
		}
		task.Updated_at = pkg.NullStringToString(update)
		resp.Task = append(resp.Task, task)

	}
	return resp, nil

}

func (c *tasklms) Getbyidstudents(id string) (model.Tasks, error) {

	task := model.Tasks{}
	if err := c.db.QueryRow(`SELECT 
	id ,
	first_name, 
	last_name ,
	gmail, 
	phone,
    is_blocked
	from tasks where id=$1`, id).Scan(

		&task.Id,
		&task.Id,
		&task.Lesson_id,
		&task.Group_id,
		&task.Task,
		&task.Score,
	); err != nil {

		return model.Tasks{}, err
	}
	return task, nil

}

func (c *tasklms) Deletecus(id string) error {

	query := `UPDATE tasks  set 
	deleted_at=date_part('epoch',CURRENT_TIMESTAMP)::int
where id=$1 AND deleted_at=0
`
	_, err := c.db.Exec(query, id)

	if err != nil {
		return err
	}

	return nil

}

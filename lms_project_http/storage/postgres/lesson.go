package storage

import (
	"database/sql"
	"fmt"
	"lms_backed_pr/model"
	pkg "lms_backed_pr/pkg/check"

	"github.com/google/uuid"
)

type lessonLms struct {
	db *sql.DB
}

func Newlesson(db *sql.DB) lessonLms {

	return lessonLms{

		db: db,
	}

}

func (c *lessonLms) Creatlesson(lesson model.Lesson) (string, error) {

	id := uuid.New()

	query := `INSERT INTO group(
	id        
	group_id     
	branch_id    
	teacher    
	type     
	created_at
	updated_at)
	VALUES($1,$2,$3,$4,$5)`

	_, err := c.db.Exec(query,
		id.String(),
		lesson.Schedule_id,
		lesson.Group_id,
		lesson.From,
		lesson.To,
		lesson.Theme)
	if err != nil {
		fmt.Println("error while creating lesson:", err)
		return "", err
	}

	return id.String(), nil

}

func (c *teacherlms) Updatelesson(lesson model.Lesson) (string, error) {

	queryy := `UPDATE student set
	id =$1
	schedule_id=$2 
	group_id =$3,    
	from=$4,
	to=$5    
	theme=$6  
    updated_at=CURRENT_TIMESTAMP,
	id=$8 WHERE id=$8 AND deleted_at=0
    `

	_, err := c.db.Exec(queryy,
		&lesson.ID,
		&lesson.Schedule_id,
		&lesson.Group_id,
		&lesson.From,
		&lesson.To,
		&lesson.Theme)

	if err != nil {
		fmt.Println("Error while updating teacher:", err)
		return "", err
	}

	return lesson.ID, nil
}

func (c *lessonLms) Getalllesson(search string) (model.GEtlesson, error) {

	var (
		respp  = model.GEtlesson{}
		filter = ""
	)

	if search != "" {
		filter += fmt.Sprintf(` and name ILIKE  '%%%v%%' `, search)
	}

	fmt.Println("filter: ", filter)

	rows, err := c.db.Query(`select 
				count(id) OVER(),
				id        
				schedule_id 
				group_id     
				from    
				to    
				theme     
				created_at::date,
				updated_at
	  FROM lesson WHERE deleted_at = 0 ` + filter + ``)

	if err != nil {
		return respp, err
	}

	for rows.Next() {

		var (
			lesson = model.Lesson{}
			update sql.NullString
		)

		if err := rows.Scan(

			&respp.Count,
			&lesson.ID,
			&lesson.Schedule_id,
			&lesson.Group_id,
			&lesson.From,
			&lesson.To,
			&lesson.Theme); err != nil {
			fmt.Println("error while scaning all infos", err)
			return respp, err
		}
		lesson.Updated_at = pkg.NullStringToString(update)
		respp.Lesson = append(respp.Lesson, lesson)

	}
	return respp, nil

}

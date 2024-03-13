package storage

import (
	"database/sql"
	"fmt"
	"lms_backed_pr/model"

	"github.com/google/uuid"
)

type groupLms struct {
	db *sql.DB
}

func Newgroup(db *sql.DB) groupLms {

	return groupLms{

		db: db,
	}

}

func (c *groupLms) Cretegroup(group model.Group) (string, error) {

	id := uuid.New()

	query := `INSERT INTO lesson(
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
		group.Id,
		group.Group_id,
		group.Branch_id,
		group.Teacher,
		group.Type)
	if err != nil {
		fmt.Println("error while creating group:", err)
		return "", err
	}

	return id.String(), nil

}

func (c *groupLms) Updategroup(group model.Group) (string, error) {

	queryy := `UPDATE group set
	id=$1,     
	group_id=$2,     
	branch_id=$3, 
	teacher=$4,    
	type=$5,
    updated_at=CURRENT_TIMESTAMP,
	id=$8 WHERE id=$8 AND deleted_at=0
    `

	_, err := c.db.Exec(queryy,
		&group.Id,
		&group.Group_id,
		&group.Branch_id,
		&group.Teacher,
		&group.Type)

	if err != nil {
		fmt.Println("Error while updating group:", err)
		return "", err
	}

	return group.Id, nil
}

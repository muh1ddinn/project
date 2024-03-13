package storage

import (
	"database/sql"
	"fmt"
	"lms_backed_pr/model"
	pkg "lms_backed_pr/pkg/check"

	"github.com/google/uuid"
)

type scheduleLms struct {
	db *sql.DB
}

func Newschedule(db *sql.DB) scheduleLms {

	return scheduleLms{

		db: db,
	}

}

func (c *teacherlms) Creatschedule(schedule model.Schedule) (string, error) {

	id := uuid.New()

	query := `INSERT INTO teacher(
	id        
	group_id 
	group_type     
	start_time    
	end_time    
	date     
	branch_id  
	teacher_id
	created_at
	updated_at)
	VALUES($1,$2,$3,$4,$5,$6,$7)`

	_, err := c.db.Exec(query,
		id.String(),
		schedule.Group_id,
		schedule.Group_type,
		schedule.Start_time,
		schedule.End_time,
		schedule.Date,
		schedule.Branch_id,
		schedule.Teacher_id)
	if err != nil {
		fmt.Println("error while creating schedule:", err)
		return "", err
	}

	return id.String(), nil

}

func (c *teacherlms) Updateschedule(schedule model.Schedule) (string, error) {

	queryy := `UPDATE student set
	group_id=$1,
	group_type=$2,    
	start_time=$3,    
	end_time=$4,    
	date=$5,    
	branch_id=$6  
	teacher_id=$7
    updated_at=CURRENT_TIMESTAMP,
	id=$8 WHERE id=$8 AND deleted_at=0
    `

	_, err := c.db.Exec(queryy,
		schedule.Group_type,
		schedule.Start_time,
		schedule.End_time,
		schedule.Date,
		schedule.Branch_id,
		schedule.Teacher_id)

	if err != nil {
		fmt.Println("Error while updating teacher:", err)
		return "", err
	}

	return schedule.Teacher_id, nil
}

func (c *scheduleLms) Getallschedule(search string) (model.Getschedule, error) {

	var (
		respp  = model.Getschedule{}
		filter = ""
	)

	if search != "" {
		filter += fmt.Sprintf(` and name ILIKE  '%%%v%%' `, search)
	}

	fmt.Println("filter: ", filter)

	rows, err := c.db.Query(`select 
				count(id) OVER(),
				id        
				group_id 
				group_type     
				start_time    
				end_time    
				date     
				branch_id  
				teacher_id 
				created_at::date,
				updated_at
	  FROM schedule WHERE deleted_at = 0 ` + filter + ``)

	if err != nil {
		return respp, err
	}

	for rows.Next() {

		var (
			schedule = model.Schedule{}
			update   sql.NullString
		)

		if err := rows.Scan(

			&respp.Count,
			&schedule.Group_type,
			&schedule.Start_time,
			&schedule.End_time,
			&schedule.Date,
			&schedule.Branch_id,
			&schedule.Teacher_id); err != nil {
			fmt.Println("error while scaning all infos", err)
			return respp, err
		}
		schedule.Updated_at = pkg.NullStringToString(update)
		respp.Schedule = append(respp.Schedule, schedule)

	}
	return respp, nil

}

package storage

import (
	"database/sql"
	"fmt"
	"lms_backed_pr/model"
	pkg "lms_backed_pr/pkg/check"

	"github.com/google/uuid"
)

type teacherlms struct {
	db *sql.DB
}

func Newteacher(db *sql.DB) teacherlms {

	return teacherlms{

		db: db,
	}

}

func (c *teacherlms) Createteacher(teacher model.Teacher) (string, error) {

	id := uuid.New()

	query := `INSERT INTO teacher(
	id        
	full_name 
	email     
	age    
	status    
	login     
	password  
	created_at
	updated_at)
	VALUES($1,$2,$3,$4,$5,$6,$7)`

	_, err := c.db.Exec(query,
		id.String(), teacher.Id,
		teacher.Full_name,
		teacher.Email,
		teacher.Age,
		teacher.Status,
		teacher.Login,
		teacher.Password)
	if err != nil {
		fmt.Println("error while creating teacher:", err)
		return "", err
	}

	return id.String(), nil

}

func (c *teacherlms) Updateteacher(teacher model.Teacher) (string, error) {

	queryy := `UPDATE student set
	full_name =$1,
	email =$2,
	age =$3,
	status =$4,
	login =$5,    
	password=$6
    updated_at=CURRENT_TIMESTAMP,
	id=$7 WHERE id=$7 AND deleted_at=0
    `

	_, err := c.db.Exec(queryy,
		teacher.Full_name,
		teacher.Email,
		teacher.Age,
		teacher.Status,
		teacher.Login,
		teacher.Password)
	if err != nil {
		fmt.Println("Error while updating teacher:", err)
		return "", err
	}

	return teacher.Id, nil
}

func (c *teacherlms) Getallteacher(search string) (model.Getteacher, error) {

	var (
		resp   = model.Getteacher{}
		filter = ""
	)

	if search != "" {
		filter += fmt.Sprintf(` and name ILIKE  '%%%v%%' `, search)
	}

	fmt.Println("filter: ", filter)

	rows, err := c.db.Query(`select 
				count(id) OVER(),
	id        
	full_name,
	email,
	age,    
	paid_sum,   
	status,    
	login,     
	password,  
	group_id,  
				created_at::date,
				updated_at
	  FROM teacher WHERE deleted_at = 0 ` + filter + ``)

	if err != nil {
		return resp, err
	}

	for rows.Next() {

		var (
			teacher = model.Teacher{}
			update  sql.NullString
		)

		if err := rows.Scan(

			&resp.Count,
			&teacher.Id,
			&teacher.Full_name,
			&teacher.Email,
			&teacher.Age,
			&teacher.Paid_sum,
			&teacher.Status,
			&teacher.Login,
			&teacher.Password,
			&update); err != nil {
			fmt.Println("error while scaning all infos", err)
			return resp, err
		}
		teacher.Updated_at = pkg.NullStringToString(update)
		resp.Teacher = append(resp.Teacher, teacher)

	}
	return resp, nil

}

func (c *teacherlms) Getbyidteacher(id string) (model.Teacher, error) {

	teacher := model.Teacher{}
	if err := c.db.QueryRow(`SELECT 
	id ,
	first_name, 
	last_name ,
	gmail, 
	phone,
    is_blocked
	from teacher where id=$1`, id).Scan(

		&teacher.Id,
		&teacher.Full_name,
		&teacher.Email,
		&teacher.Age,
		&teacher.Paid_sum,
		&teacher.Status,
		&teacher.Login,
		&teacher.Password,
	); err != nil {

		return model.Teacher{}, err
	}
	return teacher, nil

}

func (c *teacherlms) Deleteteacher(id string) error {

	query := `UPDATE teacher  set 
	deleted_at=date_part('epoch',CURRENT_TIMESTAMP)::int
where id=$1 AND deleted_at=0
`
	_, err := c.db.Exec(query, id)

	if err != nil {
		return err
	}

	return nil

}

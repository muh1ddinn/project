package storage

import (
	"database/sql"
	"fmt"
	"lms_backed_pr/model"
	pkg "lms_backed_pr/pkg/check"

	"github.com/google/uuid"
)

type studentsLms struct {
	db *sql.DB
}

func Newstudent(db *sql.DB) studentsLms {

	return studentsLms{

		db: db,
	}

}

func (c *studentsLms) Createtudent(students model.Student) (string, error) {

	id := uuid.New()

	query := `INSERT INTO student(
	id        
	full_name 
	email     
	age    
	paid_sum   
	status    
	login     
	password  
	group_id  
	created_at
	updated_at)
	VALUES($1,$2,$3,$4,$5,$6,$7,$8)`

	_, err := c.db.Exec(query,
		id.String(), students.Id,
		students.Full_name,
		students.Age,
		students.Paid_sum,
		students.Status,
		students.Login,
		students.Password,
		students.Group_id)
	if err != nil {
		fmt.Println("error while creating student:", err)
		return "", err
	}

	return id.String(), nil

}

func (c *studentsLms) UpdateStudent(students model.Student) (string, error) {

	queryy := `UPDATE student set
	full_name =$1,
	email =$2,
	age =$3,
	paid_sum=$4,   
	status =$5,
	login =$6,    
	password=$7
	group_id=$8  ,
    updated_at=CURRENT_TIMESTAMP,
	id=$9 WHERE id=$9 AND deleted_at=0
    `

	_, err := c.db.Exec(queryy,
		students.Full_name,
		students.Age,
		students.Paid_sum,
		students.Status,
		students.Login,
		students.Password,
		students.Group_id,
		students.Id)
	if err != nil {
		fmt.Println("Error while updating customer:", err)
		return "", err
	}

	return students.Id, nil
}

func (c *studentsLms) GetallStudents(search string) (model.GetAllstudent, error) {

	var (
		resp   = model.GetAllstudent{}
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
	  FROM student WHERE deleted_at = 0 ` + filter + ``)

	if err != nil {
		return resp, err
	}

	for rows.Next() {

		var (
			students = model.Student{}
			update   sql.NullString
		)

		if err := rows.Scan(

			&resp.Count,
			&students.Id,
			&students.Full_name,
			&students.Email,
			&students.Age,
			&students.Paid_sum,
			&students.Status,
			&students.Login,
			&students.Password,
			&students.Group_id,
			&update); err != nil {
			fmt.Println("error while scaning all infos", err)
			return resp, err
		}
		students.Updated_at = pkg.NullStringToString(update)
		resp.Student = append(resp.Student, students)

	}
	return resp, nil

}

func (c *studentsLms) GetByIDStudent(id string) (model.Student, error) {

	studentt := model.Student{}
	if err := c.db.QueryRow(`SELECT 
	id ,
	first_name, 
	last_name ,
	gmail, 
	phone,
    is_blocked
	from student where id=$1`, id).Scan(

		&studentt.Id,
		&studentt.Full_name,
		&studentt.Email,
		&studentt.Age,
		&studentt.Paid_sum,
		&studentt.Status,
		&studentt.Login,
		&studentt.Password,
		&studentt.Group_id,
	); err != nil {

		return model.Student{}, err
	}
	return studentt, nil

}

func (c *studentsLms) DeleteStudent(id string) error {

	query := `UPDATE student  set 
	deleted_at=date_part('epoch',CURRENT_TIMESTAMP)::int
where id=$1 AND deleted_at=0
`
	_, err := c.db.Exec(query, id)

	if err != nil {
		return err
	}

	return nil

}

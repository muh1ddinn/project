package storage

import (
	"database/sql"
	"fmt"
	"lms_backed_pr/model"
	pkg "lms_backed_pr/pkg/check"

	"github.com/google/uuid"
)

type paymentLms struct {
	db *sql.DB
}

func Newpayment(db *sql.DB) paymentLms {

	return paymentLms{

		db: db,
	}

}

func (c *paymentLms) Createpayment(payment model.Payment) (string, error) {

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
		id.String(), payment.Id,
		payment.Full_name,
		payment.Age,
		payment.Paid_sum,
		payment.Status,
		payment.Login,
		payment.Password,
		payment.Group_id)
	if err != nil {
		fmt.Println("error while creating student:", err)
		return "", err
	}

	return id.String(), nil

}

func (c *paymentLms) Updatepayment(payment model.Payment) (string, error) {

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
		payment.Full_name,
		payment.Age,
		payment.Paid_sum,
		payment.Status,
		payment.Login,
		payment.Password,
		payment.Group_id,
		payment.Id)
	if err != nil {
		fmt.Println("Error while updating customer:", err)
		return "", err
	}

	return payment.Id, nil
}

func (c *paymentLms) Getallpayment(search string) (model.GetPayment, error) {

	var (
		resp   = model.GetPayment{}
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
			payment = model.Payment{}
			update  sql.NullString
		)

		if err := rows.Scan(

			&resp.Count,
			&payment.Id,
			&payment.Full_name,
			&payment.Email,
			&payment.Age,
			&payment.Paid_sum,
			&payment.Status,
			&payment.Login,
			&payment.Password,
			&payment.Group_id,
			&update); err != nil {
			fmt.Println("error while scaning all infos", err)
			return resp, err
		}
		payment.Updated_at = pkg.NullStringToString(update)
		resp.Payment = append(resp.Payment, payment)

	}
	return resp, nil

}

func (c *paymentLms) Getbyidpayment(id string) (model.Student, error) {

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

func (c *paymentLms) Deletepayment(id string) error {

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

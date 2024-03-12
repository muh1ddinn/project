package storage

import (
	model "cars_with_sql/models"
	"cars_with_sql/pkg"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type customerRepo struct {
	db *sql.DB
}

func Newcustomer(db *sql.DB) customerRepo {
	return customerRepo{
		db: db,
	}
}
func (c *customerRepo) Createcuss(customer model.Customers) (string, error) {

	id := uuid.New()

	query := `INSERT INTO customerss (
        id,
        first_name,
        last_name,
        gmail,
        phone,
        is_blocked)
		VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := c.db.Exec(query,
		id.String(),
		customer.First_name,
		customer.Last_name,
		customer.Gmail,
		customer.Phone,
		customer.Is_blocked)
	if err != nil {
		fmt.Println("you have error while creating :", err)
		return "", err

	}

	return id.String(), nil

}

func (c *customerRepo) Updatecus(customer model.Customers) (string, error) {

	queryy := `UPDATE customerss set
            first_name=$1,
            last_name=$2,
            gmail=$3,
            phone=$4,
            is_blocked=$5,
            updated_at=CURRENT_TIMESTAMP,
			id=$6
        WHERE id=$6 AND deleted_at=0
    `

	_, err := c.db.Exec(queryy,
		customer.First_name, customer.Last_name,
		customer.Gmail, customer.Phone,
		customer.Is_blocked, customer.Id)
	if err != nil {
		fmt.Println("Error while updating customer:", err)
		return "", err
	}

	return customer.Id, nil
}

func (c *customerRepo) GETallcus(search string) (model.GetAllCustomersResponse, error) {

	var (
		resp   = model.GetAllCustomersResponse{}
		filter = ""
	)

	if search != "" {
		filter += fmt.Sprintf(` and name ILIKE  '%%%v%%' `, search)
	}

	fmt.Println("filter: ", filter)

	rows, err := c.db.Query(`select 
				count(id) OVER(),
				id ,
				first_name, 
				last_name ,
				gmail, 
				phone,
				is_blocked,
				created_at::date,
				updated_at
	  FROM customerss WHERE deleted_at = 0 ` + filter + ``)

	if err != nil {
		return resp, err
	}

	for rows.Next() {

		var (
			customerr = model.Customers{}
			update    sql.NullString
		)

		if err := rows.Scan(

			&resp.Count,
			&customerr.Id,
			&customerr.First_name,
			&customerr.Last_name,
			&customerr.Gmail,
			&customerr.Phone,
			&customerr.Is_blocked,
			&customerr.Created_at,
			&update); err != nil {
			fmt.Println("error while scaning all infos", err)
			return resp, err
		}
		customerr.Updated_at = pkg.NullStringToString(update)
		resp.Customers = append(resp.Customers, customerr)

	}
	return resp, nil

}

func (c *customerRepo) Getbyid(id string) (model.Customers, error) {

	fmt.Println("hey mann")
	custommer := model.Customers{}
	if err := c.db.QueryRow(`SELECT 
	id ,
	first_name, 
	last_name ,
	gmail, 
	phone,
    is_blocked
	from customerss where id=$1`, id).Scan(

		&custommer.Id,
		&custommer.First_name,
		&custommer.Last_name,
		&custommer.Gmail,
		&custommer.Phone,
		&custommer.Is_blocked,
	); err != nil {

		return model.Customers{}, err
	}
	return custommer, nil

}

func (c *customerRepo) Delete(id string) error {

	query := `UPDATE customerss set 
	deleted_at=date_part('epoch',CURRENT_TIMESTAMP)::int
where id=$1 AND deleted_at=0
`
	_, err := c.db.Exec(query, id)

	if err != nil {
		return err
	}

	return nil

}

//

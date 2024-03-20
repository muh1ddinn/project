package postgres

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
func (c *customerRepo) CreateCus(customer model.Customers) (string, error) {

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
func (c *customerRepo) GetAllCustomers(search string) (model.GetAllCustomersResponse, error) {
	resp := model.GetAllCustomersResponse{}
	filter := ""

	if search != "" {
		filter = fmt.Sprintf(` AND first_name ILIKE '%%%v%%' `, search)
	}

	fmt.Println("filter: ", filter)

	query := `
        SELECT 
		count(id) OVER(),
            id,
            first_name,
            last_name,
            gmail,
            phone,
            is_blocked
        FROM 
            customerss 
        WHERE 
            deleted_at = 0` + filter

	rows, err := c.db.Query(query)
	if err != nil {
		return resp, err
	}
	defer rows.Close()

	for rows.Next() {
		var customer model.Customers
		var order model.GetOrder
		var orders model.GetAllOrder
		var updatedAt sql.NullString

		err := rows.Scan(
			&resp.Count,
			&customer.Id,
			&customer.First_name,
			&customer.Last_name,
			&customer.Gmail,
			&customer.Phone,
			&customer.Is_blocked,
		)
		if err != nil {
			fmt.Println("error while scanning customer info: ", err)
			return resp, err
		}

		query := `
            SELECT 
			count(id) OVER(),
			id,
                from_date,
                to_date,
                status,
                payment_status,
                created_at,
                updated_at
            FROM 
                orrders 
            WHERE 
                customer_id = $1`

		orderRows, err := c.db.Query(query, customer.Id)
		if err != nil {
			fmt.Println("error while querying orders: ", err)
			return resp, err
		}
		defer orderRows.Close()

		for orderRows.Next() {
			err := orderRows.Scan(
				&orders.Count,
				&order.Id,
				&order.FromDate,
				&order.ToDate,
				&order.Status,
				&order.Paid,
				&order.Created_at,
				&updatedAt,
			)
			if err != nil {
				fmt.Println("error while scanning order info: ", err)
				return resp, err
			}
			order.Updated_at = pkg.NullStringToString(updatedAt)
			customer.Orders = append(customer.Orders, order)
		}
		if err := orderRows.Err(); err != nil {
			fmt.Println("error while iterating over order rows: ", err)
			return resp, err
		}

		resp.Customers = append(resp.Customers, customer)
	}
	if err := rows.Err(); err != nil {
		fmt.Println("error while iterating over rows: ", err)
		return resp, err
	}

	return resp, nil
}

func (c *customerRepo) GetByIDCustomer(id string) (model.Customers, error) {

	query := `SELECT 
	count(id) OVER(),
        id,
        first_name,
        last_name,
        gmail,
        phone,
        is_blocked
        FROM customerss WHERE id=$1 AND deleted_at=0`

	row := c.db.QueryRow(query, id)

	customer := model.Customers{
		Orders: []model.GetOrder{},
	}

	err := row.Scan(

		&customer.Id,
		&customer.First_name,
		&customer.Last_name,
		&customer.Gmail,
		&customer.Phone,
		&customer.Is_blocked)
	if err != nil {
		fmt.Println("error while getting id customer err: ", err)
		return customer, err
	}

	query = `SELECT 
        id,
        from_date,
        to_date,
        status,
        payment_status,
        created_at,
        updated_at
        FROM orrders WHERE customer_id=$1`

	rows, err := c.db.Query(query, id)
	if err != nil {
		return model.Customers{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var order model.GetOrder
		var updatedAt sql.NullString
		err := rows.Scan(

			&order.Id,
			&order.FromDate,
			&order.ToDate,
			&order.Status,
			&order.Paid,
			&order.Created_at,
			&updatedAt,
		)
		if err != nil {
			fmt.Println("error while scanning getbyid: ", err)
			return customer, err
		}
		order.Updated_at = pkg.NullStringToString(updatedAt)
		customer.Orders = append(customer.Orders, order)
	}
	if err := rows.Err(); err != nil {
		fmt.Println("error while iterating over rows: ", err)
		return customer, err
	}

	return customer, nil
}

func (c *customerRepo) getcustomerandcar(id string) (model.Customers, error) {

	query := `SELECT 
	count(id) OVER(),
        id,
        first_name,
        last_name,
        gmail,
        phone,
        is_blocked
        FROM customerss WHERE id=$1 AND deleted_at=0`

	row := c.db.QueryRow(query, id)

	customer := model.Customers{
		Orders: []model.GetOrder{},
	}

	err := row.Scan(

		&customer.Id,
		&customer.First_name,
		&customer.Last_name,
		&customer.Gmail,
		&customer.Phone,
		&customer.Is_blocked)
	if err != nil {
		fmt.Println("error while getting id customer err: ", err)
		return customer, err
	}

	query = `SELECT 
        id,
        from_date,
        to_date,
        status,
        payment_status,
        created_at,
        updated_at
        FROM orrders WHERE customer_id=$1`

	rows, err := c.db.Query(query, id)
	if err != nil {
		return model.Customers{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var order model.GetOrder
		var updatedAt sql.NullString
		err := rows.Scan(

			&order.Id,
			&order.FromDate,
			&order.ToDate,
			&order.Status,
			&order.Paid,
			&order.Created_at,
			&updatedAt,
		)
		if err != nil {
			fmt.Println("error while scanning getbyid: ", err)
			return customer, err
		}
		order.Updated_at = pkg.NullStringToString(updatedAt)
		customer.Orders = append(customer.Orders, order)
	}
	if err := rows.Err(); err != nil {
		fmt.Println("error while iterating over rows: ", err)
		return customer, err
	}

	return customer, nil
}
func (c *customerRepo) DeleteCustomer(id string) error {

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

package postgres

import (
	model "cars_with_sql/models"
	"cars_with_sql/pkg"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type orderRepo struct {
	db *sql.DB
}

func Neworder(db *sql.DB) orderRepo {
	return orderRepo{
		db: db,
	}
}

func (c *orderRepo) Create(order model.CreateOrder) (string, error) {

	id := uuid.New()

	query := `INSERT INTO orrders (
		id,customer_id, 
		cars_id, from_date, 
		to_date,
		status,payment_status,
		amount,created_at,updated_at 
 )
		VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := c.db.Exec(query,
		id.String(),

		order.CustomerId,
		order.CarId,
		order.FromDate,
		order.ToDate,
		order.Status,
		order.Paid, order.Amount)

	if err != nil {
		fmt.Println("you have error while creating :", err)
		return "", err

	}

	return id.String(), nil

}

func (o *orderRepo) GetAll(search string) (model.GetAllOrder, error) {
	response := &model.GetAllOrder{}

	query := `SELECT 
        o.id,
        c.name AS car_name,
        c.brand AS car_brand,
        cu.id AS customer_id,
        cu.first_name AS customer_first_name,
        cu.gmail AS customer_gmail,
        o.from_date,
        o.to_date,
        o.status,
        o.payment_status,
        o.created_at,
        o.updated_at
        FROM orrders o
        JOIN cars c ON o.cars_id=c.id
        JOIN customerss cu ON o.customer_id = cu.id`

	rows, err := o.db.Query(query)
	if err != nil {
		return *response, err
	}
	defer rows.Close()

	for rows.Next() {
		order := model.GetOrder{}
		car := model.Car{}
		customer := model.Customers{}
		var update sql.NullString

		err := rows.Scan(
			&order.Id,
			&car.Name,
			&car.Brand,
			&customer.Id,
			&customer.First_name,
			&customer.Gmail,
			&order.FromDate,
			&order.ToDate,
			&order.Status,
			&order.Paid,
			&order.Created_at,
			&update,
		)
		if err != nil {
			return *response, err
		}
		order.Updated_at = pkg.NullStringToString(update)

		order.Car = car
		order.Customers = customer

		response.Orders = append(response.Orders, order)
	}

	if err = rows.Err(); err != nil {
		return model.GetAllOrder{}, nil
	}
	return *response, nil
}

func (o *orderRepo) GetByID(id string) (model.GetOrder, error) {

	response := &model.GetOrder{}

	query := `SELECT 
        o.id,
        c.name AS car_name,
        c.brand AS car_brand,
        cu.id AS customer_id,
        cu.first_name AS customer_first_name,
        cu.gmail AS customer_gmail,
        o.from_date,
        o.to_date,
        o.status,
        o.payment_status,
        o.created_at,
        o.updated_at
        FROM orrders o
        JOIN cars c ON o.car_id=c.id
        JOIN customerss cu ON o.customer_id = cu.id`

	rows, err := o.db.Query(query)
	if err != nil {
		return *response, err
	}
	defer rows.Close()

	for rows.Next() {
		order := model.GetOrder{}
		car := model.Car{}
		customer := model.Customers{}
		var update sql.NullString
		err := rows.Scan(
			&order.Id,
			&car.Name,
			&car.Brand,
			&customer.Id,
			&customer.First_name,
			&customer.Gmail,
			&order.FromDate,
			&order.ToDate,
			&order.Status,
			&order.Paid,
			&order.Created_at,
			&update,
		)
		if err != nil {
			return *response, err
		}
		order.Updated_at = pkg.NullStringToString(update)

	}

	return *response, nil
}

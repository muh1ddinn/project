package storage

import (
	model "cars_with_sql/models"
	"cars_with_sql/pkg"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type carRepo struct {
	DATA *sql.DB
}

func Newwcar(DATA *sql.DB) carRepo {
	return carRepo{
		DATA: DATA,
	}
}

/*
======================================
create (body) id ,err
update(body) id,err
delete(id) err
get(id) body,err
getALL(serach) []body,count,err

======================================
*/
func (c *carRepo) Createcar(car model.Car) (string, error) {

	id := uuid.New()

	query := `INSERT INTO cars(
        id,
        name,
        brand,
        model,
		year,
        hourse_power,
        colour,
        engine_cap)
    VALUES($1,$2,$3,$4,$5,$6,$7,$8) `

	res, err := c.DATA.Exec(query,
		id.String(),
		car.Name, car.Brand,
		car.Model, car.Year,
		car.HoursePower, car.Colour, car.EngineCap)

	if err != nil {
		return "", fmt.Errorf("error executing query: %w", err)

	}
	fmt.Printf("%+v\n", res)

	return id.String(), nil
}

// //////////////////////////////////////////////////////
// /////////////////////////////////////////////////////
// /////////////////////////////////////////////////////

func (c carRepo) GetAll(req model.GetAllCarsRequest) (model.GetAllCarsResponse, error) {
	var (
		resp   = model.GetAllCarsResponse{}
		filter = ""
	)
	offset := (req.Page - 1) * req.Limit

	if req.Search != "" {
		filter += fmt.Sprintf(` and name ILIKE  '%%%v%%' `, req.Search)
	}

	filter += fmt.Sprintf(" OFFSET %v LIMIT %v", offset, req.Limit)
	fmt.Println("filter: ", filter)
	rows, err := c.DATA.Query(`select 
				count(id) OVER(),
				id, 
				name,
				brand,
				model,
				year,
				hourse_power,
				colour,
				engine_cap,
				created_at::date,
				updated_at
	  FROM cars WHERE deleted_at = 0 ` + filter + `
	  `)
	if err != nil {
		return resp, err
	}
	for rows.Next() {
		var (
			car      = model.Car{}
			updateAt sql.NullString
		)

		if err := rows.Scan(
			&resp.Count,
			&car.Id,
			&car.Name,
			&car.Brand,
			&car.Model,
			&car.Year,
			&car.HoursePower,
			&car.Colour,
			&car.EngineCap,
			&car.CreatedAt,
			&updateAt); err != nil {
			return resp, err
		}

		car.UpdatedAt = pkg.NullStringToString(updateAt)
		resp.Cars = append(resp.Cars, car)
	}
	return resp, nil
}

func (c carRepo) GetByidcar(id string) ([]model.Car, error) {
	carrr := []model.Car{}
	rows, err := c.DATA.Query(`SELECT 
    id,
    name,
    brand,
    model,
    year,
    hourse_power,
    colour,
    engine_cap FROM cars
    where id=$1`, id)
	if err != nil {
		fmt.Println("error while getting id country err: ", err)
		return carrr, err
	}

	defer rows.Close()

	for rows.Next() {
		car := model.Car{}
		if err = rows.Scan(&car.Id, &car.Name, &car.Brand, &car.Model, &car.Year, &car.HoursePower, &car.Colour, &car.EngineCap); err != nil {
			fmt.Println("error while scanning country err: ", err)
			return nil, err
		}
		carrr = append(carrr, car)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("error while iterating rows: ", err)
		return nil, err
	}

	return carrr, nil
}

////////////////////////////////////////////////////////////
/*
func (c carRepo) GetAll(search string) (model.GetAllCarsResponse, error) {
	var (
		resp   = model.GetAllCarsResponse{}
		filter = ""
	)

	if search != "" {
		filter += fmt.Sprintf(` and name ILIKE  '%%%v%%' `, search)
	}

	fmt.Println("filter: ", filter)

	rows, err := c.db.Query(`select
				count(id) OVER(),
				id,
				name,
				brand,
				model,
				year,
				hourse_power,
				colour,
				engine_cap,
				created_at::date,
				updated_at
	  FROM cars WHERE deleted_at = 0 ` + filter + ``)
	if err != nil {
		return resp, err
	}
	for rows.Next() {
		var (
			car      = model.Car{}
			updateAt sql.NullString
		)

		if err := rows.Scan(
			&resp.Count,
			&car.Id,
			&car.Name,
			&car.Brand,
			&car.Model,
			&car.Year,
			&car.HoursePower,
			&car.Colour,
			&car.EngineCap,
			&car.CreatedAt,
			&updateAt); err != nil {
			return resp, err
		}

		car.UpdatedAt = pkg.NullStringToString(updateAt)
		resp.Cars = append(resp.Cars, car)
	}
	return resp, nil
}
*/
////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////

func (c *carRepo) Deletecar(id string) error {

	query := ` UPDATE cars set
			deleted_at = date_part('epoch', CURRENT_TIMESTAMP)::int
		WHERE id = $1 AND deleted_at=0
	`

	_, err := c.DATA.Exec(query, id)

	if err != nil {
		return err
	}

	return nil

}

func (c *carRepo) Update(car model.Car) (string, error) {

	query := ` UPDATE cars set
			name=$1,
			brand=$2,
			model=$3,
			hourse_power=$4,
			colour=$5,
			engine_cap=$6,
			updated_at=CURRENT_TIMESTAMP
		WHERE id = $7 AND deleted_at=0
	`

	_, err := c.DATA.Exec(query,
		car.Name, car.Brand,
		car.Model, car.HoursePower,
		car.Colour, car.EngineCap, car.Id)

	if err != nil {
		return "", err
	}

	return car.Id, nil
}

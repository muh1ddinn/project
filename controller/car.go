package controller

import (
	model "cars_with_sql/models"
	"cars_with_sql/pkg/check"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func (c Controller) Car(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		c.CreateCars(w, r)
	case http.MethodGet:
		values := r.URL.Query()
		_, ok := values["id"]
		if !ok {
			c.GetAllCars(w, r)
		}
	case http.MethodPut:
		values := r.URL.Query()
		_, ok := values["id"]
		if ok {
			c.UpdateCars(w, r)
		}

	case http.MethodDelete:
		values := r.URL.Query()
		_, ok := values["id"]
		if ok {
			c.DeleteCar(w, r)
		}

	default:
		handleResponse(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}

func (c Controller) CreateCars(w http.ResponseWriter, r *http.Request) {

	car := model.Car{}

	if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
		errStr := fmt.Sprintf("error while request body,err:%v\n", err)

		fmt.Println(errStr)
		handleResponse(w, http.StatusBadRequest, err)
		return

	}
	if err := check.ValidateCarYear(car.Year); err != nil {
		fmt.Println("error while validating year: ", car.Year)
		handleResponse(w, http.StatusBadRequest, err)
		return
	}
	id, err := c.Store.Car().Createcar(car)
	if err != nil {
		fmt.Println("error while creating car,err:", err)
		handleResponse(w, http.StatusInternalServerError, err)
	}

	handleResponse(w, http.StatusOK, id)

}

func (c Controller) UpdateCars(w http.ResponseWriter, r *http.Request) {

	car := model.Car{}

	if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
		errStr := fmt.Sprintf("error while decoding request body,err:%v\n", err)
		fmt.Println(errStr)
		handleResponse(w, http.StatusBadRequest, errStr)
		return
	}

	if err := check.ValidateCarYear(car.Year); err != nil {
		fmt.Println("error while validating year: ", car.Year)
		handleResponse(w, http.StatusBadRequest, err)
		return
	}
	car.Id = r.URL.Query().Get("id")

	err := uuid.Validate(car.Id)
	if err != nil {
		fmt.Println("error while validating, err: ", err)
		handleResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	id, err := c.Store.Car().UpdateCar(car)
	if err != nil {
		fmt.Println("error while creating car, err: ", err)
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusOK, id)
}

func (c Controller) GetAllCars(w http.ResponseWriter, r *http.Request) {
	var (
		values  = r.URL.Query()
		search  string
		request = model.GetAllCarsRequest{}
	)
	if _, ok := values["search"]; ok {
		search = values["search"][0]
	}

	request.Search = search
	page, err := ParsePageQueryParam(r)

	if err != nil {
		fmt.Println("error while parsing limit,err:", err)
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	limit, err := ParseLimitQueryParam(r)

	if err != nil {
		fmt.Println("error while parsing limit ,err:", err)
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	request.Page = page
	request.Limit = limit

	cars, err := c.Store.Car().GetAllCars(request)
	if err != nil {
		fmt.Println("error while getting cars, err: ", err)
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(w, http.StatusOK, cars)
}

func (c Controller) DeleteCar(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	fmt.Println("id: ", id)

	err := uuid.Validate(id)
	if err != nil {
		fmt.Println("error while validating id, err: ", err)
		handleResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	err = c.Store.Car().Deletecar(id)
	if err != nil {
		fmt.Println("error while deleting car, err: ", err)
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusOK, id)
}

func (c Controller) GetbyCar(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	fmt.Println("id: ", id)

	err := uuid.Validate(id)
	if err != nil {
		fmt.Println("error while validating id, err: ", err)
		handleResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	carrs, err := c.Store.Car().GetByIDCar(id)
	if err != nil {
		fmt.Println("error while getting cars, err: ", err)
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(w, http.StatusOK, carrs)
}

/*
func (c *Controller) Createcaaa() {
	car := getCarInfo()

	id, err := c.Store.Car.Create(car)
	if err != nil {
		fmt.Println("error while creating car, err: ", err)
		return
	}
	fmt.Printf("Car created successfully with ID: %v\n", id)

}

func getCarInfo() model.Car {

	car := model.Car{}
	fmt.Println(`enter the car datas
	name,
	brand,
	model,
	year,
	hourse_power,
	colour,
	engine_cap`)

	// var (
	// 	Name, Brand, Model, Colour string
	// 	Year, HoursePower          int
	// 	EngineCap                  float32
	// )

	fmt.Scan(&car.Name, &car.Brand, &car.Model, &car.Year, &car.HoursePower, &car.Colour, &car.EngineCap)

	return car

}

func (c *Controller) Getalll() {

	carrrs, err := c.Store.Car.GetAll()
	if err != nil {
		fmt.Println("error while geting car's info, err: ", err)
		return
	}
	fmt.Printf("Car's all inormations: %v\n", carrrs)

}

func (c *Controller) Getbyidd(id string) {

	carrr, err := c.Store.Car.GetByid(id)
	if err != nil {
		fmt.Println("error while geting car's info, err: ", err)
		return
	}
	fmt.Printf("Car's id inormations: %v\n", carrr)

}

func (c *Controller) Deletee(id string) {

	err := c.Store.Car.Delete(id)
	if err != nil {
		fmt.Println("error while deleting time :", err)
		return
	}

	fmt.Printf("this %vID information was deleted\n", id)
}

func (c *Controller) Updatee() {

	car := updateCarInfo()

	id, err := c.Store.Car.Update(car)
	if err != nil {

		fmt.Println("error while updating time:", err)

		return
	}

	fmt.Printf("Car updated successfully with ID: %v\n", id)

}

func updateCarInfo() model.Car {

	car := model.Car{}
	fmt.Println(`enter the car datas
	name,
	brand,
	model,
	year,
	hourse_power,
	colour,
	engine_cap,
	which id `)

	fmt.Scan(&car.Name, &car.Brand, &car.Model, &car.Year, &car.HoursePower, &car.Colour, &car.EngineCap, &car.Id)

	return car
}
*/

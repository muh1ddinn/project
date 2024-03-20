package handler

import (
	model "cars_with_sql/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func (c Controller) Createcus(w http.ResponseWriter, r *http.Request) {

	cus := model.Customers{}

	if err := json.NewDecoder(r.Body).Decode(&cus); err != nil {
		errstr := fmt.Sprintf("error while decoding request body,err:%v\n", err)
		fmt.Print(errstr)
		handleResponse(w, http.StatusBadRequest, errstr)
		return

	}

	id, err := c.Store.Customer().CreateCus(cus)
	if err != nil {
		fmt.Println("error while creating customer,err:", err)
		return
	}
	handleResponse(w, http.StatusOK, id)

}

func (c Controller) Getallcus(w http.ResponseWriter, r *http.Request) {
	var (
		values  = r.URL.Query()
		search  string
		request = model.GetAllCustomerRequest{}
	)
	if _, ok := values["search"]; ok {
		search = values["search"][0]
	}
	request.Search = search

	cars, err := c.Store.Customer().GetAllCustomers("")
	if err != nil {
		fmt.Println("error while getting customer, err: ", err)
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(w, http.StatusOK, cars)
}

func (c Controller) GetByIDCus(w http.ResponseWriter, r *http.Request) {

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
	id := values["id"][0]

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

	cus, err := c.Store.Customer().GetByIDCustomer(id)
	if err != nil {
		fmt.Println("error while getting customer by id", err)
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}
	handleResponse(w, http.StatusOK, cus)
}

func (c Controller) Deletecus(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	fmt.Println("id: ", id)

	err := uuid.Validate(id)
	if err != nil {
		fmt.Println("error while validating id, err: ", err)
		handleResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	err = c.Store.Customer().DeleteCustomer(id)
	if err != nil {
		fmt.Println("error while deleting car, err: ", err)
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusOK, id)
}

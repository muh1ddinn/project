package controller

import (
	model "cars_with_sql/models"
	"cars_with_sql/pkg/check"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func (c Controller) Customer(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		c.Createcus(w, r)
	case http.MethodGet:
		values := r.URL.Query()
		_, ok := values["id"]
		if !ok {

			c.Getallcus(w, r)
		} else {

			c.GetByIDCus(w, r)
		}
	case http.MethodPut:
		values := r.URL.Query()
		_, ok := values["id"]
		if ok {
			fmt.Print("update")
			c.Upadatecus(w, r)
		}

	case http.MethodDelete:
		values := r.URL.Query()
		_, ok := values["id"]
		if ok {
			c.Deletecus(w, r)
		}

	default:
		handleResponse(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}

func (c Controller) Createcus(w http.ResponseWriter, r *http.Request) {

	cus := model.Customers{}

	if err := json.NewDecoder(r.Body).Decode(&cus); err != nil {
		errstr := fmt.Sprintf("error while decoding request body,err:%v\n", err)
		fmt.Print(errstr)
		handleResponse(w, http.StatusBadRequest, errstr)
		return

	}

	id, err := c.Store.Customer().Create(cus)
	if err != nil {
		fmt.Println("error while creating customer,err:", err)
		return
	}
	handleResponse(w, http.StatusOK, id)

}
func (c Controller) Upadatecus(w http.ResponseWriter, r *http.Request) {

	custommer := model.Customers{}

	if err := json.NewDecoder(r.Body).Decode(&custommer); err != nil {

		errstr := fmt.Sprintf("errpr while decoding request body,err:%v\n", err)
		fmt.Println(errstr)
		handleResponse(w, http.StatusBadRequest, errstr)
		return

	}
	if err := check.Validategmail(custommer.Gmail); err != nil {
		fmt.Println("error while validating gamil: ", custommer.Gmail)
		handleResponse(w, http.StatusConflict, err)
		return
	}
	if err := check.Validatenumber(custommer.Phone); err != nil {

		fmt.Println("error while validating number:", custommer.Phone)
		handleResponse(w, http.StatusBadRequest, err)
		return
	}

	custommer.Id = r.URL.Query().Get("id")
	err := uuid.Validate(custommer.Id)
	if err != nil {
		fmt.Println("error while validating ,err: ", err)
		handleResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	id, err := c.Store.Customer().UpdateCustomer(custommer)
	if err != nil {
		fmt.Println("error while updating customer,err:", err)
		handleResponse(w, http.StatusBadRequest, err.Error())
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

	cars, err := c.Store.Customer().GetAllCustomers(request)
	if err != nil {
		fmt.Println("error while getting cars, err: ", err)
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(w, http.StatusOK, cars)
}

func (c Controller) GetByIDCus(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	id := values["id"][0]

	cus, err := c.Store.Customer().GetByIDCustomer(id)
	if err != nil {
		fmt.Println("error while getting car by id")
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

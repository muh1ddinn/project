package handler

import (
	"cars_with_sql/api /model"
	"encoding/json"
	"fmt"
	"net/http"
)

func (c Controller) Create(w http.ResponseWriter, r *http.Request) {

	cus := model.CreateOrder{}

	if err := json.NewDecoder(r.Body).Decode(&cus); err != nil {
		errstr := fmt.Sprintf("error while decoding request body,err:%v\n", err)
		fmt.Print(errstr)
		handleResponse(w, http.StatusBadRequest, errstr)
		return

	}

	id, err := c.Store.Order().Create(cus)
	if err != nil {
		fmt.Println("error while creating customer,err:", err)
		return
	}
	handleResponse(w, http.StatusOK, id)

}

func (c Controller) GetAll(w http.ResponseWriter, r *http.Request) {
	var (
		values = r.URL.Query()
		search string
	)
	if _, ok := values["search"]; ok {
		search = values["search"][0]
	}

	cars, err := c.Store.Order().GetAll(search)
	if err != nil {
		fmt.Println("error while getting order, err: ", err)
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(w, http.StatusOK, cars)
}

func (c Controller) GetByID(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	id := values["id"][0]

	cus, err := c.Store.Order().GetByID(id)
	if err != nil {
		fmt.Println("error while getting car by id")
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}
	handleResponse(w, http.StatusOK, cus)
}

/*
func (c Controller) Deletecus(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	fmt.Println("id: ", id)

	err := uuid.Validate(id)
	if err != nil {
		fmt.Println("error while validating id, err: ", err)
		handleResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	err = c.Store.Order().
	if err != nil {
		fmt.Println("error while deleting car, err: ", err)
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusOK, id)
}

*/

package controller

import (
	"encoding/json"
	"fmt"
	"lms_backed_pr/model"
	"net/http"

	"github.com/google/uuid"
)

/*
	func (c Controller) Students(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			c.Createstudent(w, r)
		case http.MethodGet:
			values := r.URL.Query()
			_, ok := values["id"]
			if !ok {

				c.Getallstudent(w, r)
			} else {

				c.(w, r)
			}
		case http.MethodPut:
			values := r.URL.Query()
			_, ok := values["id"]
			if ok {
				fmt.Print("update")
				c.(w, r)
			}

		case http.MethodDelete:
			values := r.URL.Query()
			_, ok := values["id"]
			if ok {
				c.Deletestudent(w, r)
			}

		default:
			handleResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		}
	}
*/
func (c Controller) Createstudent(w http.ResponseWriter, r *http.Request) {

	student := model.Student{}

	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		errstr := fmt.Sprintf("error while decoding request body,err:%v\n", err)
		fmt.Print(errstr)
		handleResponse(w, http.StatusBadRequest, errstr)
		return

	}

	id, err := c.Store.Student().CreateStudent(student)
	if err != nil {
		fmt.Println("error while creating customer,err:", err)
		return
	}
	handleResponse(w, http.StatusOK, id)

}
func (c Controller) Upadatestu(w http.ResponseWriter, r *http.Request) {

	student := model.Student{}

	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {

		errstr := fmt.Sprintf("errpr while decoding request body,err:%v\n", err)
		fmt.Println(errstr)
		handleResponse(w, http.StatusBadRequest, errstr)
		return

	}

	student.Id = r.URL.Query().Get("id")
	err := uuid.Validate(student.Id)
	if err != nil {
		fmt.Println("error while validating ,err: ", err)
		handleResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	id, err := c.Store.Student().UpdateStudent(student)
	if err != nil {
		fmt.Println("error while updating customer,err:", err)
		handleResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(w, http.StatusOK, id)

}

func (c Controller) Getallstudent(w http.ResponseWriter, r *http.Request) {
	var (
		values = r.URL.Query()
		search string
	)
	if _, ok := values["search"]; ok {
		search = values["search"][0]
	}
	fmt.Print("gett all")

	cars, err := c.Store.Student().GetallStudent(search)
	if err != nil {
		fmt.Println("error while getting cars, err: ", err)
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(w, http.StatusOK, cars)
}

func (c Controller) GetByIDstudent(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	id := values["id"][0]

	cus, err := c.Store.Student().GetByIDStudent(id)
	if err != nil {
		fmt.Println("error while getting car by id")
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}
	handleResponse(w, http.StatusOK, cus)
}

func (c Controller) Deletestudent(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	fmt.Println("id: ", id)

	err := uuid.Validate(id)
	if err != nil {
		fmt.Println("error while validating id, err: ", err)
		handleResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	err = c.Store.Student().DeleteStudent(id)
	if err != nil {
		fmt.Println("error while deleting car, err: ", err)
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusOK, id)
}

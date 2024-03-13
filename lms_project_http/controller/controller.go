package controller

import (
	"encoding/json"
	"fmt"
	"lms_backed_pr/configg"
	model "lms_backed_pr/model"
	"lms_backed_pr/storage"
	"net/http"
)

type Controller struct {
	Store storage.IStorage
}

func NewController(db storage.IStorage) Controller {
	return Controller{
		Store: db,
	}
}

func handleResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	resp := model.Response{}

	if statusCode >= 100 && statusCode <= 199 {
		resp.Description = configg.ERR_INFORMATION
	} else if statusCode >= 200 && statusCode <= 299 {
		resp.Description = configg.SUCCESS
	} else if statusCode >= 300 && statusCode <= 399 {
		resp.Description = configg.ERR_REDIRECTION
	} else if statusCode >= 400 && statusCode <= 499 {
		resp.Description = configg.ERR_BADREQUEST
	} else {
		resp.Description = configg.ERR_INTERNAL_SERVER
	}
	resp.StatusCode = statusCode
	resp.Data = data

	js, err := json.Marshal(resp)
	if err != nil {
		fmt.Println("error while marshalling: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(statusCode)
	w.Write(js)
}

package main

import (
	"cars_with_sql/config"
	"cars_with_sql/controller"
	"cars_with_sql/storage/postgres"
	"fmt"
	"net/http"
)

func main() {
	cfg := config.Load()
	store, err := postgres.Neww(cfg)
	if err != nil {
		fmt.Println("error while connecting db, err: ", err)
		return
	}
	defer store.DB.Close()

	con := controller.NewController(store)

	http.HandleFunc("/car", con.Car)
	http.HandleFunc("/customers", con.Customer)

	fmt.Println("programm is running on localhost:1700...")
	http.ListenAndServe(":1700", nil)

	/*c.Createcaaa()
	//c.Getbyidd("0658679e-0d24-4db8-8394-af38f95a027e")
	//c.Deletee("020df0f2-3d4c-48be-b182-61a7ff801195")
	c.Updatee()

	*/

}

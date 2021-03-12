package main

import (
	"fmt"
	"rest/orm"
)

func main() {

	orm.CreateConnection()
	orm.CreateTables()

	user := orm.NewUser("Eduardo", "123", "di_564@hotmail.com")

	user.Save()

	users := orm.GetUsers()
	fmt.Println(users)

	user = orm.GetUser(1)
	user.Username = "TEST 1"
	user.Password = "TEST 2"
	user.Email = "TEST 3"
	user.Save()

	fmt.Println(user)

	user.Delete()

	orm.CloseConnection()

	//models.CreateConnection()
	//models.Ping()
	//models.CreateTables()
	//
	//models.CreateUser("diego test 1", "11597", "di_564@hotmail.com")
	//models.CreateUser("diego test 2", "11597", "di_564@hotmail.com")
	//models.CreateUser("diego test 3", "11597", "di_564@hotmail.com")
	//
	//user := models.GetUser(3)
	//fmt.Println(user)
	//
	//users := models.GetUsers()
	//fmt.Println(users)
	//
	//models.CloseConnection()

	//mux := mux.NewRouter()
	//models.SetDefaultUser()
	//
	//mux.HandleFunc("/api/v1/users/", handlers.GetUsers).Methods("GET")
	//mux.HandleFunc("/api/v1/users/{id:[0-9]+}", handlers.GetUser).Methods("GET")
	//mux.HandleFunc("/api/v1/users/", handlers.CreateUser).Methods("POST")
	//mux.HandleFunc("/api/v1/users/{id:[0-9]+}", handlers.UpdateUser).Methods("PUT")
	//mux.HandleFunc("/api/v1/users/{id:[0-9]+}", handlers.DeleteUser).Methods("DELETE")
	//
	//log.Println("El servidor ala esucha en el puerto 8000")
	//log.Fatal(http.ListenAndServe(":8000", mux))

}

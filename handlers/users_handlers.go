package handlers

import (
	"github.com/Abdelwhab-1/book_store_user_service-/model/data"
	"github.com/Abdelwhab-1/book_store_user_service-/service"
	"github.com/Abdelwhab-1/book_store_user_service-/utils/data"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func getid(r *http.Request) (int64, error) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	return id, err
}
func GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := getid(r)
	if err != nil {
		//TODO : should handle the error here
		log.Fatal("error " ,err)
	}
	user := &data.User{Id: id}
	user, err = services.UserService.Get(user)
	if err != nil {
		//TODO: should handle the error here
		log.Fatal(err)
	}
	datamarshal.MarshUser(w,user)
	w.WriteHeader(http.StatusOK)
}
func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := &data.User{}
	datamarshal.UnmarshUser(r , user)
	user , err := services.UserService.Create(user)
	if err != nil {
		//TODO : should handle that error
	}
	datamarshal.MarshUser(w,user)
	// TODO :  should deal with the headrs and the status code .
	w.WriteHeader(http.StatusOK)
}
func EditUser(w http.ResponseWriter, r *http.Request) {

	id, err := getid(r)
	if err != nil {
		//TODO : should handle the error here
		log.Fatal(err)
	}

	var user *data.User
	datamarshal.UnmarshUser(r, user)
	user.Id = id
	user , err = services.UserService.Update(user)
	if err != nil{
		log.Fatal(err)
		//TODO : should handle this erorr
	}
	datamarshal.MarshUser(w,user)
	w.WriteHeader(http.StatusOK)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id , err := getid(r)
	if err != nil{
		log.Fatal(err)
		// TODO: should handle this errror
	}
	if err:= services.UserService.Delete(id); err != nil {
		log.Fatal(err)
		//TODO : shuld handle this error the right way
	}
	//TODO : should sned the right status code 
	w.WriteHeader(http.StatusOK)
}

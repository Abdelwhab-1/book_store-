package services 

import (
	"github.com/Abdelwhab-1/book_store_user_service-/model/data"
)
var(
	UserService serviceinterface = &userservices{}
)
type serviceinterface interface {
	Get(user *data.User)(*data.User , error)
	Create(user *data.User)(*data.User , error)
	Update(user *data.User)(*data.User, error)
	Delete( id int64)(error)
}
type userservices struct {
	
}
func(s *userservices)Get(user *data.User)(*data.User , error){
	db := data.User_db
	dbi := data.Dbinterface{db}
	return dbi.FindUserById(user)
}
func(s *userservices)Create(user *data.User)(*data.User , error){
	db := data.User_db
	dbi := data.Dbinterface{db}
	_, user , err := dbi.CreatUser(user)
	return user,err
}

func(s *userservices)Update(user *data.User)(*data.User, error){
	db := data.User_db
	dbi := data.Dbinterface{db}
	_, user , err := dbi.UpdateUser(user)
	return user , err
}
func(s *userservices)Delete(id int64)(error){
	db := data.User_db
	dbi := data.Dbinterface{db}
	_,err := dbi.DeletUser(id)
	return err
}
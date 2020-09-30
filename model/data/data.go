package data

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)
var User_db *sql.DB
type User struct {
	Id       int64 `json:"id ometempty"`
	Password int
	Fname    string `json:"firstName"`
	Lname    string `json:"lastName"`
	Email    string `jsone:"email"`
	Phone    int    `json:"phone"`
}
type Dbinterface struct {
	Db *sql.DB
}

func (dbi Dbinterface) FindUserById(user *User) (User *User, err error) {
	row := dbi.Db.QueryRow("SELECT fname,lname,email,phone FROM users WHERE id=$1",user.id)
	err = row.Scan(&user.Fname, &user.Lname, &user.Email, &user.Phone)
	if err != nil {
		return nil, err
	}
	return User , nil
}
func (dbi Dbinterface) CreatUser(user *User) (*sql.Result,*User, error) {
	stmt, err := dbi.Db.Prepare("INSERT INTO users(password,fname,lname,email,phone)VALUES($1,$2,$3,$4,$5)")
	if err != nil {
		log.Fatal(err)
		return nil, nil ,err
	}
	result, err := stmt.Exec(user.Password, user.Fname, user.Lname, user.Email, user.Phone)
	if err != nil {
		log.Fatal(err)
		return nil, nil, err
	}
	return &result,user , err
}
func(dbi Dbinterface) UpdateUser(user *User)(*sql.Result,*User, error){
	stmt , err := dbi.Db.Prepare("UPDATE users SET password=$1 ,fname=$2 ,lname=$3 ,email=$4,phone=$5 WHERE  id = $6")
	if err != nil {
		return nil , nil ,err
	}
	res , err := stmt.Exec(user.Password,user.Fname,user.Lname,user.Email,user.Phone,user.Id)
	if err != nil {
		return nil , nil ,err
	}
	return &res , user , err
}
func(dbi Dbinterface)DeletUser(id  int64)(*sql.Result,error){
	stmt , err := dbi.Db.Prepare("DELETE FROM users  WHERE id = $1")
	if err != nil {
		return nil , err
	}
	res , err := stmt.Exec(id)
	if err != nil {
		return nil , err
	}
	return &res ,err
}
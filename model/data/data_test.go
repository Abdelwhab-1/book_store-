package data

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"log"
	"testing"
)

func newmock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	return db, mock
}

func TestFinduserbyid(t *testing.T) {
	//initiation
	u := User{
		Id:       2139850275894,
		Password: 1234793857,
		Fname:    "abdelwhab",
		Lname:    "elassfer",
		Email:    "abdo@gmail.com",
		Phone:    01063260347}

	db, mock := newmock()
	dbi := Dbinterface{Db: db}

	query := "SELECT  fname,lname,email,phone FROM users WHERE id=\\$1"
	rows := sqlmock.NewRows([]string{"fname", "lname", "email", "phone"}).
		AddRow(u.Fname, u.Lname, u.Email, u.Phone)

	mock.ExpectQuery(query).WillReturnRows(rows)
	//exeqution
	user, err := dbi.FindUserById(2139850275894)
	//validation
	if err != nil || user.Fname != u.Fname {
		fmt.Println(rows)
		t.Errorf("error : %v", err)
	}
	fmt.Println(user.Id)

}

func TestCreatuser(t *testing.T) {
    //init
	db, mock := newmock()
	dbi := Dbinterface{Db: db}
	u := &User{
		Password: 1234793857,
		Fname:    "abdelwhab",
		Lname:    "elassfer",
		Email:    "abdo@gmail.com",
		Phone:    01063260347}
	query := "INSERT INTO users\\(password,fname,lname,email,phone\\)VALUES\\(\\$1,\\$2,\\$3,\\$4,\\$5\\)"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(u.Password, u.Fname, u.Lname, u.Email, u.Phone).
		WillReturnResult(sqlmock.NewResult(1, 0))
	//Execute
	result, us ,err := dbi.CreatUser(u)
	// Validate
	if err != nil || us.Email != u.Email{
		fmt.Println(result)
		t.Fatal(err)

	}

}

func TestUpdateuser(t *testing.T){
  db , mock := newmock()
  dbi := Dbinterface{Db: db}
  u := &User{
    Password: 1234793857,
    Fname:    "abdelwhab",
    Lname:    "elassfer",
    Email:    "abdo@gmail.com",
    Phone:    01063260347}
  query := "DELETE FORM users WHERE id =$1"
  stmt := mock.ExpectPrepare(query)
  stmt.ExpectExec().WithArgs(u.Id).WillReturnResult(sqlmock.NewResult(0,1 ))
  res ,us ,err := dbi.UpdateUser(u)
  fmt.Println(res,us,err)
}

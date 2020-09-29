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
	if err != nil  {
		if user != nil{
			t.Errorf("expecting user to be nil found %v",user)
		}
	}

}

func TestFinduserbyiderr(t *testing.T) {
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
	rows := sqlmock.NewRows([]string{"id","fname", "lname", "email", "phone"}).
		AddRow(u.Id,u.Fname, u.Lname, u.Email, u.Phone)

	mock.ExpectQuery(query).WillReturnRows(rows)
	//exeqution
	user, err := dbi.FindUserById(32445)
	//validation

	if err == nil  {
		t.Errorf("should hav  error  we could not find  the user %v", user)
	}


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

func TestUpdateusererr(t *testing.T){
 db , mock := newmock()
 dbi := Dbinterface{Db: db}
 u := &User{
   Password: 1234793857,
   Fname:    "abdelwhab",
   Lname:    "elassfer",
   Email:    "abdo@gmail.com",
   Phone:    01063260347}
 query := "UPDATE users SET password=\\$1 ,fname=\\$2 ,lname=\\$3 ,email=\\$4,phone=\\$5 WHERE  id = \\$6"
 stmt := mock.ExpectPrepare(query)
 stmt.ExpectExec().WithArgs(u.Password,u.Fname,u.Lname,u.Email,u.Phone,47512783481295).WillReturnResult(sqlmock.NewResult(0,0 ))
 _ ,_ ,err := dbi.UpdateUser(u)

 if err == nil {
 	t.Errorf("should have an error %v",err)
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
	query := "UPDATE users SET password=\\$1 ,fname=\\$2 ,lname=\\$3 ,email=\\$4,phone=\\$5 WHERE  id = \\$6"
	stmt := mock.ExpectPrepare(query)
	stmt.ExpectExec().WithArgs(u.Password,u.Fname,u.Lname,u.Email,u.Phone,u.Id).WillReturnResult(sqlmock.NewResult(0,1 ))
	res ,u ,err := dbi.UpdateUser(u)
	if err != nil || u == nil || res == nil {
		t.Error("coudln't update the user enternal error ")
	}
}
func  TestDeleteuser(t *testing.T){
	u := &User{
		Password: 1234793857,
		Fname:    "abdelwhab",
		Lname:    "elassfer",
		Email:    "abdo@gmail.com",
		Phone:    01063260347}
	db , mock := newmock()
	dbi := Dbinterface{db}
	stmt := mock.ExpectPrepare("DELETE FROM users  WHERE id = \\$1")
	stmt.ExpectExec().WithArgs(u.Id).WillReturnResult(sqlmock.NewResult(0,1))
	_ ,err := dbi.DeletUser(u.Id)

	if err  != nil {
		t.Errorf("should have passed but have error %s",err)
	}
}
func  TestDeleteusererr(t *testing.T){
	u := &User{
		Id: 123456789,
		Password: 1234793857,
		Fname:    "abdelwhab",
		Lname:    "elassfer",
		Email:    "abdo@gmail.com",
		Phone:    01063260347}
	db , mock := newmock()
	dbi := Dbinterface{db}
	stmt := mock.ExpectPrepare("DELETE FROM users  WHERE id = \\$1")
	stmt.ExpectExec().WithArgs(59134589732598).WillReturnResult(sqlmock.NewResult(0,1))
	_ ,err := dbi.DeletUser(u.Id)
	if err  == nil {
		t.Errorf("should have  error  %v",err)
	}
}


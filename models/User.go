package models

import (
	"database/sql"
	"dev-platform/drivers"
	"log"
)

// models package's db obj
// all db operation should be done in models pkg
// so db is a pkg inner var
var db *sql.DB = drivers.MysqlDb

type User struct {
	Id   int    `json:"id" form:"id" primaryKey:"true"`
	Name string `json:"name" form:"name" binding:"required"`
	Email  string    `json:"email" form:"email" binding:"required"`
	Password	string	`json:"password" form:"password"`

}

// get one
func (model *User) UserGet(id int) (user User, err error) {
	// find one record
	err = db.QueryRow("SELECT `id`, `name`, `email` FROM `users` WHERE `id` = ?", id).Scan(&user.Id, &user.Name, &user.Email)

	if err != nil {
		log.Println(err.Error())
		return
	}

	return
}

// get Info
func (model *User) UserGetByEmail(email string) (user User, err error) {
	// find one record
	err = db.QueryRow("SELECT `id`, `name`, `email`, `password` FROM `users` WHERE `email` = ?", email).Scan(&user.Id, &user.Name, &user.Email, &user.Password )

	if err != nil {
		log.Println(err.Error())
		return
	}

	return
}

// get list
func (model *User) UserGetList(page int, pageSize int) (users []User, err error) {
	users = make([]User, 0)

	offset := pageSize * (page - 1)
	limit := pageSize

	rows, err := db.Query("SELECT `id`, `name` FROM `users` LIMIT ?, ?", offset, limit)
	defer rows.Close()

	if err != nil {
		log.Println(err.Error())
		return
	}

	var user User
	for rows.Next() {
		rows.Scan(&user.Id, &user.Name, &user.Email)
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		log.Println(err.Error())
	}

	return
}

// create
func (model *User) UserAdd() (id int64, err error) {
	result, err := db.Exec("INSERT INTO `users`(`name`, `Email`) VALUES (?, ?)", model.Name, model.Email)

	if nil != err {
		log.Println(err.Error())
		return
	}

	id, err = result.LastInsertId()
	return
}

// update
func (model *User) UserUpdate(id int) (afr int64, err error) {
	result, err := db.Exec("UPDATE `users` SET `name` = ?, `email` = ? WHERE `id` = ?", model.Name, model.Email, id)

	if nil != err {
		log.Println(err.Error())
		return
	}

	afr, err = result.RowsAffected()
	return
}

// delete
func (model *User) UserDelete(id int) (afr int64, err error) {
	result, err := db.Exec("DELETE FROM `users` WHERE `id` = ?", id)

	if nil != err {
		log.Println(err.Error())
		return
	}

	afr, err = result.RowsAffected()
	return
}
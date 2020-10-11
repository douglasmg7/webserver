package models

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID                uint      `db:"id"`
	Email             string    `db:"email"`
	Name              string    `db:"name"`
	ContactName       string    `db:"contact_name"`
	RG                string    `db:"rg"`
	CPF               string    `db:"cpf"`
	CNPJ              string    `db:"cnpj"`
	StateRegistration string    `db:"state_registration"`
	MobileNumber      string    `db:"mobile_number"`
	Password          string    `db:"password"`
	Permission        string    `db:"permission"`
	Status            uint      `db:"status"`
	CreatedAt         time.Time `db:"created_at"`
	ChangedAt         time.Time `db:"changed_at"`
	IsDeleted         bool      `db:"is_deleted"`
}

// Check user password
func (u *User) CheckPassword(password string) bool {
	// Compare password.
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	// Incorrect password.
	if err != nil {
		return false
	}
	// Correct password.
	return true
}

// Get user
func GetUser(db sqlx.DB, id string) (user *User, ok bool) {
	switch err := db.Get(&user, "select * from users where id=$1", id); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return
	case nil:
		ok = true
		return
	default:
		panic(err)
	}
}

// // Check user password
// func (u *User) CheckPassword(db sqlx.DB, userID uint, password string) bool {
// var userPassword string
// // Replace 3 with an ID from your database or another random
// // value to test the no rows use case.
// row := db.QueryRow("select password from users where id=$1", userID)
// switch err := row.Scan(&userPassword); err {
// case sql.ErrNoRows:
// fmt.Println("No rows were returned!")
// return false
// case nil:
// fmt.Printf("retrived password: %v", userPassword)
// if password == userPassword {
// return true
// }
// return false
// default:
// panic(err)
// }
// }

type UserGroups uint8

const (
	GroupRoot UserGroups = 1 << iota
	GroupAdmin
	GorupSeller
)

// Set.
func (g *UserGroups) Set(flag UserGroups) {
	// r := *g | flag
	// g = &r
	*g = *g | flag
}

// Clear.
func (g *UserGroups) Clear(flag UserGroups) {
	*g = *g &^ flag
}

// Toggle.
func (g *UserGroups) Toggle(flag UserGroups) {
	*g = *g ^ flag
}

// Has.
func (g UserGroups) Has(flag UserGroups) bool {
	return (g & flag) != 0
}

// Has.
func (g UserGroups) Hav() bool {
	// *g & flag
	return false
}

// Is Admin.
func (g UserGroups) IsAdmin() bool {
	return g&GroupAdmin != 0
}

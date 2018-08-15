package models

import (
	"time"

	// "github.com/sepulsa/rest_echo/api/models"
	"github.com/sepulsa/rest_echo/api/models/orm"
)

type (
	User struct {
		BaseModel
		Name  string `json:"name"`
		Email string `json:"email" valid:"email"`
	}

	UserPaginationResponse struct {
		Meta orm.PaginationResponse `json:"meta"`
		Data []User                 `json:"data"`
	}

	// just use string type, since it will be use on query at DB layer
	UserFilterable struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
)

var (
	_page = 1
	_rp   = 25
)

// Callback before update user
func (m *User) BeforeUpdate() (err error) {
	m.UpdatedAt = time.Now()
	return
}

// Callback before create user
func (m *User) BeforeCreate() (err error) {
	m.CreatedAt = time.Now()
	return
}

// Create
func Create(m *User) (*User, error) {
	var err error
	err = orm.Create(&m)
	return m, err
}

// Update
func (m *User) Update() error {
	var err error
	err = orm.Save(&m)
	return err
}

// Delete
func (m *User) Delete() error {
	var err error
	err = orm.Delete(&m)
	return err
}

// FindUserByID
func FindUserByID(id int) (User, error) {
	var (
		user User
		err  error
	)
	err = orm.FindOneByID(&user, id)
	return user, err
}

// FindAllUsers
func FindAllUsers(page int, rp int, filters interface{}) (interface{}, error) {
	var (
		users []User
		err   error
	)

	resp, err := orm.FindAllWithPage(&users, page, rp, filters)
	return resp, err
}

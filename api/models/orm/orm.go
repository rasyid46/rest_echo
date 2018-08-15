package orm

/*
 * 	Custom ORM package
 * 	Based on gorm ORM.
 * 	Make sure the every database actions within the transaction.
 * 	Do begin, rollback and commit the db transaction automatically.
 * 	Return proper error
 *
 */

import (
	"fmt"
	"math"
	"reflect"

	cgorm "rest_echo/db/gorm"

	"github.com/jinzhu/gorm"
)

type (
	DBFunc func(tx *gorm.DB) error // func type which accept *gorm.DB and return error

	PaginationResponse struct {
		Total       int         `json:"total"`
		PerPage     int         `json:"per_page"`
		CurrentPage int         `json:"current_page"`
		LastPage    int         `json:"last_page"`
		From        int         `json:"from"`
		To          int         `json:"to"`
		Data        interface{} `json:"data"`
	}
)

var (
	total_rec int
)

// Create
// Helper function to insert gorm model to database by using 'WithinTransaction'
func Create(v interface{}) error {
	return WithinTransaction(func(tx *gorm.DB) (err error) {
		// check new object
		if !cgorm.DBManager().NewRecord(v) {
			return err
		}
		if err = tx.Create(v).Error; err != nil {
			tx.Rollback() // rollback
			return err
		}
		return err
	})
}

// Save
// Helper function to save gorm model to database by using 'WithinTransaction'
func Save(v interface{}) error {
	return WithinTransaction(func(tx *gorm.DB) (err error) {
		// check new object
		if cgorm.DBManager().NewRecord(v) {
			return err
		}
		if err = tx.Save(v).Error; err != nil {
			tx.Rollback() // rollback
			return err
		}
		return err
	})
}

// Delete
// Helper function to save gorm model to database by using 'WithinTransaction'
func Delete(v interface{}) error {
	return WithinTransaction(func(tx *gorm.DB) (err error) {
		// check new object
		if err = tx.Delete(v).Error; err != nil {
			tx.Rollback() // rollback
			return err
		}
		return err
	})
}

// FindOneByID
// Helper function to find a record by using 'WithinTransaction'
func FindOneByID(v interface{}, id int) (err error) {
	return WithinTransaction(func(tx *gorm.DB) error {
		if err = tx.Last(v, id).Error; err != nil {
			tx.Rollback() // rollback db transaction
			return err
		}
		return err
	})
}

// FindAll
// Helper function to find records by using 'WithinTransaction'
func FindAll(v interface{}) (err error) {
	return WithinTransaction(func(tx *gorm.DB) error {
		if err = tx.Find(v).Error; err != nil {
			tx.Rollback() // rollback db transaction
			return err
		}
		return err
	})
}

// FindOneByQuery
// Helper function to find a record by using 'WithinTransaction'
func FindOneByQuery(v interface{}, params map[string]interface{}) (err error) {
	return WithinTransaction(func(tx *gorm.DB) error {
		if err = tx.Where(params).Last(v).Error; err != nil {
			tx.Rollback() // rollback db transaction
			return err
		}
		return err
	})
}

// FindByQuery
// Helper function to find records by using 'WithinTransaction'
func FindByQuery(v interface{}, params map[string]interface{}) (err error) {
	return WithinTransaction(func(tx *gorm.DB) error {
		if err = tx.Where(params).Find(v).Error; err != nil {
			tx.Rollback() // rollback db transaction
			return err
		}
		return err
	})
}

// FindAllWithPage
// Helper function to find all records in pagination by using 'WithinTransaction'
// v interface{}	Gorm model struct
// page int	Page number
// rp int	Record per page to be showed
// filters int	Gorm model struct for filters
func FindAllWithPage(v interface{}, page int, rp int, filters interface{}) (resp PaginationResponse, err error) {
	var (
		offset   int
		lastPage int = 1
	)

	if rp <= 0 {
		rp = 25
	}

	// tx := cgorm.DBManager().Begin()
	tx := cgorm.DBManager()

	// loop through "filters"
	refOf := reflect.ValueOf(filters).Elem()
	typeOf := refOf.Type()
	for i := 0; i < refOf.NumField(); i++ {
		f := refOf.Field(i)
		// ignore if empty
		// just make sure ModelFilterable its all in string type
		if f.Interface() != "" {
			tx = tx.Where(fmt.Sprintf("%s = ?", typeOf.Field(i).Name), f.Interface())
		}
	}

	// copy of tx
	ctx := tx

	// get total record include filters
	ctx.Find(v).Count(&total_rec)

	offset = (page * rp) - rp

	lastPage = int(math.Ceil(float64(total_rec) / float64(rp)))

	tx.Limit(rp).Offset(offset).Find(v)

	resp = PaginationResponse{
		Total:       total_rec,
		PerPage:     rp,
		CurrentPage: page,
		LastPage:    lastPage,
		From:        offset + 1,
		To:          offset + rp,
		Data:        &v,
	}
	if err != nil {
		// tx.Rollback() // rollback db transaction
		return resp, err
	}

	// tx.Commit()

	return resp, err
}

// WithinTransaction
// accept DBFunc as parameter
// call DBFunc function within transaction begin, and commit and return error from DBFunc
func WithinTransaction(fn DBFunc) (err error) {
	tx := cgorm.DBManager().Begin() // start db transaction
	defer tx.Commit()
	err = fn(tx)
	// close db transaction
	return err
}

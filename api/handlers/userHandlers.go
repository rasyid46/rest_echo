package handlers

import (
	"net/http"
	"os"
	"strconv"

	"rest_echo/api/models"
	"rest_echo/db/gorm"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"github.com/thedevsaddam/govalidator"
)

type Credentials struct {
	Password string `json:"password", db:"password"`
	Username string `json:"username", db:"username"`
}

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
}

func GetUsers(c echo.Context) error {
	rp, err := strconv.Atoi(c.QueryParam("rp"))
	page, err := strconv.Atoi(c.QueryParam("p"))
	name := c.QueryParam("name")
	email := c.QueryParam("email")

	defer c.Request().Body.Close()

	rules := govalidator.MapData{
		"rp":    []string{"numeric"},
		"page":  []string{"numeric"},
		"name":  []string{"alpha_num"},
		"email": []string{"email"},
	}

	vld := ValidateQueryStr(c, rules)
	if vld != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, vld)
	}

	result, err := models.FindAllUsers(page, rp, &models.UserFilterable{name, email})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func GetUserById(c echo.Context) error {
	id, err := strconv.Atoi(c.QueryParam("id"))

	defer c.Request().Body.Close()

	rules := govalidator.MapData{
		"id": []string{"numeric"},
	}

	vld := ValidateQueryStr(c, rules)
	if vld != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, vld)
	}

	result, err := models.FindUserByID(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func AddUser(c echo.Context) error {
	user := models.User{}

	defer c.Request().Body.Close()

	rules := govalidator.MapData{
		"name":  []string{"required"},
		"email": []string{"required", "email"},
	}

	vld := ValidateRequest(c, rules, &user)
	if vld != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, vld)
	}

	checkEmail := gorm.DBManager().Where("email = ?", user.Email).Find(&user)
	if checkEmail.RowsAffected > 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "email sudah ada")
	}

	result, err := models.Create(&user)
	if err != nil {
		log.Printf("FAILED TO CREATE : %s\n", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to create new user")
	}
	// result := &user
	return c.JSON(http.StatusCreated, result)
}

func EditUser(c echo.Context) error {
	id, err := strconv.Atoi(c.QueryParam("id"))

	defer c.Request().Body.Close()

	rules := govalidator.MapData{
		"name":  []string{},
		"email": []string{"email"},
	}

	user, err := models.FindUserByID(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	vld := ValidateRequest(c, rules, &user)
	if vld != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, vld)
	}

	c.Bind(&user)

	err = user.Update()
	if err != nil {
		log.Printf("FAILED TO UPDATE: %s\n", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to update user")
	}

	return c.JSON(http.StatusOK, user)
}

func DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.QueryParam("id"))

	defer c.Request().Body.Close()

	rules := govalidator.MapData{
		"id": []string{"required", "numeric"},
	}

	vld := ValidateQueryStr(c, rules)
	if vld != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, vld)
	}

	user, err := models.FindUserByID(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	c.Bind(&user)

	err = user.Delete()
	if err != nil {
		log.Printf("FAILED TO DELETE: %s\n", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to delete user")
	}

	return c.JSON(http.StatusOK, user)
}
func LoginUser(c echo.Context) error {
	user := models.User{}

	defer c.Request().Body.Close()

	rules := govalidator.MapData{
		"password": []string{"required"},
		"email":    []string{"required", "email"},
	}

	vld := ValidateRequest(c, rules, &user)
	if vld != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, vld)
	}
	passrequest := user.Password
	checkEmail := gorm.DBManager().Where("email = ?", user.Email).Find(&user)
	if checkEmail.RowsAffected > 0 {
		passDB := user.Password
		match := models.CheckPasswordHash(passrequest, passDB)
		if match == false {
			return echo.NewHTTPError(http.StatusBadRequest, "password salah")
		}
		return c.JSON(http.StatusCreated, match)
	} else {
		return echo.NewHTTPError(http.StatusBadRequest, "email tidak terdaftar")
	}
}

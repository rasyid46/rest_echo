# TRYING ECHO

Setup REST API with [Echo Framework](https://echo.labstack.com/guide/migration), Postgre and Messaging.

## Package Dependencies

Below is the packages used by this project

* Framework : https://github.com/labstack/echo
* ORM : https://github.com/jinzhu/gorm
* DB Driver :
    - https://github.com/lib/pq (for Postgres)
    - https://github.com/go-sql-driver/mysql (for MySQL)
* Configs:
    - main : https://github.com/spf13/viper
    - listener : https://github.com/fsnotify/fsnotify
* Request Validator : github.com/thedevsaddam/govalidator
* Unit Test : https://github.com/gavv/httpexpect
* Logger :
    - Logrus: https://github.com/sirupsen/logrus
    - Rotator: https://github.com/lestrrat-go/file-rotatelogs

## Configs

Config file located in ```config.yaml``` on the root of project. 

## ORM

Ready for Postgres and MySQL. Connection function located at ```db/gorm/gorm.go```. Those function called in ```main.go```.

### CRUD Functionality

Base CRUD function are located in ```models/orm/orm.go```. Example implementation within each model can be found at ```models/user.go```

### Models

Models located in ```models```. All models should inherit ```BaseModel struct``` in ```models/base.go```. ```BaseModel``` struct are holding default tables attribute. Another attributes specific to each model should defined as struct in each model file (i.e ```models/user.go```)

## Request Validator

Currently there is two validation function for request.

### Validate JSON Body

Function for validate this kind of request is ```ValidateRequest(c echo.Context, rules govalidator.MapData, data interface{}) map[string]interface{}```, located in ```api/handlers/generalHandlers```

### Validate URL Query String

Function for validate this kind of request is ```ValidateQueryStr(c echo.Context, rules govalidator.MapData, data interface{}) map[string]interface{}```, located in ```api/handlers/generalHandlers```

### Limitation

You cannot use single validation for both query string and/or request body

## Logger

Logrus are wrapped within ```modules/logger/logger.go``` under ```logger``` package. This wrapper also implement file rotator.

Default rotator time is ```every one day``` and only kept for ```seven days of log```.

Example implementation can be found at ```api/middlewares/logMiddlewares.go```

## Unit Test

This unit test is for testing each endpoint and response. All test file located in ```tests``` folder. You can run a test with

```bash
go test tests/*_test.go
```

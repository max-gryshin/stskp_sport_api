
# postgres connection
[pgx postgres driver](https://habr.com/ru/company/oleg-bunin/blog/461935/)
    buff size limit 4KB!
    
# pretty project for example
[photoprism](https://github.com/photoprism/photoprism)

# project layout
[project layout](https://github.com/golang-standards/project-layout)

# awesome go
[awesome go](https://awesome-go.com/)

# go gin contrib
[go gin contrib](https://github.com/gin-gonic/contrib)

#go jwt 
[go jwt](https://github.com/appleboy/gin-jwt) 

# sqlx
[sqlx](https://github.com/jmoiron/sqlx)
# Plane
## Todo 	
- API
- Documentation
- Tests - show example how to write test for routes
- Business logic
- Migrations
- Basic role system
- swag docs can't set query params with string\[string]=string

## Resolved
- routes: create, login
- middleware: jwt
- business logic
    - authorization
    - model of user, workout, workout_type, workout_value
    - password checking
- Swagger documentation: installed, configured and set basic realization
- sqlx: extension of database/sql for better life
- basic query builder (find, orderBy, offsetRows, andWhere, maxResult method in repository)
- in FindBy - set flexible field selection
- parser query parameters
#Business logic
### basic entities
 - workout
 - workout type
 - workout value
## Workout scheme
##### workout_type - `(run, gym)` and
##### workout_template - `(workout scheme of certain units for workout type)`
имеют конечное множество значений, которое можно хранить в кеше (redis).
##### workout_template - буду хранить в map'е, при необходимости вынесу в BD или кеш

# Documentation
##### Swagger
 - install
 - configure
 - basic realization
 
#API
 - routes
 - cors
 - http 2.0
 
# In progress
 - Business logic
    - CRUD user, workout, workout_type, workout_value
    - create routes post, get, patch, delete for user, workout, workout_type, workout_value
    - create methods to handle crud with db
    - describe a thing like DTO in symfony (to hide private fields as password)
    - describe methods for basic api functional 
        - look at work project
        - getById, update, sort, filter, per page (offset),        
        - get all
        - get by id
        - get by a field ?
        - post
        - patch
        - delete
        - order by
        - filter
    - algorithm
        - create route
        - create method in repository
        - test it
        - create annotation
        - generate swagger (swag init)
   
# make flexible query:
[gin-examples](https://riptutorial.com/go/example/29299/restfull-projects-api-with-gin)
 - example:
   ```
   {
      "criteria": {
         "fieldName": [">", "1"]
      },
      "limit": 0,
      "offset": 0,
      "order": {"fieldName": "DESC"}
   }
   ```
# swag annotation examples
```
/**
* @api {post} /v1/auth/login Login
* @apiGroup Users
* @apiHeader {application/json} Content-Type Accept application/json
* @apiParam {String} username User username
* @apiParam {String} password User Password
* @apiParamExample {json} Input
*    {
*      "username": "your username",
*        "password"     : "your password"        
*    }
* @apiSuccess {Object} authenticate Response
* @apiSuccess {Boolean} authenticate.success Status
* @apiSuccess {Integer} authenticate.statuscode Status Code
* @apiSuccess {String} authenticate.message Authenticate Message
* @apiSuccess {String} authenticate.token Your JSON Token
* @apiSuccessExample {json} Success
*    {
*        "authenticate": {     
*               "statuscode": 200,
*              "success": true,
*           "message": "Login Successfully",
*              "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ"
*            }
*      }
* @apiErrorExample {json} List error
*    HTTP/1.1 500 Internal Server Error
*/
```
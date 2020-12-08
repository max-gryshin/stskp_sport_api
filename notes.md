
# postgres connection
[pgx postgres driver](https://habr.com/ru/company/oleg-bunin/blog/461935/)
    buff size limit 4KB!
    
# pretty project for example
[photoprism](https://github.com/photoprism/photoprism)

# project layout
[project layout](https://github.com/golang-standards/project-layout)

# awesome go
[awesome go](https://awesome-go.com/)

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

# postgres connection
[pgx postgres driver](https://habr.com/ru/company/oleg-bunin/blog/461935/)
    buff size limit 4KB!
    
# pretty project for example
[photoprism](https://github.com/photoprism/photoprism)

# project layout
[project layout](https://github.com/golang-standards/project-layout)

# Plane
## Todo 	
- API
- Documentation
- Tests - show example how to write test for routes
- Business logic
- Migrations
- Basic role system

## Resolved
- routes: create, login
- middleware: jwt
- business logic
    - authorization
    - model of user, workout, workout_type, workout_value
    - password checking
- Swagger documentation: installed, configured and set basic realization 

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
 